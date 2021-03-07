package qtui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type __loadingwindow struct{}

func (*__loadingwindow) init() {}

type LoadingWindow struct {
	*__loadingwindow
	*widgets.QDialog
	VerticalLayout_2 *widgets.QVBoxLayout
	Widget           *widgets.QWidget
	VerticalLayout   *widgets.QVBoxLayout
	VerticalSpacer_2 *widgets.QSpacerItem
	Label            *widgets.QLabel
	ProgressBar      *widgets.QProgressBar
	VerticalSpacer   *widgets.QSpacerItem
	ButtonBox        *widgets.QDialogButtonBox
}

func NewLoadingWindow(p widgets.QWidget_ITF) *LoadingWindow {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &LoadingWindow{QDialog: widgets.NewQDialog(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *LoadingWindow) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("LoadingWindow")
	}
	w.Resize2(335, 231)
	w.SetMinimumSize(core.NewQSize2(335, 231))
	icon := gui.NewQIcon()
	icon.AddFile("../../../../../../../../Downloads/xdumpgo-new-icon.png", core.NewQSize(), gui.QIcon__Normal, gui.QIcon__Off)
	w.SetWindowIcon(icon)
	w.VerticalLayout_2 = widgets.NewQVBoxLayout2(w)
	w.VerticalLayout_2.SetObjectName("verticalLayout_2")
	w.Widget = widgets.NewQWidget(w, 0)
	w.Widget.SetObjectName("widget")
	w.VerticalLayout = widgets.NewQVBoxLayout2(w.Widget)
	w.VerticalLayout.SetObjectName("verticalLayout")
	w.VerticalSpacer_2 = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout.AddItem(w.VerticalSpacer_2)
	w.Label = widgets.NewQLabel(w.Widget, 0)
	w.Label.SetObjectName("label")
	w.Label.SetAlignment(core.Qt__AlignCenter)
	w.VerticalLayout.QLayout.AddWidget(w.Label)
	w.ProgressBar = widgets.NewQProgressBar(w.Widget)
	w.ProgressBar.SetObjectName("progressBar")
	w.ProgressBar.SetValue(24)
	w.VerticalLayout.QLayout.AddWidget(w.ProgressBar)
	w.VerticalSpacer = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout.AddItem(w.VerticalSpacer)
	w.VerticalLayout_2.QLayout.AddWidget(w.Widget)
	w.ButtonBox = widgets.NewQDialogButtonBox(w)
	w.ButtonBox.SetObjectName("buttonBox")
	w.ButtonBox.SetOrientation(core.Qt__Horizontal)
	w.ButtonBox.SetStandardButtons(widgets.QDialogButtonBox__Cancel | widgets.QDialogButtonBox__Ok)
	w.VerticalLayout_2.QLayout.AddWidget(w.ButtonBox)
	w.retranslateUi()
	w.ButtonBox.ConnectAccepted(w.Accept)
	w.ButtonBox.ConnectRejected(w.Reject)
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *LoadingWindow) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("LoadingWindow", "Loading", "", 0))
	w.Label.SetText(core.QCoreApplication_Translate("LoadingWindow", "Loading :0/0", "", 0))
	w.ProgressBar.SetFormat(core.QCoreApplication_Translate("LoadingWindow", "%p% %v/%m", "", 0))

}
