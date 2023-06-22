package controllers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rajenderK7/go-mysql/pkg/database"
	"github.com/rajenderK7/go-mysql/pkg/models"
	"github.com/rajenderK7/go-mysql/pkg/utils"
)

func CreateUser(ctx *fiber.Ctx) {
	user := models.User{}
	if err := ctx.BodyParser(&user); err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"message": "Cannot parse user data.",
			},
		)
	} else {
		db := database.GetDB()
		if err := db.Create(&user).Error; err == nil {
			ctx.Status(http.StatusOK).JSON(
				&fiber.Map{
					"message": "User created succesfully.",
				},
			)
		} else {
			ctx.Status(http.StatusInternalServerError).JSON(
				&fiber.Map{
					"message": "Something went wrong.",
				},
			)
			utils.CheckDBError(err)
		}
	}
}

func GetUsers(ctx *fiber.Ctx) {
	var users []models.User
	db := database.GetDB()
	if err := db.Find(&users).Error; err == nil {
		ctx.Status(http.StatusOK).JSON(
			&fiber.Map{
				"users": users,
			},
		)
	} else {
		ctx.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "Something went wrong.",
			},
		)
	}
}

func GetUserById(ctx *fiber.Ctx) {
	userId := ctx.Params("id")
	user := &models.User{}
	db := database.GetDB()
	if err := db.Where("user_id = ?", userId).First(&user); err == nil {
		ctx.Status(http.StatusOK).JSON(
			&fiber.Map{
				"user": user,
			},
		)
	} else {
		ctx.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "Something went wrong.",
			},
		)
	}
}
