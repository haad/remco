package consul

import (
	"github.com/HeavyHorst/remco/backends"
	"github.com/HeavyHorst/remco/log"
	"github.com/HeavyHorst/remco/template"
	"github.com/Sirupsen/logrus"
	"github.com/kelseyhightower/confd/backends/consul"
)

type Config struct {
	Nodes        []string
	Scheme       string
	ClientCert   string `toml:"client_cert"`
	ClientKey    string `toml:"client_key"`
	ClientCaKeys string `toml:"client_ca_keys"`
	template.Backend
}

func (c *Config) Connect() (backends.Store, error) {
	log.WithFields(logrus.Fields{
		"backend": "consul",
		"nodes":   c.Nodes,
	}).Info("Set backend nodes")
	client, err := consul.New(c.Nodes, c.Scheme, c.ClientCert, c.ClientKey, c.ClientCaKeys)
	if err != nil {
		return backends.Store{}, err
	}
	c.Backend.StoreClient = client
	c.Backend.Name = "consul"
	return backends.Store{
		Name:   c.Backend.Name,
		Client: client,
	}, nil
}
