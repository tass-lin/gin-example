package database

import (
	"crypto/tls"
	"gopkg.in/mgo.v2"
	"gin-example/configs"
	"net"
)

func Conn() *mgo.Session {
	// get db config
	mongoDbConfig := configs.GetMongoDbConfig()
	var mongoURI = mongoDbConfig["DB_DSN"]
	dialInfo, err := mgo.ParseURL(mongoURI)
	tlsConfig := &tls.Config{}

	dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		return conn, err
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	return session


}
