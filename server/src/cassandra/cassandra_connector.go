package cassandra

import (
	"errors"
	"fmt"
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

	//TODO: think about having gocql password authenticator
	c.session, err = cluster.CreateSession()
	if err != nil {
		return nil, err
	}
	fmt.Println("Cassandra init done...")

	return c, nil
}

func (c *Cassandra) Close() error {
	c.session.Close()
	return nil
}
