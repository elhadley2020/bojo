package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) About() revel.Result {
	return c.Render()
}

func (c App) Services() revel.Result {
	return c.Render()
}

func (c App) Blog() revel.Result {
	return c.Render()
}

func (c App) Contact() revel.Result {
	return c.Render()
}

func (c App) Pickup() revel.Result{
	return c.Render()
}

func (c App) Wtf() revel.Result{
	return c.Render()
}

