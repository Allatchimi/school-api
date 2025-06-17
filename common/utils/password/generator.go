package password

import (
	"fmt"
	"strings"
	"time"
)

func GeneratePasswordFromUserInfo(firstName string, lastName string, birthday *time.Time) string {
	var formattedBirthday = time.Now().Year()
	firstNameParts := strings.Split(firstName, " ")
	if len(firstNameParts) > 0 {
		firstName = firstNameParts[0]
	}
	lastNameParts := strings.Split(lastName, " ")
	if len(lastNameParts) > 0 {
		lastName = lastNameParts[0]
	}
	if birthday != nil {
		formattedBirthday = birthday.Year()
	}
	var password string = ""
	if len(firstName) > 0 && len(lastName) > 0 && formattedBirthday > 0 {
		password = fmt.Sprintf("%s%s%d", firstName, lastName, formattedBirthday)
	}
	return password
}
