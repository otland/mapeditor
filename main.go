package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"

	"github.com/go-gl/gl/v4.5-core/gl"
	glfw "github.com/go-gl/glfw/v3.1/glfw"

	"github.com/joseivan/mapeditor/client"
	"github.com/joseivan/mapeditor/ot"
	"github.com/joseivan/mapeditor/renderer"
)

func main() {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	fmt.Println("OpenTibia Map Editor")

	conf := &Configuration{}
	conf.Load("config.json")

	var sprLoader client.SpriteLoader
	var datLoader client.DatLoader
	var itemLoader ot.ItemLoader
	//var otbLoader ot.OtbLoader
	//var otMap ot.Map

	var group sync.WaitGroup
	group.Add(3)

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
	go func() {
		if err := itemLoader.LoadXML("items.xml"); err != nil {
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

	if glfw.Init() != nil {
		panic("Failed to initialize GLFW")
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 2)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

	window, err := glfw.CreateWindow(640, 480, "Map Editor", nil, nil)
	if err != nil {
		panic(err)
	}

	window.SetKeyCallback(keyHandler)

	window.MakeContextCurrent()

	if err := gl.Init(); err != nil {
		log.Fatal("Could not init gl")
	}

	log.Printf("OpenGL Version: %s", gl.GetString(gl.VERSION))
	log.Printf("GLSL Version: %s", gl.GetString(gl.SHADING_LANGUAGE_VERSION))
	log.Printf("Vendor: %s", gl.GetString(gl.VENDOR))
	log.Printf("Renderer: %s", gl.GetString(gl.RENDERER))

	r := renderer.Renderer{}

	r.Initialize()
	r.SetViewport(0, 0, 800, 600)

	for !window.ShouldClose() {
		r.Render()

		window.SwapBuffers()
		glfw.PollEvents()
	}

	/*
		ids := datLoader.GetSpriteIDs(420)
		sprite := sprLoader.GetSprite(ids[0])

		img := image.NewNRGBA(image.Rect(0, 0, 32, 32))
		img.Pix = sprite

		out, _ := os.Create("test.png")
		defer out.Close()

		if err := png.Encode(out, img); err != nil {
			log.Fatal(err)
		}*/
}

func keyHandler(window *glfw.Window, k glfw.Key, s int, action glfw.Action, mods glfw.ModifierKey) {
	if action != glfw.Press {
		return
	}

	switch glfw.Key(k) {
	case glfw.KeyEscape:
		window.SetShouldClose(true)
	}
}
