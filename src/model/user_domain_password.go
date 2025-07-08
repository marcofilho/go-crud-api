package model

import (
	"crypto/md5"
	"encoding/hex"
)

func (ud *userDomain) EncryptPassword() {
	has := md5.New()
	defer has.Reset()
	has.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(has.Sum(nil))
}
