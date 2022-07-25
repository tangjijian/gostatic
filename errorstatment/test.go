package errorstatment

import (
	"errors"
)

type Defines struct {
	Err1 int
}

func (d *Defines) Error() error {
	//fmtstr := `
	//	this is a test error. num is %d
	//`
	//return fmt.Sprintf(fmtstr, d.Err1)
	return errors.New("this is a test error")
}
