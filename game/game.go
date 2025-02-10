package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	Rm               ResourceManager
	Maps             map[string]Map
	Player           *Player
	Camera           Camera
	CurrentLevelName string
	Em               EnemyManager
}

func InitGame() Game {
	g := Game{}
	g.Rm = InitResourceManager(&g)
	g.Em = InitEnemyManager(&g)
	g.Em.SpawnEnemy("fox", rl.Vector2{X: 300, Y: 300})
	g.Player = InitPlayer()
	g.Camera = InitCamera(g.Player)
	g.Maps = InitMaps()
	g.CurrentLevelName = "Level_0"
	return g
}

func (g *Game) Update() {
	g.Player.Move()
  g.Em.MoveEnemy()
  g.UpdateCollisionBoxes()
	g.PlayerCollisions()
	g.Camera.MoveTo(g.Player)
}

func (g *Game) PlayerCollisions() {
	for _, e := range g.Em.Enemies {
		if rl.CheckCollisionRecs(e.CollisionBox, g.Player.CollisionBox) {
			g.Player.Pos = g.Player.LastPos
			e.Pos = e.LastPos
		}
	}
}

func (g *Game) UpdateCollisionBoxes() {
	g.Player.UpdateCollisionBox()
  g.Em.UpdateCollisionBoxes()
}

func (g *Game) DrawWithCamera() {
	g.DrawActiveTileset()
	g.Player.Draw()
	g.Em.DrawEnemies()
}

func (g *Game) DrawWithoutCamera() {
	rl.DrawText(fmt.Sprintf("Player pos:\n(%d, %d)", int(g.Player.Pos.X), int(g.Player.Pos.Y)), 5, 25, 16, rl.Blue)
	rl.DrawText(fmt.Sprintf("Fox pos:\n(%d, %d)", int(g.Em.Enemies[0].Pos.X), int(g.Em.Enemies[0].Pos.Y)), 5, 45, 16, rl.DarkBlue)
	rl.DrawFPS(5, 5)
}

func (g *Game) DrawActiveTileset() {
	//Draw Tileset
	g.DrawTilemap(g.CurrentLevelName)

}

func (g *Game) DrawTilemapImage(levelName string) {

}

func (g *Game) DrawTilemap(levelName string) {
	tm := g.Maps[levelName]
  for _, layer := range tm.Map.Layers {
    
		for y := 0; y < int(tm.Map.Height); y++ {
			for x := 0; x < int(tm.Map.Width); x++ {
				tileID := layer.Tiles[x+y*int(tm.Map.Width)].ID

				if tileID != 0 { // 0 means no tile
					tileWidth := int(tm.Map.Tilesets[0].TileWidth)
					tileHeight := int(tm.Map.Tilesets[0].TileHeight)

					// Calculate source rectangle (in the tileset image)
					srcX := (int(tileID) % (int(tm.Map.Tilesets[0].Image.Width) / tileWidth)) * tileWidth
					srcY := (int(tileID) / (int(tm.Map.Tilesets[0].Image.Width) / tileWidth)) * tileHeight

					// Calculate destination rectangle (on the screen)
					dstX := float32(x * tileWidth)
					dstY := float32(y * tileHeight)

					srcRect := rl.NewRectangle(float32(srcX), float32(srcY), float32(tileWidth), float32(tileHeight))
					dstRect := rl.NewRectangle(dstX, dstY, float32(tileWidth), float32(tileHeight))

					rl.DrawTexturePro(tm.Texture, srcRect, dstRect, rl.NewVector2(0, 0), 0, rl.White)
				}
			}
		}
	}
}

func (g *Game) Unload() {
	for _, m := range g.Maps {
		rl.UnloadTexture(m.Texture)
	}
	for _, tex := range g.Player.Textures {
		rl.UnloadTexture(tex)
	}
	for _, tex := range g.Em.Enemies {
    for _, t := range tex.Textures {
      rl.UnloadTexture(t)
    }
	}
}
