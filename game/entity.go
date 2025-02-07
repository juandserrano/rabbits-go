package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity interface {
	GetPosition() rl.Vector2
}
