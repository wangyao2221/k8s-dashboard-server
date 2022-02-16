package helper

import (
	"crypto/md5"
	"encoding/hex"
)

func (s *Service) Md5(str string) (md5Str string) {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}
