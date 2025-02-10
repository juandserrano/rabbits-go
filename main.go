package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/juandserrano/rggg-go/game"
)

func main() {
	rl.InitWindow(game.SCREEN_WIDTH, game.SCREEN_HEIGHT, "Rabbits Go Go Go!")
	defer rl.CloseWindow()
	g := game.InitGame()

	rl.SetTargetFPS(120)
	for !rl.WindowShouldClose() {
		g.Update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)

		//Camera Follows
		rl.BeginMode2D(*g.Camera.Camera2D)

		g.DrawWithCamera()
    DebugDrawCollisionBoxes(&g)
		rl.EndMode2D()

		//Draw Fixed UI -- Doesn't move with camera
		g.DrawWithoutCamera()

		rl.EndDrawing()
	}
	g.Unload()
}

func DebugDrawCollisionBoxes(g *game.Game) {
  for _, e := range g.Em.Enemies {
    rl.DrawRectangleLines(int32(e.CollisionBox.X), int32(e.CollisionBox.Y), int32(e.CollisionBox.Width), int32(e.CollisionBox.Height), rl.Red)
  }
    rl.DrawRectangleLines(int32(g.Player.CollisionBox.X), int32(g.Player.CollisionBox.Y), int32(g.Player.CollisionBox.Width), int32(g.Player.CollisionBox.Height), rl.Blue)
  

}
