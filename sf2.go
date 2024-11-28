package main

import (
	"image/color"
)
import "crypto/sha1"

func sha(array []byte) []byte {
	h := sha1.New()
	h.Write(array)
	return h.Sum(nil)
}

type SF2 struct {
	Game
}

func (game *SF2) GetName() string {
	return game.name
}

func makeSF2() *SF2 {
	var game SF2

	game.gfxROMSize = 0x600000
	game.gfx_banks = []RomSrc{
		{"sf2-5m.4a", "b9194fb337b30502c1c9501cd6c64ae4035544d4", 2, 0, 0x80000, 0x0000000, 8},
		{"sf2-7m.6a", "3759b851ac0904ec79cbb67a2264d384b6f2f9f9", 2, 0, 0x80000, 0x0000002, 8},
		{"sf2-1m.3a", "520840d727161cf09ca784919fa37bc9b54cc3ce", 2, 0, 0x80000, 0x0000004, 8},
		{"sf2-3m.5a", "2360cff890551f76775739e2d6563858bff80e41", 2, 0, 0x80000, 0x0000006, 8},

		{"sf2-6m.4c", "357c2275af9133fd0bd6fbb1fa9ad5e0b490b3a2", 2, 0, 0x80000, 0x200000, 8},
		{"sf2-8m.6c", "baa92b91cf616bc9e2a8a66adc777ffbf962a51b", 2, 0, 0x80000, 0x200002, 8},
		{"sf2-2m.3c", "2eea16673e60ba7a10bd4d8f6c217bb2441a5b0e", 2, 0, 0x80000, 0x200004, 8},
		{"sf2-4m.5c", "f787aab98668d4c2c54fc4ba677c0cb808e4f31e", 2, 0, 0x80000, 0x200006, 8},

		{"sf2-13m.4d", "5669b845f624b10e7be56bfc89b76592258ce48b", 2, 0, 0x80000, 0x400000, 8},
		{"sf2-15m.6d", "9af9df0826988872662753e9717c48d46f2974b0", 2, 0, 0x80000, 0x400002, 8},
		{"sf2-9m.3d", "a6a7f4725e52678cbd8d557285c01cdccb2c2602", 2, 0, 0x80000, 0x400004, 8},
		{"sf2-11m.5d", "f9a92d614e8877d648449de2612fc8b43c85e4c2", 2, 0, 0x80000, 0x400006, 8},
	}

	game.codeROMSize = 0x400000
	game.code_banks = []RomSrc{
		{"sf2e_30g.11e", "22558eb15e035b09b80935a32b8425d91cd79669", 1, 0, 0x20000, 0x00000, 2},
		{"sf2e_37g.11f", "bf1ccfe7cc1133f0f65556430311108722add1f2", 1, 0, 0x20000, 0x00001, 2},

		{"sf2e_31g.12e", "86a3954335310865b14ce8b4e0e4499feb14fc12", 1, 0, 0x20000, 0x40000, 2},
		{"sf2e_38g.12f", "6565946591a18eaf46f04c1aa449ee0ae9ac2901", 1, 0, 0x20000, 0x40001, 2},

		{"sf2e_28g.9e", "bbcef63f35e5bff3f373968ba1278dd6bd86b593", 1, 0, 0x20000, 0x80000, 2},
		{"sf2e_35g.9f", "507bda3e4519de237aca919cf72e543403ec9724", 1, 0, 0x20000, 0x80001, 2},

		{"sf2_29b.10e", "75f0827f4f7e9f292add46467f8d4fe19b2514c9", 1, 0, 0x20000, 0xc0000, 2},
		{"sf2_36b.10f", "b807cc495bff3f95d03b061fc629c95f965cb6d8", 1, 0, 0x20000, 0xc0000, 2},
	}
	game.paletteAddr = 0x8ACBA // ? Not sure about that but it makes sense
	game.numPalettes = 300     // ? Not sure about that

	game.areas = []Area{
		{0, 0x480000 / (1 << 15), OBJ},
		{0x500000, 0x40000 / (1 << 15), SCR1},
		{0x540000, 0x80000 / (1 << 15), SCR2},
		{0x480000, 0x80000 / (1 << 15), SCR3},
	}

	game.name = "sf2"
	return &game
}

