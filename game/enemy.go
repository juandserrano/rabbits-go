package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type EnemyManager struct {
	Enemies []*Enemy
	g       *Game
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

func InitEnemyManager(g *Game) EnemyManager {
	em := EnemyManager{}
	em.g = g
	return em
}

func (em *EnemyManager) SpawnEnemy(name string, pos rl.Vector2) {
	e := Enemy{
		Pos:          pos,
		LastPos:      pos,
		Dir:          rl.Vector2{X: 0, Y: 0},
		CollisionBox: rl.Rectangle{X: pos.X + ENEMY_CB_OFFSET_X, Y: pos.Y, Width: 40, Height: 64},
		Textures:     make(map[string]rl.Texture2D),
	}
	switch name {
	case "fox":
		e.Speed = ENEMY_FOX_SPEED
		e.Health = ENEMY_FOX_HEALTH
		e.Textures["front"] = em.g.Rm.Textures["fox"]
		e.Textures["current"] = e.Textures["front"]

		break
	default:
	}
	em.Enemies = append(em.Enemies, &e)
}

func (em *EnemyManager) MoveEnemy(){
  fox := em.Enemies[0]
  fox.LastPos = fox.Pos
	fox.Dir = rl.Vector2{X: 0, Y: 0}
	if rl.IsKeyDown(rl.KeyD) {
		fox.Dir.X += 1
	}
	if rl.IsKeyDown(rl.KeyA) {
		fox.Dir.X -= 1
	}
	if rl.IsKeyDown(rl.KeyS) {
		fox.Dir.Y += 1
	}
	if rl.IsKeyDown(rl.KeyW) {
		fox.Dir.Y -= 1
	}
	fox.Dir = rl.Vector2Normalize(fox.Dir)

	fox.Pos.X += fox.Speed * fox.Dir.X * rl.GetFrameTime()
	fox.Pos.Y += fox.Speed * fox.Dir.Y * rl.GetFrameTime()


}

func (em *EnemyManager) UpdateCollisionBoxes(){
  for i := 0; i < len(em.Enemies); i++ {
    em.Enemies[i].CollisionBox.X = em.Enemies[i].Pos.X + ENEMY_CB_OFFSET_X
    em.Enemies[i].CollisionBox.Y = em.Enemies[i].Pos.Y
  }
}

func (em *EnemyManager) DrawEnemies() {
	for _, e := range em.Enemies {
		rl.DrawTexture(e.Textures["current"], int32(e.Pos.X), int32(e.Pos.Y), rl.White)
	}

}
