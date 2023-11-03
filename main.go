package main

import (
	"flag"
	"image/color"
	"log"
	"os"

	"github.com/DB-Vincent/kubescope/kubernetes"
	page "github.com/DB-Vincent/kubescope/pages"
	"github.com/DB-Vincent/kubescope/pages/daemonsets"
	"github.com/DB-Vincent/kubescope/pages/deployments"
	"github.com/DB-Vincent/kubescope/pages/home"
	"github.com/DB-Vincent/kubescope/pages/pods"
	"github.com/DB-Vincent/kubescope/pages/replicasets"

	"gioui.org/app"
	"gioui.org/font/gofont"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/widget/material"
)

var opts kubernetes.KubeConfigOptions

type (
	C = layout.Context
	D = layout.Dimensions
)

func main() {
	flag.Parse()

	var err error
	initialConfig := kubernetes.GetKubeConfig()
	opts, err = initialConfig.CreateConfig()
	opts.GetNamespaces()
	if err != nil {
		log.Fatal(err)
	}

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
	th.Palette.ContrastBg = color.NRGBA{0, 64, 160, 255}
	th.Palette.ContrastFg = color.NRGBA{228, 226, 230, 255}
	th.Palette.Bg = color.NRGBA{0xf2, 0xf2, 0xf2, 0xff}

	var ops op.Ops

	router := page.NewRouter()
	router.Register(0, home.New(&router, &opts))
	router.Register(1, pods.New(&router, &opts))
	router.Register(2, deployments.New(&router, &opts))
	router.Register(3, daemonsets.New(&router, &opts))
	router.Register(4, replicasets.New(&router, &opts))

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
