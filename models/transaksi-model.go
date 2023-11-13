package models

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/configs"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/db"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/entities"
	"github.com/nikitamirzani323/BTANGKAS_AGEN_API/helpers"
)

func Fetch_transaksiHome(idcompany string) (helpers.Response, error) {
	var obj entities.Model_transaksi
	var arraobj []entities.Model_transaksi
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	_, _, tbl_trx_transaksi, _ := Get_mappingdatabase(idcompany)

	sql_select := `SELECT 
			A.idtransaksi , A.username_client, 
			A.roundbet , A.total_bet,  A.total_win, A.total_bonus,
			A.card_codepoin, COALESCE(B.nmpoin,'') as nmpoin, A.card_pattern,  A.card_result, A.card_win,
			A.create_transaksi, to_char(COALESCE(A.createdate_transaksi,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			A.update_transaksi, to_char(COALESCE(A.updatedate_transaksi,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + tbl_trx_transaksi + ` as A   
			LEFT JOIN ` + configs.DB_tbl_mst_listpoint + ` as B ON B.codepoin = A.card_codepoin    
			WHERE A.idcompany=$1   
			ORDER BY A.createdate_transaksi DESC   LIMIT 500 `
	log.Println("COMPANY : " + strings.ToLower(idcompany))
	row, err := con.QueryContext(ctx, sql_select, strings.ToLower(idcompany))
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idtransaksi_db, username_client_db                                                         string
			roundbet_db, total_bet_db, total_win_db, total_bonus_db                                    int
			card_codepoin_db, nmpoin_db, card_pattern_db, card_result_db, card_win_db                  string
			create_transaksi_db, createdate_transaksi_db, update_transaksi_db, updatedate_transaksi_db string
		)

		err = row.Scan(&idtransaksi_db, &username_client_db,
			&roundbet_db, &total_bet_db, &total_win_db, &total_bonus_db,
			&card_codepoin_db, &nmpoin_db, &card_pattern_db, &card_result_db, &card_win_db,
			&create_transaksi_db, &createdate_transaksi_db, &update_transaksi_db, &updatedate_transaksi_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		if create_transaksi_db != "" {
			create = create_transaksi_db + ", " + createdate_transaksi_db
		}
		if update_transaksi_db != "" {
			update = update_transaksi_db + ", " + updatedate_transaksi_db
		}

		obj.Transaksi_id = idtransaksi_db
		obj.Transaksi_date = createdate_transaksi_db
		obj.Transaksi_username = username_client_db
		obj.Transaksi_roundbet = roundbet_db
		obj.Transaksi_totalbet = total_bet_db
		obj.Transaksi_totalwin = total_win_db
		obj.Transaksi_totalbonus = total_bonus_db
		obj.Transaksi_card_codepoin = card_codepoin_db + " - " + nmpoin_db
		obj.Transaksi_card_pattern = card_pattern_db
		obj.Transaksi_card_result = card_result_db
		obj.Transaksi_card_win = card_win_db
		obj.Transaksi_create = create
		obj.Transaksi_update = update
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
func Fetch_transaksidetailHome(idtransaksi, idcompany string) (helpers.Response, error) {
	var obj entities.Model_transaksidetail
	var arraobj []entities.Model_transaksidetail
	var res helpers.Response
	msg := "Data Not Found"
	con := db.CreateCon()
	ctx := context.Background()
	start := time.Now()

	_, _, _, tbl_trx_transaksidetail := Get_mappingdatabase(idcompany)

	log.Println("INVOICE : ", idtransaksi)
	log.Println("COMPANY : ", strings.ToLower(idcompany))

	sql_select := `SELECT 
			idtransaksidetail , 
			roundbet_detail , bet,  credit_before, credit_after, win, bonus, 
			codepoin , resultcard_win, status_transaksidetail, 
			create_transaksidetail, to_char(COALESCE(createdate_transaksidetail,now()), 'YYYY-MM-DD HH24:MI:SS'), 
			update_transaksidetail, to_char(COALESCE(updatedate_transaksidetail,now()), 'YYYY-MM-DD HH24:MI:SS') 
			FROM ` + tbl_trx_transaksidetail + `  
			WHERE idtransaksi=$1   
			ORDER BY createdate_transaksidetail ASC   `
	row, err := con.QueryContext(ctx, sql_select, idtransaksi)
	helpers.ErrorCheck(err)
	for row.Next() {
		var (
			idtransaksidetail_db                                                                                               string
			roundbet_detail_db, bet_db, credit_before_db, credit_after_db, win_db, bonus_db                                    int
			codepoin_db, resultcard_win_db, status_transaksidetail_db                                                          string
			create_transaksidetail_db, createdate_transaksidetail_db, update_transaksidetail_db, updatedate_transaksidetail_db string
		)

		err = row.Scan(&idtransaksidetail_db,
			&roundbet_detail_db, &bet_db, &credit_before_db, &credit_after_db, &win_db, &bonus_db,
			&codepoin_db, &resultcard_win_db, &status_transaksidetail_db,
			&create_transaksidetail_db, &createdate_transaksidetail_db, &update_transaksidetail_db, &updatedate_transaksidetail_db)

		helpers.ErrorCheck(err)
		create := ""
		update := ""
		if create_transaksidetail_db != "" {
			create = create_transaksidetail_db + ", " + createdate_transaksidetail_db
		}
		if update_transaksidetail_db != "" {
			update = update_transaksidetail_db + ", " + updatedate_transaksidetail_db
		}
		status_css := configs.STATUS_CANCEL
		if status_transaksidetail_db == "WIN" {
			status_css = configs.STATUS_COMPLETE
		}

		obj.Transaksidetail_id = idtransaksidetail_db
		obj.Transaksidetail_date = createdate_transaksidetail_db
		obj.Transaksidetail_roundbet = roundbet_detail_db
		obj.Transaksidetail_bet = bet_db
		obj.Transaksidetail_creditbefore = credit_before_db
		obj.Transaksidetail_creditafter = credit_after_db
		obj.Transaksidetail_win = win_db
		obj.Transaksidetail_bonus = bonus_db
		obj.Transaksidetail_card_codepoin = codepoin_db
		obj.Transaksidetail_card_win = resultcard_win_db
		obj.Transaksidetail_status = status_transaksidetail_db
		obj.Transaksidetail_status_css = status_css
		obj.Transaksidetail_create = create
		obj.Transaksidetail_update = update
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
