package main

type Ghouls struct {
	Game
}

func (game Ghouls) GetName() string {
	return game.name
}

func makeGhouls() *Ghouls {
	var game Ghouls
	game.gfxROMSize = 0x300000
	game.gfx_banks = []RomSrc{
		{"dm-05.3a", "c4945b603115f32b7346d72426571dc2d361159f", 2, 0x00000, 0x80000, 0x00000, 8},
		{"dm-07.3f", "212176947933fcfef991bc80ad5bd91718689ffe", 2, 0x00000, 0x80000, 0x00002, 8},
		{"dm-06.3c", "35bc9dec5ddbf064c30c951627581c16764456ac", 2, 0x00000, 0x80000, 0x00004, 8},
		{"dm-08.3g", "7d0c4736f16577afe9966447a18f039728f6fbdf", 2, 0x00000, 0x80000, 0x00006, 8},
	}

	game.name = "ghouls"
	game.paletteAddr = 4

	game.codeROMSize = 0x100000
	game.code_banks = []RomSrc{
		{"dme_29.10h", "f21fcf88d2ebb7bc9e8885fde760a5d82f295c1a", 1, 0, 0x20000, 0x00000, 2},
		{"dme_30.10j", "3613699213db47bfeabedf87f12eb0fa7e5973b6", 1, 0, 0x20000, 0x00001, 2},
		{"dme_27.9h", "fa230bf5503487ec11d767485a18f0a55dcc13d2", 1, 0, 0x20000, 0x40000, 2},
		{"dme_28.9j", "a07786062358c89f3b4634b8822173261802290b", 1, 0, 0x20000, 0x40001, 2},
		{"dm-17.7j", "c51f1c38cdaed77ad715cedd845617a291ab2441", 2, 0, 0x80000, 0x80000, 0},
	}

	return &game
}

func (game *Ghouls) Load() bool {
	if !game.Game.Load() {
		return false
	}
	return true
}
