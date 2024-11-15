package token

import "time"

type Maker interface {
	//create new token
	CreateToken(username string, duration time.Duration) (string, error)

	//VerifyToken
	VerifyToken(token string) (*Payload, error)
}
