package monitor

import (
	"gopkg.in/mgo.v2"
	"log"
	"net"
	"net/url"
	"os"
)

// DB ...
type DB struct {
	connectionURI string
	DBName        string
	Session       *mgo.Session
}

func (db *DB) setConnectionURIFromEnvConfig() {
	errMsg := "no connection string provided"

	mongoConn, err := url.Parse(os.Getenv("MONGO_PORT"))
	if err != nil {
		log.Fatalln(errMsg)
	}

	dbHost, dbPort, err := net.SplitHostPort(mongoConn.Host)
	if err != nil {
		log.Fatalln(errMsg)
	}

	db.DBName = os.Getenv("MONGODB_DATABASE")
	if db.DBName == "" || dbHost == "" || dbPort == "" {
		log.Fatalln(errMsg)
	}

	db.connectionURI = "mongodb://" + dbHost + ":" + dbPort + "/" + db.DBName
}

// Start ...
func (db *DB) Start() {
	var err error
	db.setConnectionURIFromEnvConfig()

	db.Session, err = mgo.Dial(db.connectionURI)
	if err != nil {
		log.Fatalf("Can't connect to mongo, go error %v\n", err)
	}

	db.Session.SetSafe(&mgo.Safe{})
	db.Session.SetMode(mgo.Monotonic, true)
}

// Close ...
func (db *DB) Close() {
	db.Session.Close()
}
