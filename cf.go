// Package bluemix implements facilities for
// dealing with IBM bluemix environment.
package bluemix

import (
	"encoding/json"
	"os"
)

// vCAPService represents the contents of Cloud Foundry's VCAP_SERVICES environment variable.
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

func (vs VCAPServices) mongoDB() *MongoDBSvc {
	v := vs["mongodb-2.2"]
	if v != nil && len(v) > 0 {
		var m MongoDBSvc
		err := json.Unmarshal(v[0].Credentials, &m)
		if err != nil {
			return nil
		}
		return &m
	}

	return nil
}

// AppServices is the description of the services available to an application
// running in bluemix.
var AppServices VCAPServices

// MongoService returns the description of the MongoDB service
// available to an application running in bluemix.
// In case that there is no MongoDB service available it returns nil.
func MongoService() *MongoDBSvc {
	return AppServices.mongoDB()
}

func init() {
	AppServices := newVCAPServices()
	vs := os.Getenv("VCAP_SERVICES")
	json.Unmarshal([]byte(vs), &AppServices)
}
