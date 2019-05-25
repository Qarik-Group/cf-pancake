package cfconfig

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// SetEnvVars is a set of variable names and values
type SetEnvVars struct {
	AppName         string
	AppEnv          *AppEnv
	RequiredEnvVars map[string]string
}

// NewSetEnvVars creates SetEnvVars
func NewSetEnvVars(appName string, appEnv *AppEnv) (envVars *SetEnvVars, err error) {
	envVars = &SetEnvVars{
		AppName:         appName,
		AppEnv:          appEnv,
		RequiredEnvVars: map[string]string{},
	}

	err = envVars.discoverEnvVars()

	return
}

// UpdateEnvVars updates target app
func (envVars *SetEnvVars) UpdateEnvVars() (err error) {
	if err != nil {
		return
	}
	for name, value := range envVars.RequiredEnvVars {
		err = envVars.cfSetEnv(name, value)
		if err != nil {
			return
		}
	}
	return
}

func (envVars *SetEnvVars) discoverEnvVars() (err error) {
	vcapServices, err := envVars.AppEnv.VCAPServices()
	if err != nil {
		return
	}

	for serviceName, serviceInstances := range *vcapServices {
		namePrefix := serviceName + "_"
		serviceInstance := serviceInstances[0]
		for credentialKey, credentialValue := range serviceInstance.Credentials {
			if strValue, ok := credentialValue.(string); ok {
				envKey := strings.ToUpper(namePrefix + credentialKey)
				envVars.RequiredEnvVars[envKey] = strValue
			}
		}
	}
	return
}

func (envVars *SetEnvVars) existingEnvVars() map[string]interface{} {
	return envVars.AppEnv.Environment
}

func (envVars *SetEnvVars) cfSetEnv(name, value string) (err error) {
	cmd := exec.Command("cf", "set-env", envVars.AppName, name, value)
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	fmt.Println(out.String())
	return
}
