package modules

import (
	"bufio"
	"errors"
	"fmt"
	"git.quartzinc.dev/Zertex/XDGv2/qtui"
	"git.quartzinc.dev/Zertex/XDGv2/utils"
	"github.com/spf13/viper"
	"github.com/therecipe/qt/widgets"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
	"sync"
	"time"
)

type GeneratorModule struct {
	Index int
	Parameters map[string]*Parameter
	Patterns []*Pattern
	WordChan chan string
	UIMap map[int]int
	UILock sync.RWMutex
}

type Pattern struct {
	PatternStr  string
	Prefixes    []string
	TotalDorks int
	DorksGenerated int
}

type Parameter struct {
	Name string //`json:"Name"`
	Prefix string 	//`json:"Prefix"`
	FilePath string //`json:"FilePath"`
	Data []string
}

var	Generator *GeneratorModule

func NewGenerator() *GeneratorModule {
	Generator = &GeneratorModule{
		Parameters: map[string]*Parameter{},
	}
	var params []*Parameter
	if err := viper.UnmarshalKey("generator.Parameters", &params); err != nil {
		log.Fatal(err.Error())
	}

	for _, param := range params {
		Generator.Parameters[param.Prefix] = param
	}


	for _, pattern := range viper.GetStringSlice("generator.Patterns") {
		re, _ := regexp.Compile(`\(([^()]+)\)`)
		if strings.TrimSpace(pattern) != "" {
			if m := re.FindAllString(pattern, -1); m != nil{
				Generator.Patterns = append(Generator.Patterns, &Pattern{
					Prefixes:    m,
					TotalDorks: 1,
					PatternStr:  pattern,
				})
			}
		}
	}
	return Generator
}

func (p *Parameter) GetData() []string {
	if len(p.Data) > 0 {
		return p.Data
	}

	os.Mkdir("params", os.ModeDir)

	cur, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	if f, err := os.Open(path.Join(cur, "params", p.FilePath)); err == nil {
		scanner  := bufio.NewScanner(f)
		var data []string
		for scanner.Scan() {
			data = append(data, scanner.Text())
		}
		p.Data = data
	}

	return p.Data
}

func (gm *GeneratorModule) Combinations() error{
	for _, pattern := range gm.Patterns {
		pattern.TotalDorks = 1
		pattern.DorksGenerated = 0
		for _, prefix := range pattern.Prefixes{
			if _, exists := gm.Parameters[prefix]; exists{
				pattern.TotalDorks *= len(gm.Parameters[prefix].GetData())
			}else{
				return errors.New(fmt.Sprintf("Parameter %s on Pattern %s has not been created", prefix, pattern.PatternStr))
			}
		}
	}
	return nil
}

type Prefix struct {
	Key string
	Data string
}


func (gm *GeneratorModule) Start() error {
	err := gm.Combinations()
	if err != nil{
		return err
	}

	f, _ := os.OpenFile("dorks.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0755)
	defer f.Close()

	gm.UIMap = make(map[int]int, len(gm.Patterns))

	gm.WordChan = make(chan string)
	done := make(chan interface{})
	//defer close(gm.WordChan)

	//Save each dork
	go func() {
		for dork := range gm.WordChan {
			select {
			case <- done:
				return
			default:
			}
			f.WriteString(dork+"\r\n")
		}
	}()

	go func() {
		for {
			select {
			case <-done:
				return
			case <- time.After(time.Second):
				for index, pattern := range gm.Patterns {
					qtui.Main.TableGenerator.SetItem(index, 0, widgets.NewQTableWidgetItem2(pattern.PatternStr, 0))
					qtui.Main.TableGenerator.SetItem(index, 1, widgets.NewQTableWidgetItem2(fmt.Sprintf("%d/%d", pattern.DorksGenerated, pattern.TotalDorks), 0))
				}
			}
		}
	}()

	qtui.Main.TableGenerator.SetRowCount(len(gm.Patterns))

	genWG := sync.WaitGroup{}
	genWG.Add(len(gm.Patterns))

	for gm.Index, _ = range gm.Patterns {
		go func(i int, pattern *Pattern) {
			defer genWG.Done()

			var prefixSem []chan *Prefix
			pwg := sync.WaitGroup{}
			doneChan := make(chan interface{})

			for _, prefix := range pattern.Prefixes {
				wordChan :=  make(chan *Prefix, 1)
				prefixSem = append(prefixSem, wordChan)
				pwg.Add(1)

				go func(prefix string, pattern *Pattern) {
					defer pwg.Done()
					if len(gm.Parameters[prefix].Data) > 0 {
						for i := 0; i < pattern.TotalDorks/len(gm.Parameters[prefix].Data); i++ {
							for _, word := range gm.Parameters[prefix].Data {
								select {
								case <-utils.Done:
									return
								case <-doneChan:
									return
								default:
								}
								wordChan <- &Prefix{
									Key:  prefix,
									Data: word,
								}
							}
						}
					}
					time.Sleep(100*time.Millisecond)
				}(prefix, pattern)
			}


			go func() {
				for {
					select {
					case <- doneChan:
						return
					case <-utils.Done:
						return
					default:
					}

					dork := pattern.PatternStr
					for _, word := range prefixSem {
						if a := <- word; a != nil {
							dork = strings.ReplaceAll(dork, a.Key, a.Data)
						}
					}

					gm.WordChan <- dork
					pattern.DorksGenerated++
					if qtui.Main.GeneratorLimiterCheckbox.IsChecked() && pattern.DorksGenerated == qtui.Main.GeneratorLimiterSpinbox.Value(){
						close(utils.Done)
						for _, a := range prefixSem {
							<- a
						}
						return
					}
				}
			}()

			pwg.Wait()
			close(doneChan)
			for _, a := range prefixSem {
				close (a)
			}
		}(gm.Index, gm.Patterns[gm.Index])
	}

	genWG.Wait()
	close(done)
	return nil
}
