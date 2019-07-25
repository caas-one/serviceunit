package tools

import (
	"gitlab.yeepay.com/yce/helmcharts/yaml"
)

// TemplateGenerator generate the template file: ingress.yaml/service.yaml/deployment.yaml
type TemplatesGenerator struct {
	Path string `json:"path"`
}

// NewTemplatesGenerator gives a new instannce
func NewTemplatesGenerator(path string) *TemplatesGenerator {
	return &TemplatesGenerator{
		Path: path,
	}
}

// Do generate the ingress.yaml/service.yaml/deployment.yaml
func (tg *TemplatesGenerator) Do(root, su, name string, app *yaml.Application) error {
	tf := NewTemplatesFile(root, su, name)
	err := tf.Sync()
	return err
}
