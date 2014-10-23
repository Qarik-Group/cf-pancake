package main

import (
	"fmt"
	"os"

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
		fmt.Println(err)
		return
	}
	fmt.Println("Services:", appEnv.Services)

	exportVars := &EnvVars{
		"PG_PORT": "4000",
		"PG_HOST": "10.10.3.3",
	}
	fmt.Print(exportVars)
}

func main() {
	app := cli.NewApp()
	app.Name = "cf-pancake"
	app.Usage = "Flatten $VCAP_SERVICES into many environment variables"
	app.Action = pancakeCommandExports

	app.Run(os.Args)

}
