package main

import (
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"sync"

	"github.com/otland/mapeditor/client"
)

func main() {
	fmt.Println("OpenTibia Map Editor")

	conf := &Configuration{}
	conf.Load("config.json")

	var sprLoader client.SpriteLoader
	var datLoader client.DatLoader
	//var otbLoader ot.OtbLoader
	//var otMap ot.Map

	var group sync.WaitGroup
	group.Add(2)

	go func() {
		if err := sprLoader.Load(conf.SprFile); err != nil {
			log.Fatal(err)
		}

		group.Done()
	}()
	go func() {
		if err := datLoader.Load(conf.DatFile); err != nil {
			log.Fatal(err)
		}

		group.Done()
	}()
	/*
		go func() {
			if err := otbLoader.Load("items.otb"); err != nil {
				log.Fatal(err)
			}

			otMap.Initialize()
			if err := otMap.ReadOTBM("forgotten.otbm", &otbLoader); err != nil {
				log.Fatal(err)
			}

			group.Done()
		}()
	*/
	group.Wait()

	ids := datLoader.GetSpriteIDs(420)
	sprite := sprLoader.GetSprite(ids[0])

	img := image.NewNRGBA(image.Rect(0, 0, 32, 32))
	img.Pix = sprite

	out, _ := os.Create("test.png")
	defer out.Close()

	if err := png.Encode(out, img); err != nil {
		log.Fatal(err)
	}
}
