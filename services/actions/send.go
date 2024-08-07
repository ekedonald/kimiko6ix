package actions

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"

	"github.com/hngprojects/hng_boilerplate_golang_web/external/request"
	"github.com/hngprojects/hng_boilerplate_golang_web/internal/models"
	"github.com/hngprojects/hng_boilerplate_golang_web/services/actions/names"
	notifications "github.com/hngprojects/hng_boilerplate_golang_web/services/notification"
)

func Send(extReq request.ExternalRequest, db *gorm.DB, rdb *redis.Client, notification *models.NotificationRecord) error {
	var (
		err  error
		req  = notifications.NewNotificationObject(extReq, rdb, db, notification)
		name = GetName(notification.Name)
	)

	callEmailFunc := map[names.NotificationName]func() error{
		names.SendOTP: func() error {
			return req.SendOTP()
		},
		names.SendWelcomeMail: func() error {
			return req.SendWelcomeMail()
		},
		names.SendResetPasswordMail: func() error {
			return req.SendResetPasswordMail()
		},
		names.SendEmailVerificationMail: func() error {
			return req.SendEmailVerificationMail()
		},
		names.SendMagicLink: func() error {
			return req.SendMagicLink()
		},
	}

	err = callEmailFunc[name]()

	if err != nil {
		return err
	}

	return nil
}

func GetName(name string) names.NotificationName {
	return names.NotificationName(name)
}
