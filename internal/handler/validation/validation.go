package validation

import (
	"errors"
	"github.com/vtotbl/test_const/internal/handler/requests"
)

func SenReq(req requests.SendReq) error {
	if len(req.Urls) > 20 {
		return errors.New("urls cannot be more than 20")
	}

	return nil
}
