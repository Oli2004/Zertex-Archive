package main

import (
	"XDGv2/qtui"
	"github.com/therecipe/qt/widgets"
	"os"
)

func main() {
	app := widgets.NewQApplication(len(os.Args), os.Args)
	windows := qtui.NewMainWindow(nil)
	windows.Show()
	app.Exec()
}
