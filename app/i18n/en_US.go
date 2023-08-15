package i18n

func EnUS() map[string]string {
	return map[string]string{
		"400_bad_request":              "The request cannot be performed because of malformed or missing parameters.",
		"401_unauthorized":             "Unauthorized. Please Re-Login",
		"403_forbidden":                "The user does not have permission to :action.",
		"404_not_found":                "The resource you have specified cannot be found.",
		"500_internal_error":           "Failed to connect to the server, please try again later.",
		"invalid_username_or_password": "Invalid username or password",
		"unique":                       ":attribute (:value) sudah ada dan tidak bisa digunakan lagi.",
		"username_wrong":               "Username is not found.",
		"password_wrong":               "Username/Password not valid.",
		"extension_not_in":             "The file extension must be one of the following types: (:value).",
		"not_found":                    ":entity data with :key = :value cannot be found.",
		"success":                      "The action was successful.",
	}
}
