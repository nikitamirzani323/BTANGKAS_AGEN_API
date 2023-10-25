package entities

type Model_admin struct {
	Admin_id            string `json:"admin_id"`
	Admin_idcompany     string `json:"admin_idcompany"`
	Admin_username      string `json:"admin_username"`
	Admin_nama          string `json:"admin_nama"`
	Admin_tipe          string `json:"admin_tipe"`
	Admin_idrule        int    `json:"admin_idrule"`
	Admin_rule          string `json:"admin_rule"`
	Admin_phone1        string `json:"admin_phone1"`
	Admin_phone2        string `json:"admin_phone2"`
	Admin_lastlogin     string `json:"admin_lastlogin"`
	Admin_lastIpaddress string `json:"admin_lastipaddres"`
	Admin_status        string `json:"admin_status"`
	Admin_status_css    string `json:"admin_status_css"`
	Admin_create        string `json:"admin_create"`
	Admin_update        string `json:"admin_update"`
}
type Model_adminruleall struct {
	Adminrule_id     int    `json:"adminrule_id"`
	Adminrule_name   string `json:"adminrule_name"`
	Adminrule_rule   string `json:"adminrule_rule"`
	Adminrule_create string `json:"adminrule_create"`
	Adminrule_update string `json:"adminrule_update"`
}
type Model_adminruleshare struct {
	Adminrule_id   int    `json:"adminrule_id"`
	Adminrule_name string `json:"adminrule_name"`
}
type Model_adminsave struct {
	Username string `json:"admin_username"`
	Nama     string `json:"admin_nama"`
	Rule     string `json:"admin_rule"`
	Status   string `json:"admin_status"`
	Create   string `json:"admin_create"`
	Update   string `json:"admin_update"`
}
type Controller_admindetail struct {
	Username string `json:"admin_username" validate:"required"`
}
type Controller_adminsave struct {
	Sdata          string `json:"sdata" validate:"required"`
	Page           string `json:"page" validate:"required"`
	Admin_id       string `json:"admin_id" `
	Admin_username string `json:"admin_username" validate:"required"`
	Admin_password string `json:"admin_password" `
	Admin_nama     string `json:"admin_nama" validate:"required"`
	Admin_idrule   int    `json:"admin_idrule"`
	Admin_phone1   string `json:"admin_phone1"`
	Admin_phone2   string `json:"admin_phone2"`
	Admin_status   string `json:"admin_status"`
}
type Controller_adminrulesave struct {
	Sdata          string `json:"sdata" validate:"required"`
	Page           string `json:"page" validate:"required"`
	Adminrule_id   int    `json:"adminrule_id"`
	Adminrule_name string `json:"adminrule_name" validate:"required"`
	Adminrule_rule string `json:"adminrule_rule"`
}
