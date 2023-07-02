package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	// Based on the image in the asset pack
	screenWidth  = 1000
	screenHeight = 480
)

var (
	running     = true
	bkgColor    = rl.NewColor(147, 211, 196, 255)
	grassSprite rl.Texture2D

	playerSprite rl.Texture2D

	playerSrc    rl.Rectangle
	playerDest   rl.Rectangle
	playerMoving bool
	// Determines layer and frame ( 0 ... 3)
	playerDir                                     int
	playerUp, playerDown, playerRight, playerLeft bool
	playerFrame                                   int

	frameCount int

	tileDest   rl.Rectangle
	tileSrc    rl.Rectangle
	tileMap    []int
	srcMap     []string
	mapW, mapH int

	playerSpeed float32 = 3

	musicPaused bool
	music       rl.Music

	cam rl.Camera2D
)

func drawScence() {
	//rl.DrawTexture(grassSprite, 100, 50, rl.White)
	//
	for i := 0; i < len(tileMap); i++ {
		// Make sure tileMap is not empty
		if tileMap[i] != 0 {
			tileDest.X = tileDest.Width * float32(i%mapW)
			tileDest.Y = tileDest.Height * float32(i/mapW)
			tileSrc.X = tileSrc.Width * float32((tileMap[i]-1)%int(grassSprite.Width/int32(tileSrc.Width)))
			tileSrc.Y = tileSrc.Height * float32((tileMap[i]-1)/int(grassSprite.Width/int32(tileSrc.Width)))

			rl.DrawTexturePro(grassSprite, tileSrc, tileDest, rl.NewVector2(tileDest.Width, tileDest.Height), 0, rl.White)
		}
	}
	rl.DrawTexturePro(playerSprite, playerSrc, playerDest, rl.NewVector2(playerDest.Width, playerDest.Height), 0, rl.White)

}
func input() {
	if rl.IsKeyDown(rl.KeyW) || rl.IsKeyDown(rl.KeyUp) {
		playerMoving = true
		playerDir = 1
		playerUp = true
	}
	if rl.IsKeyDown(rl.KeyS) || rl.IsKeyDown(rl.KeyDown) {
		playerMoving = true
		playerDir = 0
		playerDown = true
	}
	if rl.IsKeyDown(rl.KeyA) || rl.IsKeyDown(rl.KeyLeft) {
		playerMoving = true
		playerDir = 2
		playerLeft = true
	}
	if rl.IsKeyDown(rl.KeyD) || rl.IsKeyDown(rl.KeyRight) {
		playerMoving = true
		playerDir = 3
		playerRight = true
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		musicPaused = !musicPaused
	}

}
func update() {
	running = !rl.WindowShouldClose()

	playerSrc.X = playerSrc.Width * float32(playerFrame)

	if playerMoving {
		if playerUp {
			playerDest.Y -= playerSpeed
		}
		if playerDown {
			playerDest.Y += playerSpeed
		}
		if playerLeft {
			playerDest.X -= playerSpeed
		}
		if playerRight {
			playerDest.X += playerSpeed
		}
		if frameCount%8 == 1 {
			playerFrame++
		}
		// Players idle animation
		// Make the players frame change every 45 ticks
	} else if frameCount%45 == 1 {
		playerFrame++

	}

	frameCount++
	if playerFrame > 3 {
		playerFrame = 0
	}

	// Players idle animation
	// it only has two frames
	if !playerMoving && playerFrame > 1 {
		playerFrame = 0
	}
	playerSrc.X = playerSrc.Width * float32(playerFrame)

	playerSrc.Y = playerSrc.Height * float32(playerDir)

	rl.UpdateMusicStream(music)
	if musicPaused {
		rl.PauseMusicStream(music)
	} else {
		rl.ResumeMusicStream(music)
	}

	cam.Target = rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height/2)))
	playerMoving = false
	playerUp, playerDown, playerRight, playerLeft = false, false, false, false
}
func render() {

	rl.BeginDrawing()

	rl.ClearBackground(bkgColor)
	rl.BeginMode2D(cam)

	drawScence()

	rl.EndMode2D()
	rl.EndDrawing()
}

func loadMap() {

	mapW = 5
	mapH = 5
	for i := 0; i < (mapW * mapH); i++ {
		tileMap = append(tileMap, 1)
	}

}

func init() {

	rl.InitWindow(screenWidth, screenHeight, "Sproutlings")
	// Set ExitKey to ESC
	rl.SetExitKey(256)
	//defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	grassSprite = rl.LoadTexture("res/Sprout Lands - Sprites - Basic pack/Tilesets/ground tiles/old tiles/Grass.png")

	tileDest = rl.NewRectangle(0, 0, 16, 16)
	tileSrc = rl.NewRectangle(0, 0, 16, 16)

	playerSprite = rl.LoadTexture("res/Sprout Lands - Sprites - Basic pack/Characters/Basic Charakter Spritesheet.png")
	playerSrc = rl.NewRectangle(0, 0, 48, 48)
	playerDest = rl.NewRectangle(200, 200, 100, 100)

	rl.InitAudioDevice()
	music = rl.LoadMusicStream("res/Avery's Farm.mp3")
	musicPaused = false
	rl.PlayMusicStream(music)

	// Note: Character stays in the center of the screen.
	cam = rl.NewCamera2D(rl.NewVector2(float32(screenWidth/2), float32(screenHeight/2)), rl.NewVector2(float32(playerDest.X-(playerDest.Width/2)), float32(playerDest.Y-(playerDest.Height/2))), 0.0, 1.0)
	cam.Zoom = 2.0

	loadMap()

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
