package shell

import (
	"gopkg.in/pipe.v2"
)

func (p *pipeline) runOn(v *VM, program *Program) error {
	pl := pipe.Line()
	return pipe.Run(pl)
}
