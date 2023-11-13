package controllers

import (
	"fmt"
	"strings"
	"time"

	"github.com/buger/jsonparser"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/helpers"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/models"
)

const Fieldtransaksi_home_redis = "AGEN_TRANSAKSI"

func Transksihome(c *fiber.Ctx) error {
	user := c.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	temp_decp := helpers.Decryption(name)
	_, _, client_company := helpers.Parsing_Decry(temp_decp, "==")

	var obj entities.Model_transaksi
	var arraobj []entities.Model_transaksi
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldtransaksi_home_redis + "_" + strings.ToLower(client_company))
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		transaksi_id, _ := jsonparser.GetString(value, "transaksi_id")
		transaksi_date, _ := jsonparser.GetString(value, "transaksi_date")
		transaksi_username, _ := jsonparser.GetString(value, "transaksi_username")
		transaksi_roundbet, _ := jsonparser.GetInt(value, "transaksi_roundbet")
		transaksi_totalbet, _ := jsonparser.GetInt(value, "transaksi_totalbet")
		transaksi_totalwin, _ := jsonparser.GetInt(value, "transaksi_totalwin")
		transaksi_totalbonus, _ := jsonparser.GetInt(value, "transaksi_totalbonus")
		transaksi_card_codepoin, _ := jsonparser.GetString(value, "transaksi_card_codepoin")
		transaksi_card_pattern, _ := jsonparser.GetString(value, "transaksi_card_pattern")
		transaksi_card_result, _ := jsonparser.GetString(value, "transaksi_card_result")
		transaksi_card_win, _ := jsonparser.GetString(value, "transaksi_card_win")
		transaksi_create, _ := jsonparser.GetString(value, "transaksi_create")
		transaksi_update, _ := jsonparser.GetString(value, "transaksi_update")

		obj.Transaksi_id = transaksi_id
		obj.Transaksi_date = transaksi_date
		obj.Transaksi_username = transaksi_username
		obj.Transaksi_roundbet = int(transaksi_roundbet)
		obj.Transaksi_totalbet = int(transaksi_totalbet)
		obj.Transaksi_totalwin = int(transaksi_totalwin)
		obj.Transaksi_totalbonus = int(transaksi_totalbonus)
		obj.Transaksi_card_codepoin = transaksi_card_codepoin
		obj.Transaksi_card_pattern = transaksi_card_pattern
		obj.Transaksi_card_result = transaksi_card_result
		obj.Transaksi_card_win = transaksi_card_win
		obj.Transaksi_update = transaksi_create
		obj.Transaksi_update = transaksi_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_transaksiHome(client_company)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldtransaksi_home_redis+"_"+strings.ToLower(client_company), result, 60*time.Minute)
		fmt.Println("TRANSAKSI MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("TRANSAKSI CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func Transksidetailhome(c *fiber.Ctx) error {
	var errors []*helpers.ErrorResponse
	client := new(entities.Controller_transaksidetail)
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

	var obj entities.Model_transaksidetail
	var arraobj []entities.Model_transaksidetail
	render_page := time.Now()
	resultredis, flag := helpers.GetRedis(Fieldtransaksi_home_redis + "_" + strings.ToLower(client_company) + "_" + client.Transaksi_id)
	jsonredis := []byte(resultredis)
	record_RD, _, _, _ := jsonparser.Get(jsonredis, "record")
	jsonparser.ArrayEach(record_RD, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		transaksidetail_id, _ := jsonparser.GetString(value, "transaksidetail_id")
		transaksidetail_date, _ := jsonparser.GetString(value, "transaksidetail_date")
		transaksidetail_roundbet, _ := jsonparser.GetInt(value, "transaksidetail_roundbet")
		transaksidetail_bet, _ := jsonparser.GetInt(value, "transaksidetail_bet")
		transaksidetail_creditbefore, _ := jsonparser.GetInt(value, "transaksidetail_creditbefore")
		transaksidetail_creditafter, _ := jsonparser.GetInt(value, "transaksidetail_creditafter")
		transaksidetail_win, _ := jsonparser.GetInt(value, "transaksidetail_win")
		transaksidetail_bonus, _ := jsonparser.GetInt(value, "transaksidetail_bonus")
		transaksidetail_card_codepoin, _ := jsonparser.GetString(value, "transaksidetail_card_codepoin")
		transaksidetail_card_win, _ := jsonparser.GetString(value, "transaksidetail_card_win")
		transaksidetail_status, _ := jsonparser.GetString(value, "transaksidetail_status")
		transaksidetail_status_css, _ := jsonparser.GetString(value, "transaksidetail_status_css")
		transaksidetail_create, _ := jsonparser.GetString(value, "transaksidetail_create")
		transaksidetail_update, _ := jsonparser.GetString(value, "transaksidetail_update")

		obj.Transaksidetail_id = transaksidetail_id
		obj.Transaksidetail_date = transaksidetail_date
		obj.Transaksidetail_roundbet = int(transaksidetail_roundbet)
		obj.Transaksidetail_bet = int(transaksidetail_bet)
		obj.Transaksidetail_creditbefore = int(transaksidetail_creditbefore)
		obj.Transaksidetail_creditafter = int(transaksidetail_creditafter)
		obj.Transaksidetail_win = int(transaksidetail_win)
		obj.Transaksidetail_bonus = int(transaksidetail_bonus)
		obj.Transaksidetail_card_codepoin = transaksidetail_card_codepoin
		obj.Transaksidetail_card_win = transaksidetail_card_win
		obj.Transaksidetail_status = transaksidetail_status
		obj.Transaksidetail_status_css = transaksidetail_status_css
		obj.Transaksidetail_create = transaksidetail_create
		obj.Transaksidetail_update = transaksidetail_update
		arraobj = append(arraobj, obj)
	})

	if !flag {
		result, err := models.Fetch_transaksidetailHome(client.Transaksi_id, client_company)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"status":  fiber.StatusBadRequest,
				"message": err.Error(),
				"record":  nil,
			})
		}
		helpers.SetRedis(Fieldtransaksi_home_redis+"_"+strings.ToLower(client_company)+"_"+client.Transaksi_id, result, 60*time.Minute)
		fmt.Println("AGEN TRANSAKSI DETAIL MYSQL")
		return c.JSON(result)
	} else {
		fmt.Println("AGEN TRANSAKSI DETAIL CACHE")
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Success",
			"record":  arraobj,
			"time":    time.Since(render_page).String(),
		})
	}
}
func _deleteredis_transaksi() {
	val_master := helpers.DeleteRedis(Fieldtransaksi_home_redis)
	fmt.Printf("Redis Delete AGEN TRANSAKSI : %d\n", val_master)

	val_master_share := helpers.DeleteRedis(Fieldtransaksi_home_redis)
	fmt.Printf("Redis Delete BACKEND LISTPOINT : %d", val_master_share)
}
