package names

import (
	"fmt"
	"reflect"

	"github.com/hngprojects/hng_boilerplate_golang_web/utility"
)

type NotificationName string

const (
	SendWelcomeMail           NotificationName = "send_welcome_mail"
	SendOTP                   NotificationName = "send_otp"
	SendResetPasswordMail     NotificationName = "send_reset_password_mail"
	SendEmailVerificationMail NotificationName = "send_email_verification_mail"
	SendMagicLink             NotificationName = "send_magic_link"
)

func Check() {
	constantName := "SendWelcomeMail"
	constantValue := reflect.ValueOf(constantName).Interface().(string)
	fmt.Println("check", constantValue)
}

func GetNames(pkgImportPath string) ([]string, error) {
	names := []string{}
	constants, err := utility.GetConstants(pkgImportPath)
	if err != nil {
		return names, err
	}

	for _, v := range constants {
		names = append(names, v)
	}

	return names, nil
}
