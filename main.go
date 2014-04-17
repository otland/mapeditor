package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/otland/mapeditor/client"
	"github.com/otland/mapeditor/ot"
)

func main() {
	fmt.Println("OpenTibia Map Editor")

	var sprLoader client.SpriteLoader
	var otbLoader ot.OtbLoader
	var otMap ot.Map

	var group sync.WaitGroup
	group.Add(2)

	go func() {
		if err := sprLoader.Load("data.spr"); err != nil {
			log.Fatal(err)
		}

		group.Done()
	}()

	otMap.Initialize()
	go func() {
		if err := otbLoader.Load("items.otb"); err != nil {
			log.Fatal(err)
		}

		if err := otMap.ReadOTBM("forgotten.otbm", &otbLoader); err != nil {
			log.Fatal(err)
		}

		group.Done()
	}()

	group.Wait()
}
