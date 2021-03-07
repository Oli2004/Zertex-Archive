package qtui

import "github.com/therecipe/qt/widgets"

func SimpleMB(parent widgets.QWidget_ITF, caption string, title string) *widgets.QMessageBox {
	msg := widgets.NewQMessageBox(parent)
	msg.SetText(caption)
	msg.SetWindowTitle(title)
	return msg
}

func NewYesNo(parent widgets.QWidget_ITF, caption string, title string) *Dialog {
	msg := NewYesNoWidget(parent)
	msg.Label.SetText(caption)
	msg.Label.SetWordWrap(true)
	msg.SetWindowTitle(title)
	return msg
}
