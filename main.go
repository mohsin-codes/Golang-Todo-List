package main

import (
	"github.com/thedevsaddam/renderer"
	mgo "gopkg.in/mgo.v2"
)

var rnd *renderer.Render
var db *mgo.Database

const (
	hostName       string = "localhost:27017"
	dbName         string = "todo-list"
	collectionName string = "todo"
	port           string = ":9000"
)

func main() {

}
