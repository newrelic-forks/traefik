package nrgoagent

import (
	"net/http"

	"github.com/newrelic/go-agent"
	"github.com/outbrain/golib/log"
)

type Newrelic struct {
	Application *newrelic.Application
	Transaction *newrelic.Transaction
}

func NewConfig(applicationName string, licenseKey string) newrelic.Config {
	return newrelic.NewConfig(applicationName, licenseKey)
}
func New(config newrelic.Config) (*Newrelic, error) {
	app, err := newrelic.NewApplication(config)
	return &Newrelic{Application: &app}, err
}

func (n *Newrelic) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	txn := ((*n.Application).StartTransaction(r.URL.Path, rw, r)).(newrelic.Transaction)
	if err := (*n.Transaction).SetName(r.Host); err != nil {
		log.Errorf("Error setting NR transaction name: %s", err.Error())
	}
	n.Transaction = &txn
	defer (*n.Transaction).End()

	next(rw, r)
}
