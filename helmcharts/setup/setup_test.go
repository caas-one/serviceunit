package setup

import (
	"fmt"
	"io/ioutil"
	"testing"
)

// TestGenerate func
func TestGenerate(t *testing.T) {
	client, err := GetK8sClientFromFiles(APIServer, "ca.crt", "client.crt", "client.key")
	if err != nil {
		t.Errorf("GetK8sClientFromFiles error: err=%s", err)
	}

	var builder Builder
	profile, err := builder.Generate(client, "newspace")
	if err != nil {
		t.Errorf("builder.Generate error: err=%s\n", err)
		return
	}
	data, err := profile.MarshalToYaml()
	if err != nil {
		t.Errorf("")
	}
	// sync to file
	ioutil.WriteFile("profile.yaml", []byte(data), 0644)

	fmt.Printf("%s", string(data))

}
