package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Pos          rl.Vector2
	LastPos      rl.Vector2
	Dir          rl.Vector2
	Textures     map[string]rl.Texture2D
	Speed        float32
	CollisionBox rl.Rectangle
}

var p Player

func InitPlayer() *Player {
	p = Player{
		Pos:   rl.Vector2{X: PLAYER_START_POS_X, Y: PLAYER_START_POS_Y},
		Speed: PLAYER_START_SPEED,
	}
	p.CollisionBox = rl.Rectangle{X: p.Pos.X, Y: p.Pos.Y, Width: 64, Height: 64}
	p.LoadTextures()
	return &p
}

func (p *Player) Move() {
	p.handleInput()
	p.LastPos = p.Pos
	p.selectTexture()
}

func (p Player) GetPosition() rl.Vector2 {
	return p.Pos
}

func (p *Player) selectTexture() {
	if p.Dir.X > 0 {
		p.Textures["current"] = p.Textures["right"]
	} else if p.Dir.X < 0 {
		p.Textures["current"] = p.Textures["left"]
	}
}

func (p *Player) handleInput() {
	p.Dir = rl.Vector2{X: 0, Y: 0}
	if rl.IsKeyDown(rl.KeyRight) {
		p.Dir.X += 1
	}
	if rl.IsKeyDown(rl.KeyLeft) {
		p.Dir.X -= 1
	}
	if rl.IsKeyDown(rl.KeyDown) {
		p.Dir.Y += 1
	}
	if rl.IsKeyDown(rl.KeyUp) {
		p.Dir.Y -= 1
	}
	p.Dir = rl.Vector2Normalize(p.Dir)
	p.Pos.X += p.Speed * p.Dir.X * rl.GetFrameTime()
	p.Pos.Y += p.Speed * p.Dir.Y * rl.GetFrameTime()
}

func (p *Player) LoadTextures() {
	img := rl.LoadImage("resources/textures/main-rabbit.png")
	if img == nil {
		panic(1)
	}
	p.Textures = map[string]rl.Texture2D{}
	rl.ImageResize(img, 64, 64)
	p.Textures["right"] = rl.LoadTextureFromImage(img)
	rl.ImageFlipHorizontal(img)
	p.Textures["left"] = rl.LoadTextureFromImage(img)
	rl.UnloadImage(img)
	p.Textures["current"] = p.Textures["right"]
}

func (p *Player) Draw() {
	rl.DrawTexture(p.Textures["current"], int32(p.Pos.X), int32(p.Pos.Y), rl.White)
}
