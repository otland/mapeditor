package renderer

import "errors"
import "github.com/go-gl/gl/v4.5-core/gl"
import "strings"

func LoadShader(code string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)
	csource := gl.Str(code)
	gl.ShaderSource(shader, 1, &csource, nil)
	gl.CompileShader(shader)	

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)	
	if status == gl.FALSE {

		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return shader, errors.New(log)
	}

	return shader, nil
}

func LoadProgram(shaders ...uint32) (uint32, error) {
	p := gl.CreateProgram()
	for _, shader := range shaders {
		gl.AttachShader(p, shader)
	}

	gl.LinkProgram(p)
	var status int32
	gl.GetProgramiv(p, gl.LINK_STATUS, &status)
	if status != gl.FALSE {
		
		var logLength int32
		gl.GetProgramiv(p, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(p, logLength, nil, gl.Str(log))
		return p, errors.New(log)
	}

	gl.ValidateProgram(p)
	gl.GetProgramiv(p, gl.VALIDATE_STATUS, &status)
	if status != gl.FALSE {

		var logLength int32
		gl.GetProgramiv(p, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(p, logLength, nil, gl.Str(log))
		return p, errors.New(log)	
	}

	return p, nil
}
