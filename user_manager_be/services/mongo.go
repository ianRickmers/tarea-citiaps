package services

import (
	"context"
	"time"
	"user_manager/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Se crea el puntero al cliente de MongoDB
// Se utiliza un puntero para poder acceder a la conexión desde otros archivos
// y no tener que crear una nueva conexión cada vez
var Client *mongo.Client

// InitMongo inicializa la conexión a MongoDB
// Se utiliza un contexto con timeout para evitar que la conexión se quede colgada
// Si no se puede conectar, se imprime un error y se termina el programa
// Se utiliza un defer para cerrar la conexión al final
func InitMongo() {
	utils.Info("Iniciando conexión a MongoDB")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		utils.Error(err, "Error al conectar a MongoDB")
		panic(err)
	}
	utils.Info("Conexión a MongoDB establecida correctamente")
}

// GetDatabase devuelve la base de datos de MongoDB
// Se utiliza un puntero para poder acceder a la base de datos desde otros archivos
// Se crea la base de datos si no existe al insertar el primer dato
func GetDatabase() *mongo.Database {
	utils.Info("Accediendo a la base de datos: user_manager")
	return Client.Database("user_manager")
}
