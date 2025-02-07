package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/juandserrano/rggg-go/game"
)

func main() {
	rl.InitWindow(game.SCREEN_WIDTH, game.SCREEN_HEIGHT, "Rabbits Go Go Go!")
	defer rl.CloseWindow()
	g := game.InitGame()

	// tm := g.Maps["Level_0"].Map
	// tsTexture := rl.LoadTexture(tm.Tilesets[0].Image.Source)
	// defer rl.UnloadTexture(tsTexture)

	rl.SetTargetFPS(120)
	for !rl.WindowShouldClose() {
		g.Update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)

		//Camera Follows
		rl.BeginMode2D(*g.Camera.Camera2D)

		//Draw Tileset
		g.DrawTilemap(g.CurrentLevelName)

		g.DrawWithCamera()
		rl.EndMode2D()

		//Draw Fixed UI -- Doesn't move with camera
		g.DrawWithoutCamera()

		rl.EndDrawing()
	}
	g.Unload()
}
