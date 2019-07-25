package tools

import (
	"testing"
	"gitlab.yeepay.com/yce/helmcharts/yaml"
	"gitlab.yeepay.com/yce/helmcharts/build"
)

func TestTemplatesGenerator(t *testing.T) {
	root := "./temp"	
	su := "bank"	
	name := "bankrouter-component-hessian"
	tg := NewTemplatesGenerator(su)
	app := yaml.NewApplication(name)
	err := tg.Do(root, su, name, app)
	if err != nil {
		t.Errorf("ChartsGenerator error: err=%s\n", err)
	}
}

func TestTemplatesGeneratorAll(t *testing.T) {
	root := "./temp"
	for key, value := range build.ServiceMap {
		tg := NewTemplatesGenerator(key)
		for _, name := range value {
			app := yaml.NewApplication(name)
			err := tg.Do(root, key, name, app)
			if err != nil {
				t.Errorf("ChartsGenerator error: err=%s\n", err)
			}
		}
	}

}