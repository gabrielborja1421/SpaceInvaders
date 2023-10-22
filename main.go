package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"pelota/scenes"

)

func main(){
	myApp := app.New()
	myWindow := myApp.NewWindow("Space Invaders")
	myWindow.CenterOnScreen()
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(800, 600))

	//Cargar y mostrar la escena principal
	mainMenuScene := scenes.NewMainMenuScene(myWindow)
	mainMenuScene.Show()
	myWindow.ShowAndRun()
}
