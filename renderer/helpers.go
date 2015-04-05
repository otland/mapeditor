package renderer

import "errors"
import "github.com/go-gl/gl/v4.5-core/gl"

func LoadShader(code string, shaderType gl.GLenum) (gl.Shader, error) {
	shader := gl.CreateShader(shaderType)
	shader.Source(string(code))
	shader.Compile()

	if shader.Get(gl.COMPILE_STATUS) != 1 {
		return shader, errors.New(shader.GetInfoLog())
	}

	return shader, nil
}

func LoadProgram(shaders ...gl.Shader) (gl.Program, error) {
	p := gl.CreateProgram()
	for _, shader := range shaders {
		p.AttachShader(shader)
	}

	p.Link()
	if p.Get(gl.LINK_STATUS) != 1 {
		return p, errors.New(p.GetInfoLog())
	}

	p.Validate()
	if p.Get(gl.VALIDATE_STATUS) != 1 {
		return p, errors.New(p.GetInfoLog())
	}

	return p, nil
}
