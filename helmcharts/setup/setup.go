package setup

import (
	"strings"

	"gitlab.yeepay.com/yce/helmcharts/yaml"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Builder can genterate the profile.yaml
type Builder struct {
}

// Generate func
func (b *Builder) Generate(client *kubernetes.Clientset, srcNamespace, dstNamespace string) (*yaml.Profile, error) {
	profile := yaml.NewProfile()
	profile.Namespace = dstNamespace

	list, err := client.AppsV1().Deployments(srcNamespace).List(meta_v1.ListOptions{})
	if err != nil {
		log.Errorf("client.AppsV1().Deployments.List error: err=%s", err)
		return nil, err
	}

	for _, dp := range list.Items {
		// If deployment(application) in ServiceUnit
		su := GetServiceUnitName(dp.Name)
		if !strings.EqualFold("", su) {
			svc, err := client.CoreV1().Services(dp.Namespace).Get(dp.Name+"-svc", meta_v1.GetOptions{})
			if err != nil {
				log.Errorf("client.CoreV1().Server.Get error: err=%s", err)
				continue
			}
			profile.AddApplication(su, dp.Name, &dp, svc)
		}

	}

	profile.TroubleShooting = *TroubleShooting()
	return profile, nil
}

// GetServiceUnitName func get the serviceUnitName of the application
func GetServiceUnitName(name string) string {
	for key, value := range ServiceMap {
		for _, app := range value {
			if strings.EqualFold(app, name) {
				return key
			}
		}
	}

	return ""
}

// TroubleShooting func
func TroubleShooting() *yaml.TroubleShooting {
	return &yaml.TroubleShooting{
		DebugImage: yaml.DebugImage{
			Repository: DebugImageRepository,
			Tag:        DebugImageTag,
			PullPolicy: ImagePullPolicy,
		},
		Resources: yaml.Resources{
			ResourceLimits: yaml.ResourceLimits{
				CPU:    ResourceLimitsCPU,
				Memory: ResourceLimitsMemory,
			},
			ResourceRequests: yaml.ResourceRequests{
				CPU:    ResourceRequestsCPU,
				Memory: ResourceRequestsMemory,
			},
		},
	}
}
