package utils

import (
	"image/png"
	"log"
	"os"

	"github.com/kolesa-team/go-webp/encoder"
	"github.com/kolesa-team/go-webp/webp"
)

func PNGToWebp() {
	file, err := os.Open("images/s2.png")
	if err != nil {
		log.Fatalln(err)
	}

	img, err := png.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}

	output, err := os.Create("output/decode.webp")
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	options, err := encoder.NewLossyEncoderOptions(encoder.PresetDefault, 75)
	if err != nil {
		log.Fatalln(err)
	}

	if err := webp.Encode(output, img, options); err != nil {
		log.Fatalln(err)
	}
}
