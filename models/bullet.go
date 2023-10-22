// models/bullet.go
package models

import (
	"time"
    // "fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Bullet struct {
    posX, posY float32
    status     bool
    Bullet     *canvas.Image
}

func NewBullet(posx float32, posy float32, img *canvas.Image) *Bullet {
    return &Bullet{
        posX:   posx,
        posY:   posy,
        status: false, // Inicialmente, la bala no está en movimiento
        Bullet: img,
    }
}

func (b *Bullet) Run(naveX float32) {

    for b.status {
        var vel =10
        // Mover la bala hacia arriba (hacia el enemigo)
        b.posY -= float32(vel) // Ajusta la velocidad de movimiento de la bala según sea necesario
        // Utiliza la posición X de la nave para ajustar la posición de la bala
        b.Bullet.Move(fyne.NewPos(naveX, b.posY))
        time.Sleep(100 * time.Millisecond)

        // Verifica si la bala alcanza al enemigo
        if b.posY<0 {
            // La bala alcanza al enemigo

            b.posY=500
            b.Bullet.Move(fyne.NewPos(naveX, b.posY))
            b.SetStatus(false)
        }
    }
}
func (b *Bullet) GetPosition() (float32) {
    return  b.posY
}


func (b *Bullet) SetStatus(status bool) {
    b.status = status
}
