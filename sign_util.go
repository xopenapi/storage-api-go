package storage

import(
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func Sign(data []byte) string {
	m := md5.New()
	m.Write(data)
	signature := hex.EncodeToString(m.Sum(nil))
	fmt.Printf("signature: %s\n", signature)
	return signature
}
