package md5

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"mime/multipart"
)

func MD5WithFile(file multipart.File) string {
	m := md5.New()
	io.Copy(m, file)
	hashBytes := m.Sum(nil)
	return hex.EncodeToString(hashBytes)
}
