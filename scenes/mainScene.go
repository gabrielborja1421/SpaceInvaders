package scenes

import (
	"fmt"
	// "time"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"pelota/models"

	
)

type MainMenuScene struct {
	window fyne.Window
}

// var p *models.Pelota
var n *models.Nave

var gameOverLabel *widget.Label
var restartButton *widget.Button

var status =true
// var e *moldels.nave
var e *models.Enemy

//var g *models.gun
var b *models.Bullet

func NewMainMenuScene(window fyne.Window) *MainMenuScene {
	return &MainMenuScene{window: window,}
}

func (s *MainMenuScene) Show() {
	// pelota := canvas.NewImageFromURI(storage.NewFileURI("./assets/player.png"))
	//nave
	nave := canvas.NewImageFromURI(storage.NewFileURI("./assets/player.png"))
	
	// enemigo
	malo := canvas.NewImageFromURI(storage.NewFileURI("./assets/enemy.png"))


	//disparo
	bala:= canvas.NewImageFromURI(storage.NewFileURI("./assets/bala.png"))


	
	// pelota
	// pelota.Resize(fyne.NewSize(50,50))
	// pelota.Move(fyne.NewPos(350,500))
	// nave
	nave.Resize(fyne.NewSize(50,50))
	nave.Move(fyne.NewPos(350,500))
	
	//enemy
	malo.Resize(fyne.NewSize(50,50))
	malo.Move(fyne.NewPos(50,50))

	//GUN
	bala.Resize(fyne.NewSize(10, 30)) 
	bala.Move(fyne.NewPos(0, 500))
	    
	//Creamos el modelo
	// p = models.NewPelota(100,100,pelota)
	//nave

	n =models.NewNave(350, nave,s.window)
	//enemy
	e = models.NewEnemy(50,50,malo)
	//bala
	b = models.NewBullet(0, 500, bala)
	botonIniciar := widget.NewButton("Start Game", s.StartGame)
	botonIniciar.Resize(fyne.NewSize(150,30))
	botonIniciar.Move(fyne.NewPos(300,10))

	botonDetener := widget.NewButton("Stop Game", s.StopGame)
	botonDetener.Resize(fyne.NewSize(150,30))
	botonDetener.Move(fyne.NewPos(300,50))
	
	// Inicializar el mensaje "Game Over"
	gameOverLabel = widget.NewLabelWithStyle("Game Over", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	gameOverLabel.Resize(fyne.NewSize(200, 50))
	gameOverLabel.Move(fyne.NewPos(250, 250))
	gameOverLabel.Hide()
  
	// Inicializar el botón de reinicio
	restartButton = widget.NewButton("Restart", s.RestartGame)
	restartButton.Resize(fyne.NewSize(150, 30))
	restartButton.Move(fyne.NewPos(300, 300))
	restartButton.Hide()
  
	// pelota
	// s.window.SetContent(container.NewWithoutLayout(pelota, botonIniciar, botonDetener)) 
	// nave
	s.window.SetContent(container.NewWithoutLayout(nave,malo,botonIniciar, botonDetener,bala))
	s.window.Canvas().SetOnTypedRune(s.OnTypedRune)
}

func (s *MainMenuScene) StartGame() {
	// pelota
	// go p.Run()

	// nave
	go n.Run()

	go e.Run()
	go s.Collicions()
	
}
func (s *MainMenuScene) OnTypedRune(r rune) {

    if r == ' ' {
        // Obtiene la posición X de la nave
        posX := n.ObtenerX()

        // Ajusta la posición de inicio de la bala en función de la posición de la nave
        go b.Run(posX)
		b.SetStatus(true)
    }

}

func (s *MainMenuScene) StopGame() {
	// p.SetStatus(false)
	n.SetStatus(false)
	e.SetStatus(false)
	b.SetStatus(false)
}
func (s *MainMenuScene) RestartGame() {
    // Reiniciar el juego ocultando el mensaje "Game Over" y el botón de reinicio
    gameOverLabel.Hide()
    restartButton.Hide()

    // Iniciar un nuevo juego (puedes reiniciar las posiciones y estados aquí)
    n.SetStatus(true)
    e.SetStatus(true)
    b.SetStatus(true)
    go s.StartGame()
}
func (s *MainMenuScene) Collicions() {
	
	for status{
		balaX := n.ObtenerX() 
		balaY := b.GetPosition()
		enemigoX, enemigoY := e.GetPosition()
		// fmt.Println("cordenadas de la nave: ",enemigoX," : ",enemigoY)
		// fmt.Println("cordenadas bala: ",balaX ," : ", balaY)

		margenDeError := float32(10) // Este valor determina el margen de error permitido
		const (
			anchoDelEnemigo float32 = 50
			altoDelEnemigo  float32 = 50
		)
		
		// Lógica para detectar colisiones con margen de error
		if balaX+margenDeError >= enemigoX && balaX-margenDeError <= enemigoX+anchoDelEnemigo &&
			balaY+margenDeError >= enemigoY && balaY-margenDeError <= enemigoY+altoDelEnemigo {
			// Colisión detectada con margen de error, toma acciones necesarias
			fmt.Println("choco")
			status=false
			s.StopGame()
		}
	}
	
}
