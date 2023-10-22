package models

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Enemy struct {
	posX, posY float32
	status bool	
	ene *canvas.Image
}

func NewEnemy(posx float32, posy float32, img *canvas.Image) *Enemy {
	return &Enemy{
		posX: posx,
		posY: posy,
		status: true,
		ene: img,
	}
}

func (e *Enemy) Run() {
	var incX=50;
	e.status = true
	for e.status {
		if e.posX <= 0 || e.posX >=750 {
			incX *= -1	
			e.posY+=50
			fmt.Println("bajo")
		}
		// fmt.Println("si estoy funcionando")
		e.posX+=float32(incX)
		e.ene.Move(fyne.NewPos(e.posX,e.posY))
		time.Sleep(700 * time.Millisecond)
		if e.posY==600 {
			e.SetStatus(false)
			fmt.Println("murio")
		}

	}	
}
func (e *Enemy) GetPosition() (float32, float32) {
    return e.posX, e.posY
}

func (e *Enemy) SetStatus(status bool) {
	e.status = status
}
func (e *Enemy) GetStatus() bool  {
	return e.status
}