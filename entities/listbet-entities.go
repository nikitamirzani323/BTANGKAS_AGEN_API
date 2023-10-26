package entities

type Model_lisbet struct {
	Lisbet_id     int     `json:"lisbet_id"`
	Lisbet_minbet float64 `json:"lisbet_minbet"`
	Lisbet_create string  `json:"lisbet_create"`
	Lisbet_update string  `json:"lisbet_update"`
}
type Model_lisbetshare struct {
	Lisbet_minbet float64 `json:"lisbet_minbet"`
}
type Model_listbet_conf struct {
	Listbetconf_id     int    `json:"listbetconf_id"`
	Listbetconf_idbet  int    `json:"listbetconf_idbet"`
	Listbetconf_nmpoin string `json:"listbetconf_nmpoin"`
	Listbetconf_poin   int    `json:"listbetconf_poin"`
	Listbetconf_create string `json:"listbetconf_create"`
	Listbetconf_update string `json:"listbetconf_update"`
}
type Controller_listbetsave struct {
	Page          string `json:"page" validate:"required"`
	Sdata         string `json:"sdata" validate:"required"`
	Lisbet_id     int    `json:"lisbet_id"`
	Lisbet_minbet int    `json:"lisbet_minbet" validate:"required"`
}
type Controller_listbetconfpoint struct {
	Lisbet_idbet int `json:"listbet_idbet" validate:"required"`
}
