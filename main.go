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

	playerSpeed float32 = 3

	musicPaused bool
	music       rl.Music
)

func drawScence() {
	rl.DrawTexture(grassSprite, 100, 50, rl.White)
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)

}
func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		playerDest.Y -= playerSpeed
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		playerDest.Y += playerSpeed
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerDest.X -= playerSpeed
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerDest.X += playerSpeed
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		musicPaused = !musicPaused
	}

}
func update() {
	running = !rl.WindowShouldClose()
	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}
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

	rl.InitAudioDevice()
	music = rl.LoadMusicStream("res/Avery's Farm.mp3")
	musicPaused = false
	rl.PlayMusicStream(music)
}

func quit() {
	rl.UnloadTexture(grassSprite)
	rl.UnloadTexture(playerSprite)
	rl.UnloadMusicStream(music)
	rl.CloseAudioDevice()
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
