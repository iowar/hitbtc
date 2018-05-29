package hitbtc

import (
	"errors"
)

var (
	ConnectError  = "[ERROR] Connection could not be established!"
	RequestError  = "[ERROR] NewRequest Error!"
	SetApiError   = "[ERROR] Set the API KEY and API SECRET!"
	PeriodError   = "[ERROR] Invalid Period!"
	UrlParseError = "[ERROR] Url Parse Error"
	ServerError   = "[SERVER ERROR] Response: %s"
)

func Error(msg string, args ...interface{}) error {
	if len(args) > 0 {
		return errors.New(spr(msg, args))
	} else {
		return errors.New(msg)
	}
}
