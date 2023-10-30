package controllers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/helpers"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/models"
)

const Fieldlistbet_home_redis = "AGEN_LISTBET"
const Fieldcompanylistbet_home_redis = "COMPANYLISTBET_BACKEND"

func Listbethome(c *fiber.Ctx) error {
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	_, _, client_company := helpers.Parsing_Decry(temp_decp, "==")

	var obj entities.Model_lisbet
	var arraobj []entities.Model_lisbet
	var objmstlisbet entities.Model_lisbetshare
	var arraobjmstlisbet []entities.Model_lisbetshare
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldlistbet_home_redis + "_" + client_company)
	jsonredis := []byte(resultredis)
	listbet_RD, _, _, _ := jsonparser.Get(jsonredis, "listbet")
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		lisbet_id, _ := jsonparser.GetInt(value, "lisbet_id")
		lisbet_minbet, _ := jsonparser.GetFloat(value, "lisbet_minbet")
		lisbet_create, _ := jsonparser.GetString(value, "lisbet_create")
		lisbet_update, _ := jsonparser.GetString(value, "lisbet_update")

		obj.Lisbet_id = int(lisbet_id)
		obj.Lisbet_minbet = float64(lisbet_minbet)
		obj.Lisbet_create = lisbet_create
		obj.Lisbet_update = lisbet_update
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(listbet_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		lisbet_minbet, _ := jsonparser.GetFloat(value, "lisbet_minbet")

		objmstlisbet.Lisbet_minbet = float64(lisbet_minbet)
		arraobjmstlisbet = append(arraobjmstlisbet, objmstlisbet)
	})
	if !flag {
		result, err := models.Fetch_listbetHome(client_company)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldlistbet_home_redis+"_"+client_company, result, 60*time.Minute)
		fmt.Println("LISTBET MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("LISTBET CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"listbet": arraobjmstlisbet,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Listbetconfpointhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_listbetconfpoint)
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

	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	_, _, client_company := helpers.Parsing_Decry(temp_decp, "==")

	var obj entities.Model_listbet_conf
	var arraobj []entities.Model_listbet_conf
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldlistbet_home_redis + "_CONFIG_" + client_company + "_" + strconv.Itoa(client.Lisbet_idbet))
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		listbetconf_id, _ := jsonparser.GetInt(value, "listbetconf_id")
		listbetconf_idbet, _ := jsonparser.GetInt(value, "listbetconf_idbet")
		listbetconf_nmpoin, _ := jsonparser.GetString(value, "listbetconf_nmpoin")
		listbetconf_poin, _ := jsonparser.GetInt(value, "listbetconf_poin")
		listbetconf_create, _ := jsonparser.GetString(value, "listbetconf_create")
		listbetconf_update, _ := jsonparser.GetString(value, "listbetconf_update")

		obj.Listbetconf_id = int(listbetconf_id)
		obj.Listbetconf_idbet = int(listbetconf_idbet)
		obj.Listbetconf_nmpoin = listbetconf_nmpoin
		obj.Listbetconf_poin = int(listbetconf_poin)
		obj.Listbetconf_create = listbetconf_create
		obj.Listbetconf_update = listbetconf_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_listbetConfPoint(client.Lisbet_idbet, client_company)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldlistbet_home_redis+"_CONFIG_"+client_company+"_"+strconv.Itoa(client.Lisbet_idbet), result, 60*time.Minute)
		fmt.Println("AGEN CONF MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("COMPANY CONF CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func ListbetSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_listbetsave)
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
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _, client_company := helpers.Parsing_Decry(temp_decp, "==")

	// admin, sData string, idrecord int, minbet float64
	result, err := models.Save_listbet(
		client_admin, client_company,
		client.Sdata, client.Lisbet_id, client.Lisbet_minbet)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_listbet(client_company, 0)
	return c.JSON(result)
}
func ListbetconfpointSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_listbetconfpoinsave)
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
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	client_admin, _, client_company := helpers.Parsing_Decry(temp_decp, "==")

	// admin, idcompany, sData string, idrecord, idbet, point int
	result, err := models.Save_ConfPoint(
		client_admin, client_company,
		client.Sdata, client.Conf_id, client.Conf_idbet, client.Conf_minbet)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_listbet(client_company, client.Conf_idbet)
	return c.JSON(result)
}
func _deleteredis_listbet(idcompany string, idbet int) {
	val_master := helpers.DeleteRedis(Fieldlistbet_home_redis + "_" + idcompany)
	fmt.Printf("Redis Delete AGEN LISTBET : %d\n", val_master)
	val_masterlistbetcof := helpers.DeleteRedis(Fieldlistbet_home_redis + "_CONFIG_" + idcompany + "_" + strconv.Itoa(idbet))
	fmt.Printf("Redis Delete AGEN LISTBET CONF POIN : %d\n", val_masterlistbetcof)

	val_super := helpers.DeleteRedis(Fieldcompanylistbet_home_redis + "_" + idcompany)
	fmt.Printf("Redis Delete SUPER LISTBET : %d\n", val_super)

}
