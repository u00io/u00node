package mainform

import (
	"github.com/u00io/nuiforms/ui"
	"github.com/u00io/u00node/forms/pagedetailswidget"
	"github.com/u00io/u00node/forms/pageswidget"
	"github.com/u00io/u00node/system"
)

type MainForm struct {
	ui.Widget
	topPanel          *ui.Panel
	centerPanel       *ui.Panel
	pagesWidget       *pageswidget.Pages
	pageDetailsWidget *pagedetailswidget.PageDetailsWidget
	bottomPanel       *ui.Panel
}

func NewMainForm() *MainForm {
	system.Instance = system.NewSystem()
	system.Instance.Start()

	var c MainForm
	c.InitWidget()

	// Top panel
	c.topPanel = ui.NewPanel()
	c.topPanel.SetYExpandable(false)
	c.topPanel.AddWidgetOnGrid(ui.NewHSpacer(), 1, 0)
	c.AddWidgetOnGrid(c.topPanel, 0, 0)

	// Center panel
	c.centerPanel = ui.NewPanel()
	c.centerPanel.SetYExpandable(true)
	c.AddWidgetOnGrid(c.centerPanel, 0, 1)

	// Pages widget
	c.pagesWidget = pageswidget.NewPagesWidget()
	c.centerPanel.AddWidgetOnGrid(c.pagesWidget, 0, 0)
	c.pagesWidget.SetXExpandable(false)
	c.pagesWidget.SetYExpandable(true)
	c.pagesWidget.SetMaxWidth(300)
	c.pagesWidget.SetOnPageSelected(func(unitId string) {
		c.pageDetailsWidget.SetUnitId(unitId)
	})

	// Content widget
	c.pageDetailsWidget = pagedetailswidget.NewPageDetailsWidget()
	c.centerPanel.AddWidgetOnGrid(c.pageDetailsWidget, 1, 0)
	c.pageDetailsWidget.SetXExpandable(true)
	c.pageDetailsWidget.SetYExpandable(true)

	// Bottom panel
	c.bottomPanel = ui.NewPanel()
	c.bottomPanel.SetYExpandable(false)
	c.bottomPanel.AddWidgetOnGrid(ui.NewLabel("Powered by U00.io"), 0, 0)
	c.AddWidgetOnGrid(c.bottomPanel, 0, 2)

	return &c
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
