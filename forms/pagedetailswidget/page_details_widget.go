package pagedetailswidget

import (
	"github.com/u00io/nuiforms/ui"
	"github.com/u00io/u00node/utils"
)

type PageDetailsWidget struct {
	ui.Widget

	unitId string

	panelButtons *ui.Panel
	panelContent *ui.Panel

	txtUrl *ui.TextBox

	contentWidget *PageContentWidget
}

func NewPageDetailsWidget() *PageDetailsWidget {
	var c PageDetailsWidget
	c.InitWidget()

	c.panelButtons = ui.NewPanel()
	c.panelButtons.SetYExpandable(false)
	c.panelButtons.SetBackgroundColor(c.BackgroundColorAccent1())
	c.AddWidgetOnGrid(c.panelButtons, 0, 0)

	c.txtUrl = ui.NewTextBox()
	c.txtUrl.SetReadOnly(true)
	c.txtUrl.SetCanBeFocused(false)
	c.txtUrl.SetEmptyText("")
	c.panelButtons.AddWidgetOnGrid(c.txtUrl, 0, 0)

	btnCopy := ui.NewButton("Copy")
	btnCopy.SetOnButtonClick(func(btn *ui.Button) {
		ui.ClipboardSetText(c.generateUrl(c.unitId))
	})
	c.panelButtons.AddWidgetOnGrid(btnCopy, 1, 0)

	btnOpen := ui.NewButton("Open")
	btnOpen.SetOnButtonClick(func(btn *ui.Button) {
		utils.OpenURL(c.generateUrl(c.unitId))
	})
	c.panelButtons.AddWidgetOnGrid(btnOpen, 2, 0)

	c.panelContent = ui.NewPanel()
	c.panelContent.SetXExpandable(true)
	c.panelContent.SetYExpandable(true)
	c.AddWidgetOnGrid(c.panelContent, 0, 1)

	c.contentWidget = NewPageContentWidget()
	c.panelContent.AddWidgetOnGrid(c.contentWidget, 0, 0)
	c.contentWidget.SetXExpandable(true)
	c.contentWidget.SetYExpandable(true)

	c.SetPanelPadding(1)
	c.SetBackgroundColor(c.BackgroundColorAccent1())

	c.SetUnitId("")

	return &c
}

func (c *PageDetailsWidget) SetUnitId(id string) {
	c.unitId = id
	c.contentWidget.SetUnitId(id)
	if id == "" {
		c.txtUrl.SetText("no page selected")
	} else {
		c.txtUrl.SetText(c.generateUrl(id))
	}
	ui.UpdateMainForm()
}

func (c *PageDetailsWidget) generateUrl(id string) string {
	return "https://u00.io/native/0x" + id
}
