package service

import (
	"strconv"

	"k8s.io/api/core/v1"
)

func transfer(svc *v1.Service) *Service {
	var nodePort int32
	if svc.Spec.Type == v1.ServiceTypeNodePort {
		for _, port := range svc.Spec.Ports {
			if port.NodePort > 0 {
				nodePort = port.NodePort
				break
			}
		}
	}

	service := NewService(svc.Namespace, svc.Name, strconv.Itoa(int(nodePort)))
	return service
}
