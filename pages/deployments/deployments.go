package deployments

import (
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

	// Kubernetes counts
	deployments []kubernetes.Deployment
}

// New constructs a Page with the provided router.
func New(router *page.Router, kubeConfig *kubernetes.KubeConfigOptions) *Page {
	deployments, err := kubeConfig.GetDeployments()
	if err != nil {
		panic(err.Error())
	}

	return &Page{
		Router:      router,
		deployments: deployments,
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
		Name: "Deployments",
		Icon: icon.DeploymentsIcon,
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
						return visList.Layout(gtx, len(p.deployments),
							func(gtx C, index int) D {
								paragraph := material.Label(th, unit.Sp(float32(fontSize)), p.deployments[index].Name)
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
