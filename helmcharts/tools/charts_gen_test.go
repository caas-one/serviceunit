package tools

import (
	"testing"
	"gitlab.yeepay.com/yce/helmcharts/yaml"
	"gitlab.yeepay.com/yce/helmcharts/build"
)

/*
func TestChartsGenerator(t *testing.T) {
	root := "./"	
	su := "bank"	
	name := "bankrouter-component-hessian"
	cg := NewChartsGenerator(su)
	app := yaml.NewApplication(name)
	err := cg.Do(root, su, name, app)
	if err != nil {
		t.Errorf("ChartsGenerator error: err=%s\n", err)
	}
}
*/

func TestChartsGeneratorAll(t *testing.T) {
	root := "./temp"
	for key, value := range build.ServiceMap {
		cg := NewChartsGenerator(key)
		for _, name := range value {
			app := yaml.NewApplication(name)
			err := cg.Do(root, key, name, app)
			if err != nil {
				t.Errorf("ChartsGenerator error: err=%s\n", err)
			}
		}
	}
}