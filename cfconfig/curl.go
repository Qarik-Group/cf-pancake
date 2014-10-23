package cfconfig

import (
	"bytes"
	"encoding/json"
	"os/exec"

	"github.com/cloudfoundry-community/go-cfenv"
)

// APIGETResource represents the structure of every `cf curl -X GET` command result
type APIGETResource struct {
	TotalResults int        `json:"total_results"`
	TotalPages   int        `json:"total_pages"`
	PrevURL      string     `json:"prev_url"`
	NextURL      string     `json:"next_url"`
	Resources    []resource `json:"resources"`
}

type resource struct {
	Metadata metadata               `json:"metadata"`
	Entity   map[string]interface{} `json:"entity"`
}

type metadata struct {
	GUID string `json:"guid"`
	URL  string `json:"url"`
}

type AppEnv struct {
	Staging     map[string]interface{} `json:"staging_env_json"`
	Running     map[string]interface{} `json:"running_env_json"`
	Environment map[string]interface{} `json:"environment_json"`
	System      map[string]interface{} `json:"system_env_json"`
}

// CurlGETResources performs a `cf curl URL` command and marshals the JSON response
func CurlGETResources(url string) (response *APIGETResource, err error) {
	cmd := exec.Command("cf", "curl", url)
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		return
	}
	response = &APIGETResource{}
	json.Unmarshal(out.Bytes(), response)

	return
}

// CurlAppEnv performs a `cf curl URL/env` command and marshals the JSON response
func CurlAppEnv(appURL string) (response *AppEnv, err error) {
	cmd := exec.Command("cf", "curl", appURL+"/env")
	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		return
	}
	response = &AppEnv{}
	json.Unmarshal(out.Bytes(), response)

	return
}

// VCAPServices extracts the system env variable `VCAP_SERVICES`
func (appEnv *AppEnv) VCAPServices() (services *cfenv.Services, err error) {
	str, err := json.Marshal(appEnv.System["VCAP_SERVICES"])
	if err != nil {
		return
	}
	services = &cfenv.Services{}
	json.Unmarshal([]byte(str), services)
	return
}
