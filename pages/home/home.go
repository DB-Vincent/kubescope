package home

import (
	"fmt"
	"image"

	"gioui.org/font"
	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"
	"gioui.org/x/outlay"

	alo "github.com/DB-Vincent/kubescope/applayout"
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

	// Connectivity check
	Connected bool

	// Kubernetes counts
	podCount        int
	deployCount     int
	daemonSetCount  int
	replicaSetCount int
	nameSpaceCount  int
}

// New constructs a Page with the provided router.
func New(router *page.Router, kubeConfig *kubernetes.KubeConfigOptions) *Page {
	connected := true
	pods, err := kubeConfig.GetPods()
	if err != nil {
		connected = false
		// panic(err.Error())
	}

	deploys, err := kubeConfig.GetDeployments()
	if err != nil {
		connected = false
		// panic(err.Error())
	}

	daemonSets, err := kubeConfig.GetDaemonSets()
	if err != nil {
		connected = false
		// panic(err.Error())
	}

	replicaSets, err := kubeConfig.GetReplicaSets()
	if err != nil {
		connected = false
		// panic(err.Error())
	}

	return &Page{
		Router:          router,
		Connected:       connected,
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

	return material.List(th, &p.List).Layout(gtx, 1, func(gtx C, _ int) D {
		if p.Connected {
			return layout.Flex{
				Alignment: layout.Start,
				Axis:      layout.Vertical,
			}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
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

					return hGrid.Layout(gtx, len(items), func(gtx layout.Context, i int) layout.Dimensions {
						return layout.Inset{Top: unit.Dp(8), Left: unit.Dp(8)}.Layout(gtx, func(gtx C) D {
							return layout.Stack{}.Layout(gtx,
								layout.Expanded(func(gtx C) D {
									gtx.Constraints.Min.X = gtx.Constraints.Max.X
									bgSize := image.Point{X: gtx.Constraints.Min.X/2 - 8, Y: gtx.Dp(unit.Dp(100))}
									return alo.Fill{alo.Rgb(0xffffff)}.Layout(gtx, bgSize)
								}),
								layout.Stacked(func(gtx C) D {
									in := layout.Inset{Top: unit.Dp(8), Right: unit.Dp(16), Bottom: unit.Dp(8), Left: unit.Dp(16)}
									return in.Layout(gtx, func(gtx C) D {
										return layout.Flex{
											Alignment: layout.Baseline,
											Axis:      layout.Vertical,
										}.Layout(gtx,
											layout.Rigid(func(gtx C) D {
												countLabel := material.H4(th, fmt.Sprintf("%d", items[i].count))
												countLabel.Font.Weight = font.Bold
												return countLabel.Layout(gtx)
											}),
											layout.Rigid(func(gtx C) D {
												return material.Body1(th, items[i].item).Layout(gtx)
											}),
										)
									})
								}),
							)
						})
					})
				}),
			)
		} else {
			return layout.Flex{
				Alignment: layout.Middle,
				Axis:      layout.Horizontal,
			}.Layout(gtx,
				layout.Rigid(func(gtx C) D {
					return material.Body1(th, "Could not connect to cluster!").Layout(gtx)
				}),
			)
		}
	})
}
