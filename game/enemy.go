package game

import rl "github.com/gen2brain/raylib-go/raylib"

type EnemyManager struct {
	Enemies []Enemy
}

type Enemy struct {
	Pos          rl.Vector2
	LastPos      rl.Vector2
	Dir          rl.Vector2
	Textures     map[string]rl.Texture2D
	Speed        float32
	CollisionBox rl.Rectangle
	Health       float32
}

func InitEnemyManager() EnemyManager {
	em := EnemyManager{}
	return em
}

func (em *EnemyManager) SpawnEnemy(name string, pos rl.Vector2) {
	e := Enemy{
		Pos:          pos,
		LastPos:      pos,
		Dir:          rl.Vector2{X: 0, Y: 0},
		Speed:        70,
		Health:       100,
		CollisionBox: rl.Rectangle{X: pos.X, Y: pos.Y, Width: 64, Height: 64},
	}
	em.Enemies = append(em.Enemies, e)
}

func (em *EnemyManager) DrawEnemies() {
	for _, e := range em.Enemies {
		rl.DrawRectangle(int32(e.Pos.X), int32(e.Pos.Y), 64, 64, rl.White)
		// rl.DrawTexture(e.Textures["current"], int32(e.Pos.X), int32(e.Pos.Y), rl.White)
	}

}
