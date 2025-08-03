package pages

import (
	"github.com/u00io/nuiforms/ui"
	"github.com/u00io/u00node/system"
)

type Pages struct {
	ui.Widget

	loadedPagesCount int

	selectedUnitId string

	panelPages       *ui.Panel
	panelCurrentPage *ui.Panel
	contentWidget    *PageContent
}

func NewPages() *Pages {
	var c Pages
	c.InitWidget()

	c.SetCellPadding(1)

	c.panelPages = ui.NewPanel()
	c.panelPages.SetXExpandable(false)
	c.panelPages.SetYExpandable(true)
	c.AddWidgetOnGrid(c.panelPages, 0, 0)

	c.panelCurrentPage = ui.NewPanel()
	c.AddWidgetOnGrid(c.panelCurrentPage, 1, 0)

	c.panelCurrentPage.SetXExpandable(true)
	c.panelCurrentPage.SetYExpandable(true)

	c.contentWidget = NewPageContent()
	c.panelCurrentPage.AddWidgetOnGrid(c.contentWidget, 0, 0)

	c.AddTimer(500, c.timerUpdate)

	return &c
}

func (c *Pages) loadPages() {
	state := system.Instance.GetState()
	if len(state.Units) != c.loadedPagesCount {
		c.panelPages.RemoveAllWidgets()
		for i, page := range state.Units {
			pageWidget := NewPageWidget(page.UnitType, page.Id)
			pageWidget.OnClick = func(unitId string) {
				c.SelectPage(unitId)
			}
			c.panelPages.AddWidgetOnGrid(pageWidget, 0, i)
		}
		c.loadedPagesCount = len(state.Units)
	}

	ws := c.panelPages.Widgets()
	for _, w := range ws {
		if pageWidget, ok := w.(*PageWidget); ok {
			pageWidget.UpdateData()
		}
	}
}

func (c *Pages) timerUpdate() {
	c.loadPages()
}

func (c *Pages) SelectPage(unitId string) {
	c.selectedUnitId = unitId
	for _, widget := range c.panelPages.Widgets() {
		if pageWidget, ok := widget.(*PageWidget); ok {
			pageWidget.SetSelected(pageWidget.id == unitId)
		}
	}
	c.contentWidget.SetUnitId(unitId)
}
