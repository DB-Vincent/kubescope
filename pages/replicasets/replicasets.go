package replicasets

import (
	"image"
	"image/color"
	"sort"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

	alo "github.com/DB-Vincent/kubescope/applayout"
	"github.com/DB-Vincent/kubescope/icon"
	"github.com/DB-Vincent/kubescope/kubernetes"
	page "github.com/DB-Vincent/kubescope/pages"
	"github.com/xeonx/timeago"
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

	// Kubernetes
	kubeConfig  *kubernetes.KubeConfigOptions
	replicasets []kubernetes.ReplicaSet
	Connected   bool

	// Refresh
	refreshBtn widget.Clickable
}

// New constructs a Page with the provided router.
func New(router *page.Router, kubeConfig *kubernetes.KubeConfigOptions) *Page {
	connected := true

	replicasets, err := kubeConfig.GetReplicaSets()
	if err != nil {
		connected = false
		// panic(err.Error())
	}

	return &Page{
		Router:      router,
		replicasets: replicasets,
		kubeConfig:  kubeConfig,
		Connected:   connected,
	}
}

var _ page.Page = &Page{}

func (p *Page) Actions() []component.AppBarAction {
	return []component.AppBarAction{
		{
			OverflowAction: component.OverflowAction{
				Name: "Refresh",
				Tag:  &p.refreshBtn,
			},
			Layout: func(gtx layout.Context, bg, fg color.NRGBA) layout.Dimensions {
				if p.refreshBtn.Clicked() {
					replicasets, err := p.kubeConfig.GetReplicaSets()
					if err != nil {
						panic(err.Error())
					}
					p.replicasets = replicasets
				}
				btn := component.SimpleIconButton(bg, fg, &p.refreshBtn, icon.RefreshIcon)
				return btn.Layout(gtx)
			},
		},
	}
}

func (p *Page) Overflow() []component.OverflowAction {
	return []component.OverflowAction{}
}

func (p *Page) NavItem() component.NavItem {
	return component.NavItem{
		Name: "ReplicaSets",
		Icon: icon.ReplicaSetsIcon,
	}
}

func (p *Page) Layout(gtx C, th *material.Theme) D {
	p.List.Axis = layout.Vertical

	return material.List(th, &p.List).Layout(gtx, 1, func(gtx C, _ int) D {
		if p.Connected {
			return layout.Flex{
				Alignment: layout.Start,
				Axis:      layout.Vertical,
			}.Layout(gtx,
				layout.Rigid(func(gtx layout.Context) D {
					var visList = layout.List{
						Axis: layout.Vertical,
						Position: layout.Position{
							Offset: 24,
						},
					}

					margins := layout.Inset{
						Top:    unit.Dp(8),
						Right:  unit.Dp(8),
						Left:   unit.Dp(8),
						Bottom: unit.Dp(8),
					}

					sort.SliceStable(p.replicasets, func(i, j int) bool {
						return p.replicasets[i].Name < p.replicasets[j].Name
					})

					return margins.Layout(gtx,
						func(gtx C) D {
							return visList.Layout(gtx, len(p.replicasets),
								func(gtx C, index int) D {
									return layout.Inset{Bottom: unit.Dp(8)}.Layout(gtx, func(gtx C) D {
										return layout.Stack{}.Layout(gtx,
											layout.Expanded(func(gtx C) D {
												gtx.Constraints.Min.X = gtx.Constraints.Max.X
												bgSize := image.Point{X: gtx.Constraints.Min.X, Y: gtx.Constraints.Min.Y}
												return alo.Fill{Col: alo.Rgb(0xffffff)}.Layout(gtx, bgSize)
											}),
											layout.Stacked(func(gtx C) D {
												in := layout.Inset{Top: unit.Dp(8), Right: unit.Dp(16), Bottom: unit.Dp(8), Left: unit.Dp(16)}
												return in.Layout(gtx, func(gtx C) D {
													return layout.Flex{
														Axis:    layout.Vertical,
														Spacing: layout.SpaceSides,
													}.Layout(gtx,
														layout.Rigid(func(gtx C) D {
															return material.H6(th, p.replicasets[index].Name).Layout(gtx)
														}),
														layout.Rigid(func(gtx C) D {
															// p.pods[index].Creation.Format(time.RFC822)
															format := timeago.NoMax(timeago.English)
															return material.Body2(th, format.Format(p.replicasets[index].Creation.Time)).Layout(gtx)
														}),
													)
												})
											}),
										)
									})
								},
							)
						},
					)

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
