package qtui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type __singlesiteanalyzer struct{}

func (*__singlesiteanalyzer) init() {}

type SingleSiteAnalyzer struct {
	*__singlesiteanalyzer
	*widgets.QWidget
	GridLayout                  *widgets.QGridLayout
	Widget_7                    *widgets.QWidget
	VerticalLayout_3            *widgets.QVBoxLayout
	GroupBox_2                  *widgets.QGroupBox
	VerticalLayout              *widgets.QVBoxLayout
	SingleSiteText              *widgets.QLineEdit
	SingleSiteSelectInjection   *widgets.QPushButton
	Widget_4                    *widgets.QWidget
	HorizontalLayout_4          *widgets.QHBoxLayout
	Label_4                     *widgets.QLabel
	SingleCountry               *widgets.QLabel
	Widget_5                    *widgets.QWidget
	HorizontalLayout_5          *widgets.QHBoxLayout
	Label_5                     *widgets.QLabel
	SingleTechnique             *widgets.QLabel
	Widget_6                    *widgets.QWidget
	HorizontalLayout_6          *widgets.QHBoxLayout
	Label_6                     *widgets.QLabel
	SingleVector                *widgets.QLabel
	GroupBox_3                  *widgets.QGroupBox
	VerticalLayout_2            *widgets.QVBoxLayout
	Widget                      *widgets.QWidget
	HorizontalLayout            *widgets.QHBoxLayout
	Label                       *widgets.QLabel
	SingleDatabaseType          *widgets.QLabel
	Widget_2                    *widgets.QWidget
	HorizontalLayout_2          *widgets.QHBoxLayout
	Label_2                     *widgets.QLabel
	SingleDatabaseVersion       *widgets.QLabel
	Widget_3                    *widgets.QWidget
	HorizontalLayout_3          *widgets.QHBoxLayout
	Label_3                     *widgets.QLabel
	SingleCurrentUser           *widgets.QLabel
	VerticalSpacer              *widgets.QSpacerItem
	GroupBox                    *widgets.QGroupBox
	VerticalLayout_5            *widgets.QVBoxLayout
	Widget_9                    *widgets.QWidget
	HorizontalLayout_8          *widgets.QHBoxLayout
	SingleGatherStructure       *widgets.QPushButton
	SingleGatherStructureCancel *widgets.QPushButton
	Widget_11                   *widgets.QWidget
	HorizontalLayout_10         *widgets.QHBoxLayout
	SingleSkipInformationSchema *widgets.QCheckBox
	SingleDatabaseStructure     *widgets.QTreeWidget
	SingleStructureProgress     *widgets.QProgressBar
	GroupBox_4                  *widgets.QGroupBox
	VerticalLayout_4            *widgets.QVBoxLayout
	Widget_8                    *widgets.QWidget
	HorizontalLayout_7          *widgets.QHBoxLayout
	SingleDumpStart             *widgets.QPushButton
	SingleDumpStop              *widgets.QPushButton
	Widget_10                   *widgets.QWidget
	HorizontalLayout_9          *widgets.QHBoxLayout
	Label_7                     *widgets.QLabel
	SpinBox                     *widgets.QSpinBox
	SingleDumpOutput            *widgets.QTableWidget
	SingleTableDumpProgress     *widgets.QProgressBar
}

