package tools

import (
	"errors"
	"strings"

	"gitlab.yeepay.com/yce/helmcharts/setup"
	"gitlab.yeepay.com/yce/helmcharts/yaml"
)

// ValuesGenerator generate the Values.yaml
type ValuesGenerator struct {
	Path      string `json:"path"`
	Namespace string `json:"path"`
}

// NewValuesGenerator gives a new/empty instannce
func NewValuesGenerator(path, namespace string) *ValuesGenerator {
	return &ValuesGenerator{
		Path:      path,
		Namespace: namespace,
	}
}

// Do func generate the Values.yaml
func (vg *ValuesGenerator) Do(root, su, name string, app *yaml.Application) error {
	values := vg.trans(app)
	vf := NewValuesFile(root, su, name)
	vf.Render(values, ValuesTempl)
	return nil
}

func (vg *ValuesGenerator) trans(app *yaml.Application) *Values {
	values := NewValues()
	values.AppName = app.Name
	values.ReplicaCount = app.Deployment.Replica

	repository, tag, err := split(app.Deployment.Image)
	if err != nil {
		log.Errorf("Split deployment images error: err=%s", err)
	}
	values.Namespace = vg.Namespace
	values.Image.Repository = repository
	values.Image.Tag = tag
	values.Image.PullPolicy = setup.ImagePullPolicy
	values.DebugImage.Repository = setup.DebugImageRepository
	values.DebugImage.Tag = setup.DebugImageTag
	values.DebugImage.PullPolicy = setup.ImagePullPolicy
	values.ImagePullSecrets = setup.ImagePullSecrets
	values.NameOverride = app.Name
	values.FullnameOverride = ""
	for _, env := range app.Deployment.Env {
		environment := Environment{
			Name:  env.Name,
			Value: env.Value,
		}
		values.Env = append(values.Env, environment)
	}
	values.Service.Type = app.Service.Type
	values.Service.Port = app.Service.Port
	values.Service.NodePort = app.Service.NodePort

	values.Resources.Limits.CPU = app.Deployment.Resources.ResourceLimits.CPU
	values.Resources.Limits.Memory = app.Deployment.Resources.ResourceLimits.Memory
	values.Resources.Requests.CPU = app.Deployment.Resources.ResourceRequests.CPU
	values.Resources.Requests.Memory = app.Deployment.Resources.ResourceRequests.Memory

	return values
}

func split(image string) (string, string, error) {
	ss := strings.Split(image, ":")
	if 2 == len(ss) {
		return ss[0], ss[1], nil
	}
	return "", "", errors.New("Images does not match repository:tag format")

}