func (game *SF2) Load() bool {
	if !game.Game.Load() {
		return false
	}

	font := game.RetrievePalette(0)

	ryu := game.RetrievePalette(1)
	//ryu_portrait := game.RetrievePalette(159)
	hon := game.RetrievePalette(2)
	bla := game.RetrievePalette(3)
	gui := game.RetrievePalette(4)
	ken := game.RetrievePalette(5)
	chu := game.RetrievePalette(6)
	zan := game.RetrievePalette(7)
	dal := game.RetrievePalette(8)

	// Fireball
	fb := game.RetrievePalette(0xE)

	dic := game.RetrievePalette(0x90)
	dic_cape := game.RetrievePalette(145)

	box := game.RetrievePalette(0xB0)
	sag := game.RetrievePalette(0xA0)
	dac := game.RetrievePalette(0xC0)

	flam := game.RetrievePalette(15)
	//elec := game.RetrievePalette(161)
	elec := game.RetrievePalette(14)
	elec2 := game.RetrievePalette(283)

	warrier := game.RetrievePalette(0x11E)
	logo := game.RetrievePalette(0x11F)
	health_bar := game.RetrievePalette(12)
	credit := game.RetrievePalette(273)
	shimo := game.RetrievePalette(274)
	shimo.colors[10] = color.RGBA{0xFF, 0xFF, 0xFF, 0xff}
	shimo.colors[12] = color.RGBA{0xFF, 0xFF, 0xFF, 0xff}

	barrel := game.RetrievePalette(89)
	car := game.RetrievePalette(217)

	game.set_sheet_color(0, ryu)
	game.set_sprite_color(0, 0xED, 3, 2, ken, OBJ)

	game.set_sheet_color(0x1, ken)
	game.set_sprite_color(0x1, 0x0, 4, 5, ryu, OBJ)
	game.set_sprite_color(0x1, 0x48, 8, 11, ryu, OBJ)
	game.set_sprite_color(0x1, 0xA1, 3, 6, ryu, OBJ)
	game.set_sprite_color(0x1, 0xD0, 1, 1, ryu, OBJ)
	game.set_sprite_color(0x1, 0x84, 4, 6, ryu, OBJ)

	game.set_sheet_color(0x2, ryu)

	game.set_sheet_color(0x3, hon)
	game.set_sprite_color(0x3, 0, 0xF, 0x8, ryu, OBJ)

	for i := 0x4; i <= 0x9; i++ {
		game.set_sheet_color(i, hon)
	}
	game.set_sprite_color(0x4, 0x01, 7, 5, flam, OBJ)
	game.set_sprite_color(0x8, 0x04, 4, 4, flam, OBJ)
	game.set_sprite_color(0x8, 0x44, 1, 3, flam, OBJ)

	for i := 0xA; i <= 0xD; i++ {
		game.set_sheet_color(i, bla)
	}

	for i := 0xE; i <= 0x10; i++ {
		game.set_sheet_color(i, hon)
	}

	game.set_sprite_color(0x10, 0x88, 8, 8, bla, OBJ)
	game.set_sprite_color(0x10, 0x4E, 2, 4, bla, OBJ)

	game.set_sprite_color(0xE, 0x67, 3, 1, fb, OBJ)
	game.set_sprite_color(0xE, 0x5B, 3, 1, fb, OBJ)
	game.set_sprite_color(0xE, 0x4F, 1, 1, fb, OBJ)

	game.set_sheet_color(17, ryu)

	for i := 18; i <= 22; i++ {
		game.set_sheet_color(i, bla)
	}
	game.set_sprite_color(18, 0x96, 1, 1, flam, OBJ)
	game.set_sprite_color(18, 0x99, 2, 1, flam, OBJ)
	game.set_sprite_color(18, 0x95, 1, 1, flam, OBJ)
	game.set_sprite_color(18, 0x98, 2, 1, flam, OBJ)
	game.set_sprite_color(18, 0xD4, 1, 1, flam, OBJ)
	game.set_sprite_color(18, 0xF6, 1, 1, flam, OBJ)
	game.set_sprite_color(18, 0xFD, 3, 1, flam, OBJ)
	game.set_sprite_color(18, 0x81, 1, 1, flam, OBJ)
	game.set_sprite_color(18, 0x8E, 1, 1, flam, OBJ)

	game.set_sprite_color(19, 0x02, 2, 1, flam, OBJ)
	game.set_sprite_color(19, 0x25, 2, 2, flam, OBJ)
	game.set_sprite_color(19, 0x37, 2, 1, flam, OBJ)
	game.set_sprite_color(19, 0x52, 3, 1, flam, OBJ)
	game.set_sprite_color(19, 0x0c, 1, 1, flam, OBJ)
	game.set_sprite_color(19, 0x0f, 2, 1, flam, OBJ)

	for i := 23; i <= 33; i++ {
		game.set_sheet_color(i, zan)
	}

	game.set_sprite_color(27, 0x00, 5, 4, gui, OBJ)

	//game.set_sprite_color(28, 0x00, 5, 4, elec, OBJ)
	game.set_sprite_color(28, 0x17, 3, 6, elec2, OBJ)
	game.set_sprite_color(28, 0x35, 5, 4, elec2, OBJ)

	game.set_sprite_color(33, 0xD0, 16, 3, flam, OBJ)
	game.set_sprite_color(33, 0x8A, 6, 5, flam, OBJ)
	game.set_sprite_color(33, 0x05, 5, 6, gui, OBJ)

	for i := 34; i <= 39; i++ {
		game.set_sheet_color(i, hon)
	}
	game.set_sprite_color(35, 0x08, 2, 1, credit, OBJ)
	game.set_sprite_color(35, 0x048, 2, 1, credit, OBJ)

	game.set_sprite_color(39, 0x00, 2, 2, elec, OBJ)
	game.set_sprite_color(39, 0x50, 2, 2, elec, OBJ)
	game.set_sprite_color(39, 0xA0, 2, 3, elec, OBJ)
	game.set_sprite_color(39, 0x08, 1, 6, elec, OBJ)
	game.set_sprite_color(39, 0x09, 1, 2, elec, OBJ)
	game.set_sprite_color(39, 0x49, 1, 1, elec, OBJ)
	game.set_sprite_color(39, 0x6F, 1, 7, elec, OBJ)
	game.set_sprite_color(39, 0x9E, 1, 4, elec, OBJ)
	game.set_sprite_color(39, 0xaD, 1, 1, elec, OBJ)

	for i := 40; i <= 53; i++ {
		game.set_sheet_color(i, dal)
	}

	game.set_sprite_color(44, 0x06, 2, 5, elec, OBJ)
	game.set_sprite_color(44, 0x0c, 4, 2, elec, OBJ)
	game.set_sprite_color(44, 0x2e, 2, 5, elec, OBJ)
	game.set_sprite_color(44, 0x9f, 1, 6, elec, OBJ)
	game.set_sprite_color(44, 0xbe, 1, 5, elec, OBJ)

	game.set_sprite_color(51, 0x60, 3, 6, chu, OBJ)
	game.set_sprite_color(51, 0x63, 1, 2, chu, OBJ)
	game.set_sprite_color(51, 0x3C, 1, 1, bla, OBJ)
	game.set_sprite_color(51, 0x4C, 3, 1, bla, OBJ)
	game.set_sprite_color(51, 0x4D, 2, 3, bla, OBJ)

	for i := 54; i <= 59; i++ {
		game.set_sheet_color(i, dac)
	}
	game.set_sprite_color(58, 0x03, 2, 4, elec, OBJ)
	game.set_sprite_color(58, 0x50, 5, 3, elec, OBJ)
	game.set_sprite_color(58, 0x70, 1, 5, elec, OBJ)
	game.set_sprite_color(58, 0x80, 3, 3, elec, OBJ)

	for i := 60; i <= 70; i++ {
		game.set_sheet_color(i, ryu)
	}
	// Chunli uses different elec?
	game.set_sprite_color(62, 0xd4, 2, 3, elec, OBJ)
	game.set_sprite_color(62, 0xe6, 2, 2, elec, OBJ)
	game.set_sprite_color(62, 0xc8, 5, 4, elec, OBJ)

	game.set_sprite_color(69, 0x0D, 1, 1, credit, OBJ)
	game.set_sprite_color(69, 0x40, 1, 1, credit, OBJ)
	game.set_sprite_color(69, 0x43, 1, 1, credit, OBJ)
	game.set_sprite_color(69, 0x4E, 1, 1, credit, OBJ)
	game.set_sprite_color(69, 0xC2, 1, 1, credit, OBJ)
	game.set_sprite_color(69, 0x54, 1, 1, credit, OBJ)
	game.set_sprite_color(69, 0x58, 1, 1, credit, OBJ)
	game.set_sprite_color(69, 0x5F, 1, 1, credit, OBJ)
	game.set_sprite_color(69, 0xA1, 1, 1, credit, OBJ)
	game.set_sprite_color(69, 0xA6, 1, 2, shimo, OBJ)
	game.set_sprite_color(69, 0xB6, 3, 1, shimo, OBJ)
	game.set_sprite_color(69, 0xBE, 2, 1, credit, OBJ)

	game.set_sprite_color(69, 0x98, 1, 1, ken, OBJ)
	game.set_sprite_color(69, 0xAA, 1, 1, ken, OBJ)
	game.set_sprite_color(69, 0xB9, 2, 1, ken, OBJ)

	game.set_sprite_color(69, 0xC2, 2, 2, credit, OBJ)
	game.set_sprite_color(69, 0xC4, 1, 1, credit, OBJ)

	game.set_sprite_color(69, 0xC5, 2, 1, chu, OBJ)
	game.set_sprite_color(69, 0xD4, 2, 2, chu, OBJ)
	game.set_sprite_color(69, 0xE0, 1, 1, chu, OBJ)
	game.set_sprite_color(69, 0xE3, 1, 1, chu, OBJ)
	game.set_sprite_color(69, 0xE6, 2, 1, chu, OBJ)
	game.set_sprite_color(69, 0xF8, 1, 1, chu, OBJ)
	game.set_sprite_color(69, 0xCC, 4, 4, chu, OBJ)

	game.set_sheet_color(71, ken)
	game.set_sprite_color(71, 0xC0, 2, 2, shimo, OBJ)
	game.set_sprite_color(71, 0x38, 2, 1, shimo, OBJ)
	game.set_sprite_color(71, 0x8D, 3, 4, flam, OBJ)
	game.set_sprite_color(71, 0xCF, 1, 4, flam, OBJ)
	game.set_sprite_color(71, 0x2D, 3, 1, chu, OBJ)

	for i := 72; i <= 77; i++ {
		game.set_sheet_color(i, box)
	}
	game.set_sprite_color(0x4A, 0xC6, 10, 2, elec, OBJ)
	game.set_sprite_color(0x4A, 0xE7, 9, 1, elec, OBJ)
	game.set_sprite_color(0x4A, 0xFC, 4, 1, elec, OBJ)

	for i := 78; i <= 83; i++ {
		game.set_sheet_color(i, sag)
	}
	game.set_sprite_color(78, 0x00, 4, 1, box, OBJ)
	game.set_sprite_color(78, 0x10, 3, 4, box, OBJ)

	game.set_sheet_color(84, chu)

	game.set_sprite_color(0x54, 0x01, 4, 7, elec, OBJ)
	game.set_sprite_color(0x54, 0x10, 6, 1, elec, OBJ)
	game.set_sprite_color(0x54, 0x15, 1, 3, elec, OBJ)
	game.set_sprite_color(0x54, 0x50, 1, 2, elec, OBJ)
	game.set_sprite_color(0x54, 0x72, 3, 1, elec, OBJ)

	for i := 85; i <= 89; i++ {
		game.set_sheet_color(i, bla)
	}

	for i := 90; i <= 98; i++ {
		game.set_sheet_color(i, dic)
	}
	game.set_sprite_color(0x60, 0x70, 4, 6, elec, OBJ)
	game.set_sprite_color(0x60, 0x61, 3, 1, elec, OBJ)
	game.set_sprite_color(0x60, 0xD0, 3, 1, elec, OBJ)
	game.set_sprite_color(0x60, 0xA3, 1, 3, elec, OBJ)

	game.set_sprite_color(0x62, 0x00, 5, 5, dic_cape, OBJ)
	game.set_sprite_color(0x62, 0x05, 1, 5, dic_cape, OBJ)

	game.set_sprite_color(0x61, 0x50, 5, 5, dic_cape, OBJ)
	game.set_sprite_color(0x61, 0x75, 1, 2, dic_cape, OBJ)

	game.set_sprite_color(0x62, 0xA0, 3, 6, dic_cape, OBJ)
	game.set_sprite_color(0x62, 0xC3, 2, 2, dic_cape, OBJ)
	game.set_sprite_color(0x62, 0xE3, 1, 2, dic_cape, OBJ)

	//game.set_sprite_color(0x61, 0xF5, 1, 1, dic_cape, OBJ)
	//game.set_sprite_color(0x62, 0xC5, 1, 2, dic_cape, OBJ)
	//game.set_sprite_color(0x62, 0x15, 5, 5, dic_cape, OBJ)

	game.set_sprite_color(0x62, 0x9A, 6, 4, dic_cape, OBJ)
	game.set_sprite_color(0x62, 0x8D, 3, 1, dic_cape, OBJ)
	game.set_sprite_color(0x62, 0xDC, 4, 1, dic_cape, OBJ)

	game.set_sprite_color(0x63, 0x00, 4, 0xF, zan, OBJ)
	game.set_sprite_color(0x63, 0x04, 4, 12, zan, OBJ)

	// Zangief electricyt
	game.set_sprite_color(0x63, 0x24, 4, 8, elec, OBJ)
	game.set_sprite_color(0x63, 0x63, 1, 4, elec, OBJ)
	game.set_sprite_color(0x63, 0x82, 1, 2, elec, OBJ)
	game.set_sprite_color(0x63, 0x28, 1, 5, elec, OBJ)
	game.set_sprite_color(0x63, 0x49, 1, 2, elec, OBJ)

	for i := 100; i <= 101; i++ {
		game.set_sheet_color(i, gui)
	}

	for i := 102; i <= 111; i++ {
		game.set_sheet_color(i, chu)
	}

	for i := 112; i <= 113; i++ {
		game.set_sheet_color(i, bla)
	}

	for i := 114; i <= 122; i++ {
		game.set_sheet_color(i, gui)
	}

	game.set_sprite_color(123, 0xC8, 8, 2, warrier, OBJ)

	game.set_sheet_color(128, font)
	game.set_sheet_color(129, font)

	game.set_sprite_color(129, 0xDE, 2, 2, health_bar, OBJ)
	game.set_sprite_color(129, 0xF0, 16, 1, health_bar, OBJ)

	game.set_sprite_color(0x82, 0x00, 2, 2, ryu, OBJ)
	game.set_sprite_color(0x82, 0x02, 2, 2, hon, OBJ)
	game.set_sprite_color(0x82, 0x04, 2, 2, bla, OBJ)
	game.set_sprite_color(0x82, 0x06, 2, 2, gui, OBJ)
	game.set_sprite_color(0x82, 0x08, 2, 2, sag, OBJ)
	game.set_sprite_color(0x82, 0x0A, 2, 2, bla, OBJ)
	game.set_sprite_color(0x82, 0x20, 2, 2, ken, OBJ)
	game.set_sprite_color(0x82, 0x22, 2, 2, chu, OBJ)
	game.set_sprite_color(0x82, 0x24, 2, 2, zan, OBJ)
	game.set_sprite_color(0x82, 0x26, 2, 2, dal, OBJ)
	game.set_sprite_color(0x82, 0x28, 2, 2, box, OBJ)
	game.set_sprite_color(0x82, 0x2A, 2, 2, dic, OBJ)
	game.set_sprite_color(0x82, 0x40, 16, 7, logo, OBJ)
	game.set_sprite_color(0x82, 0xB1, 7, 1, logo, OBJ)
	game.set_sprite_color(0x82, 0xC0, 9, 3, logo, OBJ)
	game.set_sprite_color(0x82, 0xF0, 8, 1, logo, OBJ)

	for i := 136; i <= 143; i++ {
		game.set_sheet_color(i, zan)
	}

	game.set_sprite_color(0x86, 0x08, 8, 10, barrel, OBJ)
	game.set_sprite_color(0x86, 0x40, 8, 4, flam, OBJ)

	// Car
	game.set_sheet_color(0x87, car)

	game.set_sheet_color(144, font)
	return true
}

//
//const CODE_ROMS_SIZE = 1 << 17 // 128 KiB
//const CODE_ROMS_PER_BANK = 2
//const CODE_BANKS = 4
//
//func desinterleave_code_bank(roms [][2]string, dst []byte) {
//	var files [CODE_ROMS_PER_BANK][]byte
//	for i, _ := range files {
//		content, err := ioutil.ReadFile("./roms/" + roms[i][0])
//		hash := sha(content)
//		hash_string := hex.EncodeToString(hash[:])
//		if hash_string != roms[i][1] {
//			fmt.Println(hash_string)
//			fmt.Println(roms[i][1])
//			panic("Unexpected file")
//		}
//
//		if err != nil {
//			panic(err)
//		}
//		files[i] = content
//	}
//
//	var cursor = 0
//	for i := 0; i < CODE_ROMS_SIZE; i++ {
//		for _, f := range files {
//			dst[cursor] = f[i]
//			cursor += 1
//		}
//	}
//
//}
