package renderer

const vertexShader = `
#version 330

uniform mat4 modelview_matrix;

in vec3 in_position;
in vec4 in_color;
in vec2 in_texcoord;

out vec4 geom_color;
out vec4 geom_texcoord;

void main(void)
{
	gl_Position = modelview_matrix * vec4(in_position.xy, 1, 1);
	geom_texcoord = vec4(in_texcoord, 0, 0);

	geom_color = in_color;
}
`

const geometryShader = `
#version 330

uniform mat4 modelview_matrix;

layout(points) in;
layout(triangle_strip, max_vertices = 4) out;

in vec4 geom_color[];
in vec4 geom_texcoord[];
out vec4 texcoord;
out vec4 color;

void main(void)
{
    int i;

    for(i=0; i < gl_in.length(); i++)
	{
        vec4 in_pos = gl_in[i].gl_Position;
		vec4 in_texture = geom_texcoord[i];

		color = geom_color[i];

        gl_Position = in_pos;
        texcoord = in_texture;
	    EmitVertex();

        gl_Position = in_pos + modelview_matrix * vec4(0, 32, 0, 0);
        texcoord = vec4(in_texture.x, in_texture.y + 32, 0, 0);
	    EmitVertex();

        gl_Position = in_pos + modelview_matrix * vec4(32, 0, 0, 0);
        texcoord = vec4(in_texture.x + 32, in_texture.y, 0, 0);
	    EmitVertex();

        gl_Position = in_pos + modelview_matrix * vec4(32, 32, 0, 0);
        texcoord = vec4(in_texture.x + 32, in_texture.y + 32, 0, 0);
	    EmitVertex();

        EndPrimitive();
	}
}
`

const fragmentShader = `
#version 330

//uniform sampler2D texture;

in vec4 texcoord;
in vec4 color;
out vec4 fragColor;

void main(void)
{
    fragColor = color; //texture2D(texture, texcoord.st) * color;
}
`
