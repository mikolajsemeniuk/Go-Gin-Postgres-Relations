package main

import "go-gin-postgres-relations/program/application"

// @title Go-Gin-Postgres-Relations
// @version 1.0
// @description Example usage of Go with Gin and Postgres with relations like one to many, many to many
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath
// @schemes http
func main() {
	application.Route()
	application.Listen()
}
