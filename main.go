package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/wcarlsen/seatrout/infrastructure/persistence"
	"github.com/wcarlsen/seatrout/interfaces"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {

	dbdriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	services, err := persistence.NewRepositories(dbdriver, user, password, port, host, dbname)
	if err != nil {
		panic(err)
	}
	defer services.Close()
	services.Automigrate()

	spots := interfaces.NewSpots(services.Spot)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/spots", spots.SaveSpot)
	r.GET("/spots", spots.GetSpots)
	r.GET("/spots/:spot_id", spots.GetSpot)

	log.Fatal(r.Run(":" + os.Getenv("PORT")))

}
