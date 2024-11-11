package main

import (
	"encoding/hex"
	"fmt"
	"image"
	"image/png"
	"io/ioutil"
	"log"
	"os"
)

const TILE_WIDTH = 16
const TILE_HEIGHT = 16

const SHEET_WIDTH = 16
const SHEET_HEIGHT = 16

const TILES_SIZE = TILE_WIDTH * TILE_HEIGHT / 2

const SHEET_SIZE = SHEET_WIDTH * SHEET_HEIGHT * TILES_SIZE // 32 KiB

const PALETTE_SIZE = 32

type RomSrc struct {
	filename   string
	sha        string
	word_size  int
	src_offset int
	src_length int
	dst_offset int
	dst_stride int
}

type TileType int

const (
	OBJ  TileType = 0
	SCR1 TileType = 1
	SCR2 TileType = 2
	SCR3 TileType = 3
)

var tileDimension = []int{
	16,
	8,
	16,
	32,
}

var tileBytes = []int{
	0x80,
	0x40, // not a bug, the SCR1 tiles are packed ineffieintly
	0x80,
	0x200,
}

var tilesPerAxis = []int{
	16,
	32,
	16,
	8,
}

type Area struct {
	offset    int
	numSheets int
	tileType  TileType
}

type Game struct {
	palettes map[int][]*Palette

	gfx_banks  []RomSrc
	code_banks []RomSrc
	name       string

	gfxROMSize int
	gfxROM     []byte

	codeROMSize int
	codeROM     []byte

	paletteAddr int
	numPalettes int

	areas []Area
}

var pixbit = []byte{128, 64, 32, 16, 8, 4, 2, 1}

func (game *Game) drawTile32(tile []byte, drawAtX int, drawAtY int, img *image.RGBA, palette *Palette) {
	tileDim := tileDimension[SCR3]
	offset := 0
	pens := make([]int, tileDim*tileDim)
	for y := 0; y < tileDim; y++ {
		pixels := tile[offset : offset+4] // Get four bytes
		for x := 0; x < 8; x++ {
			if pixels[0]&pixbit[x] != 0 {
				pens[x+y*tileDim] += 1
			}
			if pixels[1]&pixbit[x] != 0 {
				pens[x+y*tileDim] += 2
			}
			if pixels[2]&pixbit[x] != 0 {
				pens[x+y*tileDim] += 4
			}
			if pixels[3]&pixbit[x] != 0 {
				pens[x+y*tileDim] += 8
			}
		}
		offset += 4

		pixels = tile[offset : offset+4] // Get four bytes
		for x := 0; x < 8; x++ {
			if pixels[0]&pixbit[x] != 0 {
				pens[x+8+y*tileDim] += 1
			}
			if pixels[1]&pixbit[x] != 0 {
				pens[x+8+y*tileDim] += 2
			}
			if pixels[2]&pixbit[x] != 0 {
				pens[x+8+y*tileDim] += 4
			}
			if pixels[3]&pixbit[x] != 0 {
				pens[x+8+y*tileDim] += 8
			}
		}
		offset += 4

		pixels = tile[offset : offset+4] // Get four bytes
		for x := 0; x < 8; x++ {
			if pixels[0]&pixbit[x] != 0 {
				pens[x+16+y*tileDim] += 1
			}
			if pixels[1]&pixbit[x] != 0 {
				pens[x+16+y*tileDim] += 2
			}
			if pixels[2]&pixbit[x] != 0 {
				pens[x+16+y*tileDim] += 4
			}
			if pixels[3]&pixbit[x] != 0 {
				pens[x+16+y*tileDim] += 8
			}
		}
		offset += 4

		pixels = tile[offset : offset+4] // Get four bytes
		for x := 0; x < 8; x++ {
			if pixels[0]&pixbit[x] != 0 {
				pens[x+24+y*tileDim] += 1
			}
			if pixels[1]&pixbit[x] != 0 {
				pens[x+24+y*tileDim] += 2
			}
			if pixels[2]&pixbit[x] != 0 {
				pens[x+24+y*tileDim] += 4
			}
			if pixels[3]&pixbit[x] != 0 {
				pens[x+24+y*tileDim] += 8
			}
		}
		offset += 4
	}
	// Write pixels to img
	for y := 0; y < tileDim; y++ {
		for x := 0; x < tileDim; x++ {
			value := pens[x+y*tileDim]
			img.Set(drawAtX+x, drawAtY+y, palette.colors[value])
		}
	}
}

