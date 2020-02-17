package shell

func (p *Program) runOn(v *VM) error {
	for _, pl := range p.pipelines {
		err := pl.runOn(v, p)
		if err != nil {
			return err
		}
	}
	return nil
}
