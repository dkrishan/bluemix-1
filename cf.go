// Package bluemix implements facilities for
// dealing with IBM bluemix environment.
package bluemix

import (
	"encoding/json"
	"errors"
	"os"
)

// VCAPService represents the contents of Cloud Foundry's VCAP_SERVICES environment variable.
type VCAPServices map[string][]VCAPService

// VCAPService represents the contents of an entry in Cloud Foundry's VCAP_SERVICES environment variable.
type VCAPService struct {
	Name        string   `json:"name"`
	Label       string   `json:"label,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Plan        string   `json:"plan,omitempty"`
	Credentials json.RawMessage
}

func newVCAPServices() VCAPServices {
	return VCAPServices(make(map[string][]VCAPService))
}

type VCAPApplication struct {
	InstanceId  string `json:"instance_id"`
	InstanceIdx int    `json:"instance_index"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	StartedAt   string `json:"started_at"`
	Limits      struct {
		Memory int `json:"mem"`
	} `json:"limits"`
	Name string   `json:"name"`
	URIs []string `json:"uris"`
}

// MongoDBSvc is the description of the MongoDB service in the bluemix environment.
type MongoDBSvc struct {
	Hostname string `json:"hostname"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Db       string `json:"db"`
	URL      string `json:"url"`
}

func (vs VCAPServices) mongoDB() (*MongoDBSvc, error) {
	if v, ok := vs["mongodb-2.2"]; ok {
		var m MongoDBSvc
		err := json.Unmarshal(v[0].Credentials, &m)
		if err != nil {
			return nil, err
		}
		return &m, nil
	}
	return nil, errors.New("No mongodb service available")
}

var (
	// AppServices is the description of the services available to an application
	// running in bluemix.
	AppServices VCAPServices
	Application VCAPApplication
)

// MongoService returns the description of the MongoDB service
// available to an application running in bluemix.
// In case that there is no MongoDB service available it returns nil.
func MongoService() (*MongoDBSvc, error) {
	return AppServices.mongoDB()
}

func init() {
	AppServices = newVCAPServices()
	vs := os.Getenv("VCAP_SERVICES")
	json.Unmarshal([]byte(vs), &AppServices)
	vapp := os.Getenv("VCAP_APPLICATION")
	json.Unmarshal([]byte(vapp), &Application)
}
