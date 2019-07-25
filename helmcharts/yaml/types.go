package yaml

import (
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"

	appv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

// ResourceLimits for limits in resources block
type ResourceLimits struct {
	CPU    string `json:"cpu" yaml:"cpu"`
	Memory string `json:"memory" yaml:"memory"`
}

// ResourceRequests for request in resources block
type ResourceRequests struct {
	CPU    string `json:"cpu" yaml:"cpu"`
	Memory string `json:"memory" yaml:"memory"`
}

// Resources defines the container's resources
type Resources struct {
	ResourceLimits   `json:"limits,omitempty" yaml:"limits,omitempty"`
	ResourceRequests `json:"requests,omitempty" yaml:"requests,omitempty"`
}

// DebugImage defines the debug image on troubleshooting usage
type DebugImage struct {
	Repository string `json:"repository" yaml:"repository"`
	Tag        string `json:"tag" yaml:"tag"`
	PullPolicy string `json:"pullPolicy" yaml:"pullPolicy"`
}

// TroubleShooting define the debug images
type TroubleShooting struct {
	DebugImage DebugImage `json:"debugImage" yaml:"debugImage"`
	Resources  Resources  `json:"resources" yaml:"resources"`
}

// Service define the ports and type in k8s service
type Service struct {
	Type       string `json:"type" yaml:"type"`
	Port       string `json:"port" yaml:"port"`
	TargetPort string `json:"targetPort" yaml:"targetPort"`
	NodePort   string `json:"nodePort" yaml:"nodePort"`
}

// Environments define the env of the application(deployment)
// type Environments struct {
// 	[]Environment `json:""`
// }

// Environment define the name and value
type Environment struct {
	Name  string `json:"name" yaml:"name"`
	Value string `json:"value" yaml:"value"`
}

// Deployment defines some user-defined attribes in k8s deployment
type Deployment struct {
	Image     string        `json:"image" yaml:"image"`
	Replica   int32         `json:"replica" yaml:"replica"`
	Env       []Environment `json:"env" yaml:"env"`
	Resources Resources     `json:"resources" yaml:"resources"`
}

// Application define the application domain in a service unit
type Application struct {
	Name       string     `json:"name" yaml:"name"`
	Deployment Deployment `json:"deployment" yaml:"deployment"`
	Service    Service    `json:"service" yaml:"service"`
}

// NewApplication func
func NewApplication(name string) *Application {
	return &Application{
		Name: name,
		Deployment: Deployment{
			Env: make([]Environment, 0),
		},
	}
}

// ServiceUnit define a service unit associate with business model
type ServiceUnit struct {
	Name         string        `json:"name" yaml:"name"`
	Applications []Application `json:"applications" yaml:"applications"`
}

// NewServiceUnit func gives a new/empty ServiceUnit instance
func NewServiceUnit(name string) *ServiceUnit {
	return &ServiceUnit{
		Name:         name,
		Applications: make([]Application, 0),
	}
}

// Profile is the defination of all Application and Services witch wanted to be deployed
type Profile struct {
	Namespace       string          `json:"namespace" yaml:"namespace"`
	TroubleShooting TroubleShooting `json:"troubleShooting" yaml:"troubleShooting"`
	ServiceUnits    []ServiceUnit   `json:"serviceUnits" yaml:"serviceUnits"`
}

// NewProfile func give a new/empty instance of Profile
func NewProfile() *Profile {
	return &Profile{
		ServiceUnits: make([]ServiceUnit, 0),
	}
}

// MarshalToYaml func
func (p *Profile) MarshalToYaml() (string, error) {
	data, err := yaml.Marshal(p)
	if err != nil {
		log.Errorf("yaml.Marshal error: err=%s", err)
		return "", err
	}
	return string(data), nil
}

// UnmarshalFromYaml func
func (p *Profile) UnmarshalFromYaml(data []byte) error {
	err := yaml.Unmarshal(data, p)
	if err != nil {
		log.Errorf("yaml.Unmarshal error: err=%s", err)
		return err
	}
	return nil
}

// ExistServiceUnit func if the service unit exists
func (p *Profile) ExistServiceUnit(su string) bool {
	for _, unit := range p.ServiceUnits {
		if strings.EqualFold(unit.Name, su) {
			return true
		}
	}
	return false
}

// AddServiceUnit func
func (p *Profile) AddServiceUnit(su string) {
	unit := NewServiceUnit(su)
	p.ServiceUnits = append(p.ServiceUnits, *unit)
}

// AddApplication func
func (p *Profile) AddApplication(su, app string, dp *appv1.Deployment, svc *v1.Service) {
	application := NewApplication(app)

	// Deployment
	deployment(application, dp)

	// Service
	service(application, svc)

	if p.ExistServiceUnit(su) {
		// if service unit exist
		for index, unit := range p.ServiceUnits {
			if strings.EqualFold(unit.Name, su) {
				p.ServiceUnits[index].Applications = append(p.ServiceUnits[index].Applications, *application)
			}
		}
	} else {
		// if service unit do not exist
		unit := NewServiceUnit(su)
		unit.Applications = append(unit.Applications, *application)
		// log.Infof("=============Add ServiceUnit when Unit Not exists: unit=%s, appname=%s", unit.Name, application.Name)
		p.ServiceUnits = append(p.ServiceUnits, *unit)
	}
	return
}

func deployment(application *Application, dp *appv1.Deployment) {
	application.Name = dp.Name
	application.Deployment.Replica = *dp.Spec.Replicas
	for _, container := range dp.Spec.Template.Spec.Containers {
		if strings.EqualFold(dp.Name, container.Name) {
			application.Deployment.Image = container.Image
			for _, env := range container.Env {
				e := Environment{
					Name:  env.Name,
					Value: env.Value,
				}

				application.Deployment.Env = append(application.Deployment.Env, e)
			}
			application.Deployment.Resources = Resources{
				ResourceLimits: ResourceLimits{
					CPU:    container.Resources.Limits.Cpu().String(),
					Memory: container.Resources.Limits.Memory().String(),
				},
				ResourceRequests: ResourceRequests{
					CPU:    container.Resources.Requests.Cpu().String(),
					Memory: container.Resources.Requests.Memory().String(),
				},
			} // resources
		} // for container
	}
}

func service(application *Application, svc *v1.Service) {
	// application.Service.Type = svc.Spec.P
	application.Service.Type = string(svc.Spec.Type)
	if len(svc.Spec.Ports) == 1 {
		port := svc.Spec.Ports[0]
		application.Service.NodePort = strconv.Itoa(int(port.NodePort))
		application.Service.Port = strconv.Itoa(int(port.Port))
		application.Service.TargetPort = port.TargetPort.String()
	}
}
