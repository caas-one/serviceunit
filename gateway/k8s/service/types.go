package service

import (
	"strings"
	"sync"
)

// Service Object
type Service struct {
	Namespace  string            `json:"namespace"`
	Name       string            `json:"name"`
	Nodeport   string            `json:"nodePort"`
	Annotation map[string]string `json:"annotation"`
}

// NewService new a service instance
func NewService(namespace, name, nodePort string) *Service {
	return &Service{
		Namespace:  namespace,
		Name:       name,
		Nodeport:   nodePort,
		Annotation: make(map[string]string),
	}
}

// Services container Service in the same namespace
type Services []*Service

// ServiceMap used to store the
type ServiceMap struct {
	Map map[string]Services // key: namespace-name ====> *Service
	sync.Mutex
}

// NewServiceMap return a new ServiceMap instance
func NewServiceMap() *ServiceMap {
	return &ServiceMap{
		Map: make(map[string]Services),
	}
}

var instance *ServiceMap
var once sync.Once

// Instance returns singleton instance
func Instance() *ServiceMap {
	once.Do(func() {
		instance = new(ServiceMap)
		instance.Map = make(map[string]Services)
	})
	return instance
}

func init() {
	Instance()
}

// Add func used to add a service into servicemap
func (sm *ServiceMap) Add(namespace string, service *Service) {
	sm.Lock()
	defer sm.Unlock()
	services := sm.Map[namespace]
	services = append(services, service)
	sm.Map[namespace] = services
	log.Infof("Add a service: namespace=%s, name=%s, len(servicesSlice)=%d, len(serviceMap)=%d",
		service.Namespace, service.Name, len(services), len(sm.Map))
}

// Update func used to update the servicemap
func (sm *ServiceMap) Update(namespace string, service *Service) {
	sm.Lock()
	defer sm.Unlock()
	services := sm.Map[namespace]
	for index, svc := range services {
		if strings.EqualFold(svc.Name, service.Name) {
			services[index] = service
			log.Infof("Update a service: namespace=%s, name=%s, len(servicesSlice)=%d, len(serviceMap)=%d",
				service.Namespace, service.Name, len(services), len(sm.Map))
		}
	}
}

// Del func used to delete a service from servicemap
func (sm *ServiceMap) Del(namespace, name string) {
	sm.Lock()
	defer sm.Unlock()
	services := sm.Map[namespace]
	for index, svc := range services {
		if strings.EqualFold(svc.Name, name) {
			services = append(services[:index], services[index+1:]...)
			sm.Map[namespace] = services
			log.Infof("Update a service: namespace=%s, name=%s, len(servicesSlice)=%d, len(serviceMap)=%d",
				namespace, name, len(services), len(sm.Map))
		}
	}
}

// DeepCopy func copy all data to a new instance
func (sm *ServiceMap) DeepCopy() *ServiceMap {
	newSm := NewServiceMap()
	sm.Lock()
	defer sm.Unlock()
	for k, v := range sm.Map {
		services := make(Services, 0)
		services = append(services, v...)
		newSm.Map[k] = services
	}
	return newSm
}
