package gotostatment

import "static/errorstatment"

func MakeErr() (res int, msg error) {
	deErr := errorstatment.Defines{Err1: 1}
	msg = deErr.Error()
	return
}
