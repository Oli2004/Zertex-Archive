package qtui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
)

type __mainwindow struct{}

func (*__mainwindow) init() {}

type MainWindow struct {
	*__mainwindow
	*widgets.QMainWindow
	ActionExit                   *widgets.QAction
	ActionOpen_Settings          *widgets.QAction
	Centralwidget                *widgets.QWidget
	VerticalLayout_15            *widgets.QVBoxLayout
	GridLayout_2                 *widgets.QGridLayout
	Frame_2                      *widgets.QFrame
	VerticalLayout_13            *widgets.QVBoxLayout
	StatisticsGroup              *widgets.QGroupBox
	GridLayout                   *widgets.QGridLayout
	Label_9                      *widgets.QLabel
	Label_10                     *widgets.QLabel
	Label_8                      *widgets.QLabel
	Label_2                      *widgets.QLabel
	Label_4                      *widgets.QLabel
	Label_3                      *widgets.QLabel
	Label_16                     *widgets.QLabel
	Label_23                     *widgets.QLabel
	StatsModule                  *widgets.QLabel
	StatsCurrentTimeLcd          *widgets.QLCDNumber
	StatsRuntimeLcd              *widgets.QLCDNumber
	StatsRequestsLcd             *widgets.QLCDNumber
	StatsErrorLcd                *widgets.QLCDNumber
	StatsRPSLcd                  *widgets.QLCDNumber
	DataThreadsLcd               *widgets.QLCDNumber
	DataWorkersLcd               *widgets.QLCDNumber
	HorizontalLayout_13          *widgets.QHBoxLayout
	DataGroup                    *widgets.QGroupBox
	GridLayout_4                 *widgets.QGridLayout
	Label_6                      *widgets.QLabel
	DataProxiesLcd               *widgets.QLCDNumber
	DataInjectablesLcd           *widgets.QLCDNumber
	Label_5                      *widgets.QLabel
	DataDorksLcd                 *widgets.QLCDNumber
	DataUrlsLcd                  *widgets.QLCDNumber
	Label_7                      *widgets.QLabel
	Label_12                     *widgets.QLabel
	DumpStatsGroup               *widgets.QGroupBox
	GridLayout_3                 *widgets.QGridLayout
	Label_13                     *widgets.QLabel
	Label_15                     *widgets.QLabel
	Label_14                     *widgets.QLabel
	Label_30                     *widgets.QLabel
	DumpStatsTablesLcd           *widgets.QLCDNumber
	DumpStatsColumnsLcd          *widgets.QLCDNumber
	DumpStatsRowsLcd             *widgets.QLCDNumber
	DumpStatsRpm                 *widgets.QLCDNumber
	TabControl                   *widgets.QTabWidget
	TabWelcome                   *widgets.QWidget
	GridLayout_5                 *widgets.QGridLayout
	GroupBox_4                   *widgets.QGroupBox
	VerticalLayout_18            *widgets.QVBoxLayout
	Widget_10                    *widgets.QWidget
	HorizontalLayout_14          *widgets.QHBoxLayout
	CustomerChatbox              *widgets.QListWidget
	CustomerList                 *widgets.QListWidget
	Widget_11                    *widgets.QWidget
	HorizontalLayout_15          *widgets.QHBoxLayout
	ChatMessagebox               *widgets.QLineEdit
	ChatSendMessage              *widgets.QPushButton
	GroupBox_3                   *widgets.QGroupBox
	VerticalLayout_19            *widgets.QVBoxLayout
	NewsBox                      *widgets.QListWidget
	TabGenerator                 *widgets.QWidget
	HorizontalLayout_19          *widgets.QHBoxLayout
	Widget_16                    *widgets.QWidget
	VerticalLayout_28            *widgets.QVBoxLayout
	GroupParameters              *widgets.QGroupBox
	VerticalLayout_27            *widgets.QVBoxLayout
	ComboBoxParameters           *widgets.QComboBox
	ButtonParametersSettings     *widgets.QPushButton
	PlainTextEditParameters      *widgets.QPlainTextEdit
	GeneratorLimiterCheckbox     *widgets.QCheckBox
	GeneratorLimiterSpinbox      *widgets.QSpinBox
	GroupGenerator               *widgets.QGroupBox
	VerticalLayout_26            *widgets.QVBoxLayout
	TableGenerator               *widgets.QTableWidget
	GeneratorProgress            *widgets.QProgressBar
	Frame_17                     *widgets.QFrame
	VerticalLayout_30            *widgets.QVBoxLayout
	GroupPatterns                *widgets.QGroupBox
	VerticalLayout_29            *widgets.QVBoxLayout
	PlainTextEditPatterns        *widgets.QPlainTextEdit
	Frame_6                      *widgets.QFrame
	HorizontalLayout_18          *widgets.QHBoxLayout
	GeneratorStartBtn            *widgets.QPushButton
	GeneratorStopBtn             *widgets.QPushButton
	TabParser                    *widgets.QWidget
	HorizontalLayout_3           *widgets.QHBoxLayout
	Widget                       *widgets.QWidget
	VerticalLayout_4             *widgets.QVBoxLayout
	ParserSearchEnginesGroup     *widgets.QGroupBox
	VerticalLayout_3             *widgets.QVBoxLayout
	ParserGoogleCheckbox         *widgets.QCheckBox
	ParserBingCheckbox           *widgets.QCheckBox
	ParserAOLCheckbox            *widgets.QCheckBox
	ParserMWSCheckbox            *widgets.QCheckBox
	ParserDuckDuckGoCheckbox     *widgets.QCheckBox
	ParserEcosiaCheckbox         *widgets.QCheckBox
	ParserStartPageCheckbox      *widgets.QCheckBox
	ParserYahooCheckbox          *widgets.QCheckBox
	ParserYandexCheckbox         *widgets.QCheckBox
	VerticalSpacer_7             *widgets.QSpacerItem
	ParserSettingsGroup          *widgets.QGroupBox
	VerticalLayout_2             *widgets.QVBoxLayout
	ParserFilterUrls             *widgets.QCheckBox
	ParserPagesSpinbox           *widgets.QSpinBox
	ParserCustomParams           *widgets.QLineEdit
	VerticalSpacer               *widgets.QSpacerItem
	ParserGroup                  *widgets.QGroupBox
	VerticalLayout_12            *widgets.QVBoxLayout
	ParserDorksTextbox           *widgets.QPlainTextEdit
	ParserProgress               *widgets.QProgressBar
	Frame                        *widgets.QFrame
	VerticalLayout               *widgets.QVBoxLayout
	Widget_2                     *widgets.QWidget
	HorizontalLayout_4           *widgets.QHBoxLayout
	ParserLoadDorksBtn           *widgets.QPushButton
	ParserClearDorksBtn          *widgets.QPushButton
	Widget_3                     *widgets.QWidget
	HorizontalLayout_5           *widgets.QHBoxLayout
	ParserStartBtn               *widgets.QPushButton
	ParserStopBtn                *widgets.QPushButton
	VerticalSpacer_2             *widgets.QSpacerItem
	TabExploiter                 *widgets.QWidget
	HorizontalLayout_9           *widgets.QHBoxLayout
	Widget_12                    *widgets.QWidget
	VerticalLayout_20            *widgets.QVBoxLayout
	ExploiterTechniquesGroup     *widgets.QGroupBox
	VerticalLayout_6             *widgets.QVBoxLayout
	ExploiterErrorCheckbox       *widgets.QCheckBox
	ExploiterUnionCheckbox       *widgets.QCheckBox
	ExploiterBlindCheckbox       *widgets.QCheckBox
	ExploiterStackedCheckbox     *widgets.QCheckBox
	Label                        *widgets.QLabel
	ExploiterIntensityCombo      *widgets.QComboBox
	ExploiterHeuristsicsCheckbox *widgets.QCheckBox
	ExploiterDatabaseTypes       *widgets.QGroupBox
	VerticalLayout_7             *widgets.QVBoxLayout
	ExploiterMySQL               *widgets.QCheckBox
	ExploiterOracle              *widgets.QCheckBox
	ExploiterPostgreSQL          *widgets.QCheckBox
	ExploiterMSSQL               *widgets.QCheckBox
	VerticalSpacer_10            *widgets.QSpacerItem
	ExploiterGroup               *widgets.QGroupBox
	VerticalLayout_21            *widgets.QVBoxLayout
	ExploiterInjectablesTable    *widgets.QTableWidget
	ExploiterProgress            *widgets.QProgressBar
	Frame_3                      *widgets.QFrame
	VerticalLayout_31            *widgets.QVBoxLayout
	Widget_4                     *widgets.QWidget
	HorizontalLayout_7           *widgets.QHBoxLayout
	ExploiterLoadUrlsBtn         *widgets.QPushButton
	ExploiterClearUrlsBtn        *widgets.QPushButton
	Widget_5                     *widgets.QWidget
	HorizontalLayout_8           *widgets.QHBoxLayout
	ExploiterStartBtn            *widgets.QPushButton
	ExploiterStopBtn             *widgets.QPushButton
	VerticalSpacer_3             *widgets.QSpacerItem
	GroupBox_5                   *widgets.QGroupBox
	VerticalLayout_5             *widgets.QVBoxLayout
	Label_19                     *widgets.QLabel
	ExploiterThreads             *widgets.QSpinBox
	Label_20                     *widgets.QLabel
	ExploiterWorkers             *widgets.QSpinBox
	TabDumper                    *widgets.QWidget
	HorizontalLayout_12          *widgets.QHBoxLayout
	Widget_6                     *widgets.QWidget
	VerticalLayout_10            *widgets.QVBoxLayout
	DumperMethodsGroup           *widgets.QGroupBox
	VerticalLayout_11            *widgets.QVBoxLayout
	DumperTargetedCheckbox       *widgets.QCheckBox
	DumperTargetSettingsBtn      *widgets.QPushButton
	DumperKeepBlanksCheckbox     *widgets.QCheckBox
	DumperDIOSCheckbox           *widgets.QCheckBox
	Line_3                       *widgets.QFrame
	DumperMinRowsCheckbox        *widgets.QCheckBox
	DumperMinRowsSpinbox         *widgets.QSpinBox
	Line_2                       *widgets.QFrame
	DumperAutoSkip               *widgets.QCheckBox
	DumperAutoSkipSpinbox        *widgets.QSpinBox
	VerticalSpacer_6             *widgets.QSpacerItem
	DumperGroup                  *widgets.QGroupBox
	VerticalLayout_9             *widgets.QVBoxLayout
	DumperTableWidget            *widgets.QTableWidget
	DumperProgress               *widgets.QProgressBar
	Frame_4                      *widgets.QFrame
	VerticalLayout_32            *widgets.QVBoxLayout
	Widget_7                     *widgets.QWidget
	HorizontalLayout_10          *widgets.QHBoxLayout
	DumperLoadInjectablesBtn     *widgets.QPushButton
	DumperClearInjectablesBtn    *widgets.QPushButton
	Widget_8                     *widgets.QWidget
	HorizontalLayout_11          *widgets.QHBoxLayout
	DumperStartBtn               *widgets.QPushButton
	DumperStopBtn                *widgets.QPushButton
	DumperOpenAnalyzer           *widgets.QPushButton
	VerticalSpacer_5             *widgets.QSpacerItem
	GroupBox_6                   *widgets.QGroupBox
	VerticalLayout_8             *widgets.QVBoxLayout
	Label_21                     *widgets.QLabel
	DumperThreads                *widgets.QSpinBox
	Label_22                     *widgets.QLabel
	DumperWorkers                *widgets.QSpinBox
	TabAutoSheller               *widgets.QWidget
	HorizontalLayout_6           *widgets.QHBoxLayout
	Widget_15                    *widgets.QWidget
	VerticalLayout_24            *widgets.QVBoxLayout
	AutoShellTypesGroup          *widgets.QGroupBox
	VerticalLayout_22            *widgets.QVBoxLayout
	AutoShellASPCheckbox         *widgets.QCheckBox
	AutoShellPHPCheckbox         *widgets.QCheckBox
	Line                         *widgets.QFrame
	Label_17                     *widgets.QLabel
	AutoShellKey                 *widgets.QLineEdit
	Label_18                     *widgets.QLabel
	AutoShellFile                *widgets.QLineEdit
	VerticalSpacer_11            *widgets.QSpacerItem
	GroupBox_7                   *widgets.QGroupBox
	VerticalLayout_25            *widgets.QVBoxLayout
	AutoShellTable               *widgets.QTableWidget
	AutoShellProgress            *widgets.QProgressBar
	Frame_61                     *widgets.QFrame
	VerticalLayout_23            *widgets.QVBoxLayout
	Widget_13                    *widgets.QWidget
	HorizontalLayout_16          *widgets.QHBoxLayout
	AutoShellLoadInjectablesBtn  *widgets.QPushButton
	AutoShellClearInjectablesBtn *widgets.QPushButton
	Widget_14                    *widgets.QWidget
	HorizontalLayout_17          *widgets.QHBoxLayout
	AutoShellStartBtn            *widgets.QPushButton
	AutoShellStopBtn             *widgets.QPushButton
	VerticalSpacer_12            *widgets.QSpacerItem
	TabUtils                     *widgets.QWidget
	HorizontalLayout_2           *widgets.QHBoxLayout
	GroupBox                     *widgets.QGroupBox
	HorizontalLayout             *widgets.QHBoxLayout
	CleanerUrlsTextBox           *widgets.QPlainTextEdit
	Frame_5                      *widgets.QFrame
	VerticalLayout_14            *widgets.QVBoxLayout
	CleanerDuplicateDomains      *widgets.QCheckBox
	CleanerQueryParam            *widgets.QCheckBox
	CleanerBtn                   *widgets.QPushButton
	VerticalSpacer_8             *widgets.QSpacerItem
	GroupBox_2                   *widgets.QGroupBox
	VerticalLayout_17            *widgets.QVBoxLayout
	Widget_9                     *widgets.QWidget
	VerticalLayout_16            *widgets.QVBoxLayout
	Label_11                     *widgets.QLabel
	WebUIAddress                 *widgets.QLineEdit
	WebUILaunchBtn               *widgets.QPushButton
	VerticalSpacer_9             *widgets.QSpacerItem
	Tab                          *widgets.QWidget
	HorizontalLayout_25          *widgets.QHBoxLayout
	AutoShellTypesGroup_2        *widgets.QGroupBox
	VerticalLayout_33            *widgets.QVBoxLayout
	AntipubSavePublic            *widgets.QCheckBox
	AntipubAutoSync              *widgets.QCheckBox
	VerticalSpacer_13            *widgets.QSpacerItem
	Widget_22                    *widgets.QWidget
	VerticalLayout_39            *widgets.QVBoxLayout
	Widget_17                    *widgets.QWidget
	HorizontalLayout_21          *widgets.QHBoxLayout
	HorizontalSpacer             *widgets.QSpacerItem
	GroupBox_8                   *widgets.QGroupBox
	VerticalLayout_37            *widgets.QVBoxLayout
	Frame_7                      *widgets.QFrame
	VerticalLayout_34            *widgets.QVBoxLayout
	Label_24                     *widgets.QLabel
	AntipubLinkCount             *widgets.QLCDNumber
	Frame_8                      *widgets.QFrame
	VerticalLayout_35            *widgets.QVBoxLayout
	Label_25                     *widgets.QLabel
	AntipubDomainCount           *widgets.QLCDNumber
	Frame_9                      *widgets.QFrame
	VerticalLayout_36            *widgets.QVBoxLayout
	AntipubSizeOnDisk            *widgets.QLabel
	Widget_18                    *widgets.QWidget
	HorizontalLayout_20          *widgets.QHBoxLayout
	AntipubLinkMode              *widgets.QRadioButton
	AntipubDomainMode            *widgets.QRadioButton
	HorizontalSpacer_2           *widgets.QSpacerItem
	VerticalSpacer_4             *widgets.QSpacerItem
	AntipubProgress              *widgets.QProgressBar
	Widget_23                    *widgets.QWidget
	GridLayout_6                 *widgets.QGridLayout
	Label_26                     *widgets.QLabel
	Label_27                     *widgets.QLabel
	Label_28                     *widgets.QLabel
	Label_29                     *widgets.QLabel
	AntipubLoaded                *widgets.QLCDNumber
	AntipubPublic                *widgets.QLCDNumber
	AntipubPrivate               *widgets.QLCDNumber
	AntipubPrivateRatio          *widgets.QLCDNumber
	AutoShellTypesGroup_3        *widgets.QGroupBox
	VerticalLayout_38            *widgets.QVBoxLayout
	Widget_19                    *widgets.QWidget
	HorizontalLayout_22          *widgets.QHBoxLayout
	AntipubLoadUrls              *widgets.QPushButton
	AntipubClearUrls             *widgets.QPushButton
	Widget_20                    *widgets.QWidget
	HorizontalLayout_23          *widgets.QHBoxLayout
	AntipubStart                 *widgets.QPushButton
	AntipubStop                  *widgets.QPushButton
	Widget_21                    *widgets.QWidget
	HorizontalLayout_24          *widgets.QHBoxLayout
	AntipubLoadToDB              *widgets.QPushButton
	AntipubExportDB              *widgets.QPushButton
	VerticalSpacer_14            *widgets.QSpacerItem
	Menubar                      *widgets.QMenuBar
	MenuFile                     *widgets.QMenu
	MenuSettings                 *widgets.QMenu
	Statusbar                    *widgets.QStatusBar
}

