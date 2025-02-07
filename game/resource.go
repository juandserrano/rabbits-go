package game

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type ResourceManager struct {
	Textures map[string]rl.Texture2D
}

// var rm *ResourceManager

func InitResourceManager() ResourceManager {
	// if rm != nil {
	// 	rl.TraceLog(rl.LogWarning, "ResourceManager already initialized")
	// 	return nil
	// }
	rm := ResourceManager{
		Textures: map[string]rl.Texture2D{},
	}
	return rm
}

// func GetResourceManager() *ResourceManager {
// 	return rm
// }

func (g *Game) LoadTexture(filepath, name string) error {
	img := rl.LoadImage(filepath)
	if img == nil {
		return fmt.Errorf("error loading image [%s]:", filepath)
	}
	texture := rl.LoadTextureFromImage(img)
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
