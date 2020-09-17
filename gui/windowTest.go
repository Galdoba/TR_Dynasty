package gui

import (
	ui "github.com/VladimirMarkelov/clui"
)

func Test() {
	mainLoop()
}

func mainLoop() {
	// Every application must create a single Composer and
	// call its intialize method
	ui.InitLibrary()
	defer ui.DeinitLibrary()

	b := createView()
	_ = b
	b.SetMaxItems(50)
	b.Draw()

	ui.MainLoop()

}

func createView() *ui.TextView {

	view := ui.AddWindow(0, 0, 10, 7, "TextView Demo")
	bch := ui.CreateTextView(view, 45, 24, 1)
	ui.ActivateControl(view, bch)
	bch.AddText([]string{"Line 1", "...", "Line 3", "Line Very LOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOng"})
	return bch
}
