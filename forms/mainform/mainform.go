package mainform

import (
	"github.com/u00io/nuiforms/ui"
	"github.com/u00io/u00node/forms/addpage"
	"github.com/u00io/u00node/forms/pages"
	"github.com/u00io/u00node/system"
)

type MainForm struct {
	ui.Widget

	topPanel  *ui.Panel
	tabWidget *ui.TabWidget

	widgetPages   *pages.Pages
	widgetAddPage *addpage.AddPage

	bottomPanel *ui.Panel
}

func NewMainForm() *MainForm {
	system.Instance = system.NewSystem()
	system.Instance.Start()

	var form MainForm
	form.InitWidget()

	form.topPanel = ui.NewPanel()
	form.topPanel.SetYExpandable(false)
	//form.topPanel.AddWidgetOnGrid(ui.NewLabel("U00 Node"), 0, 0)

	form.topPanel.AddWidgetOnGrid(ui.NewHSpacer(), 1, 0)
	form.AddWidgetOnGrid(form.topPanel, 0, 0)

	form.widgetAddPage = addpage.NewAddPage()

	form.widgetPages = pages.NewPages()

	form.tabWidget = ui.NewTabWidget()
	form.tabWidget.AddPage("Pages", form.widgetPages)
	form.tabWidget.AddPage("Add Page", form.widgetAddPage)
	form.AddWidgetOnGrid(form.tabWidget, 0, 1)

	form.bottomPanel = ui.NewPanel()
	form.bottomPanel.SetYExpandable(false)
	form.bottomPanel.AddWidgetOnGrid(ui.NewLabel("Powered by U00.io"), 0, 0)
	form.AddWidgetOnGrid(form.bottomPanel, 0, 2)

	return &form
}

func Run() {
	form := ui.NewForm()
	form.SetTitle("U00 Node")
	form.SetSize(800, 600)

	form.Panel().AddWidgetOnGrid(NewMainForm(), 0, 0)

	form.Exec()

	/*sendValue := func() {
		client.WriteValue("MyItem", time.Now(), time.Now().Format("2006-01-02 15:04:05.000000"))
	}

	{
		btnSend := ui.NewButton()
		btnSend.SetPosition(10, 10)
		btnSend.SetSize(100, 30)
		btnSend.SetProp("text", "Send")
		btnSend.SetProp("onClick", func() {
			sendValue()
		})
		c.Panel().AddWidget(btnSend)
	}

	{
		btnOpenUrl := ui.NewButton()
		btnOpenUrl.SetPosition(120, 10)
		btnOpenUrl.SetSize(100, 30)
		btnOpenUrl.SetProp("text", "Open URL")
		btnOpenUrl.SetProp("onClick", func() {
			err := openURL("https://u00.io/native/" + client.Address())
			if err != nil {
				fmt.Println("Error opening URL:", err)
			}
		})
		c.Panel().AddWidget(btnOpenUrl)
	}

	{
		c.AddTimer(1000, func() {
			fmt.Println("Current Time:", time.Now().Format("2006-01-02 15:04:05.000000"))
			sendValue()
		})
	}*/
}
