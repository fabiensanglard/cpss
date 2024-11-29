package main

import (
	"fmt"
)

type Extractable interface {
	Load() bool
	DumpSheets()
	DumpPaletteToHTML()
	GetName() string
}

func main() {
	fmt.Println("Extracting...")

	var games = make([]Extractable, 0)
	//games = append(games, makeCAW())
	games = append(games, makeFFight())
	games = append(games, makeForgottenUE())
	games = append(games, makeGhouls())
	games = append(games, makePang3())
	games = append(games, makeSF2())
	games = append(games, makeSF2HF())
	games = append(games, makeSFA())
	games = append(games, makeSFA3())
	games = append(games, makeSSF())
	games = append(games, makeStrider())

	//var wg sync.WaitGroup

	for _, game := range games {
		if game.Load() {
			fmt.Println("Found game:", game.GetName())
			//wg.Add(2)
			//go func() {
			//	defer wg.Done()
			game.DumpSheets()
			//}()
			//go func() {
			//	defer wg.Done()
			game.DumpPaletteToHTML()
			//}()
		}
	}
	//wg.Wait()
}
