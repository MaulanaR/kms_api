package i18n

func IdID() map[string]string {
	return map[string]string{
		"400_bad_request":              "Permintaan tidak dapat dilakukan karena ada parameter yang salah atau tidak lengkap.",
		"401_unauthorized":             "Token otentikasi tidak valid. Silakan logout dan login ulang",
		"403_forbidden":                "Pengguna tidak memiliki izin untuk :action.",
		"404_not_found":                "The resource you have specified cannot be found.",
		"500_internal_error":           "Gagal terhubung ke server, silakan coba lagi nanti.",
		"invalid_username_or_password": "Username atau kata sandi tidak valid",
		"unique":                       ":attribute (:value) sudah ada dan tidak bisa digunakan lagi.",
		"username_wrong":               "Username tidak ditemukan.",
		"password_wrong":               "Username/Password tidak sesuai.",
		"extenstion_not_in":            "Ekstensi file harus salah satu dari: (:value).",
		"not_found":                    "Data :entity dengan :key = :value tidak ditemukan.",
		"success":                      "Aksi berhasil dilakukan.",
	}
}
