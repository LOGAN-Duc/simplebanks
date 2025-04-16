package token

import (
	"fmt"
	"github.com/o1egl/paseto"
	"time"
)

type PaseMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func NewPasetoMaker(symmetricKey string) (Maker, error) {
	if len(symmetricKey) != 32 {
		return nil, fmt.Errorf("symmetric key must be 32 bytes")
	}
	maker := &PaseMaker{
		symmetricKey: []byte(symmetricKey),
		paseto:       paseto.NewV2(),
	}
	return maker, nil
}
func (maker *PaseMaker) CreateToken(username string, duration time.Duration) (string, *Payload, error) {
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", payload, err
	}

	token, err := maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	return token, payload, err
}

// Kiem tra ma dau vao co hop la khong
func (maker *PaseMaker) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}
	err = payload.Valid()
	if err != nil {
		return nil, err
	}
	return payload, nil
}