func (game *Game) drawTile8(tile []byte, drawAtX int, drawAtY int, img *image.RGBA, palette *Palette) {
	tileDim := tileDimension[SCR1]
	offset := 0
	pens := make([]int, tileDim*tileDim)
	for y := 0; y < tileDim; y++ {
		pixels := tile[offset : offset+4] // Get four bytes
		for x := 0; x < 8; x++ {
			if pixels[0]&pixbit[x] != 0 {
				pens[x+y*tileDim] += 1
			}
			if pixels[1]&pixbit[x] != 0 {
				pens[x+y*tileDim] += 2
			}
			if pixels[2]&pixbit[x] != 0 {
				pens[x+y*tileDim] += 4
			}
			if pixels[3]&pixbit[x] != 0 {
				pens[x+y*tileDim] += 8
			}
		}
		offset += 8 // Skip one byte, unused
	}

	// Write pixels to img
	for y := 0; y < tileDim; y++ {
		for x := 0; x < tileDim; x++ {
			value := pens[x+y*tileDim]
			img.Set(drawAtX+x, drawAtY+y, palette.colors[value])
		}
	}
}

func (game *Game) drawTile16(tile []byte, drawAtX int, drawAtY int, img *image.RGBA, palette *Palette) {
	tileDim := tileDimension[OBJ]
	offset := 0
	pens := make([]int, tileDim*tileDim)
	for y := 0; y < tileDim; y++ {
		pixels := tile[offset : offset+4] // Get four bytes
		for x := 0; x < 8; x++ {
			if pixels[0]&pixbit[x] != 0 {
				pens[x+y*tileDim] += 1
			}
			if pixels[1]&pixbit[x] != 0 {
				pens[x+y*tileDim] += 2
			}
			if pixels[2]&pixbit[x] != 0 {
				pens[x+y*tileDim] += 4
			}
			if pixels[3]&pixbit[x] != 0 {
				pens[x+y*tileDim] += 8
			}
		}
		offset += 4

		pixels = tile[offset : offset+4] // Get four bytes
		for x := 0; x < 8; x++ {
			if pixels[0]&pixbit[x] != 0 {
				pens[x+8+y*tileDim] += 1
			}
			if pixels[1]&pixbit[x] != 0 {
				pens[x+8+y*tileDim] += 2
			}
			if pixels[2]&pixbit[x] != 0 {
				pens[x+8+y*tileDim] += 4
			}
			if pixels[3]&pixbit[x] != 0 {
				pens[x+8+y*tileDim] += 8
			}
		}

		offset += 4
	}
	// Write pixels to img
	for y := 0; y < tileDim; y++ {
		for x := 0; x < tileDim; x++ {
			value := pens[x+y*tileDim]
			img.Set(drawAtX+x, drawAtY+y, palette.colors[value])
		}
	}
}

func (game *Game) dumpSheet(areaID int, sheet []byte, tileType TileType, sheetID int) {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{SHEET_WIDTH * TILE_WIDTH, SHEET_HEIGHT * TILE_HEIGHT}})
	offset := 0
	bytesPerTile := tileBytes[tileType]
	tileDim := tileDimension[tileType]
	tilesPerAxis := tilesPerAxis[tileType]
	for tileY := 0; tileY < tilesPerAxis; tileY++ {
		for tileX := 0; tileX < tilesPerAxis; tileX++ {
			palette := game.GetPalette(sheetID, tileY*SHEET_WIDTH+tileX)
			drawAtX := tileX * tileDim
			drawAtY := tileY * tileDim
			tileBytes := sheet[offset : offset+bytesPerTile]
			switch tileType {
			case OBJ:
				{
					game.drawTile16(tileBytes, drawAtX, drawAtY, img, palette)
				}
			case SCR2:
				{
					game.drawTile16(tileBytes, drawAtX, drawAtY, img, palette)
				}
			case SCR1:
				{
					game.drawTile8(tileBytes, drawAtX, drawAtY, img, palette)
				}
			case SCR3:
				{
					game.drawTile32(tileBytes, drawAtX, drawAtY, img, palette)
				}
			}
			offset += bytesPerTile
		}
	}
	pngfilename := fmt.Sprintf("%s/%s/area_%d_%04d.png", "/Users/leaf/repos/cps_sheet", game.extractFolder(), areaID, sheetID)
	f, _ := os.Create(pngfilename)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Printf("err=%d", err)
		}
	}(f)
	err := png.Encode(f, img)
	if err != nil {
		fmt.Printf("err=%d", err)
		return
	}
	svgFileName := fmt.Sprintf("%s/%s/area_%d_%04d.svg", "/Users/leaf/repos/cps_sheet", game.extractFolder(), areaID, sheetID)
	png2svg(pngfilename, svgFileName, sheetID, tileType)
}

