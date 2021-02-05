package db

import (
	"time"

	"github.com/easyms/archv1/pc-service/internal"
	"gopkg.in/mgo.v2"
)

// Connection defines connection detailes
type Connection struct {
	session *mgo.Session
}

// NewConnection establishes a new connection to a particular mongo database
func NewConnection(config *internal.Config) (conn *Connection) {

	info := &mgo.DialInfo{

		// Address of the database host
		Addrs: []string{config.GetConfigProperty("DATABASE_HOST")},

		// Timeout when a failure happens while connecting to the database
		Timeout: 60 * time.Second,

		// Database name
		Database: config.GetConfigProperty("DATABASE_NAME"),

		// Database credentials
		Username: config.GetConfigProperty("DATABASE_USERNAME"),
		Password: config.GetConfigProperty("DATABASE_PASSWORD"),
	}

	session, err := mgo.DialWithInfo(info)

	// Changes the session safety mode, make the session check for errors
	session.SetSafe(&mgo.Safe{})
	if err != nil {
		panic(err)
	}

	conn = &Connection{session}
	return conn
}

// Use retrieves a connection to a mongo collection always use the default db
// since we use just one
func (conn *Connection) Use(docName string) (collection *mgo.Collection) {
	return conn.session.DB("").C(docName)
}

// Close handles closing a database connection session
func (conn *Connection) Close() {
	conn.session.Close()
	return
}