func NewMainWindow(p widgets.QWidget_ITF) *MainWindow {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &MainWindow{QMainWindow: widgets.NewQMainWindow(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *MainWindow) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("MainWindow")
	}
	w.Resize2(1266, 853)
	w.SetMinimumSize(core.NewQSize2(1266, 853))
	w.ActionExit = widgets.NewQAction(w)
	w.ActionExit.SetObjectName("actionExit")
	w.ActionOpen_Settings = widgets.NewQAction(w)
	w.ActionOpen_Settings.SetObjectName("actionOpen_Settings")
	w.ActionOpen_Settings.SetCheckable(false)
	w.Centralwidget = widgets.NewQWidget(w, 0)
	w.Centralwidget.SetObjectName("centralwidget")
	sizePolicy := widgets.NewQSizePolicy2(widgets.QSizePolicy__Expanding, widgets.QSizePolicy__Expanding, 0)
	sizePolicy.SetHorizontalStretch(0)
	sizePolicy.SetVerticalStretch(0)
	sizePolicy.SetHeightForWidth(w.Centralwidget.SizePolicy().HasHeightForWidth())
	w.Centralwidget.SetSizePolicy(sizePolicy)
	w.VerticalLayout_15 = widgets.NewQVBoxLayout2(w.Centralwidget)
	w.VerticalLayout_15.SetObjectName("verticalLayout_15")
	w.GridLayout_2 = widgets.NewQGridLayout2()
	w.GridLayout_2.SetObjectName("gridLayout_2")
	w.GridLayout_2.SetSizeConstraint(widgets.QLayout__SetNoConstraint)
	w.Frame_2 = widgets.NewQFrame(w.Centralwidget, 0)
	w.Frame_2.SetObjectName("frame_2")
	w.Frame_2.SetFrameShape(widgets.QFrame__StyledPanel)
	w.Frame_2.SetFrameShadow(widgets.QFrame__Raised)
	w.VerticalLayout_13 = widgets.NewQVBoxLayout2(w.Frame_2)
	w.VerticalLayout_13.SetObjectName("verticalLayout_13")
	w.StatisticsGroup = widgets.NewQGroupBox(w.Frame_2)
	w.StatisticsGroup.SetObjectName("statisticsGroup")
	w.StatisticsGroup.SetMinimumSize(core.NewQSize2(0, 100))
	w.StatisticsGroup.SetAlignment(int(core.Qt__AlignCenter))
	w.GridLayout = widgets.NewQGridLayout(w.StatisticsGroup)
	w.GridLayout.SetObjectName("gridLayout")
	w.Label_9 = widgets.NewQLabel(w.StatisticsGroup, 0)
	w.Label_9.SetObjectName("label_9")
	w.Label_9.SetAlignment(core.Qt__AlignCenter)
	w.GridLayout.AddWidget3(w.Label_9, 0, 0, 1, 1, 0)
	w.Label_10 = widgets.NewQLabel(w.StatisticsGroup, 0)
	w.Label_10.SetObjectName("label_10")
	w.Label_10.SetAlignment(core.Qt__AlignCenter)
	w.GridLayout.AddWidget3(w.Label_10, 0, 1, 1, 1, 0)
	w.Label_8 = widgets.NewQLabel(w.StatisticsGroup, 0)
	w.Label_8.SetObjectName("label_8")
	w.Label_8.SetAlignment(core.Qt__AlignCenter)
	w.GridLayout.AddWidget3(w.Label_8, 0, 2, 1, 1, 0)
	w.Label_2 = widgets.NewQLabel(w.StatisticsGroup, 0)
	w.Label_2.SetObjectName("label_2")
	w.Label_2.SetAlignment(core.Qt__AlignCenter)
	w.GridLayout.AddWidget3(w.Label_2, 0, 3, 1, 1, 0)
	w.Label_4 = widgets.NewQLabel(w.StatisticsGroup, 0)
	w.Label_4.SetObjectName("label_4")
	w.Label_4.SetAlignment(core.Qt__AlignCenter)
	w.GridLayout.AddWidget3(w.Label_4, 0, 4, 1, 1, 0)
	w.Label_3 = widgets.NewQLabel(w.StatisticsGroup, 0)
	w.Label_3.SetObjectName("label_3")
	w.Label_3.SetAlignment(core.Qt__AlignCenter)
	w.GridLayout.AddWidget3(w.Label_3, 0, 5, 1, 1, 0)
	w.Label_16 = widgets.NewQLabel(w.StatisticsGroup, 0)
	w.Label_16.SetObjectName("label_16")
	w.Label_16.SetAlignment(core.Qt__AlignCenter)
	w.GridLayout.AddWidget3(w.Label_16, 0, 6, 1, 1, 0)
	w.Label_23 = widgets.NewQLabel(w.StatisticsGroup, 0)
	w.Label_23.SetObjectName("label_23")
	w.Label_23.SetAlignment(core.Qt__AlignCenter)
	w.GridLayout.AddWidget3(w.Label_23, 0, 7, 1, 1, 0)
	w.StatsModule = widgets.NewQLabel(w.StatisticsGroup, 0)
	w.StatsModule.SetObjectName("statsModule")
	font := gui.NewQFont()
	font.SetPointSize(18)
	w.StatsModule.SetFont(font)
	w.StatsModule.SetAlignment(core.Qt__AlignCenter)
	w.GridLayout.AddWidget3(w.StatsModule, 1, 0, 1, 1, 0)
	w.StatsCurrentTimeLcd = widgets.NewQLCDNumber(w.StatisticsGroup)
	w.StatsCurrentTimeLcd.SetObjectName("statsCurrentTimeLcd")
	w.StatsCurrentTimeLcd.SetDigitCount(8)
	w.StatsCurrentTimeLcd.SetMode(widgets.QLCDNumber__Dec)
	w.StatsCurrentTimeLcd.SetSegmentStyle(widgets.QLCDNumber__Flat)
	w.StatsCurrentTimeLcd.SetProperty("intValue", core.NewQVariant1(0))
	w.GridLayout.AddWidget3(w.StatsCurrentTimeLcd, 1, 1, 1, 1, 0)
	w.StatsRuntimeLcd = widgets.NewQLCDNumber(w.StatisticsGroup)
	w.StatsRuntimeLcd.SetObjectName("statsRuntimeLcd")
	w.StatsRuntimeLcd.SetDigitCount(8)
	w.StatsRuntimeLcd.SetSegmentStyle(widgets.QLCDNumber__Filled)
	w.StatsRuntimeLcd.SetProperty("intValue", core.NewQVariant1(0))
	w.GridLayout.AddWidget3(w.StatsRuntimeLcd, 1, 2, 1, 1, 0)
	w.StatsRequestsLcd = widgets.NewQLCDNumber(w.StatisticsGroup)
	w.StatsRequestsLcd.SetObjectName("statsRequestsLcd")
	w.StatsRequestsLcd.SetDigitCount(10)
	w.StatsRequestsLcd.SetProperty("intValue", core.NewQVariant1(0))
	w.GridLayout.AddWidget3(w.StatsRequestsLcd, 1, 3, 1, 1, 0)
	w.StatsErrorLcd = widgets.NewQLCDNumber(w.StatisticsGroup)
	w.StatsErrorLcd.SetObjectName("statsErrorLcd")
	w.StatsErrorLcd.SetDigitCount(10)
	w.StatsErrorLcd.SetProperty("intValue", core.NewQVariant1(0))
	w.GridLayout.AddWidget3(w.StatsErrorLcd, 1, 4, 1, 1, 0)
	w.StatsRPSLcd = widgets.NewQLCDNumber(w.StatisticsGroup)
	w.StatsRPSLcd.SetObjectName("statsRPSLcd")
	w.StatsRPSLcd.SetDigitCount(10)
	w.StatsRPSLcd.SetProperty("intValue", core.NewQVariant1(0))
	w.GridLayout.AddWidget3(w.StatsRPSLcd, 1, 5, 1, 1, 0)
	w.DataThreadsLcd = widgets.NewQLCDNumber(w.StatisticsGroup)
	w.DataThreadsLcd.SetObjectName("dataThreadsLcd")
	w.DataThreadsLcd.SetDigitCount(10)
	w.DataThreadsLcd.SetProperty("intValue", core.NewQVariant1(0))
	w.GridLayout.AddWidget3(w.DataThreadsLcd, 1, 6, 1, 1, 0)
	w.DataWorkersLcd = widgets.NewQLCDNumber(w.StatisticsGroup)
	w.DataWorkersLcd.SetObjectName("dataWorkersLcd")
	w.DataWorkersLcd.SetDigitCount(10)
	w.DataWorkersLcd.SetProperty("intValue", core.NewQVariant1(0))
	w.GridLayout.AddWidget3(w.DataWorkersLcd, 1, 7, 1, 1, 0)
	w.VerticalLayout_13.QLayout.AddWidget(w.StatisticsGroup)
	w.HorizontalLayout_13 = widgets.NewQHBoxLayout()
	w.HorizontalLayout_13.SetObjectName("horizontalLayout_13")
	w.DataGroup = widgets.NewQGroupBox(w.Frame_2)
	w.DataGroup.SetObjectName("dataGroup")
	w.DataGroup.SetMinimumSize(core.NewQSize2(0, 108))
	w.DataGroup.SetAlignment(int(core.Qt__AlignCenter))
	w.GridLayout_4 = widgets.NewQGridLayout(w.DataGroup)
	w.GridLayout_4.SetObjectName("gridLayout_4")
	w.Label_6 = widgets.NewQLabel(w.DataGroup, 0)
	w.Label_6.SetObjectName("label_6")
	w.Label_6.SetAlignment(core.Qt__AlignCenter)
	w.GridLayout_4.AddWidget3(w.Label_6, 0, 2, 1, 1, 0)
	w.DataProxiesLcd = widgets.NewQLCDNumber(w.DataGroup)
	w.DataProxiesLcd.SetObjectName("dataProxiesLcd")
	w.DataProxiesLcd.SetDigitCount(8)
	w.DataProxiesLcd.SetProperty("intValue", core.NewQVariant1(0))
	w.GridLayout_4.AddWidget3(w.DataProxiesLcd, 1, 0, 1, 1, 0)
	w.DataInjectablesLcd = widgets.NewQLCDNumber(w.DataGroup)
	w.DataInjectablesLcd.SetObjectName("dataInjectablesLcd")
	w.DataInjectablesLcd.SetDigitCount(10)
	w.DataInjectablesLcd.SetProperty("intValue", core.NewQVariant1(0))
	w.GridLayout_4.AddWidget3(w.DataInjectablesLcd, 1, 3, 1, 1, 0)
	w.Label_5 = widgets.NewQLabel(w.DataGroup, 0)
	w.Label_5.SetObjectName("label_5")
	w.Label_5.SetAlignment(core.Qt__AlignCenter)
	w.GridLayout_4.AddWidget3(w.Label_5, 0, 0, 1, 1, 0)
	w.DataDorksLcd = widgets.NewQLCDNumber(w.DataGroup)
	w.DataDorksLcd.SetObjectName("dataDorksLcd")
	w.DataDorksLcd.SetDigitCount(10)
	w.DataDorksLcd.SetProperty("intValue", core.NewQVariant1(0))
	w.GridLayout_4.AddWidget3(w.DataDorksLcd, 1, 1, 1, 1, 0)
	w.DataUrlsLcd = widgets.NewQLCDNumber(w.DataGroup)
	w.DataUrlsLcd.SetObjectName("dataUrlsLcd")
	w.DataUrlsLcd.SetDigitCount(10)
	w.DataUrlsLcd.SetProperty("intValue", core.NewQVariant1(0))
	w.GridLayout_4.AddWidget3(w.DataUrlsLcd, 1, 2, 1, 1, 0)
	w.Label_7 = widgets.NewQLabel(w.DataGroup, 0)
	w.Label_7.SetObjectName("label_7")
	w.Label_7.SetAlignment(core.Qt__AlignCenter)
	w.GridLayout_4.AddWidget3(w.Label_7, 0, 1, 1, 1, 0)
	w.Label_12 = widgets.NewQLabel(w.DataGroup, 0)
	w.Label_12.SetObjectName("label_12")
	w.Label_12.SetAlignment(core.Qt__AlignCenter)
	w.GridLayout_4.AddWidget3(w.Label_12, 0, 3, 1, 1, 0)
	w.HorizontalLayout_13.QLayout.AddWidget(w.DataGroup)
	w.DumpStatsGroup = widgets.NewQGroupBox(w.Frame_2)
	w.DumpStatsGroup.SetObjectName("dumpStatsGroup")
	w.DumpStatsGroup.SetMinimumSize(core.NewQSize2(0, 108))
	w.DumpStatsGroup.SetAlignment(int(core.Qt__AlignCenter))
	w.GridLayout_3 = widgets.NewQGridLayout(w.DumpStatsGroup)
	w.GridLayout_3.SetObjectName("gridLayout_3")
	w.Label_13 = widgets.NewQLabel(w.DumpStatsGroup, 0)
	w.Label_13.SetObjectName("label_13")
	w.Label_13.SetAlignment(core.Qt__AlignCenter)
	w.GridLayout_3.AddWidget3(w.Label_13, 0, 0, 1, 1, 0)
	w.Label_15 = widgets.NewQLabel(w.DumpStatsGroup, 0)
	w.Label_15.SetObjectName("label_15")
	w.Label_15.SetAlignment(core.Qt__AlignCenter)
	w.GridLayout_3.AddWidget3(w.Label_15, 0, 1, 1, 1, 0)
	w.Label_14 = widgets.NewQLabel(w.DumpStatsGroup, 0)
	w.Label_14.SetObjectName("label_14")
	w.Label_14.SetAlignment(core.Qt__AlignCenter)
	w.GridLayout_3.AddWidget3(w.Label_14, 0, 2, 1, 1, 0)
	w.Label_30 = widgets.NewQLabel(w.DumpStatsGroup, 0)
	w.Label_30.SetObjectName("label_30")
	w.Label_30.SetAlignment(core.Qt__AlignCenter)
	w.GridLayout_3.AddWidget3(w.Label_30, 0, 3, 1, 1, 0)
	w.DumpStatsTablesLcd = widgets.NewQLCDNumber(w.DumpStatsGroup)
	w.DumpStatsTablesLcd.SetObjectName("dumpStatsTablesLcd")
	w.DumpStatsTablesLcd.SetDigitCount(10)
	w.DumpStatsTablesLcd.SetProperty("intValue", core.NewQVariant1(0))
	w.GridLayout_3.AddWidget3(w.DumpStatsTablesLcd, 1, 0, 1, 1, 0)
	w.DumpStatsColumnsLcd = widgets.NewQLCDNumber(w.DumpStatsGroup)
	w.DumpStatsColumnsLcd.SetObjectName("dumpStatsColumnsLcd")
	w.DumpStatsColumnsLcd.SetDigitCount(10)
	w.DumpStatsColumnsLcd.SetProperty("intValue", core.NewQVariant1(0))
	w.GridLayout_3.AddWidget3(w.DumpStatsColumnsLcd, 1, 1, 1, 1, 0)
	w.DumpStatsRowsLcd = widgets.NewQLCDNumber(w.DumpStatsGroup)
	w.DumpStatsRowsLcd.SetObjectName("dumpStatsRowsLcd")
	w.DumpStatsRowsLcd.SetDigitCount(10)
	w.DumpStatsRowsLcd.SetProperty("intValue", core.NewQVariant1(0))
	w.GridLayout_3.AddWidget3(w.DumpStatsRowsLcd, 1, 2, 1, 1, 0)
	w.DumpStatsRpm = widgets.NewQLCDNumber(w.DumpStatsGroup)
	w.DumpStatsRpm.SetObjectName("dumpStatsRpm")
	w.DumpStatsRpm.SetDigitCount(10)
	w.DumpStatsRpm.SetProperty("intValue", core.NewQVariant1(0))
	w.GridLayout_3.AddWidget3(w.DumpStatsRpm, 1, 3, 1, 1, 0)
	w.HorizontalLayout_13.QLayout.AddWidget(w.DumpStatsGroup)
	w.VerticalLayout_13.AddLayout(w.HorizontalLayout_13, 0)
	w.GridLayout_2.AddWidget3(w.Frame_2, 1, 0, 1, 1, 0)
	w.TabControl = widgets.NewQTabWidget(w.Centralwidget)
	w.TabControl.SetObjectName("tabControl")
	sizePolicy.SetHeightForWidth(w.TabControl.SizePolicy().HasHeightForWidth())
	w.TabControl.SetSizePolicy(sizePolicy)
	w.TabControl.SetTabPosition(widgets.QTabWidget__North)
	w.TabControl.SetTabShape(widgets.QTabWidget__Rounded)
	w.TabWelcome = widgets.NewQWidget(nil, 0)
	w.TabWelcome.SetObjectName("tabWelcome")
	w.GridLayout_5 = widgets.NewQGridLayout(w.TabWelcome)
	w.GridLayout_5.SetObjectName("gridLayout_5")
	w.GroupBox_4 = widgets.NewQGroupBox(w.TabWelcome)
	w.GroupBox_4.SetObjectName("groupBox_4")
	w.VerticalLayout_18 = widgets.NewQVBoxLayout2(w.GroupBox_4)
	w.VerticalLayout_18.SetObjectName("verticalLayout_18")
	w.Widget_10 = widgets.NewQWidget(w.GroupBox_4, 0)
	w.Widget_10.SetObjectName("widget_10")
	w.HorizontalLayout_14 = widgets.NewQHBoxLayout2(w.Widget_10)
	w.HorizontalLayout_14.SetObjectName("horizontalLayout_14")
	w.CustomerChatbox = widgets.NewQListWidget(w.Widget_10)
	w.CustomerChatbox.SetObjectName("CustomerChatbox")
	w.HorizontalLayout_14.QLayout.AddWidget(w.CustomerChatbox)
	w.CustomerList = widgets.NewQListWidget(w.Widget_10)
	w.CustomerList.SetObjectName("CustomerList")
	w.CustomerList.SetMaximumSize(core.NewQSize2(121, 16777215))
	w.HorizontalLayout_14.QLayout.AddWidget(w.CustomerList)
	w.VerticalLayout_18.QLayout.AddWidget(w.Widget_10)
	w.Widget_11 = widgets.NewQWidget(w.GroupBox_4, 0)
	w.Widget_11.SetObjectName("widget_11")
	w.HorizontalLayout_15 = widgets.NewQHBoxLayout2(w.Widget_11)
	w.HorizontalLayout_15.SetObjectName("horizontalLayout_15")
	w.ChatMessagebox = widgets.NewQLineEdit(w.Widget_11)
	w.ChatMessagebox.SetObjectName("ChatMessagebox")
	w.HorizontalLayout_15.QLayout.AddWidget(w.ChatMessagebox)
	w.ChatSendMessage = widgets.NewQPushButton(w.Widget_11)
	w.ChatSendMessage.SetObjectName("ChatSendMessage")
	w.HorizontalLayout_15.QLayout.AddWidget(w.ChatSendMessage)
	w.VerticalLayout_18.QLayout.AddWidget(w.Widget_11)
	w.GridLayout_5.AddWidget3(w.GroupBox_4, 0, 0, 1, 1, 0)
	w.GroupBox_3 = widgets.NewQGroupBox(w.TabWelcome)
	w.GroupBox_3.SetObjectName("groupBox_3")
	w.GroupBox_3.SetAlignment(int(core.Qt__AlignCenter))
	w.VerticalLayout_19 = widgets.NewQVBoxLayout2(w.GroupBox_3)
	w.VerticalLayout_19.SetObjectName("verticalLayout_19")
	w.NewsBox = widgets.NewQListWidget(w.GroupBox_3)
	w.NewsBox.SetObjectName("newsBox")
	w.VerticalLayout_19.QLayout.AddWidget(w.NewsBox)
	w.GridLayout_5.AddWidget3(w.GroupBox_3, 0, 1, 1, 1, 0)
	w.TabControl.AddTab(w.TabWelcome, "")
	w.TabGenerator = widgets.NewQWidget(nil, 0)
	w.TabGenerator.SetObjectName("tabGenerator")
	w.HorizontalLayout_19 = widgets.NewQHBoxLayout2(w.TabGenerator)
	w.HorizontalLayout_19.SetObjectName("horizontalLayout_19")
	w.Widget_16 = widgets.NewQWidget(w.TabGenerator, 0)
	w.Widget_16.SetObjectName("widget_16")
	w.VerticalLayout_28 = widgets.NewQVBoxLayout2(w.Widget_16)
	w.VerticalLayout_28.SetObjectName("verticalLayout_28")
	w.GroupParameters = widgets.NewQGroupBox(w.Widget_16)
	w.GroupParameters.SetObjectName("groupParameters")
	w.GroupParameters.SetAlignment(int(core.Qt__AlignCenter))
	w.VerticalLayout_27 = widgets.NewQVBoxLayout2(w.GroupParameters)
	w.VerticalLayout_27.SetObjectName("verticalLayout_27")
	w.ComboBoxParameters = widgets.NewQComboBox(w.GroupParameters)
	w.ComboBoxParameters.SetObjectName("comboBoxParameters")
	w.VerticalLayout_27.QLayout.AddWidget(w.ComboBoxParameters)
	w.ButtonParametersSettings = widgets.NewQPushButton(w.GroupParameters)
	w.ButtonParametersSettings.SetObjectName("buttonParametersSettings")
	w.VerticalLayout_27.QLayout.AddWidget(w.ButtonParametersSettings)
	w.PlainTextEditParameters = widgets.NewQPlainTextEdit(w.GroupParameters)
	w.PlainTextEditParameters.SetObjectName("plainTextEditParameters")
	w.VerticalLayout_27.QLayout.AddWidget(w.PlainTextEditParameters)
	w.GeneratorLimiterCheckbox = widgets.NewQCheckBox(w.GroupParameters)
	w.GeneratorLimiterCheckbox.SetObjectName("generatorLimiterCheckbox")
	w.VerticalLayout_27.QLayout.AddWidget(w.GeneratorLimiterCheckbox)
	w.GeneratorLimiterSpinbox = widgets.NewQSpinBox(w.GroupParameters)
	w.GeneratorLimiterSpinbox.SetObjectName("generatorLimiterSpinbox")
	w.GeneratorLimiterSpinbox.SetMaximum(999999999)
	w.GeneratorLimiterSpinbox.SetValue(5000)
	w.VerticalLayout_27.QLayout.AddWidget(w.GeneratorLimiterSpinbox)
	w.VerticalLayout_28.QLayout.AddWidget(w.GroupParameters)
	w.HorizontalLayout_19.QLayout.AddWidget(w.Widget_16)
	w.GroupGenerator = widgets.NewQGroupBox(w.TabGenerator)
	w.GroupGenerator.SetObjectName("groupGenerator")
	w.GroupGenerator.SetAlignment(int(core.Qt__AlignCenter))
	w.VerticalLayout_26 = widgets.NewQVBoxLayout2(w.GroupGenerator)
	w.VerticalLayout_26.SetObjectName("verticalLayout_26")
	w.TableGenerator = widgets.NewQTableWidget(w.GroupGenerator)
	if w.TableGenerator.ColumnCount() < 2 {
		w.TableGenerator.SetColumnCount(2)
	}
	__qtablewidgetitem := widgets.NewQTableWidgetItem(0)
	w.TableGenerator.SetHorizontalHeaderItem(0, __qtablewidgetitem)
	__qtablewidgetitem1 := widgets.NewQTableWidgetItem(0)
	w.TableGenerator.SetHorizontalHeaderItem(1, __qtablewidgetitem1)
	w.TableGenerator.SetObjectName("tableGenerator")
	w.VerticalLayout_26.QLayout.AddWidget(w.TableGenerator)
	w.GeneratorProgress = widgets.NewQProgressBar(w.GroupGenerator)
	w.GeneratorProgress.SetObjectName("generatorProgress")
	w.GeneratorProgress.SetValue(0)
	w.GeneratorProgress.SetInvertedAppearance(false)
	w.VerticalLayout_26.QLayout.AddWidget(w.GeneratorProgress)
	w.HorizontalLayout_19.QLayout.AddWidget(w.GroupGenerator)
	w.Frame_17 = widgets.NewQFrame(w.TabGenerator, 0)
	w.Frame_17.SetObjectName("frame_17")
	w.VerticalLayout_30 = widgets.NewQVBoxLayout2(w.Frame_17)
	w.VerticalLayout_30.SetObjectName("verticalLayout_30")
	w.GroupPatterns = widgets.NewQGroupBox(w.Frame_17)
	w.GroupPatterns.SetObjectName("groupPatterns")
	w.GroupPatterns.SetAlignment(int(core.Qt__AlignCenter))
	w.VerticalLayout_29 = widgets.NewQVBoxLayout2(w.GroupPatterns)
	w.VerticalLayout_29.SetObjectName("verticalLayout_29")
	w.PlainTextEditPatterns = widgets.NewQPlainTextEdit(w.GroupPatterns)
	w.PlainTextEditPatterns.SetObjectName("plainTextEditPatterns")
	w.VerticalLayout_29.QLayout.AddWidget(w.PlainTextEditPatterns)
	w.VerticalLayout_30.QLayout.AddWidget(w.GroupPatterns)
	w.Frame_6 = widgets.NewQFrame(w.Frame_17, 0)
	w.Frame_6.SetObjectName("frame_6")
	w.Frame_6.SetFrameShape(widgets.QFrame__StyledPanel)
	w.Frame_6.SetFrameShadow(widgets.QFrame__Raised)
	w.HorizontalLayout_18 = widgets.NewQHBoxLayout2(w.Frame_6)
	w.HorizontalLayout_18.SetObjectName("horizontalLayout_18")
	w.GeneratorStartBtn = widgets.NewQPushButton(w.Frame_6)
	w.GeneratorStartBtn.SetObjectName("generatorStartBtn")
	w.HorizontalLayout_18.QLayout.AddWidget(w.GeneratorStartBtn)
	w.GeneratorStopBtn = widgets.NewQPushButton(w.Frame_6)
	w.GeneratorStopBtn.SetObjectName("generatorStopBtn")
	w.HorizontalLayout_18.QLayout.AddWidget(w.GeneratorStopBtn)
	w.VerticalLayout_30.QLayout.AddWidget(w.Frame_6)
	w.HorizontalLayout_19.QLayout.AddWidget(w.Frame_17)
	w.TabControl.AddTab(w.TabGenerator, "")
	w.TabParser = widgets.NewQWidget(nil, 0)
	w.TabParser.SetObjectName("tabParser")
	w.HorizontalLayout_3 = widgets.NewQHBoxLayout2(w.TabParser)
	w.HorizontalLayout_3.SetObjectName("horizontalLayout_3")
	w.Widget = widgets.NewQWidget(w.TabParser, 0)
	w.Widget.SetObjectName("widget")
	w.VerticalLayout_4 = widgets.NewQVBoxLayout2(w.Widget)
	w.VerticalLayout_4.SetObjectName("verticalLayout_4")
	w.ParserSearchEnginesGroup = widgets.NewQGroupBox(w.Widget)
	w.ParserSearchEnginesGroup.SetObjectName("parserSearchEnginesGroup")
	w.ParserSearchEnginesGroup.SetAlignment(int(core.Qt__AlignCenter))
	w.VerticalLayout_3 = widgets.NewQVBoxLayout2(w.ParserSearchEnginesGroup)
	w.VerticalLayout_3.SetObjectName("verticalLayout_3")
	w.ParserGoogleCheckbox = widgets.NewQCheckBox(w.ParserSearchEnginesGroup)
	w.ParserGoogleCheckbox.SetObjectName("parserGoogleCheckbox")
	w.VerticalLayout_3.QLayout.AddWidget(w.ParserGoogleCheckbox)
	w.ParserBingCheckbox = widgets.NewQCheckBox(w.ParserSearchEnginesGroup)
	w.ParserBingCheckbox.SetObjectName("parserBingCheckbox")
	w.VerticalLayout_3.QLayout.AddWidget(w.ParserBingCheckbox)
	w.ParserAOLCheckbox = widgets.NewQCheckBox(w.ParserSearchEnginesGroup)
	w.ParserAOLCheckbox.SetObjectName("parserAOLCheckbox")
	w.VerticalLayout_3.QLayout.AddWidget(w.ParserAOLCheckbox)
	w.ParserMWSCheckbox = widgets.NewQCheckBox(w.ParserSearchEnginesGroup)
	w.ParserMWSCheckbox.SetObjectName("parserMWSCheckbox")
	w.VerticalLayout_3.QLayout.AddWidget(w.ParserMWSCheckbox)
	w.ParserDuckDuckGoCheckbox = widgets.NewQCheckBox(w.ParserSearchEnginesGroup)
	w.ParserDuckDuckGoCheckbox.SetObjectName("parserDuckDuckGoCheckbox")
	w.VerticalLayout_3.QLayout.AddWidget(w.ParserDuckDuckGoCheckbox)
	w.ParserEcosiaCheckbox = widgets.NewQCheckBox(w.ParserSearchEnginesGroup)
	w.ParserEcosiaCheckbox.SetObjectName("parserEcosiaCheckbox")
	w.VerticalLayout_3.QLayout.AddWidget(w.ParserEcosiaCheckbox)
	w.ParserStartPageCheckbox = widgets.NewQCheckBox(w.ParserSearchEnginesGroup)
	w.ParserStartPageCheckbox.SetObjectName("parserStartPageCheckbox")
	w.VerticalLayout_3.QLayout.AddWidget(w.ParserStartPageCheckbox)
	w.ParserYahooCheckbox = widgets.NewQCheckBox(w.ParserSearchEnginesGroup)
	w.ParserYahooCheckbox.SetObjectName("parserYahooCheckbox")
	w.VerticalLayout_3.QLayout.AddWidget(w.ParserYahooCheckbox)
	w.ParserYandexCheckbox = widgets.NewQCheckBox(w.ParserSearchEnginesGroup)
	w.ParserYandexCheckbox.SetObjectName("parserYandexCheckbox")
	w.VerticalLayout_3.QLayout.AddWidget(w.ParserYandexCheckbox)
	w.VerticalSpacer_7 = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout_3.AddItem(w.VerticalSpacer_7)
	w.VerticalLayout_4.QLayout.AddWidget(w.ParserSearchEnginesGroup)
	w.ParserSettingsGroup = widgets.NewQGroupBox(w.Widget)
	w.ParserSettingsGroup.SetObjectName("parserSettingsGroup")
	w.ParserSettingsGroup.SetAlignment(int(core.Qt__AlignCenter))
	w.VerticalLayout_2 = widgets.NewQVBoxLayout2(w.ParserSettingsGroup)
	w.VerticalLayout_2.SetObjectName("verticalLayout_2")
	w.ParserFilterUrls = widgets.NewQCheckBox(w.ParserSettingsGroup)
	w.ParserFilterUrls.SetObjectName("parserFilterUrls")
	w.VerticalLayout_2.QLayout.AddWidget(w.ParserFilterUrls)
	w.ParserPagesSpinbox = widgets.NewQSpinBox(w.ParserSettingsGroup)
	w.ParserPagesSpinbox.SetObjectName("parserPagesSpinbox")
	w.ParserPagesSpinbox.SetMinimum(1)
	w.ParserPagesSpinbox.SetMaximum(100)
	w.ParserPagesSpinbox.SetValue(2)
	w.VerticalLayout_2.QLayout.AddWidget(w.ParserPagesSpinbox)
	w.ParserCustomParams = widgets.NewQLineEdit(w.ParserSettingsGroup)
	w.ParserCustomParams.SetObjectName("parserCustomParams")
	w.ParserCustomParams.SetClearButtonEnabled(false)
	w.VerticalLayout_2.QLayout.AddWidget(w.ParserCustomParams)
	w.VerticalSpacer = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout_2.AddItem(w.VerticalSpacer)
	w.VerticalLayout_4.QLayout.AddWidget(w.ParserSettingsGroup)
	w.HorizontalLayout_3.AddWidget(w.Widget, 0, core.Qt__AlignLeft)
	w.ParserGroup = widgets.NewQGroupBox(w.TabParser)
	w.ParserGroup.SetObjectName("parserGroup")
	w.ParserGroup.SetAlignment(int(core.Qt__AlignCenter))
	w.VerticalLayout_12 = widgets.NewQVBoxLayout2(w.ParserGroup)
	w.VerticalLayout_12.SetObjectName("verticalLayout_12")
	w.ParserDorksTextbox = widgets.NewQPlainTextEdit(w.ParserGroup)
	w.ParserDorksTextbox.SetObjectName("parserDorksTextbox")
	w.VerticalLayout_12.QLayout.AddWidget(w.ParserDorksTextbox)
	w.ParserProgress = widgets.NewQProgressBar(w.ParserGroup)
	w.ParserProgress.SetObjectName("parserProgress")
	w.ParserProgress.SetAutoFillBackground(false)
	w.ParserProgress.SetValue(0)
	w.ParserProgress.SetAlignment(core.Qt__AlignRight | core.Qt__AlignTrailing | core.Qt__AlignVCenter)
	w.VerticalLayout_12.QLayout.AddWidget(w.ParserProgress)
	w.HorizontalLayout_3.QLayout.AddWidget(w.ParserGroup)
	w.Frame = widgets.NewQFrame(w.TabParser, 0)
	w.Frame.SetObjectName("frame")
	w.Frame.SetFrameShape(widgets.QFrame__StyledPanel)
	w.Frame.SetFrameShadow(widgets.QFrame__Raised)
	w.VerticalLayout = widgets.NewQVBoxLayout2(w.Frame)
	w.VerticalLayout.SetObjectName("verticalLayout")
	w.Widget_2 = widgets.NewQWidget(w.Frame, 0)
	w.Widget_2.SetObjectName("widget_2")
	w.HorizontalLayout_4 = widgets.NewQHBoxLayout2(w.Widget_2)
	w.HorizontalLayout_4.SetObjectName("horizontalLayout_4")
	w.ParserLoadDorksBtn = widgets.NewQPushButton(w.Widget_2)
	w.ParserLoadDorksBtn.SetObjectName("parserLoadDorksBtn")
	w.HorizontalLayout_4.QLayout.AddWidget(w.ParserLoadDorksBtn)
	w.ParserClearDorksBtn = widgets.NewQPushButton(w.Widget_2)
	w.ParserClearDorksBtn.SetObjectName("parserClearDorksBtn")
	w.HorizontalLayout_4.QLayout.AddWidget(w.ParserClearDorksBtn)
	w.VerticalLayout.AddWidget(w.Widget_2, 0, core.Qt__AlignHCenter|core.Qt__AlignTop)
	w.Widget_3 = widgets.NewQWidget(w.Frame, 0)
	w.Widget_3.SetObjectName("widget_3")
	w.HorizontalLayout_5 = widgets.NewQHBoxLayout2(w.Widget_3)
	w.HorizontalLayout_5.SetObjectName("horizontalLayout_5")
	w.ParserStartBtn = widgets.NewQPushButton(w.Widget_3)
	w.ParserStartBtn.SetObjectName("parserStartBtn")
	w.HorizontalLayout_5.QLayout.AddWidget(w.ParserStartBtn)
	w.ParserStopBtn = widgets.NewQPushButton(w.Widget_3)
	w.ParserStopBtn.SetObjectName("parserStopBtn")
	w.HorizontalLayout_5.QLayout.AddWidget(w.ParserStopBtn)
	w.VerticalLayout.AddWidget(w.Widget_3, 0, core.Qt__AlignTop)
	w.VerticalSpacer_2 = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout.AddItem(w.VerticalSpacer_2)
	w.HorizontalLayout_3.QLayout.AddWidget(w.Frame)
	w.TabControl.AddTab(w.TabParser, "")
	w.TabExploiter = widgets.NewQWidget(nil, 0)
	w.TabExploiter.SetObjectName("tabExploiter")
	w.HorizontalLayout_9 = widgets.NewQHBoxLayout2(w.TabExploiter)
	w.HorizontalLayout_9.SetObjectName("horizontalLayout_9")
	w.Widget_12 = widgets.NewQWidget(w.TabExploiter, 0)
	w.Widget_12.SetObjectName("widget_12")
	w.VerticalLayout_20 = widgets.NewQVBoxLayout2(w.Widget_12)
	w.VerticalLayout_20.SetObjectName("verticalLayout_20")
	w.ExploiterTechniquesGroup = widgets.NewQGroupBox(w.Widget_12)
	w.ExploiterTechniquesGroup.SetObjectName("exploiterTechniquesGroup")
	w.ExploiterTechniquesGroup.SetAlignment(int(core.Qt__AlignCenter))
	w.VerticalLayout_6 = widgets.NewQVBoxLayout2(w.ExploiterTechniquesGroup)
	w.VerticalLayout_6.SetObjectName("verticalLayout_6")
	w.ExploiterErrorCheckbox = widgets.NewQCheckBox(w.ExploiterTechniquesGroup)
	w.ExploiterErrorCheckbox.SetObjectName("exploiterErrorCheckbox")
	w.VerticalLayout_6.QLayout.AddWidget(w.ExploiterErrorCheckbox)
	w.ExploiterUnionCheckbox = widgets.NewQCheckBox(w.ExploiterTechniquesGroup)
	w.ExploiterUnionCheckbox.SetObjectName("exploiterUnionCheckbox")
	w.VerticalLayout_6.QLayout.AddWidget(w.ExploiterUnionCheckbox)
	w.ExploiterBlindCheckbox = widgets.NewQCheckBox(w.ExploiterTechniquesGroup)
	w.ExploiterBlindCheckbox.SetObjectName("exploiterBlindCheckbox")
	w.VerticalLayout_6.QLayout.AddWidget(w.ExploiterBlindCheckbox)
	w.ExploiterStackedCheckbox = widgets.NewQCheckBox(w.ExploiterTechniquesGroup)
	w.ExploiterStackedCheckbox.SetObjectName("exploiterStackedCheckbox")
	w.VerticalLayout_6.QLayout.AddWidget(w.ExploiterStackedCheckbox)
	w.Label = widgets.NewQLabel(w.ExploiterTechniquesGroup, 0)
	w.Label.SetObjectName("label")
	w.Label.SetAlignment(core.Qt__AlignCenter)
	w.VerticalLayout_6.QLayout.AddWidget(w.Label)
	w.ExploiterIntensityCombo = widgets.NewQComboBox(w.ExploiterTechniquesGroup)
	w.ExploiterIntensityCombo.AddItem("", core.NewQVariant1(0))
	w.ExploiterIntensityCombo.AddItem("", core.NewQVariant1(0))
	w.ExploiterIntensityCombo.AddItem("", core.NewQVariant1(0))
	w.ExploiterIntensityCombo.SetObjectName("exploiterIntensityCombo")
	w.VerticalLayout_6.QLayout.AddWidget(w.ExploiterIntensityCombo)
	w.ExploiterHeuristsicsCheckbox = widgets.NewQCheckBox(w.ExploiterTechniquesGroup)
	w.ExploiterHeuristsicsCheckbox.SetObjectName("exploiterHeuristsicsCheckbox")
	w.VerticalLayout_6.QLayout.AddWidget(w.ExploiterHeuristsicsCheckbox)
	w.VerticalLayout_20.QLayout.AddWidget(w.ExploiterTechniquesGroup)
	w.ExploiterDatabaseTypes = widgets.NewQGroupBox(w.Widget_12)
	w.ExploiterDatabaseTypes.SetObjectName("exploiterDatabaseTypes")
	w.ExploiterDatabaseTypes.SetAlignment(int(core.Qt__AlignCenter))
	w.VerticalLayout_7 = widgets.NewQVBoxLayout2(w.ExploiterDatabaseTypes)
	w.VerticalLayout_7.SetObjectName("verticalLayout_7")
	w.ExploiterMySQL = widgets.NewQCheckBox(w.ExploiterDatabaseTypes)
	w.ExploiterMySQL.SetObjectName("exploiterMySQL")
	w.VerticalLayout_7.QLayout.AddWidget(w.ExploiterMySQL)
	w.ExploiterOracle = widgets.NewQCheckBox(w.ExploiterDatabaseTypes)
	w.ExploiterOracle.SetObjectName("exploiterOracle")
	w.VerticalLayout_7.QLayout.AddWidget(w.ExploiterOracle)
	w.ExploiterPostgreSQL = widgets.NewQCheckBox(w.ExploiterDatabaseTypes)
	w.ExploiterPostgreSQL.SetObjectName("exploiterPostgreSQL")
	w.VerticalLayout_7.QLayout.AddWidget(w.ExploiterPostgreSQL)
	w.ExploiterMSSQL = widgets.NewQCheckBox(w.ExploiterDatabaseTypes)
	w.ExploiterMSSQL.SetObjectName("exploiterMSSQL")
	w.VerticalLayout_7.QLayout.AddWidget(w.ExploiterMSSQL)
	w.VerticalSpacer_10 = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout_7.AddItem(w.VerticalSpacer_10)
	w.VerticalLayout_20.QLayout.AddWidget(w.ExploiterDatabaseTypes)
	w.HorizontalLayout_9.QLayout.AddWidget(w.Widget_12)
	w.ExploiterGroup = widgets.NewQGroupBox(w.TabExploiter)
	w.ExploiterGroup.SetObjectName("exploiterGroup")
	w.ExploiterGroup.SetAlignment(int(core.Qt__AlignCenter))
	w.VerticalLayout_21 = widgets.NewQVBoxLayout2(w.ExploiterGroup)
	w.VerticalLayout_21.SetObjectName("verticalLayout_21")
	w.ExploiterInjectablesTable = widgets.NewQTableWidget(w.ExploiterGroup)
	if w.ExploiterInjectablesTable.ColumnCount() < 5 {
		w.ExploiterInjectablesTable.SetColumnCount(5)
	}
	__qtablewidgetitem2 := widgets.NewQTableWidgetItem(0)
	__qtablewidgetitem2.SetTextAlignment(int(core.Qt__AlignCenter))
	w.ExploiterInjectablesTable.SetHorizontalHeaderItem(0, __qtablewidgetitem2)
	__qtablewidgetitem3 := widgets.NewQTableWidgetItem(0)
	__qtablewidgetitem3.SetTextAlignment(int(core.Qt__AlignCenter))
	w.ExploiterInjectablesTable.SetHorizontalHeaderItem(1, __qtablewidgetitem3)
	__qtablewidgetitem4 := widgets.NewQTableWidgetItem(0)
	__qtablewidgetitem4.SetTextAlignment(int(core.Qt__AlignCenter))
	w.ExploiterInjectablesTable.SetHorizontalHeaderItem(2, __qtablewidgetitem4)
	__qtablewidgetitem5 := widgets.NewQTableWidgetItem(0)
	__qtablewidgetitem5.SetTextAlignment(int(core.Qt__AlignCenter))
	w.ExploiterInjectablesTable.SetHorizontalHeaderItem(3, __qtablewidgetitem5)
	__qtablewidgetitem6 := widgets.NewQTableWidgetItem(0)
	__qtablewidgetitem6.SetTextAlignment(int(core.Qt__AlignCenter))
	w.ExploiterInjectablesTable.SetHorizontalHeaderItem(4, __qtablewidgetitem6)
	w.ExploiterInjectablesTable.SetObjectName("exploiterInjectablesTable")
	w.ExploiterInjectablesTable.SetEditTriggers(widgets.QAbstractItemView__NoEditTriggers)
	w.VerticalLayout_21.QLayout.AddWidget(w.ExploiterInjectablesTable)
	w.ExploiterProgress = widgets.NewQProgressBar(w.ExploiterGroup)
	w.ExploiterProgress.SetObjectName("exploiterProgress")
	w.ExploiterProgress.SetAutoFillBackground(false)
	w.ExploiterProgress.SetValue(0)
	w.ExploiterProgress.SetAlignment(core.Qt__AlignRight | core.Qt__AlignTrailing | core.Qt__AlignVCenter)
	w.VerticalLayout_21.QLayout.AddWidget(w.ExploiterProgress)
	w.HorizontalLayout_9.QLayout.AddWidget(w.ExploiterGroup)
	w.Frame_3 = widgets.NewQFrame(w.TabExploiter, 0)
	w.Frame_3.SetObjectName("frame_3")
	w.Frame_3.SetFrameShape(widgets.QFrame__StyledPanel)
	w.Frame_3.SetFrameShadow(widgets.QFrame__Raised)
	w.VerticalLayout_31 = widgets.NewQVBoxLayout2(w.Frame_3)
	w.VerticalLayout_31.SetObjectName("verticalLayout_31")
	w.Widget_4 = widgets.NewQWidget(w.Frame_3, 0)
	w.Widget_4.SetObjectName("widget_4")
	w.HorizontalLayout_7 = widgets.NewQHBoxLayout2(w.Widget_4)
	w.HorizontalLayout_7.SetObjectName("horizontalLayout_7")
	w.ExploiterLoadUrlsBtn = widgets.NewQPushButton(w.Widget_4)
	w.ExploiterLoadUrlsBtn.SetObjectName("exploiterLoadUrlsBtn")
	w.HorizontalLayout_7.QLayout.AddWidget(w.ExploiterLoadUrlsBtn)
	w.ExploiterClearUrlsBtn = widgets.NewQPushButton(w.Widget_4)
	w.ExploiterClearUrlsBtn.SetObjectName("exploiterClearUrlsBtn")
	w.HorizontalLayout_7.QLayout.AddWidget(w.ExploiterClearUrlsBtn)
	w.VerticalLayout_31.QLayout.AddWidget(w.Widget_4)
	w.Widget_5 = widgets.NewQWidget(w.Frame_3, 0)
	w.Widget_5.SetObjectName("widget_5")
	w.HorizontalLayout_8 = widgets.NewQHBoxLayout2(w.Widget_5)
	w.HorizontalLayout_8.SetObjectName("horizontalLayout_8")
	w.ExploiterStartBtn = widgets.NewQPushButton(w.Widget_5)
	w.ExploiterStartBtn.SetObjectName("exploiterStartBtn")
	w.HorizontalLayout_8.QLayout.AddWidget(w.ExploiterStartBtn)
	w.ExploiterStopBtn = widgets.NewQPushButton(w.Widget_5)
	w.ExploiterStopBtn.SetObjectName("exploiterStopBtn")
	w.HorizontalLayout_8.QLayout.AddWidget(w.ExploiterStopBtn)
	w.VerticalLayout_31.QLayout.AddWidget(w.Widget_5)
	w.VerticalSpacer_3 = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout_31.AddItem(w.VerticalSpacer_3)
	w.GroupBox_5 = widgets.NewQGroupBox(w.Frame_3)
	w.GroupBox_5.SetObjectName("groupBox_5")
	w.VerticalLayout_5 = widgets.NewQVBoxLayout2(w.GroupBox_5)
	w.VerticalLayout_5.SetObjectName("verticalLayout_5")
	w.Label_19 = widgets.NewQLabel(w.GroupBox_5, 0)
	w.Label_19.SetObjectName("label_19")
	w.VerticalLayout_5.QLayout.AddWidget(w.Label_19)
	w.ExploiterThreads = widgets.NewQSpinBox(w.GroupBox_5)
	w.ExploiterThreads.SetObjectName("exploiterThreads")
	w.ExploiterThreads.SetMinimum(1)
	w.ExploiterThreads.SetMaximum(500)
	w.VerticalLayout_5.QLayout.AddWidget(w.ExploiterThreads)
	w.Label_20 = widgets.NewQLabel(w.GroupBox_5, 0)
	w.Label_20.SetObjectName("label_20")
	w.VerticalLayout_5.QLayout.AddWidget(w.Label_20)
	w.ExploiterWorkers = widgets.NewQSpinBox(w.GroupBox_5)
	w.ExploiterWorkers.SetObjectName("exploiterWorkers")
	w.ExploiterWorkers.SetMinimum(1)
	w.ExploiterWorkers.SetMaximum(100)
	w.VerticalLayout_5.QLayout.AddWidget(w.ExploiterWorkers)
	w.VerticalLayout_31.QLayout.AddWidget(w.GroupBox_5)
	w.HorizontalLayout_9.QLayout.AddWidget(w.Frame_3)
	w.TabControl.AddTab(w.TabExploiter, "")
	w.TabDumper = widgets.NewQWidget(nil, 0)
	w.TabDumper.SetObjectName("tabDumper")
	w.HorizontalLayout_12 = widgets.NewQHBoxLayout2(w.TabDumper)
	w.HorizontalLayout_12.SetObjectName("horizontalLayout_12")
	w.Widget_6 = widgets.NewQWidget(w.TabDumper, 0)
	w.Widget_6.SetObjectName("widget_6")
	w.VerticalLayout_10 = widgets.NewQVBoxLayout2(w.Widget_6)
	w.VerticalLayout_10.SetObjectName("verticalLayout_10")
	w.DumperMethodsGroup = widgets.NewQGroupBox(w.Widget_6)
	w.DumperMethodsGroup.SetObjectName("dumperMethodsGroup")
	w.DumperMethodsGroup.SetAlignment(int(core.Qt__AlignCenter))
	w.VerticalLayout_11 = widgets.NewQVBoxLayout2(w.DumperMethodsGroup)
	w.VerticalLayout_11.SetObjectName("verticalLayout_11")
	w.DumperTargetedCheckbox = widgets.NewQCheckBox(w.DumperMethodsGroup)
	w.DumperTargetedCheckbox.SetObjectName("dumperTargetedCheckbox")
	w.VerticalLayout_11.QLayout.AddWidget(w.DumperTargetedCheckbox)
	w.DumperTargetSettingsBtn = widgets.NewQPushButton(w.DumperMethodsGroup)
	w.DumperTargetSettingsBtn.SetObjectName("dumperTargetSettingsBtn")
	w.VerticalLayout_11.QLayout.AddWidget(w.DumperTargetSettingsBtn)
	w.DumperKeepBlanksCheckbox = widgets.NewQCheckBox(w.DumperMethodsGroup)
	w.DumperKeepBlanksCheckbox.SetObjectName("dumperKeepBlanksCheckbox")
	w.VerticalLayout_11.QLayout.AddWidget(w.DumperKeepBlanksCheckbox)
	w.DumperDIOSCheckbox = widgets.NewQCheckBox(w.DumperMethodsGroup)
	w.DumperDIOSCheckbox.SetObjectName("dumperDIOSCheckbox")
	w.VerticalLayout_11.QLayout.AddWidget(w.DumperDIOSCheckbox)
	w.Line_3 = widgets.NewQFrame(w.DumperMethodsGroup, 0)
	w.Line_3.SetObjectName("line_3")
	w.Line_3.SetFrameShape(widgets.QFrame__HLine)
	w.Line_3.SetFrameShadow(widgets.QFrame__Sunken)
	w.VerticalLayout_11.QLayout.AddWidget(w.Line_3)
	w.DumperMinRowsCheckbox = widgets.NewQCheckBox(w.DumperMethodsGroup)
	w.DumperMinRowsCheckbox.SetObjectName("dumperMinRowsCheckbox")
	w.VerticalLayout_11.QLayout.AddWidget(w.DumperMinRowsCheckbox)
	w.DumperMinRowsSpinbox = widgets.NewQSpinBox(w.DumperMethodsGroup)
	w.DumperMinRowsSpinbox.SetObjectName("dumperMinRowsSpinbox")
	w.DumperMinRowsSpinbox.SetButtonSymbols(widgets.QAbstractSpinBox__UpDownArrows)
	w.DumperMinRowsSpinbox.SetMinimum(1)
	w.DumperMinRowsSpinbox.SetMaximum(999999)
	w.DumperMinRowsSpinbox.SetSingleStep(0)
	w.DumperMinRowsSpinbox.SetValue(500)
	w.VerticalLayout_11.QLayout.AddWidget(w.DumperMinRowsSpinbox)
	w.Line_2 = widgets.NewQFrame(w.DumperMethodsGroup, 0)
	w.Line_2.SetObjectName("line_2")
	w.Line_2.SetFrameShape(widgets.QFrame__HLine)
	w.Line_2.SetFrameShadow(widgets.QFrame__Sunken)
	w.VerticalLayout_11.QLayout.AddWidget(w.Line_2)
	w.DumperAutoSkip = widgets.NewQCheckBox(w.DumperMethodsGroup)
	w.DumperAutoSkip.SetObjectName("dumperAutoSkip")
	w.VerticalLayout_11.QLayout.AddWidget(w.DumperAutoSkip)
	w.DumperAutoSkipSpinbox = widgets.NewQSpinBox(w.DumperMethodsGroup)
	w.DumperAutoSkipSpinbox.SetObjectName("dumperAutoSkipSpinbox")
	w.DumperAutoSkipSpinbox.SetMaximum(999999999)
	w.VerticalLayout_11.QLayout.AddWidget(w.DumperAutoSkipSpinbox)
	w.VerticalSpacer_6 = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout_11.AddItem(w.VerticalSpacer_6)
	w.VerticalLayout_10.QLayout.AddWidget(w.DumperMethodsGroup)
	w.HorizontalLayout_12.QLayout.AddWidget(w.Widget_6)
	w.DumperGroup = widgets.NewQGroupBox(w.TabDumper)
	w.DumperGroup.SetObjectName("dumperGroup")
	w.DumperGroup.SetAlignment(int(core.Qt__AlignCenter))
	w.VerticalLayout_9 = widgets.NewQVBoxLayout2(w.DumperGroup)
	w.VerticalLayout_9.SetObjectName("verticalLayout_9")
	w.DumperTableWidget = widgets.NewQTableWidget(w.DumperGroup)
	if w.DumperTableWidget.ColumnCount() < 7 {
		w.DumperTableWidget.SetColumnCount(7)
	}
	__qtablewidgetitem7 := widgets.NewQTableWidgetItem(0)
	__qtablewidgetitem7.SetTextAlignment(int(core.Qt__AlignCenter))
	w.DumperTableWidget.SetHorizontalHeaderItem(0, __qtablewidgetitem7)
	__qtablewidgetitem8 := widgets.NewQTableWidgetItem(0)
	__qtablewidgetitem8.SetTextAlignment(int(core.Qt__AlignCenter))
	w.DumperTableWidget.SetHorizontalHeaderItem(1, __qtablewidgetitem8)
	__qtablewidgetitem9 := widgets.NewQTableWidgetItem(0)
	__qtablewidgetitem9.SetTextAlignment(int(core.Qt__AlignCenter))
	w.DumperTableWidget.SetHorizontalHeaderItem(2, __qtablewidgetitem9)
	__qtablewidgetitem10 := widgets.NewQTableWidgetItem(0)
	__qtablewidgetitem10.SetTextAlignment(int(core.Qt__AlignCenter))
	w.DumperTableWidget.SetHorizontalHeaderItem(3, __qtablewidgetitem10)
	__qtablewidgetitem11 := widgets.NewQTableWidgetItem(0)
	__qtablewidgetitem11.SetTextAlignment(int(core.Qt__AlignCenter))
	w.DumperTableWidget.SetHorizontalHeaderItem(4, __qtablewidgetitem11)
	__qtablewidgetitem12 := widgets.NewQTableWidgetItem(0)
	__qtablewidgetitem12.SetTextAlignment(int(core.Qt__AlignCenter))
	w.DumperTableWidget.SetHorizontalHeaderItem(5, __qtablewidgetitem12)
	__qtablewidgetitem13 := widgets.NewQTableWidgetItem(0)
	__qtablewidgetitem13.SetTextAlignment(int(core.Qt__AlignCenter))
	w.DumperTableWidget.SetHorizontalHeaderItem(6, __qtablewidgetitem13)
	w.DumperTableWidget.SetObjectName("dumperTableWidget")
	w.DumperTableWidget.SetMinimumSize(core.NewQSize2(754, 0))
	w.DumperTableWidget.SetContextMenuPolicy(core.Qt__CustomContextMenu)
	w.VerticalLayout_9.QLayout.AddWidget(w.DumperTableWidget)
	w.DumperProgress = widgets.NewQProgressBar(w.DumperGroup)
	w.DumperProgress.SetObjectName("dumperProgress")
	w.DumperProgress.SetAutoFillBackground(false)
	w.DumperProgress.SetValue(0)
	w.DumperProgress.SetAlignment(core.Qt__AlignRight | core.Qt__AlignTrailing | core.Qt__AlignVCenter)
	w.VerticalLayout_9.QLayout.AddWidget(w.DumperProgress)
	w.HorizontalLayout_12.QLayout.AddWidget(w.DumperGroup)
	w.Frame_4 = widgets.NewQFrame(w.TabDumper, 0)
	w.Frame_4.SetObjectName("frame_4")
	w.Frame_4.SetFrameShape(widgets.QFrame__StyledPanel)
	w.Frame_4.SetFrameShadow(widgets.QFrame__Raised)
	w.VerticalLayout_32 = widgets.NewQVBoxLayout2(w.Frame_4)
	w.VerticalLayout_32.SetObjectName("verticalLayout_32")
	w.Widget_7 = widgets.NewQWidget(w.Frame_4, 0)
	w.Widget_7.SetObjectName("widget_7")
	w.HorizontalLayout_10 = widgets.NewQHBoxLayout2(w.Widget_7)
	w.HorizontalLayout_10.SetObjectName("horizontalLayout_10")
	w.DumperLoadInjectablesBtn = widgets.NewQPushButton(w.Widget_7)
	w.DumperLoadInjectablesBtn.SetObjectName("dumperLoadInjectablesBtn")
	w.HorizontalLayout_10.QLayout.AddWidget(w.DumperLoadInjectablesBtn)
	w.DumperClearInjectablesBtn = widgets.NewQPushButton(w.Widget_7)
	w.DumperClearInjectablesBtn.SetObjectName("dumperClearInjectablesBtn")
	w.HorizontalLayout_10.QLayout.AddWidget(w.DumperClearInjectablesBtn)
	w.VerticalLayout_32.QLayout.AddWidget(w.Widget_7)
	w.Widget_8 = widgets.NewQWidget(w.Frame_4, 0)
	w.Widget_8.SetObjectName("widget_8")
	w.HorizontalLayout_11 = widgets.NewQHBoxLayout2(w.Widget_8)
	w.HorizontalLayout_11.SetObjectName("horizontalLayout_11")
	w.DumperStartBtn = widgets.NewQPushButton(w.Widget_8)
	w.DumperStartBtn.SetObjectName("dumperStartBtn")
	w.HorizontalLayout_11.QLayout.AddWidget(w.DumperStartBtn)
	w.DumperStopBtn = widgets.NewQPushButton(w.Widget_8)
	w.DumperStopBtn.SetObjectName("dumperStopBtn")
	w.HorizontalLayout_11.QLayout.AddWidget(w.DumperStopBtn)
	w.VerticalLayout_32.QLayout.AddWidget(w.Widget_8)
	w.DumperOpenAnalyzer = widgets.NewQPushButton(w.Frame_4)
	w.DumperOpenAnalyzer.SetObjectName("dumperOpenAnalyzer")
	w.VerticalLayout_32.QLayout.AddWidget(w.DumperOpenAnalyzer)
	w.VerticalSpacer_5 = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout_32.AddItem(w.VerticalSpacer_5)
	w.GroupBox_6 = widgets.NewQGroupBox(w.Frame_4)
	w.GroupBox_6.SetObjectName("groupBox_6")
	w.VerticalLayout_8 = widgets.NewQVBoxLayout2(w.GroupBox_6)
	w.VerticalLayout_8.SetObjectName("verticalLayout_8")
	w.Label_21 = widgets.NewQLabel(w.GroupBox_6, 0)
	w.Label_21.SetObjectName("label_21")
	w.VerticalLayout_8.QLayout.AddWidget(w.Label_21)
	w.DumperThreads = widgets.NewQSpinBox(w.GroupBox_6)
	w.DumperThreads.SetObjectName("dumperThreads")
	w.DumperThreads.SetMinimum(1)
	w.DumperThreads.SetMaximum(500)
	w.VerticalLayout_8.QLayout.AddWidget(w.DumperThreads)
	w.Label_22 = widgets.NewQLabel(w.GroupBox_6, 0)
	w.Label_22.SetObjectName("label_22")
	w.VerticalLayout_8.QLayout.AddWidget(w.Label_22)
	w.DumperWorkers = widgets.NewQSpinBox(w.GroupBox_6)
	w.DumperWorkers.SetObjectName("dumperWorkers")
	w.DumperWorkers.SetMinimum(1)
	w.DumperWorkers.SetMaximum(100)
	w.VerticalLayout_8.QLayout.AddWidget(w.DumperWorkers)
	w.VerticalLayout_32.QLayout.AddWidget(w.GroupBox_6)
	w.HorizontalLayout_12.QLayout.AddWidget(w.Frame_4)
	w.TabControl.AddTab(w.TabDumper, "")
	w.TabAutoSheller = widgets.NewQWidget(nil, 0)
	w.TabAutoSheller.SetObjectName("tabAutoSheller")
	w.HorizontalLayout_6 = widgets.NewQHBoxLayout2(w.TabAutoSheller)
	w.HorizontalLayout_6.SetObjectName("horizontalLayout_6")
	w.Widget_15 = widgets.NewQWidget(w.TabAutoSheller, 0)
	w.Widget_15.SetObjectName("widget_15")
	w.VerticalLayout_24 = widgets.NewQVBoxLayout2(w.Widget_15)
	w.VerticalLayout_24.SetObjectName("verticalLayout_24")
	w.AutoShellTypesGroup = widgets.NewQGroupBox(w.Widget_15)
	w.AutoShellTypesGroup.SetObjectName("autoShellTypesGroup")
	w.AutoShellTypesGroup.SetAlignment(int(core.Qt__AlignCenter))
	w.VerticalLayout_22 = widgets.NewQVBoxLayout2(w.AutoShellTypesGroup)
	w.VerticalLayout_22.SetObjectName("verticalLayout_22")
	w.AutoShellASPCheckbox = widgets.NewQCheckBox(w.AutoShellTypesGroup)
	w.AutoShellASPCheckbox.SetObjectName("autoShellASPCheckbox")
	w.VerticalLayout_22.QLayout.AddWidget(w.AutoShellASPCheckbox)
	w.AutoShellPHPCheckbox = widgets.NewQCheckBox(w.AutoShellTypesGroup)
	w.AutoShellPHPCheckbox.SetObjectName("autoShellPHPCheckbox")
	w.VerticalLayout_22.QLayout.AddWidget(w.AutoShellPHPCheckbox)
	w.Line = widgets.NewQFrame(w.AutoShellTypesGroup, 0)
	w.Line.SetObjectName("line")
	w.Line.SetFrameShape(widgets.QFrame__HLine)
	w.Line.SetFrameShadow(widgets.QFrame__Sunken)
	w.VerticalLayout_22.QLayout.AddWidget(w.Line)
	w.Label_17 = widgets.NewQLabel(w.AutoShellTypesGroup, 0)
	w.Label_17.SetObjectName("label_17")
	w.VerticalLayout_22.QLayout.AddWidget(w.Label_17)
	w.AutoShellKey = widgets.NewQLineEdit(w.AutoShellTypesGroup)
	w.AutoShellKey.SetObjectName("autoShellKey")
	sizePolicy1 := widgets.NewQSizePolicy2(widgets.QSizePolicy__Fixed, widgets.QSizePolicy__Fixed, 0)
	sizePolicy1.SetHorizontalStretch(0)
	sizePolicy1.SetVerticalStretch(0)
	sizePolicy1.SetHeightForWidth(w.AutoShellKey.SizePolicy().HasHeightForWidth())
	w.AutoShellKey.SetSizePolicy(sizePolicy1)
	w.VerticalLayout_22.QLayout.AddWidget(w.AutoShellKey)
	w.Label_18 = widgets.NewQLabel(w.AutoShellTypesGroup, 0)
	w.Label_18.SetObjectName("label_18")
	w.VerticalLayout_22.QLayout.AddWidget(w.Label_18)
	w.AutoShellFile = widgets.NewQLineEdit(w.AutoShellTypesGroup)
	w.AutoShellFile.SetObjectName("autoShellFile")
	sizePolicy1.SetHeightForWidth(w.AutoShellFile.SizePolicy().HasHeightForWidth())
	w.AutoShellFile.SetSizePolicy(sizePolicy1)
	w.VerticalLayout_22.QLayout.AddWidget(w.AutoShellFile)
	w.VerticalSpacer_11 = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout_22.AddItem(w.VerticalSpacer_11)
	w.VerticalLayout_24.QLayout.AddWidget(w.AutoShellTypesGroup)
	w.HorizontalLayout_6.QLayout.AddWidget(w.Widget_15)
	w.GroupBox_7 = widgets.NewQGroupBox(w.TabAutoSheller)
	w.GroupBox_7.SetObjectName("groupBox_7")
	w.GroupBox_7.SetAlignment(int(core.Qt__AlignCenter))
	w.VerticalLayout_25 = widgets.NewQVBoxLayout2(w.GroupBox_7)
	w.VerticalLayout_25.SetObjectName("verticalLayout_25")
	w.AutoShellTable = widgets.NewQTableWidget(w.GroupBox_7)
	if w.AutoShellTable.ColumnCount() < 4 {
		w.AutoShellTable.SetColumnCount(4)
	}
	__qtablewidgetitem14 := widgets.NewQTableWidgetItem(0)
	__qtablewidgetitem14.SetTextAlignment(int(core.Qt__AlignCenter))
	w.AutoShellTable.SetHorizontalHeaderItem(0, __qtablewidgetitem14)
	__qtablewidgetitem15 := widgets.NewQTableWidgetItem(0)
	__qtablewidgetitem15.SetTextAlignment(int(core.Qt__AlignCenter))
	w.AutoShellTable.SetHorizontalHeaderItem(1, __qtablewidgetitem15)
	__qtablewidgetitem16 := widgets.NewQTableWidgetItem(0)
	__qtablewidgetitem16.SetTextAlignment(int(core.Qt__AlignCenter))
	w.AutoShellTable.SetHorizontalHeaderItem(2, __qtablewidgetitem16)
	__qtablewidgetitem17 := widgets.NewQTableWidgetItem(0)
	__qtablewidgetitem17.SetTextAlignment(int(core.Qt__AlignCenter))
	w.AutoShellTable.SetHorizontalHeaderItem(3, __qtablewidgetitem17)
	w.AutoShellTable.SetObjectName("autoShellTable")
	w.VerticalLayout_25.QLayout.AddWidget(w.AutoShellTable)
	w.AutoShellProgress = widgets.NewQProgressBar(w.GroupBox_7)
	w.AutoShellProgress.SetObjectName("autoShellProgress")
	w.AutoShellProgress.SetValue(24)
	w.VerticalLayout_25.QLayout.AddWidget(w.AutoShellProgress)
	w.HorizontalLayout_6.QLayout.AddWidget(w.GroupBox_7)
	w.Frame_61 = widgets.NewQFrame(w.TabAutoSheller, 0)
	w.Frame_61.SetObjectName("frame_61")
	w.Frame_61.SetFrameShape(widgets.QFrame__StyledPanel)
	w.Frame_61.SetFrameShadow(widgets.QFrame__Raised)
	w.VerticalLayout_23 = widgets.NewQVBoxLayout2(w.Frame_61)
	w.VerticalLayout_23.SetObjectName("verticalLayout_23")
	w.Widget_13 = widgets.NewQWidget(w.Frame_61, 0)
	w.Widget_13.SetObjectName("widget_13")
	w.HorizontalLayout_16 = widgets.NewQHBoxLayout2(w.Widget_13)
	w.HorizontalLayout_16.SetObjectName("horizontalLayout_16")
	w.AutoShellLoadInjectablesBtn = widgets.NewQPushButton(w.Widget_13)
	w.AutoShellLoadInjectablesBtn.SetObjectName("autoShellLoadInjectablesBtn")
	w.HorizontalLayout_16.QLayout.AddWidget(w.AutoShellLoadInjectablesBtn)
	w.AutoShellClearInjectablesBtn = widgets.NewQPushButton(w.Widget_13)
	w.AutoShellClearInjectablesBtn.SetObjectName("autoShellClearInjectablesBtn")
	w.HorizontalLayout_16.QLayout.AddWidget(w.AutoShellClearInjectablesBtn)
	w.VerticalLayout_23.AddWidget(w.Widget_13, 0, core.Qt__AlignHCenter|core.Qt__AlignTop)
	w.Widget_14 = widgets.NewQWidget(w.Frame_61, 0)
	w.Widget_14.SetObjectName("widget_14")
	w.HorizontalLayout_17 = widgets.NewQHBoxLayout2(w.Widget_14)
	w.HorizontalLayout_17.SetObjectName("horizontalLayout_17")
	w.AutoShellStartBtn = widgets.NewQPushButton(w.Widget_14)
	w.AutoShellStartBtn.SetObjectName("autoShellStartBtn")
	w.HorizontalLayout_17.QLayout.AddWidget(w.AutoShellStartBtn)
	w.AutoShellStopBtn = widgets.NewQPushButton(w.Widget_14)
	w.AutoShellStopBtn.SetObjectName("autoShellStopBtn")
	w.HorizontalLayout_17.QLayout.AddWidget(w.AutoShellStopBtn)
	w.VerticalLayout_23.AddWidget(w.Widget_14, 0, core.Qt__AlignTop)
	w.VerticalSpacer_12 = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout_23.AddItem(w.VerticalSpacer_12)
	w.HorizontalLayout_6.QLayout.AddWidget(w.Frame_61)
	w.TabControl.AddTab(w.TabAutoSheller, "")
	w.TabUtils = widgets.NewQWidget(nil, 0)
	w.TabUtils.SetObjectName("tabUtils")
	w.HorizontalLayout_2 = widgets.NewQHBoxLayout2(w.TabUtils)
	w.HorizontalLayout_2.SetObjectName("horizontalLayout_2")
	w.GroupBox = widgets.NewQGroupBox(w.TabUtils)
	w.GroupBox.SetObjectName("groupBox")
	w.GroupBox.SetAlignment(int(core.Qt__AlignCenter))
	w.HorizontalLayout = widgets.NewQHBoxLayout2(w.GroupBox)
	w.HorizontalLayout.SetObjectName("horizontalLayout")
	w.CleanerUrlsTextBox = widgets.NewQPlainTextEdit(w.GroupBox)
	w.CleanerUrlsTextBox.SetObjectName("cleanerUrlsTextBox")
	w.HorizontalLayout.QLayout.AddWidget(w.CleanerUrlsTextBox)
	w.Frame_5 = widgets.NewQFrame(w.GroupBox, 0)
	w.Frame_5.SetObjectName("frame_5")
	w.Frame_5.SetFrameShape(widgets.QFrame__StyledPanel)
	w.Frame_5.SetFrameShadow(widgets.QFrame__Raised)
	w.VerticalLayout_14 = widgets.NewQVBoxLayout2(w.Frame_5)
	w.VerticalLayout_14.SetObjectName("verticalLayout_14")
	w.CleanerDuplicateDomains = widgets.NewQCheckBox(w.Frame_5)
	w.CleanerDuplicateDomains.SetObjectName("cleanerDuplicateDomains")
	w.VerticalLayout_14.QLayout.AddWidget(w.CleanerDuplicateDomains)
	w.CleanerQueryParam = widgets.NewQCheckBox(w.Frame_5)
	w.CleanerQueryParam.SetObjectName("cleanerQueryParam")
	w.VerticalLayout_14.QLayout.AddWidget(w.CleanerQueryParam)
	w.CleanerBtn = widgets.NewQPushButton(w.Frame_5)
	w.CleanerBtn.SetObjectName("cleanerBtn")
	w.VerticalLayout_14.QLayout.AddWidget(w.CleanerBtn)
	w.VerticalSpacer_8 = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout_14.AddItem(w.VerticalSpacer_8)
	w.HorizontalLayout.QLayout.AddWidget(w.Frame_5)
	w.HorizontalLayout_2.QLayout.AddWidget(w.GroupBox)
	w.GroupBox_2 = widgets.NewQGroupBox(w.TabUtils)
	w.GroupBox_2.SetObjectName("groupBox_2")
	w.GroupBox_2.SetAlignment(int(core.Qt__AlignCenter))
	w.VerticalLayout_17 = widgets.NewQVBoxLayout2(w.GroupBox_2)
	w.VerticalLayout_17.SetObjectName("verticalLayout_17")
	w.Widget_9 = widgets.NewQWidget(w.GroupBox_2, 0)
	w.Widget_9.SetObjectName("widget_9")
	w.VerticalLayout_16 = widgets.NewQVBoxLayout2(w.Widget_9)
	w.VerticalLayout_16.SetObjectName("verticalLayout_16")
	w.Label_11 = widgets.NewQLabel(w.Widget_9, 0)
	w.Label_11.SetObjectName("label_11")
	w.VerticalLayout_16.QLayout.AddWidget(w.Label_11)
	w.WebUIAddress = widgets.NewQLineEdit(w.Widget_9)
	w.WebUIAddress.SetObjectName("webUIAddress")
	w.VerticalLayout_16.QLayout.AddWidget(w.WebUIAddress)
	w.WebUILaunchBtn = widgets.NewQPushButton(w.Widget_9)
	w.WebUILaunchBtn.SetObjectName("webUILaunchBtn")
	w.VerticalLayout_16.QLayout.AddWidget(w.WebUILaunchBtn)
	w.VerticalLayout_17.QLayout.AddWidget(w.Widget_9)
	w.VerticalSpacer_9 = widgets.NewQSpacerItem(20, 342, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout_17.AddItem(w.VerticalSpacer_9)
	w.HorizontalLayout_2.QLayout.AddWidget(w.GroupBox_2)
	w.TabControl.AddTab(w.TabUtils, "")
	w.Tab = widgets.NewQWidget(nil, 0)
	w.Tab.SetObjectName("tab")
	w.HorizontalLayout_25 = widgets.NewQHBoxLayout2(w.Tab)
	w.HorizontalLayout_25.SetObjectName("horizontalLayout_25")
	w.AutoShellTypesGroup_2 = widgets.NewQGroupBox(w.Tab)
	w.AutoShellTypesGroup_2.SetObjectName("autoShellTypesGroup_2")
	w.AutoShellTypesGroup_2.SetAlignment(int(core.Qt__AlignCenter))
	w.VerticalLayout_33 = widgets.NewQVBoxLayout2(w.AutoShellTypesGroup_2)
	w.VerticalLayout_33.SetObjectName("verticalLayout_33")
	w.AntipubSavePublic = widgets.NewQCheckBox(w.AutoShellTypesGroup_2)
	w.AntipubSavePublic.SetObjectName("antipubSavePublic")
	w.VerticalLayout_33.QLayout.AddWidget(w.AntipubSavePublic)
	w.AntipubAutoSync = widgets.NewQCheckBox(w.AutoShellTypesGroup_2)
	w.AntipubAutoSync.SetObjectName("antipubAutoSync")
	w.VerticalLayout_33.QLayout.AddWidget(w.AntipubAutoSync)
	w.VerticalSpacer_13 = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout_33.AddItem(w.VerticalSpacer_13)
	w.HorizontalLayout_25.QLayout.AddWidget(w.AutoShellTypesGroup_2)
	w.Widget_22 = widgets.NewQWidget(w.Tab, 0)
	w.Widget_22.SetObjectName("widget_22")
	w.VerticalLayout_39 = widgets.NewQVBoxLayout2(w.Widget_22)
	w.VerticalLayout_39.SetObjectName("verticalLayout_39")
	w.Widget_17 = widgets.NewQWidget(w.Widget_22, 0)
	w.Widget_17.SetObjectName("widget_17")
	w.HorizontalLayout_21 = widgets.NewQHBoxLayout2(w.Widget_17)
	w.HorizontalLayout_21.SetObjectName("horizontalLayout_21")
	w.HorizontalSpacer = widgets.NewQSpacerItem(276, 20, widgets.QSizePolicy__MinimumExpanding, widgets.QSizePolicy__Minimum)
	w.HorizontalLayout_21.AddItem(w.HorizontalSpacer)
	w.GroupBox_8 = widgets.NewQGroupBox(w.Widget_17)
	w.GroupBox_8.SetObjectName("groupBox_8")
	w.GroupBox_8.SetAlignment(int(core.Qt__AlignCenter))
	w.VerticalLayout_37 = widgets.NewQVBoxLayout2(w.GroupBox_8)
	w.VerticalLayout_37.SetObjectName("verticalLayout_37")
	w.Frame_7 = widgets.NewQFrame(w.GroupBox_8, 0)
	w.Frame_7.SetObjectName("frame_7")
	w.Frame_7.SetFrameShape(widgets.QFrame__StyledPanel)
	w.Frame_7.SetFrameShadow(widgets.QFrame__Raised)
	w.VerticalLayout_34 = widgets.NewQVBoxLayout2(w.Frame_7)
	w.VerticalLayout_34.SetObjectName("verticalLayout_34")
	w.Label_24 = widgets.NewQLabel(w.Frame_7, 0)
	w.Label_24.SetObjectName("label_24")
	w.VerticalLayout_34.QLayout.AddWidget(w.Label_24)
	w.AntipubLinkCount = widgets.NewQLCDNumber(w.Frame_7)
	w.AntipubLinkCount.SetObjectName("antipubLinkCount")
	w.AntipubLinkCount.SetDigitCount(20)
	w.VerticalLayout_34.QLayout.AddWidget(w.AntipubLinkCount)
	w.VerticalLayout_37.QLayout.AddWidget(w.Frame_7)
	w.Frame_8 = widgets.NewQFrame(w.GroupBox_8, 0)
	w.Frame_8.SetObjectName("frame_8")
	w.Frame_8.SetFrameShape(widgets.QFrame__StyledPanel)
	w.Frame_8.SetFrameShadow(widgets.QFrame__Raised)
	w.VerticalLayout_35 = widgets.NewQVBoxLayout2(w.Frame_8)
	w.VerticalLayout_35.SetObjectName("verticalLayout_35")
	w.Label_25 = widgets.NewQLabel(w.Frame_8, 0)
	w.Label_25.SetObjectName("label_25")
	w.VerticalLayout_35.QLayout.AddWidget(w.Label_25)
	w.AntipubDomainCount = widgets.NewQLCDNumber(w.Frame_8)
	w.AntipubDomainCount.SetObjectName("antipubDomainCount")
	w.AntipubDomainCount.SetDigitCount(20)
	w.VerticalLayout_35.QLayout.AddWidget(w.AntipubDomainCount)
	w.VerticalLayout_37.QLayout.AddWidget(w.Frame_8)
	w.Frame_9 = widgets.NewQFrame(w.GroupBox_8, 0)
	w.Frame_9.SetObjectName("frame_9")
	w.Frame_9.SetFrameShape(widgets.QFrame__StyledPanel)
	w.Frame_9.SetFrameShadow(widgets.QFrame__Raised)
	w.VerticalLayout_36 = widgets.NewQVBoxLayout2(w.Frame_9)
	w.VerticalLayout_36.SetObjectName("verticalLayout_36")
	w.AntipubSizeOnDisk = widgets.NewQLabel(w.Frame_9, 0)
	w.AntipubSizeOnDisk.SetObjectName("antipubSizeOnDisk")
	w.VerticalLayout_36.QLayout.AddWidget(w.AntipubSizeOnDisk)
	w.VerticalLayout_37.QLayout.AddWidget(w.Frame_9)
	w.Widget_18 = widgets.NewQWidget(w.GroupBox_8, 0)
	w.Widget_18.SetObjectName("widget_18")
	w.HorizontalLayout_20 = widgets.NewQHBoxLayout2(w.Widget_18)
	w.HorizontalLayout_20.SetObjectName("horizontalLayout_20")
	w.AntipubLinkMode = widgets.NewQRadioButton(w.Widget_18)
	w.AntipubLinkMode.SetObjectName("antipubLinkMode")
	w.AntipubLinkMode.SetChecked(true)
	w.HorizontalLayout_20.QLayout.AddWidget(w.AntipubLinkMode)
	w.AntipubDomainMode = widgets.NewQRadioButton(w.Widget_18)
	w.AntipubDomainMode.SetObjectName("antipubDomainMode")
	w.AntipubDomainMode.SetChecked(false)
	w.HorizontalLayout_20.QLayout.AddWidget(w.AntipubDomainMode)
	w.VerticalLayout_37.QLayout.AddWidget(w.Widget_18)
	w.HorizontalLayout_21.QLayout.AddWidget(w.GroupBox_8)
	w.HorizontalSpacer_2 = widgets.NewQSpacerItem(276, 20, widgets.QSizePolicy__MinimumExpanding, widgets.QSizePolicy__Minimum)
	w.HorizontalLayout_21.AddItem(w.HorizontalSpacer_2)
	w.VerticalLayout_39.QLayout.AddWidget(w.Widget_17)
	w.VerticalSpacer_4 = widgets.NewQSpacerItem(20, 94, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout_39.AddItem(w.VerticalSpacer_4)
	w.AntipubProgress = widgets.NewQProgressBar(w.Widget_22)
	w.AntipubProgress.SetObjectName("antipubProgress")
	w.AntipubProgress.SetValue(24)
	w.VerticalLayout_39.QLayout.AddWidget(w.AntipubProgress)
	w.Widget_23 = widgets.NewQWidget(w.Widget_22, 0)
	w.Widget_23.SetObjectName("widget_23")
	w.GridLayout_6 = widgets.NewQGridLayout(w.Widget_23)
	w.GridLayout_6.SetObjectName("gridLayout_6")
	w.Label_26 = widgets.NewQLabel(w.Widget_23, 0)
	w.Label_26.SetObjectName("label_26")
	w.Label_26.SetAlignment(core.Qt__AlignCenter)
	w.GridLayout_6.AddWidget3(w.Label_26, 0, 0, 1, 1, 0)
	w.Label_27 = widgets.NewQLabel(w.Widget_23, 0)
	w.Label_27.SetObjectName("label_27")
	w.Label_27.SetAlignment(core.Qt__AlignCenter)
	w.GridLayout_6.AddWidget3(w.Label_27, 0, 1, 1, 1, 0)
	w.Label_28 = widgets.NewQLabel(w.Widget_23, 0)
	w.Label_28.SetObjectName("label_28")
	w.Label_28.SetAlignment(core.Qt__AlignCenter)
	w.GridLayout_6.AddWidget3(w.Label_28, 0, 2, 1, 1, 0)
	w.Label_29 = widgets.NewQLabel(w.Widget_23, 0)
	w.Label_29.SetObjectName("label_29")
	w.Label_29.SetAlignment(core.Qt__AlignCenter)
	w.GridLayout_6.AddWidget3(w.Label_29, 0, 3, 1, 1, 0)
	w.AntipubLoaded = widgets.NewQLCDNumber(w.Widget_23)
	w.AntipubLoaded.SetObjectName("antipubLoaded")
	w.AntipubLoaded.SetDigitCount(10)
	w.AntipubLoaded.SetProperty("intValue", core.NewQVariant1(0))
	w.GridLayout_6.AddWidget3(w.AntipubLoaded, 1, 0, 1, 1, 0)
	w.AntipubPublic = widgets.NewQLCDNumber(w.Widget_23)
	w.AntipubPublic.SetObjectName("antipubPublic")
	w.AntipubPublic.SetDigitCount(10)
	w.AntipubPublic.SetProperty("intValue", core.NewQVariant1(0))
	w.GridLayout_6.AddWidget3(w.AntipubPublic, 1, 1, 1, 1, 0)
	w.AntipubPrivate = widgets.NewQLCDNumber(w.Widget_23)
	w.AntipubPrivate.SetObjectName("antipubPrivate")
	w.AntipubPrivate.SetDigitCount(10)
	w.AntipubPrivate.SetProperty("intValue", core.NewQVariant1(0))
	w.GridLayout_6.AddWidget3(w.AntipubPrivate, 1, 2, 1, 1, 0)
	w.AntipubPrivateRatio = widgets.NewQLCDNumber(w.Widget_23)
	w.AntipubPrivateRatio.SetObjectName("antipubPrivateRatio")
	w.AntipubPrivateRatio.SetDigitCount(10)
	w.AntipubPrivateRatio.SetProperty("intValue", core.NewQVariant1(0))
	w.GridLayout_6.AddWidget3(w.AntipubPrivateRatio, 1, 3, 1, 1, 0)
	w.VerticalLayout_39.QLayout.AddWidget(w.Widget_23)
	w.HorizontalLayout_25.QLayout.AddWidget(w.Widget_22)
	w.AutoShellTypesGroup_3 = widgets.NewQGroupBox(w.Tab)
	w.AutoShellTypesGroup_3.SetObjectName("autoShellTypesGroup_3")
	w.AutoShellTypesGroup_3.SetAlignment(int(core.Qt__AlignCenter))
	w.VerticalLayout_38 = widgets.NewQVBoxLayout2(w.AutoShellTypesGroup_3)
	w.VerticalLayout_38.SetObjectName("verticalLayout_38")
	w.Widget_19 = widgets.NewQWidget(w.AutoShellTypesGroup_3, 0)
	w.Widget_19.SetObjectName("widget_19")
	w.HorizontalLayout_22 = widgets.NewQHBoxLayout2(w.Widget_19)
	w.HorizontalLayout_22.SetObjectName("horizontalLayout_22")
	w.AntipubLoadUrls = widgets.NewQPushButton(w.Widget_19)
	w.AntipubLoadUrls.SetObjectName("antipubLoadUrls")
	w.HorizontalLayout_22.QLayout.AddWidget(w.AntipubLoadUrls)
	w.AntipubClearUrls = widgets.NewQPushButton(w.Widget_19)
	w.AntipubClearUrls.SetObjectName("antipubClearUrls")
	w.HorizontalLayout_22.QLayout.AddWidget(w.AntipubClearUrls)
	w.VerticalLayout_38.QLayout.AddWidget(w.Widget_19)
	w.Widget_20 = widgets.NewQWidget(w.AutoShellTypesGroup_3, 0)
	w.Widget_20.SetObjectName("widget_20")
	w.HorizontalLayout_23 = widgets.NewQHBoxLayout2(w.Widget_20)
	w.HorizontalLayout_23.SetObjectName("horizontalLayout_23")
	w.AntipubStart = widgets.NewQPushButton(w.Widget_20)
	w.AntipubStart.SetObjectName("antipubStart")
	w.HorizontalLayout_23.QLayout.AddWidget(w.AntipubStart)
	w.AntipubStop = widgets.NewQPushButton(w.Widget_20)
	w.AntipubStop.SetObjectName("antipubStop")
	w.HorizontalLayout_23.QLayout.AddWidget(w.AntipubStop)
	w.VerticalLayout_38.QLayout.AddWidget(w.Widget_20)
	w.Widget_21 = widgets.NewQWidget(w.AutoShellTypesGroup_3, 0)
	w.Widget_21.SetObjectName("widget_21")
	w.HorizontalLayout_24 = widgets.NewQHBoxLayout2(w.Widget_21)
	w.HorizontalLayout_24.SetObjectName("horizontalLayout_24")
	w.AntipubLoadToDB = widgets.NewQPushButton(w.Widget_21)
	w.AntipubLoadToDB.SetObjectName("antipubLoadToDB")
	w.HorizontalLayout_24.QLayout.AddWidget(w.AntipubLoadToDB)
	w.AntipubExportDB = widgets.NewQPushButton(w.Widget_21)
	w.AntipubExportDB.SetObjectName("antipubExportDB")
	w.HorizontalLayout_24.QLayout.AddWidget(w.AntipubExportDB)
	w.VerticalLayout_38.QLayout.AddWidget(w.Widget_21)
	w.VerticalSpacer_14 = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout_38.AddItem(w.VerticalSpacer_14)
	w.HorizontalLayout_25.QLayout.AddWidget(w.AutoShellTypesGroup_3)
	w.TabControl.AddTab(w.Tab, "")
	w.GridLayout_2.AddWidget3(w.TabControl, 0, 0, 1, 1, 0)
	w.VerticalLayout_15.AddLayout(w.GridLayout_2, 0)
	w.SetCentralWidget(w.Centralwidget)
	w.Menubar = widgets.NewQMenuBar(w)
	w.Menubar.SetObjectName("menubar")
	w.Menubar.SetGeometry(core.NewQRect4(0, 0, 1266, 30))
	w.MenuFile = widgets.NewQMenu(w.Menubar)
	w.MenuFile.SetObjectName("menuFile")
	w.MenuSettings = widgets.NewQMenu(w.Menubar)
	w.MenuSettings.SetObjectName("menuSettings")
	w.SetMenuBar(w.Menubar)
	w.Statusbar = widgets.NewQStatusBar(w)
	w.Statusbar.SetObjectName("statusbar")
	w.Statusbar.SetLayoutDirection(core.Qt__LeftToRight)
	w.SetStatusBar(w.Statusbar)
	w.Menubar.AddActions([]*widgets.QAction{w.MenuFile.MenuAction()})
	w.Menubar.AddActions([]*widgets.QAction{w.MenuSettings.MenuAction()})
	w.MenuFile.AddActions([]*widgets.QAction{w.ActionExit})
	w.MenuSettings.AddActions([]*widgets.QAction{w.ActionOpen_Settings})
	w.retranslateUi()
	w.TabControl.SetCurrentIndex(0)
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *MainWindow) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("MainWindow", "XDumpGO ", "", 0))
	w.ActionExit.SetText(core.QCoreApplication_Translate("MainWindow", "Exit", "", 0))
	w.ActionOpen_Settings.SetText(core.QCoreApplication_Translate("MainWindow", "Open", "", 0))
	w.StatisticsGroup.SetTitle(core.QCoreApplication_Translate("MainWindow", "Statistics", "", 0))
	w.Label_9.SetText(core.QCoreApplication_Translate("MainWindow", "Current Module", "", 0))
	w.Label_10.SetText(core.QCoreApplication_Translate("MainWindow", "Current Time", "", 0))
	w.Label_8.SetText(core.QCoreApplication_Translate("MainWindow", "Runtime", "", 0))
	w.Label_2.SetText(core.QCoreApplication_Translate("MainWindow", "Requests", "", 0))
	w.Label_4.SetText(core.QCoreApplication_Translate("MainWindow", "Errors", "", 0))
	w.Label_3.SetText(core.QCoreApplication_Translate("MainWindow", "RPS", "", 0))
	w.Label_16.SetText(core.QCoreApplication_Translate("MainWindow", "Threads", "", 0))
	w.Label_23.SetText(core.QCoreApplication_Translate("MainWindow", "Workers", "", 0))
	w.StatsModule.SetText(core.QCoreApplication_Translate("MainWindow", "Idle", "", 0))
	w.DataGroup.SetTitle(core.QCoreApplication_Translate("MainWindow", "Data", "", 0))
	w.Label_6.SetText(core.QCoreApplication_Translate("MainWindow", "Urls", "", 0))
	w.Label_5.SetText(core.QCoreApplication_Translate("MainWindow", "Proxies", "", 0))
	w.Label_7.SetText(core.QCoreApplication_Translate("MainWindow", "Dorks", "", 0))
	w.Label_12.SetText(core.QCoreApplication_Translate("MainWindow", "Injectables", "", 0))
	w.DumpStatsGroup.SetTitle(core.QCoreApplication_Translate("MainWindow", "Dumper Stats", "", 0))
	w.Label_13.SetText(core.QCoreApplication_Translate("MainWindow", "Tables", "", 0))
	w.Label_15.SetText(core.QCoreApplication_Translate("MainWindow", "Columns", "", 0))
	w.Label_14.SetText(core.QCoreApplication_Translate("MainWindow", "Rows", "", 0))
	w.Label_30.SetText(core.QCoreApplication_Translate("MainWindow", "Rows Per Minute", "", 0))
	w.GroupBox_4.SetTitle("")
	w.ChatSendMessage.SetText(core.QCoreApplication_Translate("MainWindow", "Send", "", 0))
	w.GroupBox_3.SetTitle(core.QCoreApplication_Translate("MainWindow", "News and Changelog", "", 0))
	w.TabControl.SetTabText(w.TabControl.IndexOf(w.TabWelcome), core.QCoreApplication_Translate("MainWindow", "Dashboard", "", 0))
	w.GroupParameters.SetTitle(core.QCoreApplication_Translate("MainWindow", "Dork Parameters", "", 0))
	w.ComboBoxParameters.SetCurrentText("")
	w.ButtonParametersSettings.SetText(core.QCoreApplication_Translate("MainWindow", "Settings", "", 0))
	w.GeneratorLimiterCheckbox.SetText(core.QCoreApplication_Translate("MainWindow", "Limiter", "", 0))
	w.GroupGenerator.SetTitle(core.QCoreApplication_Translate("MainWindow", "Generator", "", 0))
	___qtablewidgetitem := w.TableGenerator.HorizontalHeaderItem(0)
	___qtablewidgetitem.SetText(core.QCoreApplication_Translate("MainWindow", "Pattern", "", 0))
	___qtablewidgetitem1 := w.TableGenerator.HorizontalHeaderItem(1)
	___qtablewidgetitem1.SetText(core.QCoreApplication_Translate("MainWindow", "Dorks", "", 0))
	w.GeneratorProgress.SetFormat(core.QCoreApplication_Translate("MainWindow", "%p% %v/%m", "", 0))
	w.GroupPatterns.SetTitle(core.QCoreApplication_Translate("MainWindow", "Patterns", "", 0))
	w.GeneratorStartBtn.SetText(core.QCoreApplication_Translate("MainWindow", "Start", "", 0))
	w.GeneratorStopBtn.SetText(core.QCoreApplication_Translate("MainWindow", "Stop", "", 0))
	w.TabControl.SetTabText(w.TabControl.IndexOf(w.TabGenerator), core.QCoreApplication_Translate("MainWindow", "Dork Generator", "", 0))
	w.ParserSearchEnginesGroup.SetTitle(core.QCoreApplication_Translate("MainWindow", "Search Engines", "", 0))
	w.ParserGoogleCheckbox.SetText(core.QCoreApplication_Translate("MainWindow", "Google", "", 0))
	w.ParserBingCheckbox.SetText(core.QCoreApplication_Translate("MainWindow", "Bing", "", 0))
	w.ParserAOLCheckbox.SetText(core.QCoreApplication_Translate("MainWindow", "AOL", "", 0))
	w.ParserMWSCheckbox.SetText(core.QCoreApplication_Translate("MainWindow", "MyWebSearch", "", 0))
	w.ParserDuckDuckGoCheckbox.SetText(core.QCoreApplication_Translate("MainWindow", "DuckDuckGo", "", 0))
	w.ParserEcosiaCheckbox.SetText(core.QCoreApplication_Translate("MainWindow", "Ecosia", "", 0))
	w.ParserStartPageCheckbox.SetText(core.QCoreApplication_Translate("MainWindow", "StartPage", "", 0))
	w.ParserYahooCheckbox.SetText(core.QCoreApplication_Translate("MainWindow", "Yahoo", "", 0))
	w.ParserYandexCheckbox.SetText(core.QCoreApplication_Translate("MainWindow", "Yandex", "", 0))
	w.ParserSettingsGroup.SetTitle(core.QCoreApplication_Translate("MainWindow", "Settings", "", 0))
	w.ParserFilterUrls.SetText(core.QCoreApplication_Translate("MainWindow", "Filter Urls", "", 0))
	w.ParserPagesSpinbox.SetPrefix(core.QCoreApplication_Translate("MainWindow", "Pages: ", "", 0))
	w.ParserCustomParams.SetPlaceholderText(core.QCoreApplication_Translate("MainWindow", "foo=bar&bar=foo", "", 0))
	w.ParserGroup.SetTitle(core.QCoreApplication_Translate("MainWindow", "Parser", "", 0))
	w.ParserDorksTextbox.SetPlainText("")
	w.ParserDorksTextbox.SetPlaceholderText(core.QCoreApplication_Translate("MainWindow", ".php?id=1 .asp?kid=4", "", 0))
	w.ParserProgress.SetFormat(core.QCoreApplication_Translate("MainWindow", "%p% %v/%m", "", 0))
	w.ParserLoadDorksBtn.SetText(core.QCoreApplication_Translate("MainWindow", "Load Dorks", "", 0))
	w.ParserClearDorksBtn.SetText(core.QCoreApplication_Translate("MainWindow", "Clear Dorks", "", 0))
	w.ParserStartBtn.SetText(core.QCoreApplication_Translate("MainWindow", "Start", "", 0))
	w.ParserStopBtn.SetText(core.QCoreApplication_Translate("MainWindow", "Stop", "", 0))
	w.TabControl.SetTabText(w.TabControl.IndexOf(w.TabParser), core.QCoreApplication_Translate("MainWindow", "Dork Parsing", "", 0))
	w.ExploiterTechniquesGroup.SetTitle(core.QCoreApplication_Translate("MainWindow", "Injection Techniques", "", 0))
	w.ExploiterErrorCheckbox.SetText(core.QCoreApplication_Translate("MainWindow", "Error", "", 0))
	w.ExploiterUnionCheckbox.SetText(core.QCoreApplication_Translate("MainWindow", "Union", "", 0))
	w.ExploiterBlindCheckbox.SetText(core.QCoreApplication_Translate("MainWindow", "Blind", "", 0))
	w.ExploiterStackedCheckbox.SetText(core.QCoreApplication_Translate("MainWindow", "Stacked", "", 0))
	w.Label.SetText(core.QCoreApplication_Translate("MainWindow", "Scan Intensity", "", 0))
	w.ExploiterIntensityCombo.SetItemText(0, core.QCoreApplication_Translate("MainWindow", "Basic", "", 0))
	w.ExploiterIntensityCombo.SetItemText(1, core.QCoreApplication_Translate("MainWindow", "Intermediate", "", 0))
	w.ExploiterIntensityCombo.SetItemText(2, core.QCoreApplication_Translate("MainWindow", "Excessive", "", 0))
	w.ExploiterHeuristsicsCheckbox.SetText(core.QCoreApplication_Translate("MainWindow", "Heuristics", "", 0))
	w.ExploiterDatabaseTypes.SetTitle(core.QCoreApplication_Translate("MainWindow", "Databases Types", "", 0))
	w.ExploiterMySQL.SetText(core.QCoreApplication_Translate("MainWindow", "MySQL", "", 0))
	w.ExploiterOracle.SetText(core.QCoreApplication_Translate("MainWindow", "Oracle", "", 0))
	w.ExploiterPostgreSQL.SetText(core.QCoreApplication_Translate("MainWindow", "PostgreSQL", "", 0))
	w.ExploiterMSSQL.SetText(core.QCoreApplication_Translate("MainWindow", "MSSQL", "", 0))
	w.ExploiterGroup.SetTitle(core.QCoreApplication_Translate("MainWindow", "Exploiter", "", 0))
	___qtablewidgetitem2 := w.ExploiterInjectablesTable.HorizontalHeaderItem(0)
	___qtablewidgetitem2.SetText(core.QCoreApplication_Translate("MainWindow", "Url", "", 0))
	___qtablewidgetitem3 := w.ExploiterInjectablesTable.HorizontalHeaderItem(1)
	___qtablewidgetitem3.SetText(core.QCoreApplication_Translate("MainWindow", "Injection", "", 0))
	___qtablewidgetitem4 := w.ExploiterInjectablesTable.HorizontalHeaderItem(2)
	___qtablewidgetitem4.SetText(core.QCoreApplication_Translate("MainWindow", "SQL Version", "", 0))
	___qtablewidgetitem5 := w.ExploiterInjectablesTable.HorizontalHeaderItem(3)
	___qtablewidgetitem5.SetText(core.QCoreApplication_Translate("MainWindow", "SQL User", "", 0))
	___qtablewidgetitem6 := w.ExploiterInjectablesTable.HorizontalHeaderItem(4)
	___qtablewidgetitem6.SetText(core.QCoreApplication_Translate("MainWindow", "Webserver", "", 0))
	w.ExploiterProgress.SetFormat(core.QCoreApplication_Translate("MainWindow", "%p% %v/%m", "", 0))
	w.ExploiterLoadUrlsBtn.SetText(core.QCoreApplication_Translate("MainWindow", "Load Urls", "", 0))
	w.ExploiterClearUrlsBtn.SetText(core.QCoreApplication_Translate("MainWindow", "Clear Urls", "", 0))
	w.ExploiterStartBtn.SetText(core.QCoreApplication_Translate("MainWindow", "Start", "", 0))
	w.ExploiterStopBtn.SetText(core.QCoreApplication_Translate("MainWindow", "Stop", "", 0))
	w.GroupBox_5.SetTitle(core.QCoreApplication_Translate("MainWindow", "Threading", "", 0))
	w.Label_19.SetText(core.QCoreApplication_Translate("MainWindow", "Threads", "", 0))
	w.Label_20.SetText(core.QCoreApplication_Translate("MainWindow", "Workers", "", 0))
	w.TabControl.SetTabText(w.TabControl.IndexOf(w.TabExploiter), core.QCoreApplication_Translate("MainWindow", "Injection Testing", "", 0))
	w.DumperMethodsGroup.SetTitle(core.QCoreApplication_Translate("MainWindow", "Dumping Methods", "", 0))
	w.DumperTargetedCheckbox.SetText(core.QCoreApplication_Translate("MainWindow", "Targeted Dumping", "", 0))
	w.DumperTargetSettingsBtn.SetText(core.QCoreApplication_Translate("MainWindow", "Targetting Settings", "", 0))
	w.DumperKeepBlanksCheckbox.SetText(core.QCoreApplication_Translate("MainWindow", "Keep Blank lines", "", 0))
	w.DumperDIOSCheckbox.SetText(core.QCoreApplication_Translate("MainWindow", "Use DIOS", "", 0))
	w.DumperMinRowsCheckbox.SetText(core.QCoreApplication_Translate("MainWindow", "Minimum Rows", "", 0))
	w.DumperAutoSkip.SetText(core.QCoreApplication_Translate("MainWindow", "AutoSkip", "", 0))
	w.DumperGroup.SetTitle(core.QCoreApplication_Translate("MainWindow", "Dumper", "", 0))
	___qtablewidgetitem7 := w.DumperTableWidget.HorizontalHeaderItem(0)
	___qtablewidgetitem7.SetText(core.QCoreApplication_Translate("MainWindow", "Site", "", 0))
	___qtablewidgetitem8 := w.DumperTableWidget.HorizontalHeaderItem(1)
	___qtablewidgetitem8.SetText(core.QCoreApplication_Translate("MainWindow", "Databases", "", 0))
	___qtablewidgetitem9 := w.DumperTableWidget.HorizontalHeaderItem(2)
	___qtablewidgetitem9.SetText(core.QCoreApplication_Translate("MainWindow", "Tables", "", 0))
	___qtablewidgetitem10 := w.DumperTableWidget.HorizontalHeaderItem(3)
	___qtablewidgetitem10.SetText(core.QCoreApplication_Translate("MainWindow", "Columns", "", 0))
	___qtablewidgetitem11 := w.DumperTableWidget.HorizontalHeaderItem(4)
	___qtablewidgetitem11.SetText(core.QCoreApplication_Translate("MainWindow", "Rows", "", 0))
	___qtablewidgetitem12 := w.DumperTableWidget.HorizontalHeaderItem(5)
	___qtablewidgetitem12.SetText(core.QCoreApplication_Translate("MainWindow", "Errors", "", 0))
	___qtablewidgetitem13 := w.DumperTableWidget.HorizontalHeaderItem(6)
	___qtablewidgetitem13.SetText(core.QCoreApplication_Translate("MainWindow", "Status", "", 0))
	w.DumperProgress.SetFormat(core.QCoreApplication_Translate("MainWindow", "%p% %v/%m", "", 0))
	w.DumperLoadInjectablesBtn.SetText(core.QCoreApplication_Translate("MainWindow", "Load Injectables", "", 0))
	w.DumperClearInjectablesBtn.SetText(core.QCoreApplication_Translate("MainWindow", "Clear Injectables", "", 0))
	w.DumperStartBtn.SetText(core.QCoreApplication_Translate("MainWindow", "Start", "", 0))
	w.DumperStopBtn.SetText(core.QCoreApplication_Translate("MainWindow", "Stop", "", 0))
	w.DumperOpenAnalyzer.SetText(core.QCoreApplication_Translate("MainWindow", "Open Analyzer", "", 0))
	w.GroupBox_6.SetTitle(core.QCoreApplication_Translate("MainWindow", "Threading", "", 0))
	w.Label_21.SetText(core.QCoreApplication_Translate("MainWindow", "Threads", "", 0))
	w.Label_22.SetText(core.QCoreApplication_Translate("MainWindow", "Workers", "", 0))
	w.TabControl.SetTabText(w.TabControl.IndexOf(w.TabDumper), core.QCoreApplication_Translate("MainWindow", "Dumping", "", 0))
	w.AutoShellTypesGroup.SetTitle(core.QCoreApplication_Translate("MainWindow", "Settings", "", 0))
	w.AutoShellASPCheckbox.SetText(core.QCoreApplication_Translate("MainWindow", "ASP", "", 0))
	w.AutoShellPHPCheckbox.SetText(core.QCoreApplication_Translate("MainWindow", "PHP", "", 0))
	w.Label_17.SetText(core.QCoreApplication_Translate("MainWindow", "Shell Key", "", 0))
	w.AutoShellKey.SetText(core.QCoreApplication_Translate("MainWindow", "sdfadsgh4513sdGG435341FDGWWDFGDFHDFGDSFGDFSGDFG", "", 0))
	w.Label_18.SetText(core.QCoreApplication_Translate("MainWindow", "Shell Key", "", 0))
	w.AutoShellFile.SetText(core.QCoreApplication_Translate("MainWindow", "404.php", "", 0))
	w.GroupBox_7.SetTitle(core.QCoreApplication_Translate("MainWindow", "Auto Sheller", "", 0))
	___qtablewidgetitem14 := w.AutoShellTable.HorizontalHeaderItem(0)
	___qtablewidgetitem14.SetText(core.QCoreApplication_Translate("MainWindow", "Site", "", 0))
	___qtablewidgetitem15 := w.AutoShellTable.HorizontalHeaderItem(1)
	___qtablewidgetitem15.SetText(core.QCoreApplication_Translate("MainWindow", "EXT", "", 0))
	___qtablewidgetitem16 := w.AutoShellTable.HorizontalHeaderItem(2)
	___qtablewidgetitem16.SetText(core.QCoreApplication_Translate("MainWindow", "FPD", "", 0))
	___qtablewidgetitem17 := w.AutoShellTable.HorizontalHeaderItem(3)
	___qtablewidgetitem17.SetText(core.QCoreApplication_Translate("MainWindow", "Shell Dropped", "", 0))
	w.AutoShellLoadInjectablesBtn.SetText(core.QCoreApplication_Translate("MainWindow", "Load Injectables", "", 0))
	w.AutoShellClearInjectablesBtn.SetText(core.QCoreApplication_Translate("MainWindow", "Clear Injectables", "", 0))
	w.AutoShellStartBtn.SetText(core.QCoreApplication_Translate("MainWindow", "Start", "", 0))
	w.AutoShellStopBtn.SetText(core.QCoreApplication_Translate("MainWindow", "Stop", "", 0))
	w.TabControl.SetTabText(w.TabControl.IndexOf(w.TabAutoSheller), core.QCoreApplication_Translate("MainWindow", "Auto Sheller", "", 0))
	w.GroupBox.SetTitle(core.QCoreApplication_Translate("MainWindow", "Url Cleaner", "", 0))
	w.CleanerDuplicateDomains.SetText(core.QCoreApplication_Translate("MainWindow", "Filter Duplicate Domains", "", 0))
	w.CleanerQueryParam.SetText(core.QCoreApplication_Translate("MainWindow", "Filter Query Parameter", "", 0))
	w.CleanerBtn.SetText(core.QCoreApplication_Translate("MainWindow", "Clean", "", 0))
	w.GroupBox_2.SetTitle(core.QCoreApplication_Translate("MainWindow", "WebUI", "", 0))
	w.Label_11.SetText(core.QCoreApplication_Translate("MainWindow", "Listen Address", "", 0))
	w.WebUIAddress.SetText(core.QCoreApplication_Translate("MainWindow", ":8080", "", 0))
	w.WebUILaunchBtn.SetText(core.QCoreApplication_Translate("MainWindow", "Launch WebUI", "", 0))
	w.TabControl.SetTabText(w.TabControl.IndexOf(w.TabUtils), core.QCoreApplication_Translate("MainWindow", "Utilities", "", 0))
	w.AutoShellTypesGroup_2.SetTitle(core.QCoreApplication_Translate("MainWindow", "Settings", "", 0))
	w.AntipubSavePublic.SetText(core.QCoreApplication_Translate("MainWindow", "Save Public Lines", "", 0))
	w.AntipubAutoSync.SetText(core.QCoreApplication_Translate("MainWindow", "Auto Sync", "", 0))
	w.GroupBox_8.SetTitle(core.QCoreApplication_Translate("MainWindow", "Antipublic Databases", "", 0))
	w.Label_24.SetText(core.QCoreApplication_Translate("MainWindow", "Link Count", "", 0))
	w.Label_25.SetText(core.QCoreApplication_Translate("MainWindow", "Domain Count", "", 0))
	w.AntipubSizeOnDisk.SetText(core.QCoreApplication_Translate("MainWindow", "Size On Disk: 20KB", "", 0))
	w.AntipubLinkMode.SetText(core.QCoreApplication_Translate("MainWindow", "Link Mode", "", 0))
	w.AntipubDomainMode.SetText(core.QCoreApplication_Translate("MainWindow", "Domain Mode", "", 0))
	w.Label_26.SetText(core.QCoreApplication_Translate("MainWindow", "Loaded", "", 0))
	w.Label_27.SetText(core.QCoreApplication_Translate("MainWindow", "Public", "", 0))
	w.Label_28.SetText(core.QCoreApplication_Translate("MainWindow", "Private", "", 0))
	w.Label_29.SetText(core.QCoreApplication_Translate("MainWindow", "Private Ratio", "", 0))
	w.AutoShellTypesGroup_3.SetTitle("")
	w.AntipubLoadUrls.SetText(core.QCoreApplication_Translate("MainWindow", "Load Urls", "", 0))
	w.AntipubClearUrls.SetText(core.QCoreApplication_Translate("MainWindow", "Clear Urls", "", 0))
	w.AntipubStart.SetText(core.QCoreApplication_Translate("MainWindow", "Start", "", 0))
	w.AntipubStop.SetText(core.QCoreApplication_Translate("MainWindow", "Stop", "", 0))
	w.AntipubLoadToDB.SetText(core.QCoreApplication_Translate("MainWindow", "Load To DB", "", 0))
	w.AntipubExportDB.SetText(core.QCoreApplication_Translate("MainWindow", "Export DB", "", 0))
	w.TabControl.SetTabText(w.TabControl.IndexOf(w.Tab), core.QCoreApplication_Translate("MainWindow", "AntiPublic", "", 0))
	w.MenuFile.SetTitle(core.QCoreApplication_Translate("MainWindow", "File", "", 0))
	w.MenuSettings.SetTitle(core.QCoreApplication_Translate("MainWindow", "Settings", "", 0))

}
