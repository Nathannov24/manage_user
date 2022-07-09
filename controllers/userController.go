package controllers

import (
	"manage_user/lib/database"
	responses "manage_user/lib/response"
	"manage_user/middlewares"
	"manage_user/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var user models.Users

//register user
func RegisterUsersController(c echo.Context) error {
	c.Bind(&user)
	duplicate, _ := database.GetUserByEmail(user.Email)
	if duplicate > 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Email was used, try another email",
		})
	}

	Password, _ := database.GeneratehashPassword(user.Password)
	user.Password = Password

	_, err := database.RegisterUser(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed)
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

//login users
func LoginUsersController(c echo.Context) error {
	user := models.UserLogin{}
	c.Bind(&user)
	users, err := database.LoginUsers(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Login success",
		"token":   users.Token,
	})
}

//get user by id
func GetUsersController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed)
	}
	loginuser := middlewares.ExtractTokenUserId(c)
	if loginuser != id {
		return c.JSON(http.StatusUnauthorized, responses.StatusFailedInternal)
	}
	respon, e := database.GetUser(id)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusFailedInternal)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    respon,
	})
}

func GetAllUsersController(c echo.Context) error {
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	user, e := database.GetUserAll(limit, offset)
	if e != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusFailed("internal server error"))
	}
	var result []models.GetAllUser
	for _, v := range user {
		var data models.GetAllUser
		data.User_Name = v.User_Name
		data.Email = v.Email
		data.Name = v.Name
		data.Role = v.Role

		result = append(result, data)
	}
	return c.JSON(http.StatusOK, responses.StatusSuccessData("success get user", result))
}

//delete user by id
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed)
	}
	_, error := database.DeleteUser(id)
	if error != nil {
		return c.JSON(http.StatusInternalServerError, responses.StatusFailed)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}

//update user by id
func UpdateUserController(c echo.Context) error {
	c.Bind(&user)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed)
	}
	respon, e := database.UpdateUser(id, user)
	if e != nil {
		return c.JSON(http.StatusUnauthorized, responses.StatusFailedInternal)
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    respon,
	})
}

func SearchController(c echo.Context) error {
	id := c.QueryParam("name")
	idUser := c.QueryParam("user_name")
	idEmail := c.QueryParam("email")
	user, _ := database.Search(id, idUser, idEmail)
	if user == 0 {
		return c.JSON(http.StatusBadRequest, responses.StatusFailed("user not found"))
	}
	return c.JSON(http.StatusOK, responses.StatusSuccessData("success get user", user))
}
