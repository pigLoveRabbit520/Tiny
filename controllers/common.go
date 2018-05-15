package controllers

type ApiRes struct {
	Errcode uint        `json:"errcode"`
	Errmsg  string      `json:"errmsg"`
	Data    interface{} `json:"data"`
}
