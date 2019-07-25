package cmd

import (
	"io/ioutil"
	"strings"

	"gitlab.yeepay.com/yce/helmcharts/exec"
	"gitlab.yeepay.com/yce/helmcharts/setup"
	"gitlab.yeepay.com/yce/helmcharts/tools"
	"gitlab.yeepay.com/yce/helmcharts/yaml"
)

func setupProfile(apiServer, ca, cert, key, path, srcNamespace, dstNamespace string) error {
	client, err := setup.GetK8sClient(apiServer, []byte(ca), []byte(cert), []byte(key))
	if err != nil {
		log.Errorf("GetK8sClient error: err=%s", err)
	}

	var builder setup.Builder
	profile, err := builder.Generate(client, srcNamespace, dstNamespace)
	if err != nil {
		log.Errorf("builder.Generate error: err=%s\n", err)
		return err
	}
	data, err := profile.MarshalToYaml()
	if err != nil {
		log.Errorf("profile.MarshalToYaml error: err=%s", err)
		return err
	}

	// sync to file
	err = ioutil.WriteFile("profile.yaml", []byte(data), 0644)
	if err != nil {
		log.Errorf("ioutils.WriteFile error: err=%s", err)
		return err
	}

	return nil
}

// Build profile
func buildProfile(root, profile string) error {
	parser := yaml.NewParser()
	err := parser.Parse(profile)
	g := tools.NewGenerator(root)
	err = g.Do(&parser.Profile)
	if err != nil {
		log.Errorf("Build Charts.yaml/values.yaml/templates file error: err=%s", err)
	}

	log.Infof("Generate the Charts.yaml/values.yaml/templates file successfully.")
	return nil
}

// Install All Helm charts
func installAllHelm(root, profile string) error {
	parser := yaml.NewParser()
	err := parser.Parse(profile)
	if err != nil {
		log.Fatalf("parser.Parse profile error: err=%s", err)
	}
	exec := exec.NewExec(root)
	for _, su := range parser.Profile.ServiceUnits {
		for _, app := range su.Applications {
			shell := exec.Install(su.Name, parser.Profile.Namespace, app.Name)
			log.Infof("InstallAllCharts: root=%s, profile=%s, su=%s, app=%s, shell=%v",
				root, profile, su.Name, app.Name, shell)
			err := exec.Do(shell)
			if err != nil {
				log.Errorf("Helm install error: err=%s", err)
			}
		}
	}
	return nil
}

// Install charts in a given service unit
func installSuHelm(root, profile, unit string) error {
	parser := yaml.NewParser()
	err := parser.Parse(profile)
	if err != nil {
		log.Fatalf("parser.Parse profile error: err=%s", err)
	}
	exec := exec.NewExec(root)
	for _, su := range parser.Profile.ServiceUnits {
		if !strings.EqualFold(su.Name, unit) {
			continue
		}
		for _, app := range su.Applications {
			shell := exec.Install(su.Name, parser.Profile.Namespace, app.Name)
			log.Infof("InstallAllCharts: root=%s, profile=%s, su=%s, app=%s, shell=%v",
				root, profile, su.Name, app.Name, shell)
			err := exec.Do(shell)
			if err != nil {
				log.Errorf("Helm install error: err=%s", err)
			}
		}
	}
	return nil
}

// Install charts in a given application
func installAppHelm(root, profile, unit, application string) error {
	parser := yaml.NewParser()
	err := parser.Parse(profile)
	if err != nil {
		log.Fatalf("parser.Parse profile error: err=%s", err)
	}
	exec := exec.NewExec(root)
	for _, su := range parser.Profile.ServiceUnits {
		if !strings.EqualFold(su.Name, unit) {
			continue
		}
		for _, app := range su.Applications {
			if !strings.EqualFold(app.Name, application) {
				continue
			}
			shell := exec.Install(su.Name, parser.Profile.Namespace, app.Name)
			log.Infof("InstallAllCharts: root=%s, profile=%s, su=%s, app=%s, shell=%v",
				root, profile, su.Name, app.Name, shell)
			err := exec.Do(shell)
			if err != nil {
				log.Errorf("Helm install error: err=%s", err)
			}
		}
	}
	return nil
}

// Update all helm charts
func updateAllHelm(root, profile string) error {
	parser := yaml.NewParser()
	err := parser.Parse(profile)
	if err != nil {
		log.Fatalf("parser.Parse profile error: err=%s", err)
	}
	exec := exec.NewExec(root)
	for _, su := range parser.Profile.ServiceUnits {
		for _, app := range su.Applications {
			shell := exec.Update(su.Name, parser.Profile.Namespace, app.Name)
			log.Infof("UpdateAllCharts: root=%s, profile=%s, su=%s, app=%s, shell=%v",
				root, profile, su.Name, app.Name, shell)
			err := exec.Do(shell)
			if err != nil {
				log.Errorf("Helm install error: err=%s", err)
			}
		}
	}
	return nil

}

