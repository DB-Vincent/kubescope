package pods

import (
	"fmt"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"gioui.org/x/component"

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
	kubernetes.KubeConfigOptions
	widget.List
	*page.Router
}

// New constructs a Page with the provided router.
func New(router *page.Router) *Page {
	return &Page{
		Router: router,
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
		Name: "Pods",
		Icon: icon.GridIcon,
	}
}

func (p *Page) Layout(gtx C, th *material.Theme) D {
	opts := kubernetes.GetKubeConfig()
	opts.CreateConfig()
	err, data := opts.GetPodStatuses()
	if err != nil {
		panic(err.Error())
	}

	p.List.Axis = layout.Vertical

	return material.List(th, &p.List).Layout(gtx, 1, func(gtx C, _ int) D {
		return layout.Flex{
			Alignment: layout.Middle,
			Axis:      layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return alo.DefaultInset.Layout(
					gtx,
					material.Body1(th, fmt.Sprintf("%d pods on the cluster!", len(data))).Layout,
				)
			}),
		)
	})
}
