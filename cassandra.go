package benchmark

import (
	"fmt"
	"time"

	"github.com/gocql/gocql"
	"github.com/gocql/gocql/lz4"
)

type cassandraWriter struct {
	session  *gocql.Session
	table    string
	keyspace string
}

type CassandraConfig struct {
	Hosts      []string
	Keyspace   string
	Table      string
	DataCenter string
	Rack       string
}

func Cassandra(conf CassandraConfig) *cassandraWriter {
	cluster := gocql.NewCluster(conf.Hosts...)
	cluster.Timeout = 60 * time.Second
	cluster.Consistency = gocql.Quorum
	cluster.Compressor = lz4.LZ4Compressor{}

	session, err := cluster.CreateSession()
	if err != nil {
		panic(fmt.Errorf("failed to create session: %w", err))
	}

	return &cassandraWriter{
		session:  session,
		table:    conf.Table,
		keyspace: conf.Keyspace,
	}
}

func (c *cassandraWriter) Write(data []byte) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s.%s (id, data) VALUES (?, ?)", c.keyspace, c.table)
	if err := c.session.Query(query, gocql.TimeUUID(), data).Exec(); err != nil {
		return 0, err
	}
	return len(data), nil
}

func (c *cassandraWriter) Close() error {
	c.session.Close()
	return nil
}

func CreateTable(conf CassandraConfig) error {
	cluster := gocql.NewCluster(conf.Hosts...)
	cluster.Timeout = 5 * time.Second
	s, err := cluster.CreateSession()
	if err != nil {
		return err
	}
	defer s.Close()

	err = s.Query(fmt.Sprintf(""+
		"CREATE KEYSPACE IF NOT EXISTS %s "+
		"WITH replication = {"+
		"    'class': 'SimpleStrategy',"+
		"    'replication_factor': 1"+
		"}", conf.Keyspace)).Exec()
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`
        CREATE TABLE IF NOT EXISTS %s.%s (
            id UUID PRIMARY KEY,
            data BLOB
        )
    `, conf.Keyspace, conf.Table)
	return s.Query(query).Exec()
}
