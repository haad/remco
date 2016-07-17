package file

import (
	"github.com/HeavyHorst/remco/backends"
	"github.com/HeavyHorst/remco/log"
	"github.com/HeavyHorst/remco/template"
	"github.com/Sirupsen/logrus"
)

type Config struct {
	Filepath string
	template.Backend
}

func (c *Config) Connect() (backends.Store, error) {
	log.WithFields(logrus.Fields{
		"backend":  "file",
		"filepath": c.Filepath,
	}).Info("Set filepath")
	client, err := NewFileClient(c.Filepath)
	if err != nil {
		return backends.Store{}, err
	}
	c.Backend.StoreClient = client
	c.Backend.Name = "file"
	return backends.Store{
		Name:   "file",
		Client: client,
	}, nil
}
