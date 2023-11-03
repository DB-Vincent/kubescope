package home

import (
	"fmt"
	"image"
	"image/color"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"gioui.org/x/outlay"

	"github.com/DB-Vincent/kubescope/icon"
	"github.com/DB-Vincent/kubescope/kubernetes"
	page "github.com/DB-Vincent/kubescope/pages"
)

type (
	C = layout.Context
	D = layout.Dimensions
)

// Page holds the state for a page demonstrating the features of
// the AppBar component.
type Page struct {
	widget.List
	*page.Router

	// Kubernetes counts
	podCount        int
	deployCount     int
	daemonSetCount  int
	replicaSetCount int
	nameSpaceCount  int
}

// New constructs a Page with the provided router.
func New(router *page.Router, kubeConfig *kubernetes.KubeConfigOptions) *Page {
	pods, err := kubeConfig.GetPods()
	if err != nil {
		panic(err.Error())
	}

	deploys, err := kubeConfig.GetDeployments()
	if err != nil {
		panic(err.Error())
	}

	daemonSets, err := kubeConfig.GetDaemonSets()
	if err != nil {
		panic(err.Error())
	}

	replicaSets, err := kubeConfig.GetReplicaSets()
	if err != nil {
		panic(err.Error())
	}

	return &Page{
		Router:          router,
		podCount:        len(pods),
		deployCount:     len(deploys),
		daemonSetCount:  len(daemonSets),
		replicaSetCount: len(replicaSets),
		nameSpaceCount:  len(kubeConfig.Namespaces),
	}
}

var _ page.Page = &Page{}

func (p *Page) Actions() []component.AppBarAction {
	return []component.AppBarAction{}
}

func (p *Page) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (p *Page) NavItem() component.NavItem {
	return component.NavItem{
		Name: "Home",
		Icon: icon.HomeIcon,
	}
}

type HomepageItem struct {
	item  string
	count int
}

func (p *Page) Layout(gtx C, th *material.Theme) D {
	p.List.Axis = layout.Vertical

	// var list = layout.List{
	// 	Axis: layout.Vertical,
	// 	Position: layout.Position{
	// 		Offset: 24,
	// 	},
	// }

	var items = []HomepageItem{
		{item: "Pods", count: p.podCount},
		{item: "Deployments", count: p.deployCount},
		{item: "DaemonSets", count: p.daemonSetCount},
		{item: "ReplicaSets", count: p.replicaSetCount},
		{item: "Namespaces", count: p.nameSpaceCount},
	}

	hGrid := outlay.Flow{
		Num:  2,
		Axis: layout.Horizontal,
	}

	return material.List(th, &p.List).Layout(gtx, 1, func(gtx C, _ int) D {
		return layout.Flex{
			Alignment: layout.Start,
			Axis:      layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(func(gtx C) D {
				return hGrid.Layout(gtx, len(items), func(gtx layout.Context, i int) layout.Dimensions {
					return layout.Inset{Top: unit.Dp(8), Right: unit.Dp(8), Bottom: unit.Dp(8), Left: unit.Dp(8)}.Layout(gtx, func(gtx C) D {
						// return layout.Inset{}.Layout(gtx, func(gtx C) D {
						return layout.Stack{}.Layout(gtx,
							layout.Expanded(func(gtx C) D {
								gtx.Constraints.Min.X = gtx.Constraints.Max.X
								return fill{rgb(0xffffff)}.Layout(gtx)
							}),
							layout.Stacked(func(gtx C) D {
								sz := image.Point{X: gtx.Dp(unit.Dp(150)), Y: gtx.Dp(unit.Dp(100))}
								gtx.Constraints = layout.Exact(gtx.Constraints.Constrain(sz))

								return layout.Flex{
									Axis:    layout.Vertical,
									Spacing: layout.SpaceSides,
								}.Layout(gtx,
									layout.Rigid(func(gtx C) D {
										return material.H3(th, fmt.Sprintf("%d", items[i].count)).Layout(gtx)
									}),
									layout.Rigid(func(gtx C) D {
										return material.Body1(th, items[i].item).Layout(gtx)
									}),
								)
							}),
						)
					})
				})
			}),
		)
	})
}

func rgb(c uint32) color.NRGBA {
	return argb((0xff << 24) | c)
}

func argb(c uint32) color.NRGBA {
	return color.NRGBA{A: uint8(c >> 24), R: uint8(c >> 16), G: uint8(c >> 8), B: uint8(c)}
}

type fill struct {
	col color.NRGBA
}

func (f fill) Layout(gtx layout.Context) layout.Dimensions {
	d := image.Point{X: gtx.Dp(unit.Dp(150)), Y: gtx.Dp(unit.Dp(100))}
	dr := image.Rectangle{
		Max: image.Point{X: d.X, Y: d.Y},
	}

	paint.FillShape(gtx.Ops, f.col, clip.RRect{Rect: dr, SE: 10, SW: 10, NW: 10, NE: 10}.Op(gtx.Ops))
	return layout.Dimensions{Size: d}
}
