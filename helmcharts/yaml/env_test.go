package yaml

import (
	"fmt"
	"testing"

	"gopkg.in/yaml.v2"
)

var data = `
env:
- name: YP_APP_NAME
  value: bankchannel-hessian
- name: YP_DATA_CENTER
  value: CICD_DEFAULT
- name: YP_DEPLOY_ENV
  value: default
- name: DUBBO_APPLICATION_ENVIRONMENT
  value: default
- name: YP_JVM_RESOURCE_CPU
  value: "2"
- name: YP_JVM_RESOURCE_MEMORY
  value: 4G`

type Env struct {
	Env []Environment `json:"env" yaml:"env"`
}

func NewEnv() *Env {
	return &Env{
		Env: make([]Environment, 0),
	}
}

func TestENVUnmarshal(t *testing.T) {
	e := NewEnv()
	err := yaml.Unmarshal([]byte(data), &e)
	if err != nil {
		t.Errorf("yaml.Unmarshal error: err=%s\n", err)
		return
	}

	for _, v := range e.Env {
		fmt.Printf("key=%s, value=%s\n", v.Name, v.Value)
	}

	result, err := yaml.Marshal(e)
	if err != nil {
		t.Errorf("yaml.Marshal error: err=%s", err)
		return
	}

	fmt.Printf("yaml.Marshal result: \n%s\n", result)
}
