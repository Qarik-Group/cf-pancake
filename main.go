package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/codegangsta/cli"
)

// EnvVars is a set of variable names and values
type EnvVars map[string]string

func (vars *EnvVars) String() (out string) {
	for envVar, value := range *vars {
		out += fmt.Sprintf("export %s='%s'\n", envVar, value)
	}
	return
}

func pancakeCommandExports(c *cli.Context) {
	appEnv, err := cfenv.Current()
	if err != nil {
		fmt.Println("Requires $VCAP_SERVICES and $VCAP_APPLICATION to be set")
		fmt.Println(err)
		return
	}

	exportVars := EnvVars{}

	for serviceName, serviceInstances := range appEnv.Services {
		namePrefix := serviceName + "_"
		serviceInstance := serviceInstances[0]
		for credentialkey, credentialValue := range serviceInstance.Credentials {
			envKey := strings.ToUpper(namePrefix + credentialkey)
			exportVars[envKey] = credentialValue
		}

	}

	fmt.Print(&exportVars)
}

func main() {
	app := cli.NewApp()
	app.Name = "cf-pancake"
	app.Usage = "Flatten $VCAP_SERVICES into many environment variables"
	app.Commands = []cli.Command{
		{
			Name:      "exports",
			ShortName: "e",
			Usage:     "Output `export KEY=VALUE` to STDOUT based on local $VCAP_SERVICES",
			Action:    pancakeCommandExports,
		},
	}

	app.Run(os.Args)

}
