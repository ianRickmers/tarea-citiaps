package services

import (
	"context"
	"user_manager/models"
	"user_manager/utils"

	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Crea un nuevo usuario en la base de datos
// y devuelve un error si ocurre algún problema.
// La función toma un contexto (para la gestión de la cancelación)
// y un puntero a un modelo de usuario como parámetros.
func CreateUser(ctx context.Context, user *models.User) error {
	utils.Info("Insertando nuevo usuario")
	collection := GetDatabase().Collection("users")
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		utils.Error(err, "Error al insertar usuario")
		return err
	}
	utils.Info("Usuario insertado correctamente: " + user.ID.Hex())
	return nil
}

// Obtiene todos los usuarios de la base de datos
// y devuelve una lista de usuarios o un error si ocurre algún problema.
// La función toma un contexto
func GetUsers(ctx context.Context) ([]models.User, error) {
	utils.Info("Obteniendo todos los usuarios desde la base de datos")

	collection := GetDatabase().Collection("users")

	// Se utiliza un cursor para iterar sobre los resultados
	// Se utiliza bson.D{} (filtro vacío) para obtener todos los documentos
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		utils.Error(err, "Error al obtener usuarios")
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []models.User
	for cursor.Next(ctx) {
		var user models.User
		// Se decodifica el cursor en la estructura de usuario
		// y se añade a la lista de usuarios
		if err := cursor.Decode(&user); err != nil {
			utils.Error(err, "Error al decodificar usuario")
			continue // o puedes guardar el error si quieres
		}

		users = append(users, user)
	}

	utils.Info(fmt.Sprintf("Usuarios recuperados: %d", len(users)))
	return users, nil
}

// GetUserByID obtiene un usuario por su ID
// y devuelve un puntero al usuario o un error si ocurre algún problema.
// La función toma un contexto y un ID de usuario como parámetros.
func GetUserByID(ctx context.Context, id string) (*models.User, error) {
	utils.Info("Buscando usuario con ID: " + id)
	collection := GetDatabase().Collection("users")

	// Se convierte el ID de string a ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		utils.Error(err, "ID inválido al buscar usuario")
		return nil, err
	}

	var user models.User
	err = collection.FindOne(ctx, bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		utils.Error(err, "Usuario no encontrado con ese ID")
		return nil, err
	}

	utils.Info("Usuario encontrado: " + user.ID.Hex())
	return &user, nil
}
