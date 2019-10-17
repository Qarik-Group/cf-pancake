package flatten

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/cloudfoundry-community/go-cfenv"
)

// EnvVars is a set of variable names and values
type EnvVars map[string]string

func (vars *EnvVars) String() (out string) {
	for envVar, value := range *vars {
		out += fmt.Sprintf("export %s='%s'\n", envVar, value)
	}
	return
}

// VCAPServices converts $VCAP_SERVICES into a mapping of flattened credentials
func VCAPServices(vcapServices *cfenv.Services) (exportVars EnvVars) {
	exportVars = EnvVars{}
	if vcapServices == nil {
		return
	}

	for serviceName, serviceInstances := range *vcapServices {
		namePrefix := serviceName + "_"
		serviceInstance := serviceInstances[0]
		for credentialKey, credentialValue := range serviceInstance.Credentials {
			if strValue, ok := credentialValue.(string); ok {
				exportVars[cleanEnvVarName(namePrefix+credentialKey)] = strValue
				exportVars[cleanEnvVarName(serviceInstance.Name+"_"+credentialKey)] = strValue

				for _, tag := range serviceInstance.Tags {
					exportVars[cleanEnvVarName(tag+"_"+credentialKey)] = strValue
				}

			}
		}

	}
	return exportVars
}

func cleanEnvVarName(envVar string) string {
	keyToUnderscoreRE := regexp.MustCompile(`[^A-Za-z0-9]+`)
	envVar = keyToUnderscoreRE.ReplaceAllString(strings.ToUpper(envVar), "_")
	nonLetterPrefixRE := regexp.MustCompile(`^[^a-zA-Z]`)
	if nonLetterPrefixRE.MatchString(envVar) {
		envVar = "_" + envVar
	}
	return envVar
}
