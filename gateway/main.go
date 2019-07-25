package main

import (
	"flag"
	"fmt"
	"os"

	mylog "github.com/maxwell92/gokits/log"
	"gitlab.yeepay.com/yce/nodeport/k8s"
	"gitlab.yeepay.com/yce/nodeport/k8s/service"
	"gitlab.yeepay.com/yce/nodeport/nginx"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/client-go/tools/cache"
)

// for flag opt
var (
	h         bool
	ca        string
	cert      string
	key       string
	apiserver string
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")

	flag.StringVar(&ca, "ca", "/etc/controller/dev/ca.crt", "the path of ca file")
	flag.StringVar(&cert, "cert", "/etc/controller/dev/client.crt", "the path of cert file")
	flag.StringVar(&key, "key", "/etc/controller/dev/client.key", "the path of key file")
	flag.StringVar(&apiserver, "apiserver", "10.151.160.73:6443", "the address and port of the apiserver")

	flag.Usage = usage
}

var log = mylog.Log

func main() {
	flag.Parse()
	log.Infof("apiserver=%s, ca=%s, cert=%s, key=%s", apiserver, ca, cert, key)
	client, err := k8s.GetK8sClientFromFiles(apiserver, ca, cert, key)
	if err != nil {
		log.Errorf("k8s.GetK8sClient error: err=%s", err)
		return
	}

	ch := make(chan struct{})

	listWatcher := cache.NewListWatchFromClient(client.CoreV1().RESTClient(), "services", v1.NamespaceAll, fields.Everything())

	informer := service.NewInformer(listWatcher, ch)
	go informer.Run()

	nginx := nginx.NewNginxController(client, ch)
	go nginx.Run()

	select {}
}

func usage() {
	fmt.Fprintf(os.Stderr, `nodeport version: nodeport/0.1.1
Usage: nodeport --ca=[caFile] --cert=[certFile] --key=[keyFile] --apiserver=[host:port]

Options:
`)
	flag.PrintDefaults()
}
