package server

import (
	"github.com/fatih/color"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func GetMainScreenAndInstances() (app *tview.Application, instanceOne, instanceTwo *tview.TextView) {
	root := tview.NewFlex()

	app = tview.NewApplication()
	app.SetRoot(root, true).EnableMouse(true)

	instanceOne, instanceTwo = tview.NewTextView(), tview.NewTextView()

	instanceOne.SetTextColor(tcell.Color(color.FgHiBlue)).SetScrollable(true).SetChangedFunc(func() {
		app.Draw()
	}).SetBorderPadding(1, 0, 2, 2)

	instanceTwo.SetTextColor(tcell.Color(color.FgHiWhite)).SetScrollable(true).SetChangedFunc(func() {
		app.Draw()
	}).SetBorderPadding(1, 0, 2, 2)

	root.AddItem(instanceOne, 0, 1, false)
	root.AddItem(instanceTwo, 0, 1, false)

	root.SetBorder(true)
	root.SetBorderAttributes(tcell.AttrNone)

	return app, instanceOne, instanceTwo
}
