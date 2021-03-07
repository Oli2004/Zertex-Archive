package qtui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type __settings struct{}

func (*__settings) init() {}

type Settings struct {
	*__settings
	*widgets.QDialog
	VerticalLayout      *widgets.QVBoxLayout
	GroupBox_4          *widgets.QGroupBox
	HorizontalLayout    *widgets.QHBoxLayout
	Widget_2            *widgets.QWidget
	VerticalLayout_4    *widgets.QVBoxLayout
	ThreadsSpinbox      *widgets.QSpinBox
	TimeoutsSpinbox     *widgets.QSpinBox
	Widget_3            *widgets.QWidget
	VerticalLayout_3    *widgets.QVBoxLayout
	AutoThreadsCheckBox *widgets.QCheckBox
	BatchModeCheckBox   *widgets.QCheckBox
	GroupBox            *widgets.QGroupBox
	HorizontalLayout_2  *widgets.QHBoxLayout
	ProxiesTextbox      *widgets.QPlainTextEdit
	Widget              *widgets.QWidget
	VerticalLayout_2    *widgets.QVBoxLayout
	ProxyTypeComboBox   *widgets.QComboBox
	LoadProxiesButton   *widgets.QPushButton
	LoadProxiesFromAPI  *widgets.QCheckBox
	LineEdit            *widgets.QLineEdit
	LoadCustomerProxies *widgets.QPushButton
	VerticalSpacer      *widgets.QSpacerItem
	ButtonBox           *widgets.QDialogButtonBox
}

