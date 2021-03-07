package qtui

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

type __authform struct{}

func (*__authform) init() {}

type AuthForm struct {
	*__authform
	*widgets.QWidget
	VerticalLayout   *widgets.QVBoxLayout
	TabWidget        *widgets.QTabWidget
	TabLogin         *widgets.QWidget
	VerticalLayout_2 *widgets.QVBoxLayout
	LoginUsername    *widgets.QLineEdit
	LoginPassword    *widgets.QLineEdit
	LoginRememberMe  *widgets.QCheckBox
	VerticalSpacer   *widgets.QSpacerItem
	LoginBtn         *widgets.QPushButton
	TabRegister      *widgets.QWidget
	VerticalLayout_3 *widgets.QVBoxLayout
	RegisterUsername *widgets.QLineEdit
	RegisterEmail    *widgets.QLineEdit
	RegisterPassword *widgets.QLineEdit
	RegisterToken    *widgets.QLineEdit
	VerticalSpacer_2 *widgets.QSpacerItem
	RegisterBtn      *widgets.QPushButton
	TabRedeem        *widgets.QWidget
	VerticalLayout_6 *widgets.QVBoxLayout
	RedeemUsername   *widgets.QLineEdit
	RedeemPassword   *widgets.QLineEdit
	RedeemToken      *widgets.QLineEdit
	VerticalSpacer_3 *widgets.QSpacerItem
	RedeemBtn        *widgets.QPushButton
}

