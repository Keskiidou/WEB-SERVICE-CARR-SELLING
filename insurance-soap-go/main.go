package main

import (
	"insurance-soap-go/initializers"
	"insurance-soap-go/server"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()

}

func main() {
	server.StartServer()
}
