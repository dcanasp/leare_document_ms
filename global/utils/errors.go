package utils

import logs "global/logging"

func ErrorFail(err error, msg string) {
	if err != nil {
		logs.E.Panicf("%s: %s", msg, err)
	}
}

func Error(err error, msg string) {
	if err != nil {
		logs.E.Printf("%s: %s", msg, err)
	}
}
