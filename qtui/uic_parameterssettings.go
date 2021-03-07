package qtui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type __parameterssettings struct{}

func (*__parameterssettings) init() {}

type ParametersSettings struct {
	*__parameterssettings
	*widgets.QDialog
	VerticalLayout_4   *widgets.QVBoxLayout
	Widget             *widgets.QWidget
	HorizontalLayout   *widgets.QHBoxLayout
	ParametersgroupBox *widgets.QGroupBox
	VerticalLayout_2   *widgets.QVBoxLayout
	ParametersList     *widgets.QListWidget
	Widget_2           *widgets.QWidget
	VerticalLayout_3   *widgets.QVBoxLayout
	VerticalSpacer     *widgets.QSpacerItem
	AddParams          *widgets.QPushButton
	RemoveParams       *widgets.QPushButton
	VerticalSpacer_2   *widgets.QSpacerItem
	ButtonBox          *widgets.QDialogButtonBox
}

func NewParametersSettings(p widgets.QWidget_ITF) *ParametersSettings {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &ParametersSettings{QDialog: widgets.NewQDialog(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *ParametersSettings) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("ParametersSettings")
	}
	w.Resize2(561, 351)
	w.SetMinimumSize(core.NewQSize2(561, 351))
	w.VerticalLayout_4 = widgets.NewQVBoxLayout2(w)
	w.VerticalLayout_4.SetObjectName("verticalLayout_4")
	w.Widget = widgets.NewQWidget(w, 0)
	w.Widget.SetObjectName("widget")
	w.HorizontalLayout = widgets.NewQHBoxLayout2(w.Widget)
	w.HorizontalLayout.SetObjectName("horizontalLayout")
	w.ParametersgroupBox = widgets.NewQGroupBox(w.Widget)
	w.ParametersgroupBox.SetObjectName("ParametersgroupBox")
	w.VerticalLayout_2 = widgets.NewQVBoxLayout2(w.ParametersgroupBox)
	w.VerticalLayout_2.SetObjectName("verticalLayout_2")
	w.ParametersList = widgets.NewQListWidget(w.ParametersgroupBox)
	w.ParametersList.SetObjectName("ParametersList")
	w.VerticalLayout_2.QLayout.AddWidget(w.ParametersList)
	w.HorizontalLayout.QLayout.AddWidget(w.ParametersgroupBox)
	w.Widget_2 = widgets.NewQWidget(w.Widget, 0)
	w.Widget_2.SetObjectName("widget_2")
	w.VerticalLayout_3 = widgets.NewQVBoxLayout2(w.Widget_2)
	w.VerticalLayout_3.SetObjectName("verticalLayout_3")
	w.VerticalSpacer = widgets.NewQSpacerItem(20, 96, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout_3.AddItem(w.VerticalSpacer)
	w.AddParams = widgets.NewQPushButton(w.Widget_2)
	w.AddParams.SetObjectName("AddParams")
	w.VerticalLayout_3.QLayout.AddWidget(w.AddParams)
	w.RemoveParams = widgets.NewQPushButton(w.Widget_2)
	w.RemoveParams.SetObjectName("RemoveParams")
	w.VerticalLayout_3.QLayout.AddWidget(w.RemoveParams)
	w.VerticalSpacer_2 = widgets.NewQSpacerItem(20, 95, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout_3.AddItem(w.VerticalSpacer_2)
	w.HorizontalLayout.QLayout.AddWidget(w.Widget_2)
	w.VerticalLayout_4.QLayout.AddWidget(w.Widget)
	w.ButtonBox = widgets.NewQDialogButtonBox(w)
	w.ButtonBox.SetObjectName("ButtonBox")
	w.ButtonBox.SetLayoutDirection(core.Qt__LeftToRight)
	w.ButtonBox.SetOrientation(core.Qt__Horizontal)
	w.ButtonBox.SetStandardButtons(widgets.QDialogButtonBox__Cancel | widgets.QDialogButtonBox__Ok)
	w.ButtonBox.SetCenterButtons(true)
	w.VerticalLayout_4.QLayout.AddWidget(w.ButtonBox)
	w.retranslateUi()
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *ParametersSettings) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("ParametersSettings", "Parameters Settings", "", 0))
	w.ParametersgroupBox.SetTitle(core.QCoreApplication_Translate("ParametersSettings", "Parameters", "", 0))
	w.AddParams.SetText(core.QCoreApplication_Translate("ParametersSettings", "Add", "", 0))
	w.RemoveParams.SetText(core.QCoreApplication_Translate("ParametersSettings", "Remove", "", 0))

}
