package models

import (
	"context"
	"database/sql"
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

func Fetch_listbetHome(idcompany string) (helpers.ResponseListbet, error) {
	var obj entities.Model_lisbet
	var arraobj []entities.Model_lisbet
	var objmstlisbet entities.Model_lisbetshare
	var arraobjmstlisbet []entities.Model_lisbetshare
	var res helpers.ResponseListbet
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
			ORDER BY minbet_listbet ASC   `

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

	sql_selectmasterlistbet := `SELECT 
			minbet_listbet  
			FROM ` + configs.DB_tbl_mst_listbet + ` 
			ORDER BY minbet_listbet ASC    
	`
	rowcurr, errcurr := con.QueryContext(ctx, sql_selectmasterlistbet)
	helpers.ErrorCheck(errcurr)
	for rowcurr.Next() {
		var (
			minbet_listbet_db float64
		)

		errcurr = rowcurr.Scan(&minbet_listbet_db)

		helpers.ErrorCheck(errcurr)

		objmstlisbet.Lisbet_minbet = minbet_listbet_db
		arraobjmstlisbet = append(arraobjmstlisbet, objmstlisbet)
		msg = "Success"
	}
	defer rowcurr.Close()

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = arraobj
	res.ListBet = arraobjmstlisbet
	res.Time = time.Since(start).String()

	return res, nil
}
func Fetch_listbetConfPoint(idbet int, idcompany string) (helpers.Response, error) {
	var obj entities.Model_listbet_conf
	var arraobj []entities.Model_listbet_conf
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	_, tbl_mst_config, _, _ := Get_mappingdatabase(idcompany)

	sql_select := `SELECT 
			A.idconf_conf, A.idbet_listbet, A.idpoin, A.poin_conf,  B.nmpoin,
			A.create_conf, to_char(COALESCE(A.createdate_conf,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			A.update_conf, to_char(COALESCE(A.updatedate_conf,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + tbl_mst_config + ` as A   
			JOIN ` + configs.DB_tbl_mst_listpoint + ` as B ON B.idpoin = A.idpoin    
			WHERE A.idbet_listbet=$1 
			AND A.idcompany=$2  
			ORDER BY B.display_listpoint ASC   `

	row, err := con.QueryContext(ctx, sql_select, idbet, idcompany)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idconf_conf_db, idbet_listbet_db, idpoin_db, poin_conf_db              int
			nmpoin_db                                                              string
			create_conf_db, createdate_conf_db, update_conf_db, updatedate_conf_db string
		)

		err = row.Scan(&idconf_conf_db, &idbet_listbet_db, &idpoin_db, &poin_conf_db, &nmpoin_db,
			&create_conf_db, &createdate_conf_db, &update_conf_db, &updatedate_conf_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		if create_conf_db != "" {
			create = create_conf_db + ", " + createdate_conf_db
		}
		if update_conf_db != "" {
			update = update_conf_db + ", " + updatedate_conf_db
		}
		obj.Listbetconf_id = idconf_conf_db
		obj.Listbetconf_idbet = idbet_listbet_db
		obj.Listbetconf_nmpoin = nmpoin_db
		obj.Listbetconf_poin = poin_conf_db
		obj.Listbetconf_create = create
		obj.Listbetconf_update = update
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
func Save_listbet(admin, idcompany, sData string, idrecord, minbet int) (helpers.Response, error) {
	var res helpers.Response
	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()
	flag := false
	tbl_mst_listbet, _, _, _ := Get_mappingdatabase(idcompany)

	if sData == "New" {
		flag = CheckDB(tbl_mst_listbet, "minbet_listbet", strconv.Itoa(minbet))
		if !flag {
			sql_insert := `
			insert into
			` + tbl_mst_listbet + ` (
				idbet_listbet , idcompany,  minbet_listbet,  
				create_listbet, createdate_listbet 
			) values (
				$1, $2, $3,  
				$4, $5   
			)
			`

			field_column := tbl_mst_listbet + tglnow.Format("YYYY")
			idrecord_counter := Get_counter(field_column)
			flag_insert, msg_insert := Exec_SQL(sql_insert, tbl_mst_listbet, "INSERT",
				tglnow.Format("YY")+strconv.Itoa(idrecord_counter), idcompany, minbet,
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
func Save_ConfPoint(admin, idcompany, sData string, idrecord, idbet, point int) (helpers.Response, error) {
	var res helpers.Response

	msg := "Failed"
	tglnow, _ := goment.New()
	render_page := time.Now()

	_, tbl_mst_config, _, _ := Get_mappingdatabase(idcompany)

	sql_update := `
			UPDATE 
			` + tbl_mst_config + `  
			SET poin_conf=$1,  
			update_conf=$2, updatedate_conf=$3       
			WHERE idconf_conf=$4 AND idcompany=$5 AND idbet_listbet=$6       
		`

	flag_update, msg_update := Exec_SQL(sql_update, tbl_mst_config, "UPDATE",
		point,
		admin, tglnow.Format("YYYY-MM-DD HH:mm:ss"), idrecord, idcompany, idbet)

	if flag_update {
		msg = "Succes"
	} else {
		fmt.Println(msg_update)
	}

	res.Status = fiber.StatusOK
	res.Message = msg
	res.Record = nil
	res.Time = time.Since(render_page).String()

	return res, nil
}
func _Get_infomasterpoint(idpoin int) string {
	con := db.CreateCon()
	ctx := context.Background()
	nmpoin := ""
	sql_select := `SELECT
			nmpoin    
			FROM ` + configs.DB_tbl_mst_listpoint + `  
			WHERE idpoin='` + strconv.Itoa(idpoin) + `'     
		`

	row := con.QueryRowContext(ctx, sql_select)
	switch e := row.Scan(&nmpoin); e {
	case sql.ErrNoRows:
	case nil:
	default:
		helpers.ErrorCheck(e)
	}

	return nmpoin
}
