package helper

import "fmt"

var Salt = "password-salt"

func SaltString(value string) string {
	return fmt.Sprintf("%s.%s", value, Salt)
}