// update charts in a given service unit
func updateSuHelm(root, profile, unit string) error {
	parser := yaml.NewParser()
	err := parser.Parse(profile)
	if err != nil {
		log.Fatalf("parser.Parse profile error: err=%s", err)
	}
	exec := exec.NewExec(root)
	for _, su := range parser.Profile.ServiceUnits {
		if !strings.EqualFold(su.Name, unit) {
			continue
		}
		for _, app := range su.Applications {
			shell := exec.Update(su.Name, parser.Profile.Namespace, app.Name)
			log.Infof("Update serviceUnit: su=%s, app=%s, shell=%s", su.Name, app.Name, shell)
			err := exec.Do(shell)
			if err != nil {
				log.Errorf("Helm install error: err=%s", err)
			}
		}
	}
	return nil
}

// update all helm charts
func updateAppHelm(root, profile, unit, application string) error {
	parser := yaml.NewParser()
	err := parser.Parse(profile)
	if err != nil {
		log.Fatalf("parser.Parse profile error: err=%s", err)
	}
	exec := exec.NewExec(root)
	for _, su := range parser.Profile.ServiceUnits {
		if !strings.EqualFold(su.Name, unit) {
			continue
		}
		for _, app := range su.Applications {
			if !strings.EqualFold(app.Name, application) {
				continue
			}
			shell := exec.Update(su.Name, parser.Profile.Namespace, app.Name)
			log.Infof("Update application: su=%s, app=%s, shell=%s", su.Name, application, shell)
			err := exec.Do(shell)
			if err != nil {
				log.Errorf("Helm install error: err=%s", err)
				return err
			}
		}
	}
	return nil
}

// clean: remove the profile.yaml and output directory
func clean(output, profile string) error {
	exec := exec.NewExec(output)
	shell := exec.Clean(profile, output)
	log.Infof("Shell: output=%s, profile=%s", output, profile)
	err := exec.Do(shell)
	if err != nil {
		log.Errorf("Bash rm -fr error: profile=%s, dir=%s", profile, output)
		return err
	}
	return nil
}

// Update all helm charts
func deleteAllHelm(root, profile string) error {
	parser := yaml.NewParser()
	err := parser.Parse(profile)
	if err != nil {
		log.Fatalf("parser.Parse profile error: err=%s", err)
	}
	exec := exec.NewExec(root)
	for _, su := range parser.Profile.ServiceUnits {
		for _, app := range su.Applications {
			shell := exec.Delete(app.Name)
			log.Infof("DeleteAllCharts: root=%s, profile=%s, su=%s, app=%s, shell=%v",
				root, profile, su.Name, app.Name, shell)
			err := exec.Do(shell)
			if err != nil {
				log.Errorf("Helm delete error: err=%s", err)
			}
		}
	}
	return nil

}

// update charts in a given service unit
func deleteSuHelm(root, profile, unit string) error {
	parser := yaml.NewParser()
	err := parser.Parse(profile)
	if err != nil {
		log.Fatalf("parser.Parse profile error: err=%s", err)
	}
	exec := exec.NewExec(root)
	for _, su := range parser.Profile.ServiceUnits {
		log.Infof("Delete Su: su=%s, unit=%s", su.Name, unit)
		if !strings.EqualFold(su.Name, unit) {
			continue
		}
		for _, app := range su.Applications {
			shell := exec.Delete(app.Name)
			log.Infof("Delete serviceUnit: su=%s, app=%s, shell=%s", su.Name, app.Name, shell)
			err := exec.Do(shell)
			if err != nil {
				log.Errorf("Helm delete error: err=%s", err)
			}
		}
	}
	return nil
}

// update all helm charts
func deleteAppHelm(root, profile, unit, application string) error {
	parser := yaml.NewParser()
	err := parser.Parse(profile)
	if err != nil {
		log.Fatalf("parser.Parse profile error: err=%s", err)
	}
	exec := exec.NewExec(root)
	for _, su := range parser.Profile.ServiceUnits {
		if !strings.EqualFold(su.Name, unit) {
			continue
		}
		for _, app := range su.Applications {
			if !strings.EqualFold(app.Name, application) {
				continue
			}
			shell := exec.Delete(app.Name)
			log.Infof("Delete application: su=%s, app=%s, shell=%s", su.Name, application, shell)
			err := exec.Do(shell)
			if err != nil {
				log.Errorf("Helm delete error: err=%s", err)
				return err
			}
		}
	}
	return nil
}
