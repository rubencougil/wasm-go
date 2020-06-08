package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"net/http"
	"syscall/js"
)

type RGBOffset struct {
	R, G, B float64
}

type Images struct {
	OriginalBase64  string
	ConvertedBase64 string
}

var OriginalImg image.Image

func main() {
	c := make(chan struct{}, 0)

	convert := func(this js.Value, args []js.Value) interface{} {
		go func() {
			res := convert(this, args).(Images)
			doc := js.Global().Get("document")
			imgOriginal := doc.Call("getElementById", "imgOriginal")
			imgOriginal.Set("src", "data:image/png;base64, "+res.OriginalBase64)
			img := doc.Call("getElementById", "img")
			img.Set("src", "data:image/png;base64, "+res.ConvertedBase64)
		}()
		return nil
	}

	fmt.Println("Hello, WebAssembly!")
	js.Global().Set("convert", js.FuncOf(convert))
	<-c
}

func convert(this js.Value, args []js.Value) interface{} {
	if OriginalImg == nil {
		resp, err := http.Get("https://cacophony.org.nz/sites/default/files/gopher.png")
		if err != nil {
			log.Fatalf("Cannot GET image fro URL %s", err.Error())
		}
		defer resp.Body.Close()

		OriginalImg, err = png.Decode(resp.Body)
		if err != nil {
			log.Fatalf("Error decoding PNG: %s", err.Error())
		}
	}

	newImg := modifyRGB(OriginalImg, RGBOffset{
		R: args[0].Float(),
		G: args[0].Float(),
		B: args[0].Float(),
	})

	var out bytes.Buffer
	err := png.Encode(&out, newImg)
	if err != nil {
		log.Fatal(err.Error())
	}

	in := new(bytes.Buffer)
	_ = png.Encode(in, OriginalImg)
	return Images{
		OriginalBase64:  base64.StdEncoding.EncodeToString(in.Bytes()),
		ConvertedBase64: base64.StdEncoding.EncodeToString(out.Bytes()),
	}
}

func modifyRGB(img image.Image, rgbOffset RGBOffset) *image.RGBA {
	size := img.Bounds().Size()
	newImg := image.NewRGBA(image.Rect(0, 0, size.X, size.Y))
	for x := 0; x < size.X; x++ {
		for y := 0; y < size.Y; y++ {
			pixel := img.At(x, y)
			originalColor := color.RGBAModel.Convert(pixel).(color.RGBA)

			r := float64(originalColor.R) * rgbOffset.R
			g := float64(originalColor.G) * rgbOffset.G
			b := float64(originalColor.B) * rgbOffset.B

			grey := uint8((r + g + b) / 3)
			c := color.RGBA{
				R: grey, G: grey, B: grey, A: originalColor.A,
			}
			newImg.Set(x, y, c)
		}
	}
	return newImg
}
