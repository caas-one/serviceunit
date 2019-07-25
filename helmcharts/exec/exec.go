package exec

import (
	"io/ioutil"
	"os/exec"
	"strings"

	mlog "github.com/maxwell92/log"
)

var log = mlog.Log

// Exec struct
type Exec struct {
	Path      string
	Namespace string
}

// NewExec func
func NewExec(path string) *Exec {
	return &Exec{
		Path: path,
	}
}

// Do func
func (e *Exec) Do(shell []string) error {
	cmd := exec.Command(shell[0], shell[1:]...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Errorf("cmd.StdoutPipe error: err=%s", err)
		return err
	}
	defer stdout.Close()
	if err := cmd.Start(); err != nil {
		log.Errorf("cmd.Start error: err=%s", err)
		return err
	}
	_, err = ioutil.ReadAll(stdout)
	if err != nil {
		log.Errorf("ioutil.ReadAll error: err=%s", err)
		return err
	}
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	return nil
}

// Install helm install command
func (e *Exec) Install(su, namespace, name string) []string {
	command := "/usr/local/bin/helm install " + e.Path + "/" + su + "/" + name + " --name=" + name + " --namespace=" + namespace
	return strings.Split(command, " ")
}

// Update helm update command
func (e *Exec) Update(su, namespace, name string) []string {
	command := "/usr/local/bin/helm update " + e.Path + "/" + su + "/" + name + " --name=" + name + " --namespace=" + namespace
	return strings.Split(command, " ")
}

// Clean func
func (e *Exec) Clean(profile, dir string) []string {
	command := "rm -fr " + profile + " " + dir
	log.Infof("Clean: profile=%s, dir=%s", profile, dir)
	return strings.Split(command, " ")
}

// Delete func
func (e *Exec) Delete(app string) []string {
	command := "/usr/local/bin/helm delete --purge " + app
	log.Infof("Helm delete with --purge: app=%s", app)
	return strings.Split(command, " ")
}
