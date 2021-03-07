package qtui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type __whitelistadd struct{}

func (*__whitelistadd) init() {}

type WhitelistAdd struct {
	*__whitelistadd
	*widgets.QDialog
	VerticalLayout   *widgets.QVBoxLayout
	Label            *widgets.QLabel
	GroupName        *widgets.QLineEdit
	Widget           *widgets.QWidget
	HorizontalLayout *widgets.QHBoxLayout
	HorizontalSpacer *widgets.QSpacerItem
	CreateBtn        *widgets.QPushButton
	CancelBtn        *widgets.QPushButton
}

func NewWhitelistAdd(p widgets.QWidget_ITF) *WhitelistAdd {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &WhitelistAdd{QDialog: widgets.NewQDialog(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *WhitelistAdd) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("WhitelistAdd")
	}
	w.Resize2(400, 114)
	w.VerticalLayout = widgets.NewQVBoxLayout2(w)
	w.VerticalLayout.SetObjectName("verticalLayout")
	w.Label = widgets.NewQLabel(w, 0)
	w.Label.SetObjectName("label")
	w.Label.SetAlignment(core.Qt__AlignCenter)
	w.VerticalLayout.QLayout.AddWidget(w.Label)
	w.GroupName = widgets.NewQLineEdit(w)
	w.GroupName.SetObjectName("groupName")
	w.VerticalLayout.QLayout.AddWidget(w.GroupName)
	w.Widget = widgets.NewQWidget(w, 0)
	w.Widget.SetObjectName("widget")
	w.HorizontalLayout = widgets.NewQHBoxLayout2(w.Widget)
	w.HorizontalLayout.SetObjectName("horizontalLayout")
	w.HorizontalSpacer = widgets.NewQSpacerItem(201, 20, widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Minimum)
	w.HorizontalLayout.AddItem(w.HorizontalSpacer)
	w.CreateBtn = widgets.NewQPushButton(w.Widget)
	w.CreateBtn.SetObjectName("createBtn")
	w.HorizontalLayout.QLayout.AddWidget(w.CreateBtn)
	w.CancelBtn = widgets.NewQPushButton(w.Widget)
	w.CancelBtn.SetObjectName("cancelBtn")
	w.HorizontalLayout.QLayout.AddWidget(w.CancelBtn)
	w.VerticalLayout.QLayout.AddWidget(w.Widget)
	w.retranslateUi()
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *WhitelistAdd) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("WhitelistAdd", "Add Whitelist Group", "", 0))
	w.Label.SetText(core.QCoreApplication_Translate("WhitelistAdd", "Name", "", 0))
	w.GroupName.SetPlaceholderText(core.QCoreApplication_Translate("WhitelistAdd", "Tertiary", "", 0))
	w.CreateBtn.SetText(core.QCoreApplication_Translate("WhitelistAdd", "Create", "", 0))
	w.CancelBtn.SetText(core.QCoreApplication_Translate("WhitelistAdd", "Cancel", "", 0))

}
