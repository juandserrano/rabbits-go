package game

import rl "github.com/gen2brain/raylib-go/raylib"

type Camera struct {
	*rl.Camera2D
}

func InitCamera(player *Player) Camera {
	return Camera{
		&rl.Camera2D{
			Target:   player.Pos,
			Offset:   rl.Vector2{X: DEFAULT_CAMERA_OFFSET_X, Y: DEFAULT_CAMERA_OFFSET_Y},
			Zoom:     DEFAULT_CAMERA_ZOOM,
			Rotation: DEFAULT_CAMERA_ROTATION,
		},
	}
}

func (c *Camera) MoveTo(entity Entity) {
	c.Target = entity.GetPosition()
}
