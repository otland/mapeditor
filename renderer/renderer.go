package renderer

import "github.com/go-gl/gl/v4.5-core/gl"

type Renderer struct {
	prog uint32
	vao  uint32
	vbo  uint32

	mvmLocation int32

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
	gl.GenVertexArrays(1, &r.vao)
	gl.BindVertexArray(r.vao)

	// Create VBO
	gl.GenBuffers(1, &r.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, r.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(r.vertices)*4*9, gl.Ptr(r.vertices), gl.DYNAMIC_DRAW)

	prog, err := LoadProgram(vs, gs, fs)

	if err != nil {
		panic(err)
	}

	gl.UseProgram(prog)

	r.prog = prog
	r.mvmLocation = gl.GetUniformLocation(r.prog, gl.Str("modelview_matrix"))

	pos := uint32(gl.GetAttribLocation(r.prog, gl.Str("in_position")))
	color := uint32(gl.GetAttribLocation(r.prog, gl.Str("in_color")))
	texcoord := uint32(gl.GetAttribLocation(r.prog, gl.Str("in_texcoord")))

	gl.EnableVertexAttribArray(pos)
	gl.EnableVertexAttribArray(color)
	gl.EnableVertexAttribArray(texcoord)

	// FIXME: size math
	gl.VertexAttribPointer(pos, 3, gl.FLOAT, false, 9*4, nil)
	gl.VertexAttribPointer(color, 4, gl.FLOAT, false, 9*4, gl.PtrOffset(3*4))
	gl.VertexAttribPointer(texcoord, 2, gl.FLOAT, false, 9*4, gl.PtrOffset(7*4))

}

func (r *Renderer) SetViewport(x, y, w, h float32) {
	m := Matrix4{}
	m.Ortho(x, x+w, y+h, y, -1.0, 1.0)
	fl := (*[16]float32)(&m)
	
	gl.UniformMatrix4fv(r.mvmLocation, 1, false, &fl[0])
}

func (r *Renderer) Render() {
	gl.Clear(gl.COLOR_BUFFER_BIT)

	gl.UseProgram(r.prog)

	gl.BindVertexArray(r.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, r.vbo)

	//TODO Verify force cast in length
	gl.DrawArrays(gl.POINTS, 0, int32(len(r.vertices)))

	// FIXME: unbind vao/vbo/prog?
}
