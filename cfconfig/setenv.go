package cfconfig

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/cloudfoundry-community/go-cfenv"
)

// SetEnvVars is a set of variable names and values
type SetEnvVars struct {
	EnvVars     map[string]string
	EnvPrefixes []string
}

// NewSetEnvVars creates SetEnvVars
func NewSetEnvVars(vcapServices *cfenv.Services) (envVars *SetEnvVars) {
	envVars = &SetEnvVars{
		EnvVars:     map[string]string{},
		EnvPrefixes: []string{},
	}

	for serviceName, serviceInstances := range *vcapServices {
		envVars.EnvPrefixes = append(envVars.EnvPrefixes, strings.ToUpper(serviceName))

		namePrefix := serviceName + "_"
		serviceInstance := serviceInstances[0]
		for credentialkey, credentialValue := range serviceInstance.Credentials {
			envKey := strings.ToUpper(namePrefix + credentialkey)
			envVars.EnvVars[envKey] = credentialValue
		}

	}
	return
}

// UpdateEnvVars updates target app
func (envVars *SetEnvVars) UpdateEnvVars(appName string) (err error) {
	for name, value := range envVars.EnvVars {
		err = envVars.setupEnvVars(appName, name, value)
		if err != nil {
			return
		}
	}
	return
}

func (envVars *SetEnvVars) setupEnvVars(appName, name, value string) (err error) {
	cmd := exec.Command("cf", "set-env", appName, name, value)
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	fmt.Println(out.String())
	return
}
