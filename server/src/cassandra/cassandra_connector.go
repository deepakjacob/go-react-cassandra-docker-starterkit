package cassandra

import (
	"errors"
	"time"

	"github.com/gocql/gocql"
)

type Config struct {
	KeySpaceName    string
	Host            string
	TimeoutDuration time.Duration
}

type Cassandra struct {
	config  *Config
	session *gocql.Session
}

func (c *Cassandra) GetSession() *gocql.Session {
	return c.session
}

func (c *Cassandra) SetConfig(host string, keySpace string, timeout time.Duration) error {
	if host == "" || keySpace == "" || timeout == 0 {
		return errors.New("Missing: host, keyspace or timeout")
	}

	c.config = &Config{
		Host:            host,
		KeySpaceName:    keySpace,
		TimeoutDuration: timeout,
	}
	return nil
}

func (c *Cassandra) Open() (*Cassandra, error) {
	var err error

	cluster := gocql.NewCluster(c.config.Host)
	cluster.Keyspace = c.config.KeySpaceName
	cluster.Timeout = c.config.TimeoutDuration
	cluster.Consistency = gocql.Quorum
	cluster.DisableInitialHostLookup = true
	//TODO: think about having gocql password authenticator
	c.session, err = cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Cassandra) Close() error {
	c.session.Close()
	return nil
}
