package main

import (
	"bytes"
	"flag"
	"fmt"
	"image/png"
	"log/slog"
	"os"
	"path"

	"github.com/skip2/go-qrcode"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	url := flag.String("url", "", "URL to encode in the QR code")
	size := flag.Int("size", 1024, "width/height length of the .png file in px")
	filePath := flag.String("filePath", path.Join(cwd, "qr.png"), "file path of the .png file")
	flag.Parse()

	content, err := qrcode.Encode(*url, qrcode.Highest, *size)
	if err != nil {
		panic(err)
	}

	img, err := png.Decode(bytes.NewBuffer(content))
	if err != nil {
		panic(err)
	}

	file, err := os.Create(*filePath)
	if err != nil {
		panic(err)
	}

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}

	slog.Info(fmt.Sprintf("written file '%s'", *filePath), "size", *size)
}
