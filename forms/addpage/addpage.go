package addpage

import (
	"image/color"

	"github.com/u00io/nuiforms/ui"
	"github.com/u00io/u00node/system"
)

type AddPage struct {
	ui.Widget

	selectedCategory string
	selectedUnitType string
	// nameFilter       string

	panelCategories     *ui.Panel
	panelUnitTypes      *ui.Panel
	panelConfig         *ui.Panel
	panelConfigButtons  *ui.Panel
	lblSelectedUnitType *ui.Label
}

func NewAddPage() *AddPage {
	var c AddPage
	c.InitWidget()
	c.SetXExpandable(true)
	c.SetYExpandable(true)

	c.panelCategories = ui.NewPanel()
	c.panelCategories.SetBackgroundColor(color.RGBA{R: 20, G: 20, B: 20, A: 255})
	c.panelCategories.SetMinWidth(300)
	c.panelCategories.SetMaxWidth(300)
	c.panelCategories.SetYExpandable(true)
	c.panelCategories.SetAllowScroll(true, true)
	c.AddWidgetOnGrid(c.panelCategories, 0, 0)

	c.panelUnitTypes = ui.NewPanel()
	c.panelUnitTypes.SetBackgroundColor(color.RGBA{R: 20, G: 20, B: 20, A: 255})
	c.panelUnitTypes.SetMinWidth(300)
	c.panelUnitTypes.SetMaxWidth(300)
	c.panelUnitTypes.SetAllowScroll(true, true)
	c.panelUnitTypes.SetYExpandable(true)
	c.AddWidgetOnGrid(c.panelUnitTypes, 1, 0)

	c.panelConfig = ui.NewPanel()
	c.panelConfig.SetBackgroundColor(color.RGBA{R: 20, G: 20, B: 20, A: 255})
	//c.panelConfig.SetAllowScroll(true, true)
	c.panelConfig.SetYExpandable(true)
	configWidget := NewUnitConfig()
	configWidget.SetMinWidth(300)
	c.panelConfig.AddWidgetOnGrid(configWidget, 0, 1)
	c.AddWidgetOnGrid(c.panelConfig, 2, 0)

	c.panelConfigButtons = ui.NewPanel()
	c.panelConfigButtons.SetBackgroundColor(color.RGBA{R: 20, G: 20, B: 20, A: 255})
	c.panelConfigButtons.SetAllowScroll(false, false)
	c.panelConfigButtons.SetYExpandable(false)
	c.panelConfigButtons.AddWidgetOnGrid(ui.NewButton("Save"), 0, 0)
	c.lblSelectedUnitType = ui.NewLabel("Selected Unit Type: None")
	c.panelConfigButtons.AddWidgetOnGrid(c.lblSelectedUnitType, 1, 0)
	c.panelConfigButtons.AddWidgetOnGrid(ui.NewHSpacer(), 2, 0)
	c.panelConfigButtons.SetMinHeight(120)

	c.panelConfig.AddWidgetOnGrid(c.panelConfigButtons, 0, 0)

	c.loadCategories()
	c.loadUnitTypes()

	c.SelectCategory("All")
	c.SelectUnitType("")

	return &c
}

func (c *AddPage) SelectCategory(category string) {
	c.selectedCategory = category
	for _, widget := range c.panelCategories.Widgets() {
		if catWidget, ok := widget.(*CategoryWidget); ok {
			catWidget.SetSelected(catWidget.categoryName == category)
		}
	}
	c.loadUnitTypes()
}

func (c *AddPage) SelectUnitType(unitType string) {
	c.selectedUnitType = unitType
	for _, widget := range c.panelUnitTypes.Widgets() {
		if unitWidget, ok := widget.(*UnitTypeWidget); ok {
			unitWidget.SetSelected(unitWidget.unitTypeName == unitType)
		}
	}
	c.lblSelectedUnitType.SetText(unitType)
	// load editor
}

func (c *AddPage) loadCategories() {
	ui.MainForm.UpdateBlockPush()
	defer ui.MainForm.UpdateBlockPop()
	ui.MainForm.LayoutingBlockPush()
	defer ui.MainForm.LayoutingBlockPop()

	c.panelCategories.RemoveAllWidgets()
	categories := system.Registry.UnitCategories

	widgetAll := NewCategoryWidget("All")
	widgetAll.OnClick = func(clickedCategory string) {
		c.SelectCategory("All")
	}
	c.panelCategories.AddWidgetOnGrid(widgetAll, 0, c.panelCategories.NextGridY())

	for _, category := range categories {
		widget := NewCategoryWidget(category.Name)
		widget.OnClick = func(clickedCategory string) {
			c.SelectCategory(clickedCategory)
		}
		c.panelCategories.AddWidgetOnGrid(widget, 0, c.panelCategories.NextGridY())
	}
	//c.panelCategories.AddWidgetOnGrid(ui.NewVSpacer(), 0, c.panelCategories.NextGridY())
}

func (c *AddPage) loadUnitTypes() {
	ui.MainForm.UpdateBlockPush()
	defer ui.MainForm.UpdateBlockPop()
	ui.MainForm.LayoutingBlockPush()
	defer ui.MainForm.LayoutingBlockPop()

	c.SelectUnitType("")
	c.panelUnitTypes.RemoveAllWidgets()
	unitTypes := system.Registry.UnitTypes
	for _, record := range unitTypes {
		inFilter := true
		if c.selectedCategory != "" && c.selectedCategory != "All" {
			inFilter = false
			for _, category := range record.Categories {
				if category == c.selectedCategory {
					inFilter = true
					break
				}
			}
		}

		if !inFilter {
			continue
		}

		widget := NewUnitTypeWidget(record.TypeName)
		widget.OnClick = func(clickedItem string) {
			c.SelectUnitType(clickedItem)
		}
		c.panelUnitTypes.AddWidgetOnGrid(widget, 0, c.panelUnitTypes.NextGridY())
	}
	//c.panelRight.AddWidgetOnGrid(ui.NewVSpacer(), 0, c.panelRight.NextGridY())
}
