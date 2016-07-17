package md5

import (
	"crypto/md5"
	"encoding/base64"
)

func Base64MD5(v string) string {
	m := md5.New()
	m.Write([]byte(v))
	return base64.StdEncoding.EncodeToString(m.Sum(nil))
}
