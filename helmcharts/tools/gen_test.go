package tools

import (
	"testing"

	"gitlab.yeepay.com/yce/helmcharts/yaml"
)

func TestGenerator(t *testing.T) {
	root := "./tmp"
	parser := yaml.NewParser()
	err := parser.Parse("./profile.yaml")
	if err != nil {

	}
	g := NewGenerator(root)
	err = g.Do(&parser.Profile)
	if err != nil {
		t.Errorf("Generator Do error: err=%s\n", err)
	}
}
