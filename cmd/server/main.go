package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"api/config"
	"api/controllers"
	dbConn "api/db/sqlc"
	"api/routes"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

var (
	server *gin.Engine
	db     *dbConn.Queries

	SignatureController controllers.SignatureController
	SignatureRoutes     routes.SignatureRoutes
)

func init() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	conn, err := sql.Open(config.PostgreDriver, config.PostgresSource)
	if err != nil {
		log.Fatalf("could not connect to postgres database: %v", err)
	}

	db = dbConn.New(conn)

	fmt.Println("PostgreSQL connected successfully...")

	SignatureController = *controllers.NewSignatureController(db)
	SignatureRoutes = routes.NewSignatureRoutes(SignatureController)

	server = gin.Default()
}

func main() {
	config, err := config.LoadConfig(".")

	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	router := server.Group("/api")

	router.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": "Evertything is OK"})
	})

	SignatureRoutes.SignatureRoute(router)
	log.Fatal(server.Run(":" + config.Port))
}
