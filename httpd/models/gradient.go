package models

import (
	"bytes"
	"gradientApi/gradient"
	"image/color"
	"image/png"
)

type Gradient struct{
	Colors 	[]color.RGBA 	`json:"colors"`
	X 		int				`json:"x"`
	Y 		int				`json:"y"`
}

func (g *Gradient)Validate() bool {
	if(g.X < 64 || g.Y < 64 || g.X > 4096 || g.Y > 4096){
		return false
	}
	return true
}

func (g *Gradient)ToByteArray() ([]byte, error) {
	m, err := gradient.Create(g.X, g.Y, g.Colors)
	if err != nil{
		return nil, err
	}
	b := new(bytes.Buffer)
	png.Encode(b, m)
	return b.Bytes(), nil
}