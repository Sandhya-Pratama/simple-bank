package token

import (
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
)

type PasetoMarker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

// new paseto marker created
func NewPasetoMaker(synmetricKey string) (Maker, error) {
	if len(synmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size must be %d characters", chacha20poly1305.KeySize)
	}

	maker := &PasetoMarker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(synmetricKey),
	}

	return maker, nil
}

func (maker *PasetoMarker) CreateToken(username string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}

	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	return token, err
}

func (marker *PasetoMarker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}
	err := marker.paseto.Decrypt(token, marker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}
	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil

}
