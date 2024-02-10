package assets

import (
	"embed"
	"image"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed Assets/*
var assets embed.FS

// Gopher Part
var GopherIdle = GetSingleImage("Assets/GopherAssets/mainchar.png")
var GopherRight = GetSingleImage("Assets/GopherAssets/right.png")
var GopherLeft = GetSingleImage("Assets/GopherAssets/left.png")

// Chars part
var AmongUsChar = GetSingleImage("Assets/chars/impostor.png")

// Background Part
var GopherWalkBackground = GetSingleImage("Assets/Backgrounds/GopherWalk.png")
var GopherWalkBackground2 = GetSingleImage("Assets/Backgrounds/GopherWalk2.png")
var GopherJumpBackground = GetSingleImage("Assets/Backgrounds/GopherJump.png")

// Icon Part
var SoftwareIcon = GetSingleImage("Assets/Icons/myapp.png")

// Other Part
var OrangeCat = GetSingleImage("Assets/Others/main.png")
var Door = GetSingleImage("Assets/Others/door.png")

// Font part

//go:embed Assets/fonts/Honk-Regular.ttf
var Font_ttf []byte

func GetSingleImage(name string) *ebiten.Image {
	file, err := assets.Open(name)
	if err != nil {
		panic(err)
	}

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}
