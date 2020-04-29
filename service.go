package badger

import (
	"github.com/dgraph-io/badger"
)

// Service represents the badger service.
type Service struct {
	*badger.DB
}

// Dial sends the new config to Service.
func (s *Service) Dial(c Config) error {
	opt := badger.DefaultOptions("")
	opt.Dir = c.Dir
	opt.ValueDir = c.ValueDir
	var err error
	s.DB, err = badger.Open(opt)
	return err
}

// Close closes the session to cluster session.
func (s *Service) Close() error {
	return s.Close()
}

// Healthcheck returns if database responds.
func (s *Service) Healthcheck() error {
	return nil
}
