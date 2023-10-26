package models

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/configs"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/db"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/helpers"
	"github.com/nleeper/goment"
)

const database_listbet_local = configs.DB_tbl_mst_listbet

func Fetch_listbetHome(idcompany string) (helpers.Response, error) {
	var obj entities.Model_lisbet
	var arraobj []entities.Model_lisbet
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	tbl_mst_listbet, _, _, _ := Get_mappingdatabase(idcompany)

	sql_select := `SELECT 
			idbet_listbet , minbet_listbet,  
			create_listbet, to_char(COALESCE(createdate_listbet,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			update_listbet, to_char(COALESCE(updatedate_listbet,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + tbl_mst_listbet + `  
			WHERE idcompany=$1   
			ORDER BY createdate_listbet DESC   `

	row, err := con.QueryContext(ctx, sql_select, idcompany)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idbet_listbet_db                                                                   int
			minbet_listbet_db                                                                  float64
			create_listbet_db, createdate_listbet_db, update_listbet_db, updatedate_listbet_db string
		)

		err = row.Scan(&idbet_listbet_db, &minbet_listbet_db,
			&create_listbet_db, &createdate_listbet_db, &update_listbet_db, &updatedate_listbet_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		if create_listbet_db != "" {
			create = create_listbet_db + ", " + createdate_listbet_db
		}
		if update_listbet_db != "" {
			update = update_listbet_db + ", " + updatedate_listbet_db
		}

		obj.Lisbet_id = idbet_listbet_db
		obj.Lisbet_minbet = minbet_listbet_db
		obj.Lisbet_create = create
		obj.Lisbet_update = update
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
func Save_listbet(admin, idcompany, sData string, idrecord int, minbet float64) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	tbl_mst_listbet, _, _, _ := Get_mappingdatabase(idcompany)

	if sData == "New" {
		sql_insert := `
			insert into
			` + tbl_mst_listbet + ` (
				idbet , minbet_listbet,  
				create_listbet, createdate_listbet 
			) values (
				$1, $2, 
				$3, $4  
			)
			`

		field_column := tbl_mst_listbet + tglnow.Format("YYYY")
		idrecord_counter := Get_counter(field_column)
		flag_insert, msg_insert := Exec_SQL(sql_insert, tbl_mst_listbet, "INSERT",
			tglnow.Format("YY")+strconv.Itoa(idrecord_counter), minbet,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"))

		if flag_insert {
			msg = "Succes"
		} else {
			fmt.Println(msg_insert)
		}
	} else {
		sql_update := `
				UPDATE 
				` + database_listbet_local + `  
				SET minbet_listbet=$1, 
				update_listbet=$2, updatedate_listbet=$3       
				WHERE idbet=$4     
			`

		flag_update, msg_update := Exec_SQL(sql_update, database_listbet_local, "UPDATE",
			minbet,
			admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord)

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
