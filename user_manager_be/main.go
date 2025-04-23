package main

import (
	"context"
	"log"
	"time"

	"user_manager/routes"
	"user_manager/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Inicializar conexión Mongo
	services.InitMongo()

	// Desconectar al final
	defer func() {
		if err := services.Client.Disconnect(ctx); err != nil {
			log.Fatal("Error al cerrar la conexión con Mongo:", err)
		}
	}()

	r := gin.Default()

	r.Use(cors.Default())

	// Registrar rutas
	routes.RegisterRoutes(r)

	r.Run(":8081")
}
