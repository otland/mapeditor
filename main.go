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
	var datLoader client.DatLoader
	var otbLoader ot.OtbLoader
	var otMap ot.Map

	var group sync.WaitGroup
	group.Add(3)

	go func() {
		if err := sprLoader.Load("data.spr"); err != nil {
			log.Fatal(err)
		}

		group.Done()
	}()
	go func() {
		if err := datLoader.Load("data.dat"); err != nil {
			log.Fatal(err)
		}

		group.Done()
	}()

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

	group.Wait()
}
