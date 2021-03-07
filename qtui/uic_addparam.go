package qtui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type __addparamdialog struct{}

func (*__addparamdialog) init() {}

type AddParamDialog struct {
	*__addparamdialog
	*widgets.QDialog
	VerticalLayout       *widgets.QVBoxLayout
	NameLabel            *widgets.QLabel
	NameInputField       *widgets.QLineEdit
	PrefixLabel          *widgets.QLabel
	PrefixInputField     *widgets.QLineEdit
	Widget               *widgets.QWidget
	HorizontalLayout     *widgets.QHBoxLayout
	SelectButton         *widgets.QPushButton
	SelectFileInputField *widgets.QLineEdit
	VerticalSpacer       *widgets.QSpacerItem
	ButtonBox            *widgets.QDialogButtonBox
}

func NewAddParamDialog(p widgets.QWidget_ITF) *AddParamDialog {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &AddParamDialog{QDialog: widgets.NewQDialog(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *AddParamDialog) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("AddParamDialog")
	}
	w.Resize2(378, 254)
	w.VerticalLayout = widgets.NewQVBoxLayout2(w)
	w.VerticalLayout.SetObjectName("verticalLayout")
	w.NameLabel = widgets.NewQLabel(w, 0)
	w.NameLabel.SetObjectName("NameLabel")
	w.NameLabel.SetAlignment(core.Qt__AlignCenter)
	w.VerticalLayout.QLayout.AddWidget(w.NameLabel)
	w.NameInputField = widgets.NewQLineEdit(w)
	w.NameInputField.SetObjectName("NameInputField")
	w.VerticalLayout.QLayout.AddWidget(w.NameInputField)
	w.PrefixLabel = widgets.NewQLabel(w, 0)
	w.PrefixLabel.SetObjectName("PrefixLabel")
	w.PrefixLabel.SetAlignment(core.Qt__AlignCenter)
	w.VerticalLayout.QLayout.AddWidget(w.PrefixLabel)
	w.PrefixInputField = widgets.NewQLineEdit(w)
	w.PrefixInputField.SetObjectName("PrefixInputField")
	w.VerticalLayout.QLayout.AddWidget(w.PrefixInputField)
	w.Widget = widgets.NewQWidget(w, 0)
	w.Widget.SetObjectName("widget")
	w.HorizontalLayout = widgets.NewQHBoxLayout2(w.Widget)
	w.HorizontalLayout.SetObjectName("horizontalLayout")
	w.HorizontalLayout.SetContentsMargins(0, -1, 0, -1)
	w.SelectButton = widgets.NewQPushButton(w.Widget)
	w.SelectButton.SetObjectName("SelectButton")
	w.HorizontalLayout.QLayout.AddWidget(w.SelectButton)
	w.SelectFileInputField = widgets.NewQLineEdit(w.Widget)
	w.SelectFileInputField.SetObjectName("SelectFileInputField")
	w.HorizontalLayout.QLayout.AddWidget(w.SelectFileInputField)
	w.VerticalLayout.QLayout.AddWidget(w.Widget)
	w.VerticalSpacer = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout.AddItem(w.VerticalSpacer)
	w.ButtonBox = widgets.NewQDialogButtonBox(w)
	w.ButtonBox.SetObjectName("ButtonBox")
	w.ButtonBox.SetOrientation(core.Qt__Horizontal)
	w.ButtonBox.SetStandardButtons(widgets.QDialogButtonBox__Cancel | widgets.QDialogButtonBox__Ok)
	w.VerticalLayout.QLayout.AddWidget(w.ButtonBox)
	w.retranslateUi()
	w.ButtonBox.ConnectAccepted(w.Accept)
	w.ButtonBox.ConnectRejected(w.Reject)
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *AddParamDialog) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("AddParamDialog", "Add Paramter", "", 0))
	w.NameLabel.SetText(core.QCoreApplication_Translate("AddParamDialog", "Name", "", 0))
	w.PrefixLabel.SetText(core.QCoreApplication_Translate("AddParamDialog", "Prefix", "", 0))
	w.SelectButton.SetText(core.QCoreApplication_Translate("AddParamDialog", "Select File", "", 0))

}
