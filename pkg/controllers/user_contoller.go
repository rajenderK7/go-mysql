package controllers

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rajenderK7/go-mysql/pkg/database"
	"github.com/rajenderK7/go-mysql/pkg/models"
)

func CreateUser(ctx *fiber.Ctx) error {
	user := models.User{}
	if err := ctx.BodyParser(&user); err != nil {
		ctx.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{
				"message": "Cannot parse user data",
			},
		)
		return err
	}
	db := database.GetDB()
	if err := db.Create(&user).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "Failed to create user",
			},
		)
	}
	return ctx.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "User created succesfully",
		},
	)
}

func GetUsers(ctx *fiber.Ctx) error {
	var users []models.User
	db := database.GetDB()
	if err := db.Find(&users).Error; err != nil {

		return ctx.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "Something went wrong",
			},
		)
	}
	return ctx.Status(http.StatusOK).JSON(
		&fiber.Map{
			"users": users,
		},
	)
}

func GetUserById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	userID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"message": "Invalid user ID",
			},
		)
	}
	user := models.User{}
	db := database.GetDB()
	if err := db.First(&user, int32(userID)).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "Failed to find user",
			},
		)
	}
	return ctx.Status(http.StatusOK).JSON(
		&fiber.Map{
			"user": user,
		},
	)
}

func UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	userID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user_id",
		})
	}
	updateReq := models.User{}
	if err := ctx.BodyParser(&updateReq); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"message": "Error while parsing",
			},
		)
	}
	db := database.GetDB()
	existingUser := models.User{}
	if err := db.First(&existingUser, userID).Error; err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"message": "Cannot process user data",
			},
		)
	}
	if updateReq.Name != "" {
		existingUser.Name = updateReq.Name
	}
	if updateReq.Phone != "" {
		existingUser.Phone = updateReq.Phone
	}
	if err := db.Save(&existingUser).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "Failed to update user",
			},
		)
	}
	return ctx.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "User update successfully",
			"user":    existingUser,
		},
	)
}

func DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	userID, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user_id",
		})
	}
	db := database.GetDB()
	if err := db.Delete(&models.User{}, int32(userID)).Error; err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(
			&fiber.Map{
				"message": "Failed to delete user",
			},
		)
	}
	return ctx.Status(http.StatusOK).JSON(
		&fiber.Map{
			"message": "User deleted successfully",
		},
	)
}
