package home

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

func (p *Page) Layout(gtx C, th *material.Theme) D {
	p.List.Axis = layout.Vertical

	return material.List(th, &p.List).Layout(gtx, 1, func(gtx C, _ int) D {
		return layout.Flex{
			Alignment: layout.Start,
			Axis:      layout.Vertical,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) D {
				return alo.DefaultInset.Layout(gtx, material.H3(th, "Welcome!").Layout)
			}),
			layout.Rigid(func(gtx layout.Context) D {
				return alo.DefaultInset.Layout(gtx, material.Body1(th, fmt.Sprintf("You're running %d Pods, %d Deployments, %d DaemonSets and %d ReplicaSets in %d Namespaces", p.podCount, p.deployCount, p.daemonSetCount, p.replicaSetCount, p.nameSpaceCount)).Layout)
			}),
		)
	})
}
