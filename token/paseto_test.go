package token

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
	"simplebanks/util"
	"testing"
	"time"
)

func TestPoseMaker(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))

	require.NoError(t, err)

	username := util.RandomOwner()
	duration := time.Minute

	issuedAt := time.Now()
	expiredAt := issuedAt.Add(duration)

	token, payload, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.Equal(t, payload.Username, username)
	require.NotZero(t, payload.Id)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, duration)
	require.WithinDuration(t, expiredAt, payload.ExpiredAt, duration)
}
func TestExpiredPaseToken(t *testing.T) {
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)
	username := util.RandomOwner()
	duration := -time.Minute

	token, payload, err := maker.CreateToken(username, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)
	require.NotEmpty(t, payload)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)

	require.EqualError(t, err, ErrExpiredToken.Error())
}
func TestPaseInvalidTokenAlgNone(t *testing.T) {
	username := util.RandomOwner()
	duration := time.Minute
	payload, err := NewPayload(username, duration)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}
