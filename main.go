package main

import (
	"flag"
	"log"
	"os"

	page "github.com/DB-Vincent/kubescope/pages"
	"github.com/DB-Vincent/kubescope/pages/about"
	"github.com/DB-Vincent/kubescope/pages/home"
	"github.com/DB-Vincent/kubescope/pages/pods"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

func main() {
	flag.Parse()
	go func() {
		w := app.NewWindow(
			app.Title("KubeScope"),
		)
		if err := draw(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func draw(w *app.Window) error {
	th := material.NewTheme()
	th.Shaper = text.NewShaper(text.WithCollection(gofont.Collection()))
	var ops op.Ops

	router := page.NewRouter()
	router.Register(0, home.New(&router))
	router.Register(1, pods.New(&router))
	router.Register(2, about.New(&router))

	for {
		select {
		case e := <-w.Events():
			switch e := e.(type) {
			case system.DestroyEvent:
				return e.Err
			case system.FrameEvent:
				gtx := layout.NewContext(&ops, e)
				router.Layout(gtx, th)
				e.Frame(gtx.Ops)
			}
		}
	}
}
