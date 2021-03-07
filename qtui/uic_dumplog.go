package qtui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type __singlesitedumplogwidget struct{}

func (*__singlesitedumplogwidget) init() {}

type SingleSiteDumpLogWidget struct {
	*__singlesitedumplogwidget
	*widgets.QWidget
	VerticalLayout *widgets.QVBoxLayout
	DumpLogTable   *widgets.QTableWidget
}

func NewSingleSiteDumpLogWidget(p widgets.QWidget_ITF) *SingleSiteDumpLogWidget {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &SingleSiteDumpLogWidget{QWidget: widgets.NewQWidget(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *SingleSiteDumpLogWidget) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("SingleSiteDumpLogWidget")
	}
	w.Resize2(392, 403)
	w.VerticalLayout = widgets.NewQVBoxLayout2(w)
	w.VerticalLayout.SetObjectName("verticalLayout")
	w.DumpLogTable = widgets.NewQTableWidget(w)
	w.DumpLogTable.SetObjectName("dumpLogTable")
	w.VerticalLayout.QLayout.AddWidget(w.DumpLogTable)
	w.retranslateUi()
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *SingleSiteDumpLogWidget) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("SingleSiteDumpLogWidget", "Dump Log", "", 0))

}
