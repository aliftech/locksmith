package helpers

func Message(responseStatus string, data string) string {
	if data == "" {

	}

	messages := map[string]string{
		"email_exist":          "Email already registered",
		"user_not_found":       "User not found.",
		"parse_reqbody_failed": "Failed to parse request body.",
		"validation_error":     "Validation Error!",
		"signup_failed":        "Failed to create new account! " + data,
		"signup_success":       "New account have been created.",
	}

	if msg, exist := messages[responseStatus]; exist {
		return msg
	}

	return "Response not found."
}
