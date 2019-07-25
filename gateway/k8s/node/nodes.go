package node

import (
	"errors"

	mlog "github.com/maxwell92/log"
	v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var log = mlog.Log

// Node defines the node struct
type Node struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
}

// List give the available nodes (Ready)
type List []Node

// GetRandomReadyNodeList return a subset of all available nodes, less than the half of the total ready nodes
func GetRandomReadyNodeList(client *kubernetes.Clientset) (List, error) {
	list, err := GetAllReadyNodeList(client)
	if err != nil {
		log.Errorf("GetAllReadyNodeList error: err=%s", err)
		return nil, err
	}

	if len(list) == 0 {
		return nil, errors.New("No available nodes")
	}
	subset := make(List, 0)
	for index, node := range list {
		if index <= len(list)/2 {
			subset = append(subset, node)
		}
	}
	return subset, nil
}

// GetAllReadyNodeList return all of the Ready Nodes via NodeList, otherwise, return a nil obj and an error
func GetAllReadyNodeList(client *kubernetes.Clientset) (List, error) {
	nodes, err := client.CoreV1().Nodes().List(meta_v1.ListOptions{})
	if err != nil {
		log.Errorf("Get NodeList error: err=%s", err)
		return nil, err
	}

	list := make(List, 0)
	for _, n := range nodes.Items {
		ready := false
		// Ready nodes
		for _, condition := range n.Status.Conditions {
			if v1.NodeReady == condition.Type {
				// log.Infof("Node is ready: name=%s, condition=%s", n.Name, condition.Type)
				ready = true
			}
		}
		if !ready {
			continue
		}
		// Address
		for _, addr := range n.Status.Addresses {
			if addr.Type == v1.NodeInternalIP {
				var node Node
				node.Name = n.Name
				node.IP = addr.Address
				list = append(list, node)
			}
		}
	}
	return list, err
}
