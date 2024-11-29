package main

import (
	_ "embed"
	"fmt"
	"strings"
)
import "io/ioutil"
import "os"
import b64 "encoding/base64"

func png2svg(in string, out string, bank int, tileType TileType) {
	payload, err := ioutil.ReadFile(in)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(out)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	img := b64.StdEncoding.EncodeToString(payload)
	sheetId := fmt.Sprintf("%04x", bank<<8)

	var svg string
	switch tileType {
	case OBJ:
		{
			svg = string(svg_template_16)
		}
	case SCR2:
		{
			svg = string(svg_template_16)
		}
	case SCR3:
		{
			svg = string(svg_template_32)
		}
	case SCR1:
		{
			svg = string(svg_template_8)
		}
	}

	svg = strings.Replace(svg, "%%IMG_MARKER%%", img, -1)
	svg = strings.Replace(svg, "%%SHEET_ID_MARKER%%", sheetId, -1)
	f.WriteString(svg)
}

//go:embed 16.svg
var svg_template_16 []byte

//go:embed 32.svg
var svg_template_32 []byte

//go:embed 8.svg
var svg_template_8 []byte
