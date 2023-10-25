package models

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/configs"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/db"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/helpers"
	"github.com/nleeper/goment"
)

const database_admincompany_local = configs.DB_tbl_mst_company_admin
const database_adminrulecompany_local = configs.DB_tbl_mst_company_adminrule

func Fetch_adminHome(idcompany string) (helpers.ResponseAdmin, error) {
	var obj entities.Model_admin
	var arraobj []entities.Model_admin
	var res helpers.ResponseAdmin
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			company_idadmin , companyrule_adminrule, tipeadmincompany, company_username,
			company_ipaddress, to_char(COALESCE(company_lastloginadmin,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			company_name, company_phone1, company_phone2, company_status,    
			createadmin_company, to_char(COALESCE(createadmindate_company,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			updateadmin_company, to_char(COALESCE(updateadmindate_company,now()), 'YYYY-MM-DD HH24:MI:SS')     
			FROM ` + database_admincompany_local + ` 
			WHERE idcompany=$1 
			ORDER BY company_lastloginadmin DESC 
		`

	row, err := con.QueryContext(ctx, sql_select, idcompany)

	var no int = 0
	helpers.ErrorCheck(err)
	for row.Next() {
		no += 1
		var (
			companyrule_adminrule_db                                                                                                  int
			company_idadmin_db, tipeadmincompany_db, company_username_db                                                              string
			company_ipaddress_db, company_lastloginadmin_db, company_name_db, company_phone1_db, company_phone2_db, company_status_db string
			createadmin_company_db, createadmindate_company_db, updateadmin_company_db, updateadmindate_company_db                    string
		)

		err = row.Scan(
			&company_idadmin_db, &companyrule_adminrule_db, &tipeadmincompany_db, &company_username_db,
			&company_ipaddress_db, &company_lastloginadmin_db, &company_name_db, &company_phone1_db, &company_phone2_db, &company_status_db,
			&createadmin_company_db, &createadmindate_company_db, &updateadmin_company_db, &updateadmindate_company_db)

		helpers.ErrorCheck(err)
		status_css := configs.STATUS_CANCEL
		if company_status_db == "Y" {
			status_css = configs.STATUS_COMPLETE
		}
		if company_lastloginadmin_db == "0000-00-00 00:00:00" {
			company_lastloginadmin_db = ""
		}
		create := ""
		update := ""
		if createadmin_company_db != "" {
			create = createadmin_company_db + ", " + createadmindate_company_db
		}
		if updateadmin_company_db != "" {
			update = updateadmin_company_db + ", " + updateadmindate_company_db
		}

		obj.Admin_id = company_idadmin_db
		obj.Admin_idcompany = idcompany
		obj.Admin_idrule = companyrule_adminrule_db
		obj.Admin_rule = _Get_adminrule(idcompany, companyrule_adminrule_db)
		obj.Admin_tipe = tipeadmincompany_db
		obj.Admin_username = company_username_db
		obj.Admin_lastIpaddress = company_ipaddress_db
		obj.Admin_lastlogin = company_lastloginadmin_db
		obj.Admin_nama = company_name_db
		obj.Admin_phone1 = company_phone1_db
		obj.Admin_phone2 = company_phone2_db
		obj.Admin_status = company_status_db
		obj.Admin_status_css = status_css
		obj.Admin_create = create
		obj.Admin_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	var objRule entities.Model_adminruleshare
	var arraobjRule []entities.Model_adminruleshare
	sql_listrule := `SELECT 
		companyrule_adminrule,companyrule_name 	
		FROM ` + database_adminrulecompany_local + ` 
		WHERE idcompany=$1  AND companyrule_name != 'MASTER' 
	`
	row_listrule, err_listrule := con.QueryContext(ctx, sql_listrule, idcompany)

	helpers.ErrorCheck(err_listrule)
	for row_listrule.Next() {
		var (
			companyrule_adminrule_db int
			companyrule_name_db      string
		)

		err = row_listrule.Scan(&companyrule_adminrule_db, &companyrule_name_db)

		helpers.ErrorCheck(err)

		objRule.Adminrule_id = companyrule_adminrule_db
		objRule.Adminrule_name = companyrule_name_db
		arraobjRule = append(arraobjRule, objRule)
		msg = "Success"
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Listrule = arraobjRule
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_adminDetail(username string) (helpers.ResponseAdmin, error) {
	var obj entities.Model_adminsave
	var arraobj []entities.Model_adminsave
	var res helpers.ResponseAdmin
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_detail := `SELECT 
		idadmin, name, statuslogin  
		createadmin, to_char(COALESCE(createdateadmin,now()), 'YYYY-MM-DD HH24:MI:SS'), 
		updateadmin, to_char(COALESCE(updatedateadmin,now()), 'YYYY-MM-DD HH24:MI:SS')  
		FROM ` + database_admincompany_local + `
		WHERE username = $1 
	`
	var (
		idadmin_db, name_db, statuslogin_db                                    string
		createadmin_db, createdateadmin_db, updateadmin_db, updatedateadmin_db string
	)
	rows := con.QueryRowContext(ctx, sql_detail, username)

	switch err := rows.Scan(
		&idadmin_db, &name_db, &statuslogin_db,
		&createadmin_db, &createdateadmin_db, &updateadmin_db, &updatedateadmin_db); err {
	case sql.ErrNoRows:

	case nil:
		if createdateadmin_db == "0000-00-00 00:00:00" {
			createdateadmin_db = ""
		}
		if updatedateadmin_db == "0000-00-00 00:00:00" {
			updatedateadmin_db = ""
		}
		create := ""
		update := ""
		if createdateadmin_db != "" {
			create = createadmin_db + ", " + createdateadmin_db
		}
		if updateadmin_db != "" {
			create = updateadmin_db + ", " + updatedateadmin_db
		}

		obj.Username = username
		obj.Nama = name_db
		obj.Rule = idadmin_db
		obj.Status = statuslogin_db
		obj.Create = create
		obj.Update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	default:
		helpers.ErrorCheck(err)
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_adminruleHome(idcompany string) (helpers.Response, error) {
	var obj entities.Model_adminruleall
	var arraobj []entities.Model_adminruleall
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	sql_select := `SELECT 
			companyrule_adminrule , companyrule_name, companyrule_rule,  
			create_companyrule, to_char(COALESCE(createdate_companyrule,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			update_companyrule, to_char(COALESCE(updatedate_companyrule,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + database_adminrulecompany_local + ` 
			WHERE idcompany=$1 
			ORDER BY createdate_companyrule DESC   
		`

	row, err := con.QueryContext(ctx, sql_select, idcompany)

	var no int = 0
	helpers.ErrorCheck(err)
	for row.Next() {
		no += 1
		var (
			companyrule_adminrule_db                                                                           int
			companyrule_name_db, companyrule_rule_db                                                           string
			create_companyrule_db, createdate_companyrule_db, update_companyrule_db, updatedate_companyrule_db string
		)

		err = row.Scan(&companyrule_adminrule_db, &companyrule_name_db, &companyrule_rule_db,
			&create_companyrule_db, &createdate_companyrule_db, &update_companyrule_db, &updatedate_companyrule_db)

		helpers.ErrorCheck(err)

		create := ""
		update := ""
		if create_companyrule_db != "" {
			create = create_companyrule_db + ", " + createdate_companyrule_db
		}
		if update_companyrule_db != "" {
			update = update_companyrule_db + ", " + updatedate_companyrule_db
		}

		obj.Adminrule_id = companyrule_adminrule_db
		obj.Adminrule_name = companyrule_name_db
		obj.Adminrule_rule = companyrule_rule_db
		obj.Adminrule_create = create
		obj.Adminrule_update = update
		arraobj = append(arraobj, obj)
		msg = "Success"
	}
	defer row.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.Time = time.Since(start).String()

	return res, nil
}

func Save_adminHome(admin, idrecord, idcompany, username, password, nama, phone1, phone2, status, sData string, idrule int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDB(database_admincompany_local, "username", username)
		if !flag {
			sql_insert := `
				insert into
				` + database_admincompany_local + ` (
					company_idadmin , companyrule_adminrule, idcompany, tipeadmincompany, 
					company_username, company_password, company_lastloginadmin, company_name, company_phone1, company_phone2, company_status, 
					createadmin_company, createadmindate_company
				) values (
					$1, $2, $3, $4, 
					$5, $6, $7, $8, $9, $10, $11, 
					$12, $13 
				)
			`

			startjoin := tglnow.Format("YYYY-MM-DD HH:mm:ss")
			hashpass := helpers.HashPasswordMD5(password)
			field_column := database_admincompany_local + tglnow.Format("YYYY")
			idrecord_counter := Get_counter(field_column)
			flag_insert, msg_insert := Exec_SQL(sql_insert, database_admincompany_local, "INSERT",
				strings.ToUpper(idcompany)+tglnow.Format("YY")+strconv.Itoa(idrecord_counter), idrule, idcompany, "ADMIN",
				username, hashpass, startjoin, nama, phone1, phone2, status,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

			if flag_insert {
				msg = "Succes"
			} else {
				fmt.Println(msg_insert)
			}
		} else {
			msg = "Duplicate Entry"
		}
	} else {
		if password == "" {
			sql_update := `
				UPDATE 
				` + database_admincompany_local + `  
				SET companyrule_adminrule=$1, company_name=$2, 
				company_phone1=$3, company_phone2=$4, company_status=$5,
				updateadmin_company=$6, updateadmindate_company=$7  
				WHERE company_idadmin=$8 AND idcompany=$9 
			`

			flag_update, msg_update := Exec_SQL(sql_update, database_admincompany_local, "UPDATE",
				idrule, nama, phone1, phone2, status,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord, idcompany)

			if flag_update {
				msg = "Succes"
			} else {
				fmt.Println(msg_update)
			}
		} else {
			hashpass := helpers.HashPasswordMD5(password)
			sql_update2 := `
				UPDATE 
				` + configs.DB_tbl_admin + `   
				SET companyrule_adminrule=$1, company_password=$2, company_name=$3, 
				company_phone1=$4, company_phone2=$5, company_status=$6,
				updateadmin_company=$7, updateadmindate_company=$8  
				WHERE company_idadmin=$9 AND idcompany=$10 
			`
			flag_update, msg_update := Exec_SQL(sql_update2, database_admincompany_local, "UPDATE",
				idrule, hashpass, nama, phone1, phone2, status,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord, idcompany)

			if flag_update {
				msg = "Succes"
			} else {
				fmt.Println(msg_update)
			}
		}
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func Save_adminrule(admin, idcompany, name, rule, sData string, idrecord int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false

	if sData == "New" {
		flag = CheckDBTwoField(database_adminrulecompany_local, "idcompany ", idcompany, "companyrule_name", strings.ToUpper(name))
		if strings.ToUpper(name) != "MASTER" {
			flag = false
		}
		if !flag {
			sql_insert := `
				insert into
				` + database_adminrulecompany_local + ` (
					companyrule_adminrule, idcompany,  companyrule_name, companyrule_rule, 
					create_companyrule,createdate_companyrule
				) values (
					$1, $2, $3, $4, 
					$5, $6
				) 
			`

			field_column := database_adminrulecompany_local + tglnow.Format("YYYY")
			idrecord_counter := Get_counter(field_column)
			flag_insert, msg_insert := Exec_SQL(sql_insert, database_adminrulecompany_local, "INSERT",
				tglnow.Format("YY")+strconv.Itoa(idrecord_counter), idcompany, name, rule,
				admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

			if flag_insert {
				msg = "Succes"
			} else {
				fmt.Println(msg_insert)
			}
		} else {
			msg = "Duplicate Entry"
		}
	} else {
		sql_update := `
				UPDATE 
				` + database_adminrulecompany_local + `   
				SET companyrule_rule=$1 
				WHERE companyrule_adminrule=$2 AND idcompany=$3 
			`
		flag_update, msg_update := Exec_SQL(sql_update, database_adminrulecompany_local, "UPDATE", rule, idrecord, idcompany)

		if flag_update {
			msg = "Succes"
		} else {
			fmt.Println(msg_update)
		}
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}

func _Get_adminrule(idcompany string, idrecord int) string {
	con := db.CreateCon()
	ctx := context.Background()
	companyrule_name := ""
	sql_select := `SELECT
			companyrule_name    
			FROM ` + database_adminrulecompany_local + `  
			WHERE companyrule_adminrule=` + strconv.Itoa(idrecord) + ` AND idcompany='` + idcompany + `'    
		`

	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&companyrule_name); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}

	return companyrule_name
}
