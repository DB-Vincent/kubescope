package replicasets

import (
	"image/color"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

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

	// Kubernetes
	kubeConfig  *kubernetes.KubeConfigOptions
	replicasets []kubernetes.ReplicaSet

	// Refresh
	refreshBtn widget.Clickable
}

// New constructs a Page with the provided router.
func New(router *page.Router, kubeConfig *kubernetes.KubeConfigOptions) *Page {
	replicasets, err := kubeConfig.GetReplicaSets()
	if err != nil {
		panic(err.Error())
	}

	return &Page{
		Router:      router,
		replicasets: replicasets,
		kubeConfig:  kubeConfig,
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

	var fontSize unit.Sp = 16
	var visList = layout.List{
		Axis: layout.Vertical,
		Position: layout.Position{
			Offset: 24,
		},
	}

	margins := layout.Inset{
		Top:   unit.Dp(8),
		Right: unit.Dp(8),
		Left:  unit.Dp(8),
	}

	return material.List(th, &p.List).Layout(gtx, 1, func(gtx C, _ int) D {
		return layout.Flex{
			Alignment: layout.Start,
			Axis:      layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) D {
				return margins.Layout(gtx,
					func(gtx C) D {
						return visList.Layout(gtx, len(p.replicasets),
							func(gtx C, index int) D {
								paragraph := material.Label(th, unit.Sp(float32(fontSize)), p.replicasets[index].Name)
								paragraph.Alignment = text.Start
								return paragraph.Layout(gtx)
							},
						)
					},
				)

			}),
		)
	})
}
