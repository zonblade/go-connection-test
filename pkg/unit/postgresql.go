package unit

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"gotest/pkg/loader"
)

func TestPostgresqlConnection() {

	dbData := loader.GetDbConfig("postgres")

	for _, data := range dbData {
		testPostgres(
			data.Username,
			data.Password,
			data.Name,
			data.Host,
			data.Port,
		)
	}
}

func testPostgres(user string, password string, name string, host string, port string){

	// create pq connection
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host,
		user,
		password,
		name,
		port,
	)

	fmt.Printf("Opening... Database: %s , Host: %s\n", name, host)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	pgDb, err := db.DB()
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	err = pgDb.Ping()

	if err != nil {
		fmt.Println("Error pinging the database:", err)
		return
	}

	_ = pgDb.Close()

	fmt.Println("Database connection is alive and working.")
	fmt.Println()
}
