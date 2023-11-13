# Go & Raylib Tutorial

<https://www.youtube.com/playlist?list=PLVotA8ycjnCsy30WQCwVU5RrZkt4lLgY5>

## Steps

1. get sprites, install go/raylib, create window
2. keyboard input, draw image, player, walking
3. animations, font, music & audio
4. tileset, world, collision
5. inventory
6. farming
7. trading
8. saving, loading, menu

### 1

```shell
mkdir sprouts
go mod init main
apt-get install libgl1-mesa-dev libxi-dev libxcursor-dev libxrandr-dev libxinerama-dev
go get -v -u github.com/gen2brain/raylib-go/raylib
```

Example

```go
package main

import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
 rl.InitWindow(800, 450, "raylib [core] example - basic window")
 defer rl.CloseWindow()
 rl.SetTargetFPS(60)

 for !rl.WindowShouldClose() {
  rl.BeginDrawing()

  rl.ClearBackground(rl.RayWhite)
  rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

  rl.EndDrawing()
 }
}
```

```shell
go build
./main
```

## Build for RGB30 (JELOS).

Build for [RGB30](https://powkiddy.com/products/pre-sale-powkiddy-rgb30-rk3566-handheld-game-console-built-in-wifi)
with [JELOS](https://github.com/JustEnoughLinuxOS/distribution).

```shell
CGO_ENABLED=1 CC=aarch64-linux-gnu-gcc GOOS=linux GOARCH=arm64 go build -tags drm
```

## Resources

- <https://go.dev/>
- <https://www.raylib.com/index.html>
- <https://github.com/gen2brain/raylib-go>
- <https://cupnooble.itch.io/sprout-lands-asset-pack>
- <https://code.visualstudio.com/>
- <https://soundcloud.com/harry-makes/averys-farm9>
- <https://www.raylib.com/examples/textures/loader.html?name=textures_rectangle>
- <https://www.raylib.com/examples/textures/loader.html?name=textures_sprite_anim>
