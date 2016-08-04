package data

import (
	"log"

	"github.com/kcwinner/visualKPIs/common"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session

//DialDb Starts the connection to the database
func DialDb() error {
	var err error

	log.Println("Dialing mongodb:", common.AppConfig.DatabaseServer)
	session, err = mgo.Dial(common.AppConfig.DatabaseServer)

	return err
}

//GetDb Gets a copy of the database session
func GetDb() *mgo.Session {
	return session.Copy()
}

//CloseDb Closes the connection to the database
func CloseDb() {
	session.Close()
	log.Println("Closed database connection")
}
