package qtui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type __dumperwhitelist struct{}

func (*__dumperwhitelist) init() {}

type DumperWhitelist struct {
	*__dumperwhitelist
	*widgets.QDialog
	VerticalLayout       *widgets.QVBoxLayout
	Widget               *widgets.QWidget
	HorizontalLayout     *widgets.QHBoxLayout
	GroupBox             *widgets.QGroupBox
	VerticalLayout_2     *widgets.QVBoxLayout
	WhitelistGroups      *widgets.QListWidget
	Widget_3             *widgets.QWidget
	HorizontalLayout_3   *widgets.QHBoxLayout
	WhitelistAddGroup    *widgets.QPushButton
	WhitelistRemoveGroup *widgets.QPushButton
	GroupBox_2           *widgets.QGroupBox
	VerticalLayout_3     *widgets.QVBoxLayout
	WhitelistData        *widgets.QPlainTextEdit
	Widget_4             *widgets.QWidget
	HorizontalLayout_4   *widgets.QHBoxLayout
	WhitelistOk          *widgets.QPushButton
	WhitelistCancel      *widgets.QPushButton
	Label                *widgets.QLabel
	Blacklist            *widgets.QPlainTextEdit
}

func NewDumperWhitelist(p widgets.QWidget_ITF) *DumperWhitelist {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &DumperWhitelist{QDialog: widgets.NewQDialog(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *DumperWhitelist) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("DumperWhitelist")
	}
	w.Resize2(449, 425)
	w.SetMinimumSize(core.NewQSize2(434, 258))
	w.VerticalLayout = widgets.NewQVBoxLayout2(w)
	w.VerticalLayout.SetObjectName("verticalLayout")
	w.Widget = widgets.NewQWidget(w, 0)
	w.Widget.SetObjectName("widget")
	w.HorizontalLayout = widgets.NewQHBoxLayout2(w.Widget)
	w.HorizontalLayout.SetObjectName("horizontalLayout")
	w.GroupBox = widgets.NewQGroupBox(w.Widget)
	w.GroupBox.SetObjectName("groupBox")
	w.GroupBox.SetAlignment(int(core.Qt__AlignCenter))
	w.VerticalLayout_2 = widgets.NewQVBoxLayout2(w.GroupBox)
	w.VerticalLayout_2.SetObjectName("verticalLayout_2")
	w.WhitelistGroups = widgets.NewQListWidget(w.GroupBox)
	w.WhitelistGroups.SetObjectName("whitelistGroups")
	w.VerticalLayout_2.QLayout.AddWidget(w.WhitelistGroups)
	w.Widget_3 = widgets.NewQWidget(w.GroupBox, 0)
	w.Widget_3.SetObjectName("widget_3")
	w.HorizontalLayout_3 = widgets.NewQHBoxLayout2(w.Widget_3)
	w.HorizontalLayout_3.SetObjectName("horizontalLayout_3")
	w.WhitelistAddGroup = widgets.NewQPushButton(w.Widget_3)
	w.WhitelistAddGroup.SetObjectName("whitelistAddGroup")
	w.HorizontalLayout_3.QLayout.AddWidget(w.WhitelistAddGroup)
	w.WhitelistRemoveGroup = widgets.NewQPushButton(w.Widget_3)
	w.WhitelistRemoveGroup.SetObjectName("whitelistRemoveGroup")
	w.HorizontalLayout_3.QLayout.AddWidget(w.WhitelistRemoveGroup)
	w.VerticalLayout_2.QLayout.AddWidget(w.Widget_3)
	w.HorizontalLayout.QLayout.AddWidget(w.GroupBox)
	w.GroupBox_2 = widgets.NewQGroupBox(w.Widget)
	w.GroupBox_2.SetObjectName("groupBox_2")
	w.GroupBox_2.SetAlignment(int(core.Qt__AlignCenter))
	w.VerticalLayout_3 = widgets.NewQVBoxLayout2(w.GroupBox_2)
	w.VerticalLayout_3.SetObjectName("verticalLayout_3")
	w.WhitelistData = widgets.NewQPlainTextEdit(w.GroupBox_2)
	w.WhitelistData.SetObjectName("whitelistData")
	w.VerticalLayout_3.QLayout.AddWidget(w.WhitelistData)
	w.Widget_4 = widgets.NewQWidget(w.GroupBox_2, 0)
	w.Widget_4.SetObjectName("widget_4")
	w.HorizontalLayout_4 = widgets.NewQHBoxLayout2(w.Widget_4)
	w.HorizontalLayout_4.SetObjectName("horizontalLayout_4")
	w.WhitelistOk = widgets.NewQPushButton(w.Widget_4)
	w.WhitelistOk.SetObjectName("whitelistOk")
	w.HorizontalLayout_4.QLayout.AddWidget(w.WhitelistOk)
	w.WhitelistCancel = widgets.NewQPushButton(w.Widget_4)
	w.WhitelistCancel.SetObjectName("whitelistCancel")
	w.HorizontalLayout_4.QLayout.AddWidget(w.WhitelistCancel)
	w.VerticalLayout_3.QLayout.AddWidget(w.Widget_4)
	w.HorizontalLayout.QLayout.AddWidget(w.GroupBox_2)
	w.VerticalLayout.QLayout.AddWidget(w.Widget)
	w.Label = widgets.NewQLabel(w, 0)
	w.Label.SetObjectName("label")
	w.VerticalLayout.QLayout.AddWidget(w.Label)
	w.Blacklist = widgets.NewQPlainTextEdit(w)
	w.Blacklist.SetObjectName("blacklist")
	w.VerticalLayout.QLayout.AddWidget(w.Blacklist)
	w.retranslateUi()
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *DumperWhitelist) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("DumperWhitelist", "Dumper Whitelist Settings", "", 0))
	w.GroupBox.SetTitle(core.QCoreApplication_Translate("DumperWhitelist", "Whitelist Groups", "", 0))
	w.WhitelistAddGroup.SetText(core.QCoreApplication_Translate("DumperWhitelist", "Add", "", 0))
	w.WhitelistRemoveGroup.SetText(core.QCoreApplication_Translate("DumperWhitelist", "Remove", "", 0))
	w.GroupBox_2.SetTitle(core.QCoreApplication_Translate("DumperWhitelist", "Whitelist Data", "", 0))
	w.WhitelistData.SetPlainText("")
	w.WhitelistOk.SetText(core.QCoreApplication_Translate("DumperWhitelist", "OK", "", 0))
	w.WhitelistCancel.SetText(core.QCoreApplication_Translate("DumperWhitelist", "Cancel", "", 0))
	w.Label.SetText(core.QCoreApplication_Translate("DumperWhitelist", "Blacklist", "", 0))
	w.Blacklist.SetPlainText(core.QCoreApplication_Translate("DumperWhitelist", "id\n"+"last\n"+"time\n"+"", "", 0))

}
