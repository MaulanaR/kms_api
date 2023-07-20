package app

import (
	"time"
)

func RemoveExpiredToken() {
	tx, err := DB().Conn("main")
	if err == nil {
		tx.Delete("access_token").
			Where("expired_at < ?", time.Now()). // sesuai umur refresh token
			Limit(10000)
	}
}
