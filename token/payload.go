package token

import (
	"fmt"
	"time"
)

type Payload struct {
	Username string
	IssuedAt time.Time
	ExpiredAt time.Time
}

func NewPayload(username string, duration time.Duration) *Payload {
	return &Payload{
		Username: username,
		IssuedAt: time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
}

func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAt) {
		return fmt.Errorf("token has expired")
	}

	return nil
}