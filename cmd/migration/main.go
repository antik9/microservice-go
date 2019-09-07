// The migration to database
package main

import (
	"github.com/antik9/microservice-go/internal/backends/database"
)

func main() {
	db.InitialMigration()
}
