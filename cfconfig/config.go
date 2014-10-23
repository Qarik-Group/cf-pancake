package cfconfig

import (
	"encoding/json"
	"io/ioutil"
	"os/user"
	"path/filepath"
)

// CfConfig presents the ~/.cf/config.json configuration for `cf` CLI of the current user
type CfConfig struct {
	Target             string
	AccessToken        string   `json:"AccessToken"`
	RefreshToken       string   `json:"RefreshToken"`
	OrganizationFields nameGUID `json:"OrganizationFields"`
	SpaceFields        nameGUID `json:"SpaceFields"`
}

type nameGUID struct {
	Name string `json:"Name"`
	GUID string `json:"Guid"`
}

// DefaultCfConfigPath returns the current user's ~/.cf/config.json file path
func DefaultCfConfigPath() (path string, err error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Abs(usr.HomeDir + "/.cf/config.json")

}

// LoadCfConfig loads a cf CLI config file; such as ~/.cf/config.json
func LoadCfConfig(configPath string) (cfconfig *CfConfig, err error) {
	cfconfig = &CfConfig{}
	contents, err := ioutil.ReadFile(configPath)
	if err != nil {
		return cfconfig, err
	}
	json.Unmarshal(contents, cfconfig)
	return
}
