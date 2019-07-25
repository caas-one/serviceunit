package tools

import "gitlab.yeepay.com/yce/helmcharts/yaml"

// GeneratorI Gernerator interface
type GeneratorI interface {
	Do(*yaml.Profile) error
}

// Generator struct
type Generator struct {
	Root string
}

// NewGenerator give a new/empty Generator instance
func NewGenerator(root string) *Generator {
	return &Generator{
		Root: root,
	}
}

// Do func
func (g *Generator) Do(p *yaml.Profile) error {
	for _, unit := range p.ServiceUnits {
		for _, app := range unit.Applications {
			cg := NewChartsGenerator(unit.Name)
			tg := NewTemplatesGenerator(unit.Name)
			vg := NewValuesGenerator(unit.Name, p.Namespace)

			err := cg.Do(g.Root, unit.Name, app.Name, &app)
			err = tg.Do(g.Root, unit.Name, app.Name, &app)
			err = vg.Do(g.Root, unit.Name, app.Name, &app)
			if err != nil {
				log.Errorf("Generator.Do error: err=%s", err)
				return err
			}
		}
	}
	return nil
}
