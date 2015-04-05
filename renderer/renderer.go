package renderer

import "github.com/go-gl/gl/v4.5-core/gl"

type Renderer struct {
	prog gl.Program
	vao  gl.VertexArray
	vbo  gl.Buffer

	mvmLocation gl.UniformLocation

	vertices []mapVertex
}

type mapVertex struct {
	x float32
	y float32
	z float32

	r float32
	g float32
	b float32
	a float32

	s float32
	t float32
}

func (r *Renderer) Initialize() {
	r.vertices = []mapVertex{
		{200, 200, 0, 1, 1, 1, 1, 0, 0},
		{218, 200, 0, 1, 0, 0, 1, 0, 0},
	}

	vs, err := LoadShader(vertexShader, gl.VERTEX_SHADER)
	if err != nil {
		panic(err)
	}

	gs, err := LoadShader(geometryShader, gl.GEOMETRY_SHADER)
	if err != nil {
		panic(err)
	}

	fs, err := LoadShader(fragmentShader, gl.FRAGMENT_SHADER)
	if err != nil {
		panic(err)
	}

	gl.ClearColor(0, 1, 0, 0)

	// Create the VAO
	// GL 3+ allows us to store the vertex layout in a vertex array object (VAO).
	r.vao = gl.GenVertexArray()
	r.vao.Bind()

	// Create VBO
	r.vbo = gl.GenBuffer()
	r.vbo.Bind(gl.ARRAY_BUFFER)
	gl.BufferData(gl.ARRAY_BUFFER, len(r.vertices)*4*9, r.vertices, gl.DYNAMIC_DRAW)

	prog, err := LoadProgram(vs, gs, fs)

	if err != nil {
		panic(err)
	}

	prog.Use()

	r.prog = prog
	r.mvmLocation = r.prog.GetUniformLocation("modelview_matrix")

	pos := r.prog.GetAttribLocation("in_position")
	color := r.prog.GetAttribLocation("in_color")
	texcoord := r.prog.GetAttribLocation("in_texcoord")

	// FIXME: size math
	pos.AttribPointer(3, gl.FLOAT, false, 9*4, nil)
	color.AttribPointer(4, gl.FLOAT, false, 9*4, uintptr(3*4))
	texcoord.AttribPointer(2, gl.FLOAT, false, 9*4, uintptr(7*4))

	pos.EnableArray()
	color.EnableArray()
	texcoord.EnableArray()
}

func (r *Renderer) SetViewport(x, y, w, h float32) {
	m := Matrix4{}
	m.Ortho(x, x+w, y+h, y, -1.0, 1.0)
	fl := (*[16]float32)(&m)

	r.mvmLocation.UniformMatrix4f(false, fl)
}

func (r *Renderer) Render() {
	gl.Clear(gl.COLOR_BUFFER_BIT)

	r.prog.Use()

	r.vao.Bind()
	r.vbo.Bind(gl.ARRAY_BUFFER)

	gl.DrawArrays(gl.POINTS, 0, len(r.vertices))

	// FIXME: unbind vao/vbo/prog?
}
