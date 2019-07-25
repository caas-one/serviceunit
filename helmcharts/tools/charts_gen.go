package tools

import "gitlab.yeepay.com/yce/helmcharts/yaml"

// ChartsGenerator generate the Charts.yaml
type ChartsGenerator struct {
	Path string `json:"path"`
}

// NewChartsGenerator gives a new ChartsGenerator instance
func NewChartsGenerator(path string) *ChartsGenerator {
	return &ChartsGenerator{
		Path: path,
	}
}

// Do func generates the Charts.yaml
func (cg *ChartsGenerator) Do(root, su, name string, app *yaml.Application) error {
	cf := NewChartsFile(root, su, name)
	chart := NewCharts(name)
	cf.Render(chart, ChartsTempl)
	return nil
}
