package models

import (
	"context"
)

type Colleague struct{}

func (Colleague) AuthenticationByUserName(ctx context.Context, identiKey string, password string) (int64, bool, error) {
	return int64(1), true, nil
}
