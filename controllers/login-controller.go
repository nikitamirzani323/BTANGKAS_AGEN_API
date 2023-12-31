package controllers

import (
	"fmt"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/helpers"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/models"
)

func CheckLogin(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Login)
	validate := validator.New()
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	err := validate.Struct(client)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element helpers.ErrorResponse
			element.Field = err.StructField()
			element.Tag = err.Tag()
			errors = append(errors, &element)
		}
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": "validation",
			"record":  errors,
		})
	}

	result, ruleadmin, company, err := models.Login_Model(client.Username, client.Password, client.Ipaddress)

	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	if !result {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": "Username or Password Not Found",
			})

	} else {
		dataclient := client.Username + "==" + ruleadmin + "==" + company
		dataclient_encr, keymap := helpers.Encryption(dataclient)
		dataclient_encr_final := dataclient_encr + "|" + strconv.Itoa(keymap)
		t, err := helpers.GenerateNewAccessToken(dataclient_encr_final)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{
			"status": fiber.StatusOK,
			"token":  t,
		})

	}
}
func Home(c *fiber.Ctx) error {
	client := new(entities.Home)
	if err := c.BodyParser(client); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_username, idruleadmin, client_company := helpers.Parsing_Decry(temp_decp, "==")
	fmt.Printf("USERNAME : %s\n", client_username)
	fmt.Printf("RULE : %s\n", idruleadmin)
	fmt.Printf("COMPANY : %s\n", client_company)
	fmt.Printf("PAGE : %s\n", client.Page)

	if idruleadmin == "MASTER" {
		c.Status(fiber.StatusOK)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "MASTER",
			"record":  nil,
		})
	} else {
		ruleadmin := models.Get_AdminRule("ruleadmingroup", idruleadmin, client_company)
		flag := models.Get_listitemsearch(ruleadmin, ",", client.Page)
		if !flag {
			c.Status(fiber.StatusForbidden)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusForbidden,
				"message": "Anda tidak bisa akses halaman ini",
				"record":  nil,
			})
		} else {
			c.Status(fiber.StatusOK)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusOK,
				"message": "ADMIN",
				"record":  nil,
			})
		}
	}

}
