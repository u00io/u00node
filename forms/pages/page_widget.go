package pages

import (
	"image/color"

	"github.com/u00io/nui/nuikey"
	"github.com/u00io/nui/nuimouse"
	"github.com/u00io/nuiforms/ui"
)

type PageWidget struct {
	ui.Widget
	id           string
	categoryName string
	selected     bool
	OnClick      func(clickedCategory string)

	lblCategory *ui.Label
	lblUnitId   *ui.Label
}

func NewPageWidget(categoryName string, id string) *PageWidget {
	var c PageWidget
	c.InitWidget()

	c.SetPanelPadding(1)

	c.id = id
	c.categoryName = categoryName

	c.lblCategory = ui.NewLabel(categoryName)
	c.lblCategory.SetMouseCursor(nuimouse.MouseCursorPointer)
	c.lblCategory.SetOnMouseDown(func(button nuimouse.MouseButton, x int, y int, mods nuikey.KeyModifiers) bool {
		if button == nuimouse.MouseButtonLeft {
			c.Click()
		}
		return true
	})
	c.AddWidgetOnGrid(c.lblCategory, 0, 0)

	c.lblUnitId = ui.NewLabel(id)
	c.lblUnitId.SetMouseCursor(nuimouse.MouseCursorPointer)
	c.lblUnitId.SetOnMouseDown(func(button nuimouse.MouseButton, x int, y int, mods nuikey.KeyModifiers) bool {
		if button == nuimouse.MouseButtonLeft {
			c.Click()
		}
		return true
	})
	c.AddWidgetOnGrid(c.lblUnitId, 0, 1)

	c.SetYExpandable(false)
	c.SetMinWidth(300)
	c.SetMinHeight(120)
	c.SetMaxHeight(120)
	c.SetSelected(false)
	c.SetMouseCursor(nuimouse.MouseCursorPointer)
	c.SetOnMouseDown(func(button nuimouse.MouseButton, x int, y int, mods nuikey.KeyModifiers) bool {
		if button == nuimouse.MouseButtonLeft {
			c.Click()
		}
		return true
	})
	return &c
}

func (c *PageWidget) Click() {
	if c.OnClick != nil {
		c.OnClick(c.id)
	}
}

func (c *PageWidget) SetSelected(selected bool) {
	c.selected = selected
	if selected {
		backColor := c.BackgroundColorAccent1()
		c.SetBackgroundColor(backColor)
		c.lblCategory.SetBackgroundColor(backColor)
		c.lblUnitId.SetBackgroundColor(backColor)
	} else {
		var backColor color.Color
		c.SetBackgroundColor(backColor)
		c.lblCategory.SetBackgroundColor(backColor)
		c.lblUnitId.SetBackgroundColor(backColor)
	}
}

func (c *PageWidget) IsSelected() bool {
	return c.selected
}

func (c *PageWidget) UpdateData() {
}
