package scylla

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

// S Singleton session
var session gocqlx.Session

func NewSession(keyspace string, scyllaHosts ...string) (*gocqlx.Session, error) {
	var cluster = gocql.NewCluster(scyllaHosts...)
	cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy("GCE_ASIA_SOUTHEAST_1")

	sess, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		return nil, fmt.Errorf("Cannot connect to Scylla clusters ")
	} else {
		err = sess.ExecStmt(`CREATE KEYSPACE IF NOT EXISTS ` + keyspace + ` WITH replication = {'class': 'NetworkTopologyStrategy', 'replication_factor': 3}`)
		setSession(sess)
	}
	return &session, nil
}

func GetSession() gocqlx.Session {
	if session.Session == nil {
		panic("plz check connection db")
	}
	return session
}

func setSession(s gocqlx.Session) {
	session = s
}
