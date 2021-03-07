package qtui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type __dialog struct{}

func (*__dialog) init() {}

type Dialog struct {
	*__dialog
	*widgets.QDialog
	VerticalLayout   *widgets.QVBoxLayout
	VerticalSpacer   *widgets.QSpacerItem
	Label            *widgets.QLabel
	VerticalSpacer_2 *widgets.QSpacerItem
	ButtonBox        *widgets.QDialogButtonBox
}

func NewYesNoWidget(p widgets.QWidget_ITF) *Dialog {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &Dialog{QDialog: widgets.NewQDialog(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *Dialog) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("Dialog")
	}
	w.Resize2(420, 189)
	w.SetMinimumSize(core.NewQSize2(420, 189))
	w.SetMaximumSize(core.NewQSize2(420, 189))
	w.VerticalLayout = widgets.NewQVBoxLayout2(w)
	w.VerticalLayout.SetObjectName("verticalLayout")
	w.VerticalSpacer = widgets.NewQSpacerItem(20, 49, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout.AddItem(w.VerticalSpacer)
	w.Label = widgets.NewQLabel(w, 0)
	w.Label.SetObjectName("label")
	w.Label.SetAlignment(core.Qt__AlignCenter)
	w.VerticalLayout.QLayout.AddWidget(w.Label)
	w.VerticalSpacer_2 = widgets.NewQSpacerItem(20, 48, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout.AddItem(w.VerticalSpacer_2)
	w.ButtonBox = widgets.NewQDialogButtonBox(w)
	w.ButtonBox.SetObjectName("buttonBox")
	w.ButtonBox.SetOrientation(core.Qt__Horizontal)
	w.ButtonBox.SetStandardButtons(widgets.QDialogButtonBox__Cancel | widgets.QDialogButtonBox__Ok)
	w.VerticalLayout.QLayout.AddWidget(w.ButtonBox)
	w.retranslateUi()
	w.ButtonBox.ConnectAccepted(w.Accept)
	w.ButtonBox.ConnectRejected(w.Reject)
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *Dialog) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("Dialog", "Dialog", "", 0))
	w.Label.SetText(core.QCoreApplication_Translate("Dialog", "TextLabel", "", 0))

}
