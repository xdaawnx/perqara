package handler

import (
	"net/http"
	"perqara_api/models"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type User struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Sex     string `json:"sex" validate:"required,oneof=female male"`
}

// GetUsers gets all users
// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} map[string][]models.User
// @Failure 500 {object} map[string]string
// @Router /users [get]
func GetUsers(c echo.Context) error {
	modelUser := models.User{}
	users, err := modelUser.GetAllUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string][]models.User{"data": users})
}

// GetUserByID gets a user by ID
// @Summary Get a user by ID
// @Description Get a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} map[string]models.User
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func GetUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}
	modelUser := models.User{}

	user, err := modelUser.GetUserByID(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]models.User{"data": user})
}

// CreateUser creates a new user
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User"
// @Success 201 {object} map[string]models.User
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users [post]
func CreateUser(c echo.Context) error {
	var user User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}

	validate := validator.New(validator.WithRequiredStructEnabled())
	if err := validate.Struct(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	modelUser := models.User{Name: user.Name, Address: user.Address, Sex: user.Sex}

	newUser, err := modelUser.CreateUser()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusCreated, map[string]models.User{"data": newUser})
}

// UpdateUser updates a user
// @Summary Update a user
// @Description Update a user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body models.User true "User"
// @Success 200 {object} map[string]models.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [put]
func UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}

	var user User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	validate := validator.New(validator.WithRequiredStructEnabled())

	if err := validate.Struct(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
	}
	modelUser := models.User{Name: user.Name, Address: user.Address, Sex: user.Sex}

	err = modelUser.UpdateUser(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]models.User{"data": modelUser})
}

// DeleteUser deletes a user
// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 204
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /users/{id} [delete]
func DeleteUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid user ID")
	}
	modelUser := models.User{}

	err = modelUser.DeleteUser(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
