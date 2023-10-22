package models

import (
	"fmt"
	// "time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	// "golang.org/x/sys/windows"
	"github.com/eiannone/keyboard"
	"os"
)

type Nave struct {
	posX float32
	status bool	
	nav *canvas.Image
	window     fyne.Window
}

func NewNave(posx float32, img *canvas.Image,window fyne.Window) *Nave {
	return &Nave{
		posX: posx,
		status: true,
		nav: img,
		window: window,
	}
}
func (p *Nave) Run() {
    p.status = true

    err := keyboard.Open()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer keyboard.Close()

    fmt.Println("Presiona 'a' para izquierda, 'd' para derecha (Presiona 'Esc' para salir):")

    for {
        char, key, err := keyboard.GetKey()
        if err != nil {
            fmt.Println(err)
            break
        }

        if key == keyboard.KeyEsc {
            break
        }

        switch char {
        case 'a':
            // fmt.Println("Izquierda")
            if p.posX > 0 {
                p.posX -= 10
                p.nav.Move(fyne.NewPos(p.posX, 500))
            }

        case 'd':
            // fmt.Println("Derecha")
            if p.posX < 750 {
                p.posX += 10
                p.nav.Move(fyne.NewPos(p.posX, 500))
            }
        }
    }
}
func(p *Nave) ObtenerX()float32{
	return float32(p.posX)
}
func (p *Nave) SetStatus(status bool) {
	p.status = status
}
