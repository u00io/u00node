package pages

import (
	"github.com/u00io/nuiforms/ui"
)

type Pages struct {
	ui.Widget

	lvItems *ui.Table
	lvPage  *ui.Table
}

func NewPages() *Pages {
	var c Pages
	c.InitWidget()

	c.lvItems = ui.NewTable()
	c.lvItems.SetColumnCount(2)
	c.lvItems.SetColumnWidth(0, 100)
	c.lvItems.SetColumnWidth(1, 200)
	c.lvItems.SetColumnName(0, "Page Name")
	c.lvItems.SetColumnName(1, "ID")
	c.lvItems.SetMaxWidth(300)
	c.AddWidgetOnGrid(c.lvItems, 0, 0)

	c.lvPage = ui.NewTable()
	c.lvPage.SetColumnCount(2)
	c.lvPage.SetColumnWidth(0, 200)
	c.lvPage.SetColumnWidth(1, 600)
	c.lvPage.SetColumnName(0, "Name")
	c.lvPage.SetColumnName(1, "Value")
	c.AddWidgetOnGrid(c.lvPage, 1, 0)

	return &c
}
