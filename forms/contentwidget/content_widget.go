package contentwidget

import (
	"github.com/u00io/nuiforms/ui"
	"github.com/u00io/u00node/forms/addpage"
	"github.com/u00io/u00node/forms/pagedetailswidget"
)

type ContentWidget struct {
	ui.Widget
	panelContent *ui.Panel
}

func NewContentWidget() *ContentWidget {
	var c ContentWidget
	c.InitWidget()

	c.panelContent = ui.NewPanel()
	c.panelContent.SetXExpandable(true)
	c.panelContent.SetYExpandable(true)
	c.AddWidgetOnGrid(c.panelContent, 0, 1)

	c.SetPanelPadding(1)
	c.SetBackgroundColor(c.BackgroundColorAccent1())

	return &c
}

func (c *ContentWidget) SetContent(typeOfContent string, id string) {
	ui.MainForm.UpdateBlockPush()
	defer ui.MainForm.UpdateBlockPop()

	ui.MainForm.LayoutingBlockPush()
	defer ui.MainForm.LayoutingBlockPop()

	if typeOfContent == "page" {
		c.panelContent.RemoveAllWidgets()
		contentWidget := pagedetailswidget.NewPageDetailsWidget()
		contentWidget.SetUnitId(id)
		c.panelContent.AddWidgetOnGrid(contentWidget, 0, 0)
		contentWidget.SetXExpandable(true)
		contentWidget.SetYExpandable(true)
	}

	if typeOfContent == "addpage" {
		c.panelContent.RemoveAllWidgets()
		addPageWidget := addpage.NewAddPage()
		c.panelContent.AddWidgetOnGrid(addPageWidget, 0, 0)
		addPageWidget.SetXExpandable(true)
		addPageWidget.SetYExpandable(true)
	}
}
