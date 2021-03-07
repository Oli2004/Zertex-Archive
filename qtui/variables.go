package qtui

import (
	"github.com/therecipe/qt/widgets"
)

var (
	Application              *widgets.QApplication
	Main                     *MainWindow
	SettingsWindow           *Settings
	DumperWhitelistWindow    *DumperWhitelist
	ParametersSettingsWindow *ParametersSettings
	AuthWindow               *AuthForm
	SingleSiteWindow         *SingleSiteAnalyzer
	DumpLogMap               map[int]*SingleSiteDumpLogWidget
)

func init() {
	DumpLogMap = make(map[int]*SingleSiteDumpLogWidget)
}
