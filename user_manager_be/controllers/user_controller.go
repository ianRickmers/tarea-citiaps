package controllers

import (
	"net/http"
	"time"
	"user_manager/models"
	"user_manager/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUserController maneja la creación de un nuevo usuario.
func CreateUserController(c *gin.Context) {

	var user models.User

	// Se bindea el JSON a la estructura de usuario
	// y se valida si hay errores
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Se asigna un ID único y la fecha de creación
	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()

	// Se crea el usuario en la base de datos
	if err := services.CreateUser(c, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

// GetUsersController maneja la obtención de todos los usuarios.
func GetUsersController(c *gin.Context) {
	var users []models.User

	// Se obtienen los usuarios de la base de datos
	users, err := services.GetUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}

// GetUserByIDController maneja la obtención de un usuario por su ID.
func GetUserByIDController(c *gin.Context) {
	id := c.Param("id")

	// Se obtiene el usuario de la base de datos
	user, err := services.GetUserByID(c, id)

	// Se verifica si hubo un error al obtener el usuario
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Se verifica si el usuario no existe
	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
