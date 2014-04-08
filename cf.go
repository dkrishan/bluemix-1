package bluemix

import (
	"encoding/json"
	"os"
)

type VCAPService struct {
	Name        string   `json:"name"`
	Label       string   `json:"label,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Plan        string   `json:"plan,omitempty"`
	Credentials json.RawMessage
}

type VCAPServices map[string][]VCAPService

func newVCAPServices() VCAPServices {
	return VCAPServices(make(map[string][]VCAPService))
}

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

var AppServices VCAPServices

func MongoService() *MongoDBSvc {
	return AppServices.mongoDB()
}

func init() {
	AppServices := newVCAPServices()
	vs := os.Getenv("VCAP_SERVICES")
	json.Unmarshal([]byte(vs), &AppServices)
}
