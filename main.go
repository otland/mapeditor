package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	fmt.Println("OpenTibia Map Editor")

	var sprLoader SpriteLoader

	// Load data files
	var group sync.WaitGroup
	group.Add(1)
	go func() {
		if err := sprLoader.load("data.spr"); err != nil {
			log.Fatal(err)
		}
		group.Done()
	}()
	group.Wait()

	sprLoader.close()
}
