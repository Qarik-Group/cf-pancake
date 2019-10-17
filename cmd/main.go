package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cloudfoundry-community/go-cfenv"
	"github.com/codegangsta/cli"
	"github.com/starkandwayne/cf-pancake/cfconfig"
	"github.com/starkandwayne/cf-pancake/flatten"
)

func pancakeCommandExports(c *cli.Context) {
	varsAndValues := exportVars(c)
	fmt.Print(&varsAndValues)
}

func pancakeCommandEnvVarList(c *cli.Context) {
	varsAndValues := exportVars(c)
	for envVar := range varsAndValues {
		fmt.Println(envVar)
	}
}

func exportVars(c *cli.Context) flatten.EnvVars {
	appEnv, err := cfenv.Current()
	if err != nil {
		fmt.Println("Requires $VCAP_SERVICES and $VCAP_APPLICATION to be set")
		log.Fatal(err)
	}

	return flatten.VCAPServices(&appEnv.Services)
}

func pancakeCommandSetEnv(c *cli.Context) {
	appName := c.Args().First()
	if appName == "" {
		fmt.Println("USAGE: cf-pancake set-env APPNAME")
		return
	}

	configPath, err := cfconfig.DefaultCfConfigPath()
	if err != nil {
		log.Fatal(err)
	}
	config, err := cfconfig.LoadCfConfig(configPath)
	if err != nil {
		log.Fatal(err)
	}
	if config.SpaceFields.GUID == "" {
		fmt.Println("Not targeting a space. Run `cf target -o ORG -s SPACE` first.")
	}

	appFindURL := fmt.Sprintf("/v2/apps?q=space_guid:%s&q=name:%s", config.SpaceFields.GUID, appName)
	resources, err := cfconfig.CurlGETResources(appFindURL)
	if err != nil {
		log.Fatal(err)
	}
	if len(resources.Resources) == 0 {
		log.Fatalf("No application '%s' found in current org/space", appName)
	}
	appURL := resources.Resources[0].Metadata.URL
	appEnv, err := cfconfig.CurlAppEnv(appURL)
	if err != nil {
		log.Fatal(err)
	}

	setEnvVars, err := cfconfig.NewSetEnvVars(appName, appEnv)
	if err != nil {
		log.Fatal(err)
	}

	err = setEnvVars.UpdateEnvVars()
	if err != nil {
		log.Fatal(err)
	}
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
		{
			Name:   "envvars",
			Usage:  "The list of environment variables to be produced from $VCAP_SERVICES [used for testing]",
			Action: pancakeCommandEnvVarList,
		},
		{
			Name:      "set-env",
			ShortName: "s",
			Usage:     "Updates `cf set-env` for an application based on its bound services",
			Action:    pancakeCommandSetEnv,
		},
	}

	app.Run(os.Args)

}
