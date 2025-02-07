package game

import (
	"fmt"
	"os"
	"path"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/lafriks/go-tiled"
)

type Map struct {
	Map *tiled.Map
	// Renderer *render.Renderer
	Texture rl.Texture2D
}

func InitMaps() map[string]Map {
	maps := map[string]Map{}

	// Create maps and renderer
	files, err := os.ReadDir(TILEMAP_DIR)
	if err != nil {
		fmt.Printf("error getting tilemap directory entries: %s", err.Error())
	}
	for _, file := range files {
		if path.Ext(file.Name()) == ".tmx" {
			err = addToMaps(maps, TILEMAP_DIR+"/"+file.Name())
			if err != nil {
				fmt.Printf("error adding to maps: %s", err.Error())
			}
		}
	}
	return maps
}

func addToMaps(maps map[string]Map, filepath string) error {
	filename := strings.TrimSuffix(path.Base(filepath), ".tmx")
	gameMap, err := tiled.LoadFile(filepath)
	if err != nil {
		return err
	}
	tsTexture := rl.LoadTexture(gameMap.Tilesets[0].Image.Source)

	// renderer, err := render.NewRenderer(gameMap)
	// if err != nil {
	// 	return err
	// }
	m := Map{
		// Renderer: renderer,
		Map:     gameMap,
		Texture: tsTexture,
	}
	maps[filename] = m
	return nil
}
