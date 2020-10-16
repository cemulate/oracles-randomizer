package randomizer

import (
	"image/color"
	"os"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

func newPlayerOptionsCanvas() fyne.CanvasObject {
	sel := widget.NewSelect([]string{"Seasons", "Ages"}, func(s string) {})
	sel.PlaceHolder = "(Select game)"
	check1 := widget.NewCheck("Treewarp (+t)", func(b bool) {})
	check2 := widget.NewCheck("Hard logic (+h)", func(b bool) {})
	check3 := widget.NewCheck("Dungeon shuffle (+d)", func(b bool) {})
	check4 := widget.NewCheck("Portal shuffle (+p)", func(b bool) {})
	return fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		sel, check1, check2, check3, check4)
}

func runGUI() {
	os.Setenv("FYNE_THEME", "light")

	myApp := app.New()
	myWindow := myApp.NewWindow("Oracles Randomizer " + version)

	titleText := canvas.NewText("Oracles Randomizer "+version, color.Black)
	titleContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		layout.NewSpacer(), titleText, layout.NewSpacer())

	seedLabel := widget.NewLabel("Seed:")
	seedLabelLayout := fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		layout.NewSpacer(), seedLabel, layout.NewSpacer())
	seedEntry := widget.NewEntry()
	seedContainer := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		seedLabelLayout, seedEntry)

	tab1 := widget.NewTabItem("Player 1", newPlayerOptionsCanvas())
	tab2 := widget.NewTabItem("Player 2", newPlayerOptionsCanvas())
	tabs := widget.NewTabContainer(tab1, tab2)

	button1 := widget.NewButton("Add player", func() {})
	button2 := widget.NewButton("Remove player", func() {})
	button3 := widget.NewButton("Generate", func() {})
	buttons := fyne.NewContainerWithLayout(layout.NewHBoxLayout(),
		button1, button2, button3)

	myWindow.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(),
		titleContainer, seedContainer, layout.NewSpacer(), tabs, buttons))
	myWindow.ShowAndRun()
}
