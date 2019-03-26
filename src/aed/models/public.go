package models

import (
	"crypto/md5"
	"fmt"
	"io"
)

func GetMD5(words string) string {
	var str string
	const salt = "ejwqi4123IU5Rerqwe098-*/*35`1.*@#$"

	w:=md5.New()
	io.WriteString(w,words+salt)
	str=fmt.Sprintf("%x",w.Sum(nil))
	return str
}

