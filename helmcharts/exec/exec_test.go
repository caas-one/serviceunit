package exec

import "testing"

func TestDo(t *testing.T) {
	su := "accounting"
	namespace := "newspace"
	name := "accountfront-manage-hessian"
	exec := NewExec("./")
	err := exec.Do(su, namespace, name)
	if err != nil {
		t.Errorf("exec.Do error: err=%s\n", err)
	}
}
