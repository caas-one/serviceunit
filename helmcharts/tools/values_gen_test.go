package tools

import (
	"os"
	"testing"
	"text/template"

	"gitlab.yeepay.com/yce/helmcharts/build"
	"gitlab.yeepay.com/yce/helmcharts/yaml"
)

var values = Values{
	AppName:      "bankchannel-hessian",
	ReplicaCount: 3,
	Image: Image{
		Repository: "artifact.paas.yp/yeepay-docker-dev-local/bankchannel-hessian",
		Tag:        "201905131745_7888d381",
		PullPolicy: "Always",
	},
	DebugImage: Image{
		Repository: "artifact.paas.yp/yeepay-docker-dev-local/troubleshooting",
		Tag:        "201809271441",
		PullPolicy: "Always",
	},
	ImagePullSecrets: "myregistrykey",
	NameOverride:     "bankchannel-hessian",
	Env: []Environment{
		Environment{
			Name:  "YP_APP_NAME",
			Value: "bankchannel-hessian",
		},
		Environment{
			Name:  "YP_DATA_CENTER",
			Value: "CICD_DEFAULT",
		},
		Environment{
			Name:  "YP_DEPLOY_ENV",
			Value: "product",
		},
		Environment{
			Name:  "DUBBO_APPLICATION_ENVIRONMENT",
			Value: "product",
		},
		Environment{
			Name:  "YP_JVM_RESOURCE_CPU",
			Value: "2",
		},
		Environment{
			Name:  "YP_JVM_RESOURCE_MEMORY",
			Value: "4G",
		},
	},
	Resources: Resources{
		Limits: LimitsRequests{
			CPU:    "2",
			Memory: "4G",
		},
		Requests: LimitsRequests{
			CPU:    "200m",
			Memory: "2G",
		},
	},
	Service: Service{
		Type:     "NodePort",
		Port:     "8080",
		NodePort: "30387",
	},
}

func TestValuesGeneratorTemplate(t *testing.T) {
	temp := template.Must(template.New("values").Parse(ValuesTempl))
	err := temp.Execute(os.Stdout, values)
	if err != nil {
		t.Errorf("template Execute error: err=%s\n", err)
		return
	}
}

func TestValuesGenerator(t *testing.T) {
	root := "./temp"
	su := "bank"
	name := "bankrouter-component-hessian"
	vg := NewValuesGenerator(su)
	app := yaml.NewApplication(name)
	err := vg.Do(root, su, name, app)
	if err != nil {
		t.Errorf("ValuesGenerator error: err=%s\n", err)
	}
}

func TestValuesGeneratorAll(t *testing.T) {
	root := "./temp"
	for key, value := range build.ServiceMap {
		vg := NewValuesGenerator(key)
		for _, name := range value {
			app := yaml.NewApplication(name)
			err := vg.Do(root, key, name, app)
			if err != nil {
				t.Errorf("ChartsGenerator error: err=%s\n", err)
			}
		}
	}
}
