package modules

import (
	"fmt"
	"git.quartzinc.dev/Zertex/XDGv2/injection"
	"git.quartzinc.dev/Zertex/XDGv2/manager"
	"git.quartzinc.dev/Zertex/XDGv2/qtui"
	"git.quartzinc.dev/Zertex/XDGv2/utils"
	"github.com/spf13/viper"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

type AnalyzerModule struct {
	Injectable *injection.Injection
}

var Analyzer *AnalyzerModule

func NewAnaylzer() *AnalyzerModule {
	return &AnalyzerModule{}
}

func (a *AnalyzerModule) AddDBToUI(dbIndex int) {
	dbItem := qtui.SingleSiteWindow.SingleDatabaseStructure.TopLevelItem(dbIndex)
	dbItem.SetText(0, core.QCoreApplication_Translate("SingleSiteAnalyzer", fmt.Sprintf("%s - %d Tables", a.Injectable.Structure[dbIndex].Name, len(a.Injectable.Structure[dbIndex].Tables)), "", 0))
	//___qtreewidgetitem2 := ___qtreewidgetitem1.Child(0)
}

func (a *AnalyzerModule) Mount(inj *injection.Injection) {
	a.Injectable = inj
	qtui.SingleSiteWindow.SingleDatabaseStructure.Reset()
	qtui.SingleSiteWindow.SingleDatabaseStructure.Clear()
	qtui.SingleSiteWindow.SingleDumpOutput.Clear()
	qtui.SingleSiteWindow.SingleDumpOutput.SetRowCount(0)
	qtui.SingleSiteWindow.SingleSiteText.SetText(inj.ToFormattedString())
	qtui.SingleSiteWindow.SingleVector.SetText(injection.GetVectorName(inj.Vector))
	qtui.SingleSiteWindow.SingleTechnique.SetText(injection.TechniqueString(inj.Technique))
	qtui.SingleSiteWindow.SingleDatabaseType.SetText(injection.DBMSString(inj.DBType))
	qtui.SingleSiteWindow.SingleCountry.SetText("Loading...")
	c, err := inj.GetCountry()
	if err == nil {
		qtui.SingleSiteWindow.SingleCountry.SetText(c.Long)
	} else {
		qtui.SingleSiteWindow.SingleCountry.SetText("Failed")
	}
	qtui.SingleSiteWindow.SingleCurrentUser.SetText("Loading...")
	u, err := inj.GetDBUser()
	if err == nil {
		qtui.SingleSiteWindow.SingleCurrentUser.SetText(u)
	} else {
		qtui.SingleSiteWindow.SingleCurrentUser.SetText("Failed")
	}
	qtui.SingleSiteWindow.SingleDatabaseVersion.SetText("Loading...")
	v, err := inj.GetDBVersion()
	if err == nil {
		qtui.SingleSiteWindow.SingleDatabaseVersion.SetText(v)
	} else {
		qtui.SingleSiteWindow.SingleDatabaseVersion.SetText("Failed")
	}
}

func (a *AnalyzerModule) GetStructure() error {
	utils.Done = make(chan interface{})
	utils.Kill = make(chan interface{})
	manager.PManager.ResetCtx()
	manager.PManager.Client.Transport = manager.PManager.CreateProxyTransport()
	if len(a.Injectable.Structure) > 0 {
		// qtui.NewYesNo(qtui.SingleSiteWindow, "You've already gathered the structure of this site's backend, are you sure you want to clear it and re-gather?", "Just a moment")
		go qtui.SimpleMB(qtui.SingleSiteWindow, "You've already gathered the structure of this backend.", "Hold up").Show()
		return nil
	}

	qtui.SingleSiteWindow.SingleStructureProgress.SetMinimum(0)
	qtui.SingleSiteWindow.SingleStructureProgress.SetMaximum(0)
	qtui.SingleSiteWindow.SingleStructureProgress.SetValue(0)
	defer func() {
		qtui.SingleSiteWindow.SingleStructureProgress.SetMaximum(1)
		qtui.SingleSiteWindow.SingleStructureProgress.SetValue(1)
	}()

	dbC, err := a.Injectable.GetDatabaseCount()
	if err != nil {
		return err
	}

	done := make(chan interface{})
	defer close(done)

	uiLock := sync.RWMutex{}

	dbSem := make(chan interface{}, 5)
	tbSem := make(chan interface{}, 5)

	modifier := 0

	if viper.GetBool("analyzer.skipinfoschema") {
		modifier = 1
	}
	a.Injectable.Structure = make([]injection.StructureDatabase, dbC-modifier)
	strucWg := sync.WaitGroup{}
	for i:=modifier; i<dbC; i++ {
		database, err := a.Injectable.GetDatabase(i)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		if viper.GetBool("analyzer.skipinfoschema") && strings.Contains(database, "information_schema") {
			continue
		}
		a.Injectable.Structure[i-modifier].Name = database
		uiLock.Lock()
		dbItem := widgets.NewQTreeWidgetItem3(qtui.SingleSiteWindow.SingleDatabaseStructure, 0)
		uiLock.Unlock()
		dbItem.SetText(0, database)
		dbItem.SetText(1, "Loading...")
		dbItem.SetFlags(core.Qt__ItemIsUserCheckable | core.Qt__ItemIsSelectable | core.Qt__ItemIsEnabled)
		dbItem.SetCheckState(2, core.Qt__Unchecked)

		strucWg.Add(1)
		dbSem <- 0
		go func(database string, dbIndex int, dbItem *widgets.QTreeWidgetItem) {
			defer func() {
				strucWg.Done()
				<-dbSem
			}()
			tC, err := a.Injectable.GetTableCount(database)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			a.Injectable.Structure[dbIndex].Tables = make([]injection.StructureTable, tC)
			//uiLock.Lock()
			dbItem.SetText(1, fmt.Sprintf("%d Tables", tC))
			//uiLock.Unlock()
			twg := sync.WaitGroup{}
			for j :=0; j < tC; j++ {
				table, err := a.Injectable.GetTable(database, j)
				if err != nil {
					fmt.Println(err.Error())
					continue
				}
				uiLock.Lock()
				tbItem := widgets.NewQTreeWidgetItem6(dbItem, 0)
				uiLock.Unlock()
				tbItem.SetText(0, table)
				tbItem.SetText(1, "Loading...")
				tbItem.SetFlags(core.Qt__ItemIsUserCheckable | core.Qt__ItemIsSelectable | core.Qt__ItemIsEnabled)
				tbItem.SetCheckState(2, core.Qt__Unchecked)
				a.Injectable.Structure[dbIndex].Tables[j].Name = table
				rC, err := a.Injectable.GetRowCount(database, table)
				a.Injectable.Structure[dbIndex].Tables[j].RowCount = rC
				//uiLock.Lock()
				tbItem.SetText(1, fmt.Sprintf("%d Rows", rC))
				//uiLock.Unlock()
				twg.Add(1)
				tbSem <- 0
				go func(database, table string, dbIndex, tbIndex int, tbItem *widgets.QTreeWidgetItem) {
					defer func() {
						twg.Done()
						<- tbSem
					}()

					cC, err := a.Injectable.GetColumnCount(database, table)
					if err != nil {
						fmt.Println(err.Error())
						return
					}
					a.Injectable.Structure[dbIndex].Tables[tbIndex].Columns = make([]injection.StructureColumn, cC)
					for k:=0; k<cC; k++ {
						column, err := a.Injectable.GetColumn(database, table, k)
						if err != nil {
							fmt.Println(err.Error())
							continue
						}
						uiLock.Lock()
						cItem := widgets.NewQTreeWidgetItem6(tbItem, 0)
						uiLock.Unlock()
						cItem.SetText(0, column)
						cItem.SetText(1, "N/A")
						cItem.SetFlags(core.Qt__ItemIsUserCheckable | core.Qt__ItemIsSelectable | core.Qt__ItemIsEnabled)
						cItem.SetCheckState(2, core.Qt__Unchecked)
						a.Injectable.Structure[dbIndex].Tables[tbIndex].Columns[k].Name = column
						//_type, err := a.Injectable.GetColumnType(database, table, k)
						a.Injectable.Structure[dbIndex].Tables[tbIndex].Columns[k].Type = "N/A"
					}
				}(database, table, dbIndex, j, tbItem)
			}
			twg.Wait()
		}(database, i-modifier, dbItem)
	}

	strucWg.Wait()
	return nil
}

func (a *AnalyzerModule) DumpSelection() {
	rowCh := make(chan Row)
	defer close(rowCh)

	go func() {
		var err error
		outputMap := map[string]*os.File{}
		defer func() {
			for _, f := range outputMap {
				f.Close()
			}
		}()
		sitePath := path.Join(func() string {v,_:=os.Getwd();return v}(), "output", a.Injectable.Base.Hostname())
		os.MkdirAll(sitePath, 0755)
		fmt.Println("Waiting for rows...")
		for row := range rowCh{
			if _, ok := outputMap[fmt.Sprintf("%s-%s", a.Injectable.Base.Hostname(), row.Table)]; !ok {
				outputMap[fmt.Sprintf("%s-%s", a.Injectable.Base.Hostname(), row.Table)], err = os.Create(path.Join(sitePath, fmt.Sprintf("%s.csv", row.Table)))
				if err != nil {
					panic(err)
				}
				outputMap[fmt.Sprintf("%s-%s", a.Injectable.Base.Hostname(), row.Table)].WriteString(row.Header + "\r\n")
			}
			outputMap[fmt.Sprintf("%s-%s", a.Injectable.Base.Hostname(), row.Table)].WriteString(row.Row + "\r\n")
		}
	}()

	uiMap := map[int][]int{}

	dumpCh := make(chan Dump)
	UILock := sync.RWMutex{}

	wwg := sync.WaitGroup{}
	for i:=0; i < viper.GetInt("analyzer.workers"); i++ {
		wwg.Add(1)
		go func() {
			for dump := range dumpCh {
				//a.Injectable.Structure[dump.DBIndex].Tables[dump.TBIndex].Rows[dump.Index] = make([]string, len(dump.Columns))
				/*var out []string
				for _, col := range dump.Columns {
					data, _ := a.Injectable.DumpColumn(dump.Database, dump.Table, col, dump.OrderColumn, dump.Index)
					out = append(out, data)
				}*/
				data, err := a.Injectable.DumpMultiColumn(dump.Database, dump.Table, dump.Columns, dump.OrderColumn, dump.Index, ",")
				if err != nil {
					continue
				}

				a.Injectable.Structure[dump.DBIndex].Tables[dump.TBIndex].Rows = append(a.Injectable.Structure[dump.DBIndex].Tables[dump.TBIndex].Rows, strings.Split(data, ","))

				a.Injectable.Structure[dump.DBIndex].Tables[dump.TBIndex].Index++

				/*var win *qtui.SingleSiteDumpLogWidget
				ok := false
				UILock.Lock()
				if win, ok = qtui.DumpLogMap[dump.UIIndex]; ok {
					win.SetWindowTitle(fmt.Sprintf("%s - %s [%d/%d]", a.Injectable.Structure[dump.DBIndex].Name, a.Injectable.Structure[dump.DBIndex].Tables[dump.TBIndex].Name, a.Injectable.Structure[dump.DBIndex].Tables[dump.TBIndex].Index, a.Injectable.Structure[dump.DBIndex].Tables[dump.TBIndex].RowCount))
					row := win.DumpLogTable.CurrentRow()
					win.DumpLogTable.SetRowCount(row+1)
					for j, k := range out {
						win.DumpLogTable.SetItem(row, j, widgets.NewQTableWidgetItem2(k, 0))
					}
				}
				UILock.Unlock()*/

				rowCh <- Row{
					Site:   a.Injectable.Base.Hostname(),
					Table:  dump.Table,
					Row:    data,
					Header: strings.Join(dump.Columns, ","),
				}
				dump.Dwg.Done()
			}
		}()
	}

	go func() {
		for  {
			select {
			case <-time.After(time.Second):
				for uiIndex, indexes := range uiMap {
					if len(indexes) > 0 {
						qtui.SingleSiteWindow.SingleDumpOutput.SetItem(uiIndex, 2, widgets.NewQTableWidgetItem2(fmt.Sprintf("%d/%d", a.Injectable.Structure[indexes[0]].Tables[indexes[1]].Index, a.Injectable.Structure[indexes[0]].Tables[indexes[1]].RowCount), 0))
						qtui.SingleSiteWindow.SingleDumpOutput.SetItem(uiIndex, 3, widgets.NewQTableWidgetItem2(fmt.Sprintf("%d", a.Injectable.Errors), 0))
						if a.Injectable.Structure[indexes[0]].Tables[indexes[1]].Index == a.Injectable.Structure[indexes[0]].Tables[indexes[1]].RowCount {
							delete(uiMap, uiIndex)
						}
					}
				}
			}
		}
	}()

	dwg := sync.WaitGroup{}
	for dbIndex, dbItem := range a.Injectable.Structure {
		if dbItem.Selected {
			for tbIndex, tbItem := range dbItem.Tables {
				if tbItem.Selected {
					UILock.Lock()
					uiIndex := qtui.SingleSiteWindow.SingleDumpOutput.RowCount()
					qtui.SingleSiteWindow.SingleDumpOutput.SetRowCount(uiIndex+1)
					uiMap[uiIndex] = []int{dbIndex, tbIndex}
					win := qtui.NewSingleSiteDumpLogWidget(nil)

					//win.UILock = sync.RWMutex{}
					win.SetWindowTitle(fmt.Sprintf("%s - %s [%d/%d]", a.Injectable.Structure[dbIndex].Name, a.Injectable.Structure[dbIndex].Tables[tbIndex].Name, tbIndex, a.Injectable.Structure[dbIndex].Tables[tbIndex].RowCount))

					win.DumpLogTable.SetColumnCount(len(tbItem.Columns))
					for cIndex, cItem := range tbItem.Columns {
						it := widgets.NewQTableWidgetItem2(cItem.Name, 0)
						win.DumpLogTable.SetHorizontalHeaderItem(cIndex, it)
					}

					qtui.DumpLogMap[uiIndex] = win
					UILock.Unlock()
					fmt.Println("Setting up table on index", uiIndex, dbItem.Name, tbItem.Name)
					qtui.SingleSiteWindow.SingleDumpOutput.SetItem(uiIndex, 0, widgets.NewQTableWidgetItem2(dbItem.Name, 0))
					qtui.SingleSiteWindow.SingleDumpOutput.SetItem(uiIndex, 1, widgets.NewQTableWidgetItem2(tbItem.Name, 0))
					qtui.SingleSiteWindow.SingleDumpOutput.SetItem(uiIndex, 2, widgets.NewQTableWidgetItem2(fmt.Sprintf("%d/%d", 0, tbItem.RowCount), 0))

					selected := []string{}
					for _, cItem := range tbItem.Columns {
						if cItem.Selected {
							selected = append(selected, cItem.Name)
						}
					}
					//a.Injectable.Structure[dbIndex].Tables[tbIndex].Rows = make([]injection.StructureRow, a.Injectable.Structure[dbIndex].Tables[tbIndex].RowCount)

					for rIndex := 0; rIndex < tbItem.RowCount; rIndex ++ {
						dwg.Add(1)
						dumpCh <- Dump{
							Database:    dbItem.Name,
							DBIndex:     dbIndex,
							Table:       tbItem.Name,
							TBIndex:     tbIndex,
							OrderColumn: tbItem.Columns[0].Name,
							Columns:     selected,
							Index:       rIndex,
							UIIndex:     uiIndex,
							Dwg:         &dwg,
						}
					}
				}
			}
		}
	}
	dwg.Wait()
	close(dumpCh)

	f := make(chan interface{})
	go func() {
		wwg.Wait()
		close(f)
	}()
	select {
	case <- f:
		return
	case <-utils.Done:
		return
	}
}