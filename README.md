# A tool to reconstruct Capcom CPS-1 paper sheets

## How to build

```
go build -o cpss cpss/*.go 
```

## How to add a driver

`cpss` needs to know how to de-interlave a ROM and which region are OBJ, SCROLL1, SCROLL2, and SCROLL3. This is what drivers are for. You can see the numerous examples. To add new game, lookup mame's cps1.cpp.

## How to find palettes

The GFXROM only contains the pens. For the ink, you need to find the palette location which is in the 68000 ROM. A good way to do that is to play the game in Mame, set a breakpoint and lookup the palettes. 


A palette is a series of 16 16-bit values. e.g: 