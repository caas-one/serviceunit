package yaml

import (
	"fmt"
	"testing"
)

/*
var config = Profile{
	Namespaces: newspace,
	TroubleShooting: TroubleShooting{
		DebugImage: DebugImage{
			Repository: artifact.paas.yp/yeepay-docker-dev-local/troubleshooting,
		}
	}
}
*/

func TestParse(t *testing.T) {
	parser := NewParser()
	err := parser.Parse("../profile.yaml")
	if err != nil {
		t.Errorf("parser.Parse error: err=%s\n", err)
	}

	fmt.Printf("%v\n", parser)
}
