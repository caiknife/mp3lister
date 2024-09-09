package pay

import (
	"github.com/go-pay/gopay/apple"
	"github.com/pkg/errors"
)

type Apple struct {
	*apple.Client
}

func NewApple(iss, bid, kid, privateKey string, isProduction bool) (*Apple, error) {
	c, err := apple.NewClient(iss, bid, kid, privateKey, isProduction)
	if err != nil {
		err = errors.WithMessage(err, "apple new client")
		return nil, err
	}
	return &Apple{Client: c}, err
}
