package wooba

import (
	. "github.com/Panda-CRM/helpers-go"
	"log"
	"os"
	"testing"
	"time"
)

const (
	DATE_TEST_FORMAT    = "2006-01-02T15:04:05"
	ENV_WOOBA_URL       = "WOOBA_URL"
	ENV_WOOBA_USERLOGIN = "WOOBA_USERLOGIN"
	ENV_WOOBA_PASSWORD  = "WOOBA_PASSWORD"
	ENV_WOOBA_DATE_TEST = "WOOBA_DATE_TEST"
)

var (
	WOOBA_URL       = ""
	WOOBA_USERLOGIN = ""
	WOOBA_PASSWORD  = ""
	WOOBA_DATE_TEST = ""
)

func init() {
	getEnvIntegrationConfig()
}

func getEnvIntegrationConfig() {
	log.Print("[INTEGRATIONS] Lendo configurações da integração wooba")
	woobaUrl := os.Getenv(ENV_WOOBA_URL)
	woobaUserlogin := os.Getenv(ENV_WOOBA_USERLOGIN)
	woobaPassword := os.Getenv(ENV_WOOBA_PASSWORD)
	woobaDateTest := os.Getenv(ENV_WOOBA_DATE_TEST)
	if len(woobaUrl) > 0 {
		WOOBA_URL = woobaUrl
	}
	if len(woobaUserlogin) > 0 {
		WOOBA_USERLOGIN = woobaUserlogin
	}
	if len(woobaPassword) > 0 {
		WOOBA_PASSWORD = woobaPassword
	}
	if len(woobaDateTest) > 0 {
		WOOBA_DATE_TEST = woobaDateTest
	}
}

func TestWooba_VendasString(t *testing.T) {
	wooba := Wooba{
		Url:          WOOBA_URL,
		UserLogin:    WOOBA_USERLOGIN,
		UserPassword: WOOBA_PASSWORD,
	}
	date, _ := time.Parse(DATE_TEST_FORMAT, WOOBA_DATE_TEST)
	_, err := wooba.VendasString(date, date)
	Ok(t, err)
}

func TestWooba_CredencialInvalida(t *testing.T) {
	wooba := Wooba{
		Url:          WOOBA_URL,
		UserLogin:    "",
		UserPassword: "",
	}
	date, _ := time.Parse(DATE_TEST_FORMAT, WOOBA_DATE_TEST)
	_, err := wooba.VendasString(date, date)
	NotOk(t, err)
}
