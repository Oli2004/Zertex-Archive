package utils

import (
	"github.com/paulbellamy/ratecounter"
	"os"
	"time"
)

var CancelChan chan os.Signal
var RequestCounter int
var ErrorCounter int
var RateCounter = ratecounter.NewRateCounter(1 * time.Second)
var RPMCounter = ratecounter.NewRateCounter(time.Minute)
var StartTime time.Time
var GlobalSem chan interface{}
var WorkerSem chan interface{}
var GlobalIndex int
var Kill chan interface{}
var Done chan interface{}
var Module string = "Idle"

type DatabaseSystem interface {
	BuildInjection(vector string)
}