package controllers

import (
	"fmt"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/helpers"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/models"
)

const Fieldadmincompany_home_redis = "AGEN_LISTADMIN"
const Fieldadminrulecompany_home_redis = "AGEN_LISTADMINRULE"

func Adminhome(c *fiber.Ctx) error {
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	_, _, client_company := helpers.Parsing_Decry(temp_decp, "==")

	var obj entities.Model_admin
	var arraobj []entities.Model_admin
	var obj_listruleadmin entities.Model_adminruleshare
	var arraobj_listruleadmin []entities.Model_adminruleshare
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldadmincompany_home_redis + "_" + client_company)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	listruleadmin_RD, _, _, _ := jsonparser.Get(jsonredis, "listruleadmin")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		admin_id, _ := jsonparser.GetString(value, "admin_id")
		admin_idcompany, _ := jsonparser.GetString(value, "admin_idcompany")
		admin_username, _ := jsonparser.GetString(value, "admin_username")
		admin_nama, _ := jsonparser.GetString(value, "admin_nama")
		admin_tipe, _ := jsonparser.GetString(value, "admin_tipe")
		admin_idrule, _ := jsonparser.GetInt(value, "admin_idrule")
		admin_rule, _ := jsonparser.GetString(value, "admin_rule")
		admin_phone1, _ := jsonparser.GetString(value, "admin_phone1")
		admin_phone2, _ := jsonparser.GetString(value, "admin_phone2")
		admin_lastlogin, _ := jsonparser.GetString(value, "admin_lastlogin")
		admin_lastipaddres, _ := jsonparser.GetString(value, "admin_lastipaddres")
		admin_status, _ := jsonparser.GetString(value, "admin_status")
		admin_status_css, _ := jsonparser.GetString(value, "admin_status_css")
		admin_create, _ := jsonparser.GetString(value, "admin_create")
		admin_update, _ := jsonparser.GetString(value, "admin_update")

		obj.Admin_id = admin_id
		obj.Admin_idcompany = admin_idcompany
		obj.Admin_username = admin_username
		obj.Admin_nama = admin_nama
		obj.Admin_tipe = admin_tipe
		obj.Admin_idrule = int(admin_idrule)
		obj.Admin_rule = admin_rule
		obj.Admin_phone1 = admin_phone1
		obj.Admin_phone2 = admin_phone2
		obj.Admin_lastlogin = admin_lastlogin
		obj.Admin_lastIpaddress = admin_lastipaddres
		obj.Admin_status = admin_status
		obj.Admin_status_css = admin_status_css
		obj.Admin_create = admin_create
		obj.Admin_update = admin_update
		arraobj = append(arraobj, obj)
	})
	jsonparser.ArrayEach(listruleadmin_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		adminrule_id, _ := jsonparser.GetInt(value, "adminrule_id")
		adminrule_name, _ := jsonparser.GetString(value, "adminrule_name")

		obj_listruleadmin.Adminrule_id = int(adminrule_id)
		obj_listruleadmin.Adminrule_name = adminrule_name
		arraobj_listruleadmin = append(arraobj_listruleadmin, obj_listruleadmin)
	})
	if !flag {
		result, err := models.Fetch_adminHome(client_company)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldadmincompany_home_redis+"_"+client_company, result, 60*time.Minute)
		fmt.Println("ADMIN MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("ADMIN CACHE")
		return c.JSON(fiber.Map{
			"status":        fiber.StatusOK,
			"message":       "Success",
			"record":        arraobj,
			"listruleadmin": arraobj_listruleadmin,
			"company":       client_company,
			"time":          time.Since(render_page).String(),
		})
	}
}
func Adminrulehome(c *fiber.Ctx) error {
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	_, _, client_company := helpers.Parsing_Decry(temp_decp, "==")

	var obj entities.Model_adminruleall
	var arraobj []entities.Model_adminruleall
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldadminrulecompany_home_redis + "_" + client_company)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		adminrule_id, _ := jsonparser.GetInt(value, "adminrule_id")
		adminrule_name, _ := jsonparser.GetString(value, "adminrule_name")
		adminrule_rule, _ := jsonparser.GetString(value, "adminrule_rule")
		adminrule_create, _ := jsonparser.GetString(value, "adminrule_create")
		adminrule_update, _ := jsonparser.GetString(value, "adminrule_update")

		obj.Adminrule_id = int(adminrule_id)
		obj.Adminrule_name = adminrule_name
		obj.Adminrule_rule = adminrule_rule
		obj.Adminrule_create = adminrule_create
		obj.Adminrule_update = adminrule_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_adminruleHome(client_company)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldadminrulecompany_home_redis+"_"+client_company, result, 60*time.Minute)
		fmt.Println("ADMIN RULE MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("ADMIN RULE CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func AdminDetail(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_admindetail)
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

	result, err := models.Fetch_adminDetail(client.Username)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}
	return c.JSON(result)
}
func AdminSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_adminsave)
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

	//admin, idrecord, idcompany, username, password, nama, phone1, phone2, status, sData string, idrule int
	result, err := models.Save_adminHome(
		client_admin,
		client.Admin_id, client_company, client.Admin_username, client.Admin_password,
		client.Admin_nama, client.Admin_phone1, client.Admin_phone2, client.Admin_status, client.Sdata,
		client.Admin_idrule)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_admincompany(client_company)
	return c.JSON(result)
}
func AdminruleSave(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_adminrulesave)
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

	//admin, idcompany, name, rule, sData string, idrecord int
	result, err := models.Save_adminrule(
		client_admin,
		client_company, client.Adminrule_name, client.Adminrule_rule,
		client.Sdata, client.Adminrule_id)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusBadRequest,
			"message": err.Error(),
			"record":  nil,
		})
	}

	_deleteredis_admincompany(client_company)
	return c.JSON(result)
}
func _deleteredis_admincompany(idcompany string) {
	val_master := helpers.DeleteRedis(Fieldadmincompany_home_redis + "_" + idcompany)
	fmt.Printf("Redis Delete AGEN ADMIN : %d\n", val_master)

	val_master_rule := helpers.DeleteRedis(Fieldadminrulecompany_home_redis + "_" + idcompany)
	fmt.Printf("Redis Delete AGEN ADMIN RULE : %d\n", val_master_rule)

}
