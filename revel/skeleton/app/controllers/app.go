package controllers

import (
	"github.com/netassist-ua/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}
