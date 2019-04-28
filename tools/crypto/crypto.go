package crypto

import (
	"crypto/md5"
    "encoding/hex"
)

func Md5(salt, pwdStr string) string {
    h := md5.New()
    h.Write([]byte(salt + pwdStr))
    return hex.EncodeToString(h.Sum(nil))
}