func (game *Game) desinterleave(roms []RomSrc, dstROM []byte) bool {

	for _, rom := range roms {
		path := "./roms/" + rom.filename
		content, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Println("Unable to open '", rom.filename, "' for ", game.name)
			return false
		}
		hash := sha(content)
		hash_string := hex.EncodeToString(hash[:])
		if hash_string != rom.sha {
			fmt.Println(hash_string)
			fmt.Println(rom.filename)
			fmt.Println("File ", rom.filename, " bad sha. Got (", hash_string, ") but expected (", rom.sha, ")")
			return false
		}

		for j := 0; j < rom.src_length/rom.word_size; j++ {
			srcOffset := rom.src_offset + j*rom.word_size
			src := content[srcOffset : srcOffset+rom.word_size]

			dstOffset := rom.dst_offset + j*rom.dst_stride
			dst := dstROM[dstOffset : dstOffset+rom.word_size]

			for w := 0; w < rom.word_size; w++ {
				dst[w] = src[w]
			}
		}
	}
	return true
}

func (game *Game) Load() bool {
	fmt.Println("\nLoading GFX...", game.name)
	game.gfxROM = make([]byte, game.gfxROMSize)
	success := game.desinterleave(game.gfx_banks, game.gfxROM)
	if !success {
		return false
	}

	game.palettes = make(map[int][]*Palette)

	if game.code_banks != nil {
		fmt.Println("Loading Code...", game.name)
		game.codeROM = make([]byte, game.codeROMSize)
		game.desinterleave(game.code_banks, game.codeROM)
	}

	return true
}

func (game *Game) DumpSheets() {
	game.ensureExtractFolder()

	sheetCounter := 0
	for i := 0; i < len(game.areas); i++ {
		area := game.areas[i]
		for sheetId := 0; sheetId < area.numSheets; sheetId++ {
			offset := area.offset + sheetId*SHEET_SIZE
			sheet := game.gfxROM[offset : offset+SHEET_SIZE]
			game.dumpSheet(i, sheet, area.tileType, sheetCounter)
			sheetCounter += 1
		}
	}
}

func (game *Game) DumpPaletteToHTML() {
	if game.codeROM == nil {
		return
	}
	filename := fmt.Sprintf("%s/palettes.html", game.extractFolder())
	f, _ := os.Create(filename)
	defer f.Close()
	numPalettes := game.numPalettes
	fmt.Println("Found ", numPalettes, " palettes")
	for i := 0; i < numPalettes; i++ {
		palette := game.RetrievePalette(i)
		f.WriteString(fmt.Sprintf("Palette %d<br/>\n", i))
		f.WriteString(palette.toHTML())
	}

	// Also save palette to disk
	filename = fmt.Sprintf("%s/code.bin", game.extractFolder())
	os.WriteFile(filename, game.codeROM, 0666)
}

func (game *Game) w(sheetID int, p *Palette) {
	sheet := make([]*Palette, 32*32) // Able to hold 16x16(OBJ,SCR2) 8x8 (SCR1), and 32x32 (SCR3) sheets

	for i, _ := range sheet {
		sheet[i] = p
	}

	game.palettes[sheetID] = sheet
}

func (game *Game) s(sheetID int, tileID int, width int, height int, palette *Palette, tileType TileType) {
	_, hasSheet := game.palettes[sheetID]
	if !hasSheet {
		game.w(sheetID, &greyPalette)
	}

	sheet := game.palettes[sheetID]

	tilesPerAxis := tilesPerAxis[tileType]
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			sheet[tileID+i*tilesPerAxis+j] = palette
		}
	}
}

func (game *Game) u(sheetID int, tileID int, palette *Palette, tileType TileType) {
	game.s(sheetID, tileID, 1, 1, palette, tileType)
}

func (game *Game) GetPalette(sheetID int, tileNumber int) *Palette {
	// If this sheet is unknown to us, just return default grey palette
	sheet, hasSheet := game.palettes[sheetID]
	if !hasSheet {
		return &greyPalette
	}

	// If the tileID in that sheet nil?
	if sheet[tileNumber] == nil {
		return &greyPalette
	}

	return sheet[tileNumber]
}

func (game *Game) RetrievePalette(paletteId int) *Palette {
	if game.codeROM == nil {
		return &greyPalette
	}

	base := game.paletteAddr + paletteId*PALETTE_SIZE
	paletteSlice := game.codeROM[base : base+PALETTE_SIZE]
	return PaletteFrom(paletteSlice)
}

func (game *Game) ensureExtractFolder() {
	var folder = game.extractFolder()
	err := os.MkdirAll(folder, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func (game *Game) extractFolder() string {
	return fmt.Sprintf("pics/%s", game.name)
}
