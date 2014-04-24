package bluemix

import (
	"encoding/json"
	"log"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	v := newVCAPServices()
	err := json.Unmarshal(VCAP_SERVICES, &v)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMongoService(t *testing.T) {
	mongo, err := MongoService()
	if err != nil {
		log.Println(AppServices)
		t.Fatal(err)
	}
	expurl := "mongodb://4d24a6cf-8509-48b2-8937-1c2544de6f10:a5a93a60-83be-492b-98f8-9b930226c3e9@50.23.230.141:10165/db"
	if mongo.URL != expurl {
		t.Errorf("wanted %s got %s", mongo.URL, expurl)
	}
	if mongo.Port != 10165 {
		t.Errorf("wanted %v got %v", mongo.Port, 10001)
	}
}

func TestVCAPApplication(t *testing.T) {
	var vapp VCAPApplication
	err := json.Unmarshal(VCAP_APPLICATION, &vapp)
	if err != nil {
		t.Fatal(err)
	}
	if vapp.Limits.Memory != 512 {
		t.Errorf("Limits: wanted %d, got %d", 512, vapp.Limits.Memory)
	}
	if Application.Limits.Memory != 512 {
		t.Errorf("Limits: wanted %d, got %d", 512, Application.Limits.Memory)
	}
}

var VCAP_SERVICES = []byte(`
{
  "elephantsql": [
    {
      "name": "elephantsql-c6c60",
      "label": "elephantsql",
      "tags": [
        "postgres",
        "postgresql",
        "relational"
      ],
      "plan": "turtle",
      "credentials": {
        "uri": "postgres://seilbmbd:PHxTPJSbkcDakfK4cYwXHiIX9Q8p5Bxn@babar.elephantsql.com:5432/seilbmbd"
      }
    }
  ],
  "sendgrid": [
    {
      "name": "mysendgrid",
      "label": "sendgrid",
      "tags": [
        "smtp"
      ],
      "plan": "free",
      "credentials": {
        "hostname": "smtp.sendgrid.net",
        "username": "QvsXMbJ3rK",
        "password": "HCHMOYluTv"
      }
    }
  ],
"mongodb-2.2":[{"name":"mongodb-ye","label":"mongodb-2.2","tags":[],"plan":"100","credentials":{"hostname":"50.23.230.141","host":"50.23.230.141","port":10165,"username":"4d24a6cf-8509-48b2-8937-1c2544de6f10","password":"a5a93a60-83be-492b-98f8-9b930226c3e9","name":"e8a0108e-3c4d-4657-b9d4-6c8d6a6da27e","db":"db","url":"mongodb://4d24a6cf-8509-48b2-8937-1c2544de6f10:a5a93a60-83be-492b-98f8-9b930226c3e9@50.23.230.141:10165/db"}}] 
}
`)

var VCAP_APPLICATION = []byte(`{"instance_id":"451f045fd16427bb99c895a2649b7b2a","instance_index":0,"host":"0.0.0.0","port":61857,"started_at":"2013-08-12 00:05:29 +0000","started_at_timestamp":1376265929,"start":"2013-08-12 00:05:29 +0000","state_timestamp":1376265929,"limits":{"mem":512,"disk":1024,"fds":16384},"application_version":"c1063c1c-40b9-434e-a797-db240b587d32","application_name":"styx-james","application_uris":["styx-james.a1-app.cf-app.com"],"version":"c1063c1c-40b9-434e-a797-db240b587d32","name":"styx-james","uris":["styx-james.a1-app.cf-app.com"],"users":null}
`)
