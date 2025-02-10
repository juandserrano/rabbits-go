package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ResourceManager struct {
	Textures map[string]rl.Texture2D
}

func InitResourceManager(g *Game) ResourceManager {
	rm := ResourceManager{
		Textures: map[string]rl.Texture2D{},
	}
  g.Rm = rm
  g.LoadTexture("resources/textures/fox-front.png", "fox")

	return rm
}

func (g *Game) LoadTexture(filepath, name string) error {
	img := rl.LoadImage(filepath)
	if img == nil {
		return fmt.Errorf("error loading image [%s]:", filepath)
	}
  rl.ImageResize(img, 64, 64)
	texture := rl.LoadTextureFromImage(img)
  rl.UnloadImage(img)
	err := g.Rm.AddTexture(texture, name)
	if err != nil {
		return err
	}

	return nil
}

func (rm *ResourceManager) AddTexture(texture rl.Texture2D, name string) error {
	_, ok := rm.Textures[name]
	if ok {
		return fmt.Errorf("texture with name [%s] already exist. Texture not added to ResourceManager", name)
	}
	rm.Textures[name] = texture
	return nil
}

func (rm *ResourceManager) RemoveTexture(name string) error {
	val, ok := rm.Textures[name]
	if !ok {
		return fmt.Errorf("texture with name [%s] does not exist.", name)
	}
	rl.UnloadTexture(val)
	delete(rm.Textures, name)
	return nil
}
