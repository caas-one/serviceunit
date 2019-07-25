package nginx

import (
	"strings"

	"gitlab.yeepay.com/yce/nodeport/k8s/node"
	"gitlab.yeepay.com/yce/nodeport/k8s/service"
	"k8s.io/client-go/kubernetes"
)

func transform(client *kubernetes.Clientset) *Nginx {
	sm := service.Instance().DeepCopy()
	nginx := NewNginx()
	for k, v := range sm.Map {
		// for each the namespaces, once namespace will match a server block
		server := NewServer()
		server.Namespace = k
		for _, service := range v {
			location := NewLocation()
			location.Path = trimSvcSuffix(service.Name)
			location.UpstreamName = metaNamespaceKeyFunc(service.Namespace, service.Name)
			server.Locations = append(server.Locations, *location)

			list, err := node.GetRandomReadyNodeList(client)
			if err != nil {
				log.Errorf("node.GetRandomReadyNodeList error: err=%s", err)
				continue
			}
			upstream := NewUpstreamReal(location.UpstreamName, list, service.Nodeport)
			nginx.Upstreams = append(nginx.Upstreams, *upstream)
		}
		nginx.Servers = append(nginx.Servers, *server)
	}

	// fmt.Printf("nginx: %v\n", nginx)
	return nginx
}

func metaNamespaceKeyFunc(namespace, name string) string {
	return namespace + "-" + name
}

func trimSvcSuffix(service string) string {
	return strings.Replace(service, "-svc", "", -1)
}
