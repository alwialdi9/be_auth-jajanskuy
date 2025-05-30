package handlers

import (
	"github.com/alwialdi9/be_auth-jajanskuy/connection"
	"github.com/alwialdi9/be_auth-jajanskuy/models"
	"github.com/alwialdi9/be_auth-jajanskuy/utils"
	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {
	type signUpReq struct {
		Username  string `json:"username" validate:"required,min=3,max=20"`
		Email     string `json:"email" validate:"required,email"`
		Password  string `json:"password" validate:"required,min=8,max=100"`
		FirstName string `json:"first_name" validate:"required,min=2,max=50"`
		LastName  string `json:"last_name" validate:"required,min=2,max=50"`
	}

	var req signUpReq
	if err := c.BodyParser(&req); err != nil {
		return utils.JsonErrorResponse(c, fiber.StatusOK, "Invalid request body", err.Error())
	}

	validationReq := utils.ValidateStruct(req)
	if validationReq != nil {
		return utils.JsonErrorResponse(c, fiber.StatusOK, "Validation failed", validationReq)
	}

	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return utils.JsonErrorResponse(c, fiber.StatusOK, "Failed to hash password", err.Error())
	}

	user := models.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
	}

	errorCreate := models.CreateUser(connection.DB, &user)
	if errorCreate != nil {
		return utils.JsonErrorResponse(c, fiber.StatusOK, "Failed to create user", errorCreate.Error())
	}

	return utils.JsonResponse(c, fiber.StatusCreated, "User created successfully", fiber.Map{
		"user": fiber.Map{
			"id":         user.ID,
			"username":   user.Username,
			"email":      user.Email,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
		},
		"message": "User created successfully",
		"success": true,
	})
}
