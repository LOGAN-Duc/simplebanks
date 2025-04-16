package db

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/require"
	db "simplebanks/db/sqlc"
	"simplebanks/util"
	"testing"
	"time"
)

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func createRandomUser(t *testing.T) db.User {
	arg := db.CreateUserParams{
		Username:       util.RandomOwner(),
		HashedPassword: util.RandomString(6),
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
	}
	user, err := testStore.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)

	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testStore.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)
	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}

func TestUpdateUserFullName(t *testing.T) {
	oldUser := createRandomUser(t)
	newFullName := util.RandomOwner()
	var updateUse, err = testStore.UpdateUser(context.Background(), db.UpdateUserParams{
		Username: oldUser.Username,
		FullName: sql.NullString{
			String: newFullName,
			Valid:  true,
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, updateUse)
	require.Equal(t, newFullName, updateUse.FullName)
	require.Equal(t, oldUser.Email, updateUse.Email)
	require.Equal(t, oldUser.HashedPassword, updateUse.HashedPassword)
}

func TestUpdateUserEmail(t *testing.T) {
	oldUser := createRandomUser(t)
	newEmail := util.RandomEmail()
	var updateUse, err = testStore.UpdateUser(context.Background(), db.UpdateUserParams{
		Username: oldUser.Username,
		Email: sql.NullString{
			String: newEmail,
			Valid:  true,
		},
	})
	require.NoError(t, err)
	require.NotEmpty(t, updateUse)
	require.Equal(t, newEmail, updateUse.Email)
	require.Equal(t, oldUser.FullName, updateUse.FullName)
	require.Equal(t, oldUser.HashedPassword, updateUse.HashedPassword)
}

//func TestUpdateUserPassword(t *testing.T) {
//	oldUser := createRandomUser(t)
//	newPw := util.RandomString(6)
//	newHashPassword, err := util.HashPassword(newPw)
//	require.NoError(t, err)
//	updateUse, err := testStore.UpdateUser(context.Background(), UpdateUserParams{
//		Username: oldUser.Username,
//		HashedPassword: sql.NullString{
//			String: newHashPassword,
//			Valid:  true,
//		},
//	})
//	require.NoError(t, err)
//	require.NotEqual(t, oldUser.HashedPassword, updateUse.HashedPassword)
//	require.Equal(t, newHashPassword, updateUse.HashedPassword)
//	require.Equal(t, oldUser.FullName, updateUse.FullName)
//	require.Equal(t, oldUser.Email, updateUse.Email)
//}
