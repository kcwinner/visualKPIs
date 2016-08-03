package data

import (
  "gopkg.in/mgo.v2"
  "log"
)

var session *mgo.Session
var Database *mgo.Database

func DialDb() error {
  var err error

  log.Println("Dialing mongodb: localhost")
  session, err = mgo.Dial("localhost")

  return err
}

func GetDb() *mgo.Session {
  return session.Copy()
}

func CloseDb() {
  session.Close()
  log.Println("Closed database connection")
}
