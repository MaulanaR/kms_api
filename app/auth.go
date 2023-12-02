package app

import (
	"time"
)

func RemoveExpiredToken() {
	tx, err := DB().Conn("main")
	if err == nil {
		tx.Where("expired_at < ?", time.Now()).Delete("access_token")
	}
}
