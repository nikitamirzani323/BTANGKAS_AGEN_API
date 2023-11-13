package entities

type Model_transaksi struct {
	Transaksi_id            string `json:"transaksi_id"`
	Transaksi_date          string `json:"transaksi_date"`
	Transaksi_username      string `json:"transaksi_username"`
	Transaksi_roundbet      int    `json:"transaksi_roundbet"`
	Transaksi_totalbet      int    `json:"transaksi_totalbet"`
	Transaksi_totalwin      int    `json:"transaksi_totalwin"`
	Transaksi_totalbonus    int    `json:"transaksi_totalbonus"`
	Transaksi_card_codepoin string `json:"transaksi_card_codepoin"`
	Transaksi_card_pattern  string `json:"transaksi_card_pattern"`
	Transaksi_card_result   string `json:"transaksi_card_result"`
	Transaksi_card_win      string `json:"transaksi_card_win"`
	Transaksi_create        string `json:"transaksi_create"`
	Transaksi_update        string `json:"transaksi_update"`
}
type Model_transaksidetail struct {
	Transaksidetail_id            string `json:"transaksidetail_id"`
	Transaksidetail_date          string `json:"transaksidetail_date"`
	Transaksidetail_roundbet      int    `json:"transaksidetail_roundbet"`
	Transaksidetail_bet           int    `json:"transaksidetail_bet"`
	Transaksidetail_creditbefore  int    `json:"transaksidetail_creditbefore"`
	Transaksidetail_creditafter   int    `json:"transaksidetail_creditafter"`
	Transaksidetail_win           int    `json:"transaksidetail_win"`
	Transaksidetail_bonus         int    `json:"transaksidetail_bonus"`
	Transaksidetail_card_codepoin string `json:"transaksidetail_card_codepoin"`
	Transaksidetail_card_win      string `json:"transaksidetail_card_win"`
	Transaksidetail_status        string `json:"transaksidetail_status"`
	Transaksidetail_status_css    string `json:"transaksidetail_status_css"`
	Transaksidetail_create        string `json:"transaksidetail_create"`
	Transaksidetail_update        string `json:"transaksidetail_update"`
}

type Controller_transaksidetail struct {
	Transaksi_id string `json:"transaksi_id" validate:"required"`
}