func NewSingleSiteAnalyzer(p widgets.QWidget_ITF) *SingleSiteAnalyzer {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &SingleSiteAnalyzer{QWidget: widgets.NewQWidget(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *SingleSiteAnalyzer) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("SingleSiteAnalyzer")
	}
	w.Resize2(1129, 513)
	w.SetMinimumSize(core.NewQSize2(1129, 513))
	w.GridLayout = widgets.NewQGridLayout(w)
	w.GridLayout.SetObjectName("gridLayout")
	w.Widget_7 = widgets.NewQWidget(w, 0)
	w.Widget_7.SetObjectName("widget_7")
	w.VerticalLayout_3 = widgets.NewQVBoxLayout2(w.Widget_7)
	w.VerticalLayout_3.SetObjectName("verticalLayout_3")
	w.GroupBox_2 = widgets.NewQGroupBox(w.Widget_7)
	w.GroupBox_2.SetObjectName("groupBox_2")
	w.GroupBox_2.SetAlignment(int(core.Qt__AlignCenter))
	w.VerticalLayout = widgets.NewQVBoxLayout2(w.GroupBox_2)
	w.VerticalLayout.SetObjectName("verticalLayout")
	w.SingleSiteText = widgets.NewQLineEdit(w.GroupBox_2)
	w.SingleSiteText.SetObjectName("singleSiteText")
	w.SingleSiteText.SetMinimumSize(core.NewQSize2(191, 0))
	w.VerticalLayout.QLayout.AddWidget(w.SingleSiteText)
	w.SingleSiteSelectInjection = widgets.NewQPushButton(w.GroupBox_2)
	w.SingleSiteSelectInjection.SetObjectName("singleSiteSelectInjection")
	w.VerticalLayout.QLayout.AddWidget(w.SingleSiteSelectInjection)
	w.Widget_4 = widgets.NewQWidget(w.GroupBox_2, 0)
	w.Widget_4.SetObjectName("widget_4")
	w.HorizontalLayout_4 = widgets.NewQHBoxLayout2(w.Widget_4)
	w.HorizontalLayout_4.SetObjectName("horizontalLayout_4")
	w.Label_4 = widgets.NewQLabel(w.Widget_4, 0)
	w.Label_4.SetObjectName("label_4")
	w.HorizontalLayout_4.QLayout.AddWidget(w.Label_4)
	w.SingleCountry = widgets.NewQLabel(w.Widget_4, 0)
	w.SingleCountry.SetObjectName("singleCountry")
	w.SingleCountry.SetLayoutDirection(core.Qt__LeftToRight)
	w.SingleCountry.SetAlignment(core.Qt__AlignRight | core.Qt__AlignTrailing | core.Qt__AlignVCenter)
	w.HorizontalLayout_4.QLayout.AddWidget(w.SingleCountry)
	w.VerticalLayout.QLayout.AddWidget(w.Widget_4)
	w.Widget_5 = widgets.NewQWidget(w.GroupBox_2, 0)
	w.Widget_5.SetObjectName("widget_5")
	w.HorizontalLayout_5 = widgets.NewQHBoxLayout2(w.Widget_5)
	w.HorizontalLayout_5.SetObjectName("horizontalLayout_5")
	w.Label_5 = widgets.NewQLabel(w.Widget_5, 0)
	w.Label_5.SetObjectName("label_5")
	w.HorizontalLayout_5.QLayout.AddWidget(w.Label_5)
	w.SingleTechnique = widgets.NewQLabel(w.Widget_5, 0)
	w.SingleTechnique.SetObjectName("singleTechnique")
	w.SingleTechnique.SetLayoutDirection(core.Qt__LeftToRight)
	w.SingleTechnique.SetAlignment(core.Qt__AlignRight | core.Qt__AlignTrailing | core.Qt__AlignVCenter)
	w.HorizontalLayout_5.QLayout.AddWidget(w.SingleTechnique)
	w.VerticalLayout.QLayout.AddWidget(w.Widget_5)
	w.Widget_6 = widgets.NewQWidget(w.GroupBox_2, 0)
	w.Widget_6.SetObjectName("widget_6")
	w.HorizontalLayout_6 = widgets.NewQHBoxLayout2(w.Widget_6)
	w.HorizontalLayout_6.SetObjectName("horizontalLayout_6")
	w.Label_6 = widgets.NewQLabel(w.Widget_6, 0)
	w.Label_6.SetObjectName("label_6")
	w.HorizontalLayout_6.QLayout.AddWidget(w.Label_6)
	w.SingleVector = widgets.NewQLabel(w.Widget_6, 0)
	w.SingleVector.SetObjectName("singleVector")
	w.SingleVector.SetLayoutDirection(core.Qt__LeftToRight)
	w.SingleVector.SetAlignment(core.Qt__AlignRight | core.Qt__AlignTrailing | core.Qt__AlignVCenter)
	w.HorizontalLayout_6.QLayout.AddWidget(w.SingleVector)
	w.VerticalLayout.QLayout.AddWidget(w.Widget_6)
	w.VerticalLayout_3.QLayout.AddWidget(w.GroupBox_2)
	w.GroupBox_3 = widgets.NewQGroupBox(w.Widget_7)
	w.GroupBox_3.SetObjectName("groupBox_3")
	w.VerticalLayout_2 = widgets.NewQVBoxLayout2(w.GroupBox_3)
	w.VerticalLayout_2.SetObjectName("verticalLayout_2")
	w.Widget = widgets.NewQWidget(w.GroupBox_3, 0)
	w.Widget.SetObjectName("widget")
	w.HorizontalLayout = widgets.NewQHBoxLayout2(w.Widget)
	w.HorizontalLayout.SetObjectName("horizontalLayout")
	w.Label = widgets.NewQLabel(w.Widget, 0)
	w.Label.SetObjectName("label")
	w.HorizontalLayout.QLayout.AddWidget(w.Label)
	w.SingleDatabaseType = widgets.NewQLabel(w.Widget, 0)
	w.SingleDatabaseType.SetObjectName("singleDatabaseType")
	w.SingleDatabaseType.SetAlignment(core.Qt__AlignRight | core.Qt__AlignTrailing | core.Qt__AlignVCenter)
	w.HorizontalLayout.QLayout.AddWidget(w.SingleDatabaseType)
	w.VerticalLayout_2.QLayout.AddWidget(w.Widget)
	w.Widget_2 = widgets.NewQWidget(w.GroupBox_3, 0)
	w.Widget_2.SetObjectName("widget_2")
	w.HorizontalLayout_2 = widgets.NewQHBoxLayout2(w.Widget_2)
	w.HorizontalLayout_2.SetObjectName("horizontalLayout_2")
	w.Label_2 = widgets.NewQLabel(w.Widget_2, 0)
	w.Label_2.SetObjectName("label_2")
	w.HorizontalLayout_2.QLayout.AddWidget(w.Label_2)
	w.SingleDatabaseVersion = widgets.NewQLabel(w.Widget_2, 0)
	w.SingleDatabaseVersion.SetObjectName("singleDatabaseVersion")
	w.SingleDatabaseVersion.SetLayoutDirection(core.Qt__LeftToRight)
	w.SingleDatabaseVersion.SetAlignment(core.Qt__AlignRight | core.Qt__AlignTrailing | core.Qt__AlignVCenter)
	w.HorizontalLayout_2.QLayout.AddWidget(w.SingleDatabaseVersion)
	w.VerticalLayout_2.QLayout.AddWidget(w.Widget_2)
	w.Widget_3 = widgets.NewQWidget(w.GroupBox_3, 0)
	w.Widget_3.SetObjectName("widget_3")
	w.HorizontalLayout_3 = widgets.NewQHBoxLayout2(w.Widget_3)
	w.HorizontalLayout_3.SetObjectName("horizontalLayout_3")
	w.Label_3 = widgets.NewQLabel(w.Widget_3, 0)
	w.Label_3.SetObjectName("label_3")
	w.HorizontalLayout_3.QLayout.AddWidget(w.Label_3)
	w.SingleCurrentUser = widgets.NewQLabel(w.Widget_3, 0)
	w.SingleCurrentUser.SetObjectName("singleCurrentUser")
	w.SingleCurrentUser.SetLayoutDirection(core.Qt__LeftToRight)
	w.SingleCurrentUser.SetAlignment(core.Qt__AlignRight | core.Qt__AlignTrailing | core.Qt__AlignVCenter)
	w.HorizontalLayout_3.QLayout.AddWidget(w.SingleCurrentUser)
	w.VerticalLayout_2.QLayout.AddWidget(w.Widget_3)
	w.VerticalSpacer = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout_2.AddItem(w.VerticalSpacer)
	w.VerticalLayout_3.QLayout.AddWidget(w.GroupBox_3)
	w.GridLayout.AddWidget3(w.Widget_7, 0, 0, 1, 1, 0)
	w.GroupBox = widgets.NewQGroupBox(w)
	w.GroupBox.SetObjectName("groupBox")
	w.GroupBox.SetAlignment(int(core.Qt__AlignCenter))
	w.VerticalLayout_5 = widgets.NewQVBoxLayout2(w.GroupBox)
	w.VerticalLayout_5.SetObjectName("verticalLayout_5")
	w.Widget_9 = widgets.NewQWidget(w.GroupBox, 0)
	w.Widget_9.SetObjectName("widget_9")
	w.HorizontalLayout_8 = widgets.NewQHBoxLayout2(w.Widget_9)
	w.HorizontalLayout_8.SetObjectName("horizontalLayout_8")
	w.SingleGatherStructure = widgets.NewQPushButton(w.Widget_9)
	w.SingleGatherStructure.SetObjectName("singleGatherStructure")
	w.HorizontalLayout_8.QLayout.AddWidget(w.SingleGatherStructure)
	w.SingleGatherStructureCancel = widgets.NewQPushButton(w.Widget_9)
	w.SingleGatherStructureCancel.SetObjectName("singleGatherStructureCancel")
	w.HorizontalLayout_8.QLayout.AddWidget(w.SingleGatherStructureCancel)
	w.VerticalLayout_5.QLayout.AddWidget(w.Widget_9)
	w.Widget_11 = widgets.NewQWidget(w.GroupBox, 0)
	w.Widget_11.SetObjectName("widget_11")
	w.HorizontalLayout_10 = widgets.NewQHBoxLayout2(w.Widget_11)
	w.HorizontalLayout_10.SetObjectName("horizontalLayout_10")
	w.SingleSkipInformationSchema = widgets.NewQCheckBox(w.Widget_11)
	w.SingleSkipInformationSchema.SetObjectName("singleSkipInformationSchema")
	w.SingleSkipInformationSchema.SetChecked(true)
	w.HorizontalLayout_10.QLayout.AddWidget(w.SingleSkipInformationSchema)
	w.VerticalLayout_5.QLayout.AddWidget(w.Widget_11)
	w.SingleDatabaseStructure = widgets.NewQTreeWidget(w.GroupBox)
	w.SingleDatabaseStructure.SetObjectName("singleDatabaseStructure")
	w.SingleDatabaseStructure.SetEditTriggers(widgets.QAbstractItemView__NoEditTriggers)
	w.SingleDatabaseStructure.SetUniformRowHeights(true)
	w.SingleDatabaseStructure.SetAnimated(true)
	w.SingleDatabaseStructure.SetHeaderHidden(false)
	w.SingleDatabaseStructure.Header().SetCascadingSectionResizes(true)
	w.SingleDatabaseStructure.Header().SetMinimumSectionSize(60)
	w.SingleDatabaseStructure.Header().SetDefaultSectionSize(170)
	w.VerticalLayout_5.QLayout.AddWidget(w.SingleDatabaseStructure)
	w.SingleStructureProgress = widgets.NewQProgressBar(w.GroupBox)
	w.SingleStructureProgress.SetObjectName("singleStructureProgress")
	w.SingleStructureProgress.SetValue(24)
	w.VerticalLayout_5.QLayout.AddWidget(w.SingleStructureProgress)
	w.GridLayout.AddWidget3(w.GroupBox, 0, 1, 1, 1, 0)
	w.GroupBox_4 = widgets.NewQGroupBox(w)
	w.GroupBox_4.SetObjectName("groupBox_4")
	w.GroupBox_4.SetAlignment(int(core.Qt__AlignCenter))
	w.VerticalLayout_4 = widgets.NewQVBoxLayout2(w.GroupBox_4)
	w.VerticalLayout_4.SetObjectName("verticalLayout_4")
	w.Widget_8 = widgets.NewQWidget(w.GroupBox_4, 0)
	w.Widget_8.SetObjectName("widget_8")
	w.HorizontalLayout_7 = widgets.NewQHBoxLayout2(w.Widget_8)
	w.HorizontalLayout_7.SetObjectName("horizontalLayout_7")
	w.SingleDumpStart = widgets.NewQPushButton(w.Widget_8)
	w.SingleDumpStart.SetObjectName("singleDumpStart")
	w.HorizontalLayout_7.QLayout.AddWidget(w.SingleDumpStart)
	w.SingleDumpStop = widgets.NewQPushButton(w.Widget_8)
	w.SingleDumpStop.SetObjectName("singleDumpStop")
	w.HorizontalLayout_7.QLayout.AddWidget(w.SingleDumpStop)
	w.VerticalLayout_4.QLayout.AddWidget(w.Widget_8)
	w.Widget_10 = widgets.NewQWidget(w.GroupBox_4, 0)
	w.Widget_10.SetObjectName("widget_10")
	w.HorizontalLayout_9 = widgets.NewQHBoxLayout2(w.Widget_10)
	w.HorizontalLayout_9.SetObjectName("horizontalLayout_9")
	w.Label_7 = widgets.NewQLabel(w.Widget_10, 0)
	w.Label_7.SetObjectName("label_7")
	sizePolicy := widgets.NewQSizePolicy2(widgets.QSizePolicy__Preferred, widgets.QSizePolicy__Preferred, 0)
	sizePolicy.SetHorizontalStretch(0)
	sizePolicy.SetVerticalStretch(0)
	sizePolicy.SetHeightForWidth(w.Label_7.SizePolicy().HasHeightForWidth())
	w.Label_7.SetSizePolicy(sizePolicy)
	w.HorizontalLayout_9.QLayout.AddWidget(w.Label_7)
	w.SpinBox = widgets.NewQSpinBox(w.Widget_10)
	w.SpinBox.SetObjectName("spinBox")
	w.HorizontalLayout_9.QLayout.AddWidget(w.SpinBox)
	w.VerticalLayout_4.QLayout.AddWidget(w.Widget_10)
	w.SingleDumpOutput = widgets.NewQTableWidget(w.GroupBox_4)
	if w.SingleDumpOutput.ColumnCount() < 3 {
		w.SingleDumpOutput.SetColumnCount(3)
	}
	__qtablewidgetitem := widgets.NewQTableWidgetItem(0)
	w.SingleDumpOutput.SetHorizontalHeaderItem(0, __qtablewidgetitem)
	__qtablewidgetitem1 := widgets.NewQTableWidgetItem(0)
	w.SingleDumpOutput.SetHorizontalHeaderItem(1, __qtablewidgetitem1)
	__qtablewidgetitem2 := widgets.NewQTableWidgetItem(0)
	w.SingleDumpOutput.SetHorizontalHeaderItem(2, __qtablewidgetitem2)
	w.SingleDumpOutput.SetObjectName("singleDumpOutput")
	w.VerticalLayout_4.QLayout.AddWidget(w.SingleDumpOutput)
	w.SingleTableDumpProgress = widgets.NewQProgressBar(w.GroupBox_4)
	w.SingleTableDumpProgress.SetObjectName("singleTableDumpProgress")
	w.SingleTableDumpProgress.SetValue(24)
	w.VerticalLayout_4.QLayout.AddWidget(w.SingleTableDumpProgress)
	w.GridLayout.AddWidget3(w.GroupBox_4, 0, 2, 1, 1, 0)
	w.retranslateUi()
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *SingleSiteAnalyzer) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("SingleSiteAnalyzer", "Single Site Analyzer", "", 0))
	w.GroupBox_2.SetTitle(core.QCoreApplication_Translate("SingleSiteAnalyzer", "Injeciton Information", "", 0))
	w.SingleSiteText.SetPlaceholderText(core.QCoreApplication_Translate("SingleSiteAnalyzer", "https://site.com/?id=1", "", 0))
	w.SingleSiteSelectInjection.SetText(core.QCoreApplication_Translate("SingleSiteAnalyzer", "Select Injection", "", 0))
	w.Label_4.SetText(core.QCoreApplication_Translate("SingleSiteAnalyzer", "Country:", "", 0))
	w.SingleCountry.SetText(core.QCoreApplication_Translate("SingleSiteAnalyzer", "N/A", "", 0))
	w.Label_5.SetText(core.QCoreApplication_Translate("SingleSiteAnalyzer", "Technique:", "", 0))
	w.SingleTechnique.SetText(core.QCoreApplication_Translate("SingleSiteAnalyzer", "N/A", "", 0))
	w.Label_6.SetText(core.QCoreApplication_Translate("SingleSiteAnalyzer", "Vector:", "", 0))
	w.SingleVector.SetText(core.QCoreApplication_Translate("SingleSiteAnalyzer", "N/A", "", 0))
	w.GroupBox_3.SetTitle(core.QCoreApplication_Translate("SingleSiteAnalyzer", "Server Information", "", 0))
	w.Label.SetText(core.QCoreApplication_Translate("SingleSiteAnalyzer", "Database:", "", 0))
	w.SingleDatabaseType.SetText(core.QCoreApplication_Translate("SingleSiteAnalyzer", "N/A", "", 0))
	w.Label_2.SetText(core.QCoreApplication_Translate("SingleSiteAnalyzer", "Version:", "", 0))
	w.SingleDatabaseVersion.SetText(core.QCoreApplication_Translate("SingleSiteAnalyzer", "N/A", "", 0))
	w.Label_3.SetText(core.QCoreApplication_Translate("SingleSiteAnalyzer", "User:", "", 0))
	w.SingleCurrentUser.SetText(core.QCoreApplication_Translate("SingleSiteAnalyzer", "N/A", "", 0))
	w.GroupBox.SetTitle(core.QCoreApplication_Translate("SingleSiteAnalyzer", "Database Structure", "", 0))
	w.SingleGatherStructure.SetText(core.QCoreApplication_Translate("SingleSiteAnalyzer", "Gather Structure", "", 0))
	w.SingleGatherStructureCancel.SetText(core.QCoreApplication_Translate("SingleSiteAnalyzer", "Cancel", "", 0))
	w.SingleSkipInformationSchema.SetText(core.QCoreApplication_Translate("SingleSiteAnalyzer", "Skip Information Schema", "", 0))
	___qtreewidgetitem := w.SingleDatabaseStructure.HeaderItem()
	___qtreewidgetitem.SetText(2, core.QCoreApplication_Translate("SingleSiteAnalyzer", "Select", "", 0))
	___qtreewidgetitem.SetText(1, core.QCoreApplication_Translate("SingleSiteAnalyzer", "Info", "", 0))
	___qtreewidgetitem.SetText(0, core.QCoreApplication_Translate("SingleSiteAnalyzer", "Name", "", 0))
	w.GroupBox_4.SetTitle(core.QCoreApplication_Translate("SingleSiteAnalyzer", "Table Dumping", "", 0))
	w.SingleDumpStart.SetText(core.QCoreApplication_Translate("SingleSiteAnalyzer", "Start", "", 0))
	w.SingleDumpStop.SetText(core.QCoreApplication_Translate("SingleSiteAnalyzer", "Stop", "", 0))
	w.Label_7.SetText(core.QCoreApplication_Translate("SingleSiteAnalyzer", "Workers", "", 0))
	___qtablewidgetitem := w.SingleDumpOutput.HorizontalHeaderItem(0)
	___qtablewidgetitem.SetText(core.QCoreApplication_Translate("SingleSiteAnalyzer", "Table Name", "", 0))
	___qtablewidgetitem1 := w.SingleDumpOutput.HorizontalHeaderItem(1)
	___qtablewidgetitem1.SetText(core.QCoreApplication_Translate("SingleSiteAnalyzer", "Rows", "", 0))
	___qtablewidgetitem2 := w.SingleDumpOutput.HorizontalHeaderItem(2)
	___qtablewidgetitem2.SetText(core.QCoreApplication_Translate("SingleSiteAnalyzer", "Errors", "", 0))

}
