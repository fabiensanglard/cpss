# A tool to reconstruct Capcom CPS-1 paper sheets

Home page: [https://fabiensanglard.net/cpss/](https://fabiensanglard.net/cpss/)
## How to build

```
go build -o cpss src/*.go 
```

## How to add a driver

`cpss` needs to know how to deinterleave the GFXROM and also which region are OBJ, SCROLL1, SCROLL2, and SCROLL3.
This is what drivers are for. You can see the numerous examples. To add new game, lookup mame's cps1.cpp.

## How to find the palette base

The GFXROM only contains the pens. For the ink, you need to find the palette location which is in the 68000 ROM.
A good way to do that is to play the game in Mame, set a breakpoint and lookup the palettes. 

A palette is a series of 16 16-bit values encoded as _RGB (4-bit per channel). e.g Ryu from SF2:

```
0x0111, 0x0FD9, 0x0FB8, 0x0E97, 0x0C86, 0x0965 ...
```

If you lookup with an hex editor for 0x01110FD90E94, you get an offset and that is the palette location.

```
hexdump -ve '1/1 "%02X"' pics/sf2/code.bin | grep -b -o  01110FD90FB80E970C860965
175372:01110FD90FB80E970C860965
175628:01110FD90FB80E970C860965
1137076:01110FD90FB80E970C860965
```

That is how sf2 driver sets the palette base to 0x8ACBA (568506). Because Ryu is the second palette and palette is 32 bytes long

## How to find more palettes

`cpss` also generate an HTML page, [palettes.html](https://fabiensanglard.net/cpss/palettes.html) so you can visually inspect colors and speed the discovery process.