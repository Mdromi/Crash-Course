package validators

import "regexp"

var rxMail = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func IsEmailValid(email string) bool {
	if len(email) < 3 || len(email) > 254 || !rxMail.MatchString(email) {
		return false
	}
	return true
}
