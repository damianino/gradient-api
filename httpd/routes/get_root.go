package routes

import (
	"bytes"
	"gradientApi/gradient"
	"image/color"
	"image/png"
	"math/rand"
	"net/http"
	"strconv"
)

func RootGet(w http.ResponseWriter, r *http.Request) {
	stopN := rand.Intn(2) + 2
	var stops []color.RGBA
	for i := 0; i < stopN; i++ {
		stops = append(stops, color.RGBA{R: uint8(rand.Intn(256)), G: uint8(rand.Intn(256)), B: uint8(rand.Intn(256)), A: 255})
	}

	m, err := gradient.Create(256, 256, stops)
	b := new(bytes.Buffer)
	png.Encode(b, m)
	
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(b.Len()))
	if _, err := w.Write(b.Bytes()); err != nil {
		w.WriteHeader(500)
		return
	}

}
