package setup

import (
	"io/ioutil"
	"os"

	mlog "github.com/maxwell92/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var log = mlog.Log

// GetK8sClient return a kubernetes client instance
func GetK8sClient(apiServer string, ca, cert, key []byte) (*kubernetes.Clientset, error) {
	cfg := &rest.Config{
		Host: "https://" + apiServer,
		TLSClientConfig: rest.TLSClientConfig{
			CAData:   ca,
			CertData: cert,
			KeyData:  key,
		},
	}

	client, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Errorf("kubernetes NewForConfig error: err=%s", err)
		return nil, err
	}

	return client, nil
}

// GetK8sClientFromFiles read ca, cert, key files
func GetK8sClientFromFiles(apiServer, ca, cert, key string) (*kubernetes.Clientset, error) {
	cafile, err := os.Open(ca)
	certfile, err := os.Open(cert)
	keyfile, err := os.Open(key)
	if err != nil {
		log.Fatalf("os.Open file error: err=%s", err)
	}

	caData, err := ioutil.ReadAll(cafile)
	certData, err := ioutil.ReadAll(certfile)
	keyData, err := ioutil.ReadAll(keyfile)

	if err != nil {
		log.Fatalf("ioutil.ReadAll file error: err=%s", err)
	}

	return GetK8sClient(apiServer, caData, certData, keyData)

}
