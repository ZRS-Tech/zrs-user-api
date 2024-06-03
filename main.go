package main

import (
	"zrs.user.api/helper"
)

// Connection mongoDB with helper class
// var collection = helper.ConnectDB()

// var client *mongo.Client

func main() {
	helper.ConnectDB()
}
