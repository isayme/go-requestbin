package mongo

import (
	"time"

	"github.com/globalsign/mgo"
	logger "github.com/isayme/go-logger"
	"github.com/isayme/go-requestbin/app/conf"
)

// Session mongo session
type Session struct {
	config  *conf.Mongo
	session *mgo.Session
}

// NewSession new mongo session
func NewSession(config *conf.Mongo) (*Session, error) {
	dailInfo := &mgo.DialInfo{
		Addrs:          config.Addrs,
		Timeout:        time.Duration(config.Timeout) * time.Second,
		Database:       config.Database,
		ReplicaSetName: config.ReplicaSetName,
		Source:         config.Source,
		PoolLimit:      config.PoolLimit,
		Username:       config.Username,
		Password:       config.Password,
	}

	session, err := mgo.DialWithInfo(dailInfo)
	if err != nil {
		return nil, err
	}
	logger.Debugf("mongodb %v connected", config.Addrs)

	return &Session{
		config:  config,
		session: session,
	}, nil
}

func (s *Session) Copy() *Session {
	return &Session{
		config:  s.config,
		session: s.session.Copy(),
	}
}

func (s *Session) GetCollection(name string) *mgo.Collection {
	return s.session.DB("").C(name)
}

func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}
