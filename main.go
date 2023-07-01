package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	// Based on the image in the asset pack
	screenWidth  = 1000
	screenHeight = 480
)

var (
	running      = true
	bkgColor     = rl.NewColor(147, 211, 196, 255)
	grassSprite  rl.Texture2D
	playerSprite rl.Texture2D

	playerSrc  rl.Rectangle
	playerDest rl.Rectangle
)

func drawScence() {
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)

}
func input() {}
func update() {
	running = !rl.WindowShouldClose()
}
func render() {

	rl.BeginDrawing()

	rl.ClearBackground(bkgColor)
	rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
	drawScence()

	rl.EndDrawing()
}

func init() {

	rl.InitWindow(screenWidth, screenHeight, "Sproutlings")
	rl.SetExitKey(0)
	//defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	grassSprite = rl.LoadTexture("res/Sprout Lands - Sprites - Basic pack/Tilesets/ground tiles/old tiles/Grass.png")
	playerSprite = rl.LoadTexture("res/Sprout Lands - Sprites - Basic pack/Characters/Basic Charakter Spritesheet.png")
	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 100, 100)
}

func quit() {
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)
	rl.CloseWindow()
}

func main() {

	for running {
		input()
		update()
		render()
	}
	quit()
}
