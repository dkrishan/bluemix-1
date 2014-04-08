package bluemix

import (
	"encoding/json"
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
	v := newVCAPServices()
	err := json.Unmarshal(VCAP_SERVICES, &v)
	if err != nil {
		t.Fatal(err)
	}
	mongo := v.mongoDB()
	if mongo == nil {
		t.Fatal("unexpected configuration")
	}
	expurl := "mongodb://be879069-b273-4656-b5fb-3daa5c508044:f268582e-0a52-42a8-9b97-66889a9cb662@10.0.116.49:10001/db"
	if mongo.URL != expurl {
		t.Errorf("wanted %s got %s", mongo.URL, expurl)
	}
	if mongo.Port != 10001 {
		t.Errorf("wanted %v got %v", mongo.Port, 10001)
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
   "mongodb-2.2": [
      {
         "name": "mongodb-76b24",
         "credentials": {
            "hostname": "10.0.116.49",
            "host": "10.0.116.49",
            "port": 10001,
            "username": "be879069-b273-4656-b5fb-3daa5c508044",
            "password": "f268582e-0a52-42a8-9b97-66889a9cb662",
            "name": "76ea370c-8678-4c51-b3cf-a0cd722ed93a",
            "db": "db",
            "url": "mongodb://be879069-b273-4656-b5fb-3daa5c508044:f268582e-0a52-42a8-9b97-66889a9cb662@10.0.116.49:10001/db"
         }
      }
   ]
}
`)
