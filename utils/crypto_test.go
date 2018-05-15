package utils

import "testing"

func TestEncrypt(t *testing.T) {
	KEY := "LKHlhb899Y09olUi"
	var plainTxt = "salamander"
	enTxt, _ := Encrypt(plainTxt, []byte(KEY))
	res, err := Decrypt(enTxt, []byte(KEY))
	if err != nil {
		t.Error(err.Error())
		return
	}
	if res == plainTxt {
		t.Log("加密解密成功")
	} else {
		t.Error("加密解密失败：" + res)
	}
}