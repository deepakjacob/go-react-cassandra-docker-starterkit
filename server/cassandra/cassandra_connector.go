package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

func init() {
	var err error
	cluster := gocql.NewCluster("jd_cassandra")
	cluster.Keyspace = "employee"
	Session, err := cluster.CreateSession()
	if err != nil {
		panic(err)
	}
	fmt.Println("Cassandra init done...")
}