func NewAuthForm(p widgets.QWidget_ITF) *AuthForm {
	var par *widgets.QWidget
	if p != nil {
		par = p.QWidget_PTR()
	}
	w := &AuthForm{QWidget: widgets.NewQWidget(par, 0)}
	w.setupUI()
	w.init()
	return w
}
func (w *AuthForm) setupUI() {
	if w.ObjectName() == "" {
		w.SetObjectName("AuthForm")
	}
	w.Resize2(354, 365)
	w.SetMinimumSize(core.NewQSize2(354, 246))
	w.SetMaximumSize(core.NewQSize2(354, 365))
	w.VerticalLayout = widgets.NewQVBoxLayout2(w)
	w.VerticalLayout.SetObjectName("verticalLayout")
	w.TabWidget = widgets.NewQTabWidget(w)
	w.TabWidget.SetObjectName("tabWidget")
	w.TabWidget.SetMovable(false)
	w.TabLogin = widgets.NewQWidget(nil, 0)
	w.TabLogin.SetObjectName("tabLogin")
	w.VerticalLayout_2 = widgets.NewQVBoxLayout2(w.TabLogin)
	w.VerticalLayout_2.SetObjectName("verticalLayout_2")
	w.LoginUsername = widgets.NewQLineEdit(w.TabLogin)
	w.LoginUsername.SetObjectName("loginUsername")
	w.VerticalLayout_2.QLayout.AddWidget(w.LoginUsername)
	w.LoginPassword = widgets.NewQLineEdit(w.TabLogin)
	w.LoginPassword.SetObjectName("loginPassword")
	w.LoginPassword.SetEchoMode(widgets.QLineEdit__PasswordEchoOnEdit)
	w.VerticalLayout_2.QLayout.AddWidget(w.LoginPassword)
	w.LoginRememberMe = widgets.NewQCheckBox(w.TabLogin)
	w.LoginRememberMe.SetObjectName("loginRememberMe")
	w.LoginRememberMe.SetLayoutDirection(core.Qt__RightToLeft)
	w.VerticalLayout_2.QLayout.AddWidget(w.LoginRememberMe)
	w.VerticalSpacer = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout_2.AddItem(w.VerticalSpacer)
	w.LoginBtn = widgets.NewQPushButton(w.TabLogin)
	w.LoginBtn.SetObjectName("loginBtn")
	w.VerticalLayout_2.QLayout.AddWidget(w.LoginBtn)
	w.TabWidget.AddTab(w.TabLogin, "")
	w.TabRegister = widgets.NewQWidget(nil, 0)
	w.TabRegister.SetObjectName("tabRegister")
	w.VerticalLayout_3 = widgets.NewQVBoxLayout2(w.TabRegister)
	w.VerticalLayout_3.SetObjectName("verticalLayout_3")
	w.RegisterUsername = widgets.NewQLineEdit(w.TabRegister)
	w.RegisterUsername.SetObjectName("registerUsername")
	w.VerticalLayout_3.QLayout.AddWidget(w.RegisterUsername)
	w.RegisterEmail = widgets.NewQLineEdit(w.TabRegister)
	w.RegisterEmail.SetObjectName("registerEmail")
	w.VerticalLayout_3.QLayout.AddWidget(w.RegisterEmail)
	w.RegisterPassword = widgets.NewQLineEdit(w.TabRegister)
	w.RegisterPassword.SetObjectName("registerPassword")
	w.RegisterPassword.SetEchoMode(widgets.QLineEdit__PasswordEchoOnEdit)
	w.VerticalLayout_3.QLayout.AddWidget(w.RegisterPassword)
	w.RegisterToken = widgets.NewQLineEdit(w.TabRegister)
	w.RegisterToken.SetObjectName("registerToken")
	w.VerticalLayout_3.QLayout.AddWidget(w.RegisterToken)
	w.VerticalSpacer_2 = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout_3.AddItem(w.VerticalSpacer_2)
	w.RegisterBtn = widgets.NewQPushButton(w.TabRegister)
	w.RegisterBtn.SetObjectName("registerBtn")
	w.VerticalLayout_3.QLayout.AddWidget(w.RegisterBtn)
	w.TabWidget.AddTab(w.TabRegister, "")
	w.TabRedeem = widgets.NewQWidget(nil, 0)
	w.TabRedeem.SetObjectName("tabRedeem")
	w.VerticalLayout_6 = widgets.NewQVBoxLayout2(w.TabRedeem)
	w.VerticalLayout_6.SetObjectName("verticalLayout_6")
	w.RedeemUsername = widgets.NewQLineEdit(w.TabRedeem)
	w.RedeemUsername.SetObjectName("redeemUsername")
	w.VerticalLayout_6.QLayout.AddWidget(w.RedeemUsername)
	w.RedeemPassword = widgets.NewQLineEdit(w.TabRedeem)
	w.RedeemPassword.SetObjectName("redeemPassword")
	w.RedeemPassword.SetEchoMode(widgets.QLineEdit__PasswordEchoOnEdit)
	w.VerticalLayout_6.QLayout.AddWidget(w.RedeemPassword)
	w.RedeemToken = widgets.NewQLineEdit(w.TabRedeem)
	w.RedeemToken.SetObjectName("redeemToken")
	w.VerticalLayout_6.QLayout.AddWidget(w.RedeemToken)
	w.VerticalSpacer_3 = widgets.NewQSpacerItem(20, 40, widgets.QSizePolicy__Minimum, widgets.QSizePolicy__Expanding)
	w.VerticalLayout_6.AddItem(w.VerticalSpacer_3)
	w.RedeemBtn = widgets.NewQPushButton(w.TabRedeem)
	w.RedeemBtn.SetObjectName("redeemBtn")
	w.VerticalLayout_6.QLayout.AddWidget(w.RedeemBtn)
	w.TabWidget.AddTab(w.TabRedeem, "")
	w.VerticalLayout.QLayout.AddWidget(w.TabWidget)
	w.retranslateUi()
	w.TabWidget.SetCurrentIndex(0)
	core.QMetaObject_ConnectSlotsByName(w)

}
func (w *AuthForm) retranslateUi() {
	w.SetWindowTitle(core.QCoreApplication_Translate("AuthForm", "Please Authorize", "", 0))
	w.LoginUsername.SetPlaceholderText(core.QCoreApplication_Translate("AuthForm", "Username", "", 0))
	w.LoginPassword.SetPlaceholderText(core.QCoreApplication_Translate("AuthForm", "Password", "", 0))
	w.LoginRememberMe.SetText(core.QCoreApplication_Translate("AuthForm", "Remember Me", "", 0))
	w.LoginBtn.SetText(core.QCoreApplication_Translate("AuthForm", "Login", "", 0))
	w.TabWidget.SetTabText(w.TabWidget.IndexOf(w.TabLogin), core.QCoreApplication_Translate("AuthForm", "Login", "", 0))
	w.RegisterUsername.SetPlaceholderText(core.QCoreApplication_Translate("AuthForm", "Username", "", 0))
	w.RegisterEmail.SetPlaceholderText(core.QCoreApplication_Translate("AuthForm", "E-Mail", "", 0))
	w.RegisterPassword.SetPlaceholderText(core.QCoreApplication_Translate("AuthForm", "Password", "", 0))
	w.RegisterToken.SetPlaceholderText(core.QCoreApplication_Translate("AuthForm", "Token", "", 0))
	w.RegisterBtn.SetText(core.QCoreApplication_Translate("AuthForm", "Register", "", 0))
	w.TabWidget.SetTabText(w.TabWidget.IndexOf(w.TabRegister), core.QCoreApplication_Translate("AuthForm", "Register", "", 0))
	w.RedeemUsername.SetPlaceholderText(core.QCoreApplication_Translate("AuthForm", "Username", "", 0))
	w.RedeemPassword.SetPlaceholderText(core.QCoreApplication_Translate("AuthForm", "Password", "", 0))
	w.RedeemToken.SetPlaceholderText(core.QCoreApplication_Translate("AuthForm", "Token", "", 0))
	w.RedeemBtn.SetText(core.QCoreApplication_Translate("AuthForm", "Redeem and Login", "", 0))
	w.TabWidget.SetTabText(w.TabWidget.IndexOf(w.TabRedeem), core.QCoreApplication_Translate("AuthForm", "Redeem", "", 0))

}