func NewSettings(p widgets.QWidget_ITF) *Settings {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &Settings{QDialog: widgets.NewQDialog(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *Settings) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("Settings")
	}
	w.Resize2(517, 455)
	w.VerticalLayout = widgets.NewQVBoxLayout2(w)
	w.VerticalLayout.SetObjectName("verticalLayout")
	w.GroupBox_4 = widgets.NewQGroupBox(w)
	w.GroupBox_4.SetObjectName("groupBox_4")
	w.GroupBox_4.SetAlignment(int(core.Qt__AlignCenter))
	w.HorizontalLayout = widgets.NewQHBoxLayout2(w.GroupBox_4)
	w.HorizontalLayout.SetObjectName("horizontalLayout")
	w.Widget_2 = widgets.NewQWidget(w.GroupBox_4, 0)
	w.Widget_2.SetObjectName("widget_2")
	w.VerticalLayout_4 = widgets.NewQVBoxLayout2(w.Widget_2)
	w.VerticalLayout_4.SetObjectName("verticalLayout_4")
	w.ThreadsSpinbox = widgets.NewQSpinBox(w.Widget_2)
	w.ThreadsSpinbox.SetObjectName("threadsSpinbox")
	w.ThreadsSpinbox.SetMouseTracking(true)
	w.ThreadsSpinbox.SetMinimum(1)
	w.ThreadsSpinbox.SetMaximum(500)
	w.VerticalLayout_4.QLayout.AddWidget(w.ThreadsSpinbox)
	w.TimeoutsSpinbox = widgets.NewQSpinBox(w.Widget_2)
	w.TimeoutsSpinbox.SetObjectName("timeoutsSpinbox")
	w.TimeoutsSpinbox.SetMouseTracking(true)
	w.TimeoutsSpinbox.SetMinimum(5)
	w.TimeoutsSpinbox.SetMaximum(20)
	w.TimeoutsSpinbox.SetValue(10)
	w.VerticalLayout_4.QLayout.AddWidget(w.TimeoutsSpinbox)
	w.HorizontalLayout.QLayout.AddWidget(w.Widget_2)
	w.Widget_3 = widgets.NewQWidget(w.GroupBox_4, 0)
	w.Widget_3.SetObjectName("widget_3")
	w.VerticalLayout_3 = widgets.NewQVBoxLayout2(w.Widget_3)
	w.VerticalLayout_3.SetObjectName("verticalLayout_3")
	w.AutoThreadsCheckBox = widgets.NewQCheckBox(w.Widget_3)
	w.AutoThreadsCheckBox.SetObjectName("autoThreadsCheckBox")
	w.VerticalLayout_3.QLayout.AddWidget(w.AutoThreadsCheckBox)
	w.BatchModeCheckBox = widgets.NewQCheckBox(w.Widget_3)
	w.BatchModeCheckBox.SetObjectName("batchModeCheckBox")
	w.VerticalLayout_3.QLayout.AddWidget(w.BatchModeCheckBox)
	w.HorizontalLayout.QLayout.AddWidget(w.Widget_3)
	w.VerticalLayout.QLayout.AddWidget(w.GroupBox_4)
	w.GroupBox = widgets.NewQGroupBox(w)
	w.GroupBox.SetObjectName("groupBox")
	w.GroupBox.SetAlignment(int(core.Qt__AlignCenter))
	w.HorizontalLayout_2 = widgets.NewQHBoxLayout2(w.GroupBox)
	w.HorizontalLayout_2.SetObjectName("horizontalLayout_2")
	w.ProxiesTextbox = widgets.NewQPlainTextEdit(w.GroupBox)
	w.ProxiesTextbox.SetObjectName("proxiesTextbox")
	w.HorizontalLayout_2.QLayout.AddWidget(w.ProxiesTextbox)
	w.Widget = widgets.NewQWidget(w.GroupBox, 0)
	w.Widget.SetObjectName("widget")
	w.VerticalLayout_2 = widgets.NewQVBoxLayout2(w.Widget)
	w.VerticalLayout_2.SetObjectName("verticalLayout_2")
	w.ProxyTypeComboBox = widgets.NewQComboBox(w.Widget)
	w.ProxyTypeComboBox.AddItem("", core.NewQVariant1(0))
	w.ProxyTypeComboBox.AddItem("", core.NewQVariant1(0))
	w.ProxyTypeComboBox.AddItem("", core.NewQVariant1(0))
	w.ProxyTypeComboBox.AddItem("", core.NewQVariant1(0))
	w.ProxyTypeComboBox.SetObjectName("proxyTypeComboBox")
	w.VerticalLayout_2.QLayout.AddWidget(w.ProxyTypeComboBox)
	w.LoadProxiesButton = widgets.NewQPushButton(w.Widget)
	w.LoadProxiesButton.SetObjectName("loadProxiesButton")
	w.LoadProxiesButton.SetEnabled(false)
	w.VerticalLayout_2.QLayout.AddWidget(w.LoadProxiesButton)
	w.LoadProxiesFromAPI = widgets.NewQCheckBox(w.Widget)
	w.LoadProxiesFromAPI.SetObjectName("loadProxiesFromAPI")
	w.VerticalLayout_2.QLayout.AddWidget(w.LoadProxiesFromAPI)
	w.LineEdit = widgets.NewQLineEdit(w.Widget)
	w.LineEdit.SetObjectName("lineEdit")
	w.VerticalLayout_2.QLayout.AddWidget(w.LineEdit)
	w.LoadCustomerProxies = widgets.NewQPushButton(w.Widget)
	w.LoadCustomerProxies.SetObjectName("loadCustomerProxies")
	w.VerticalLayout_2.QLayout.AddWidget(w.LoadCustomerProxies)
	w.VerticalSpacer = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout_2.AddItem(w.VerticalSpacer)
	w.HorizontalLayout_2.QLayout.AddWidget(w.Widget)
	w.VerticalLayout.QLayout.AddWidget(w.GroupBox)
	w.ButtonBox = widgets.NewQDialogButtonBox(w)
	w.ButtonBox.SetObjectName("buttonBox")
	w.ButtonBox.SetOrientation(core.Qt__Horizontal)
	w.ButtonBox.SetStandardButtons(widgets.QDialogButtonBox__Cancel | widgets.QDialogButtonBox__Save)
	w.VerticalLayout.QLayout.AddWidget(w.ButtonBox)
	w.retranslateUi()
	w.ButtonBox.ConnectAccepted(w.Accept)
	w.ButtonBox.ConnectRejected(w.Reject)
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *Settings) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("Settings", "Settings", "", 0))
	w.GroupBox_4.SetTitle(core.QCoreApplication_Translate("Settings", "Core", "", 0))
	if true {
		w.ThreadsSpinbox.SetToolTip(core.QCoreApplication_Translate("Settings", "Set your Threads", "", 0))
	}
	w.ThreadsSpinbox.SetSuffix("")
	w.ThreadsSpinbox.SetPrefix(core.QCoreApplication_Translate("Settings", "Threads: ", "", 0))
	if true {
		w.TimeoutsSpinbox.SetToolTip(core.QCoreApplication_Translate("Settings", "Set your Threads", "", 0))
	}
	w.TimeoutsSpinbox.SetSuffix("")
	w.TimeoutsSpinbox.SetPrefix(core.QCoreApplication_Translate("Settings", "Timeouts: ", "", 0))
	w.AutoThreadsCheckBox.SetText(core.QCoreApplication_Translate("Settings", "Auto Threads", "", 0))
	w.BatchModeCheckBox.SetText(core.QCoreApplication_Translate("Settings", "Batch Mode", "", 0))
	w.GroupBox.SetTitle(core.QCoreApplication_Translate("Settings", "Proxies", "", 0))
	w.ProxyTypeComboBox.SetItemText(0, core.QCoreApplication_Translate("Settings", "NONE", "", 0))
	w.ProxyTypeComboBox.SetItemText(1, core.QCoreApplication_Translate("Settings", "HTTP", "", 0))
	w.ProxyTypeComboBox.SetItemText(2, core.QCoreApplication_Translate("Settings", "SOCKS4", "", 0))
	w.ProxyTypeComboBox.SetItemText(3, core.QCoreApplication_Translate("Settings", "SOCKS5", "", 0))
	w.ProxyTypeComboBox.SetItemText(4, core.QCoreApplication_Translate("Settings", "Proxy Type", "", 0))
	w.LoadProxiesButton.SetText(core.QCoreApplication_Translate("Settings", "Load", "", 0))
	w.LoadProxiesFromAPI.SetText(core.QCoreApplication_Translate("Settings", "Load from API", "", 0))
	w.LineEdit.SetText(core.QCoreApplication_Translate("Settings", "http://proxysource.com/list.txt", "", 0))
	w.LoadCustomerProxies.SetText(core.QCoreApplication_Translate("Settings", "Customer Pool", "", 0))

}
