package user_service

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	user_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/model"
	mock_repository "github.com/aziemp66/dot-indonesia-technical-test/mock/repository"
	mock_util "github.com/aziemp66/dot-indonesia-technical-test/mock/util"
	util_error "github.com/aziemp66/dot-indonesia-technical-test/util/error"
	util_jwt "github.com/aziemp66/dot-indonesia-technical-test/util/jwt"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUserServiceLogin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mock_repository.NewMockUserRepository(ctrl)
	jwtMock := mock_util.NewMockJWTManager(ctrl)
	passwordMock := mock_util.NewMockPasswordManager(ctrl)
	redisMock := mock_util.NewMockRedisManager(ctrl)

	service := NewUserService(repoMock, jwtMock, passwordMock, redisMock)

	reqEmail := "test@example.com"
	reqPassword := "password123"

	repoRes := user_model.User{
		ID:       uuid.New(),
		Email:    reqEmail,
		Password: "secure_password",
		Name:     "John Doe",
	}

	t.Run("should login user successfully", func(t *testing.T) {
		resToken := "generated_token"

		repoMock.EXPECT().GetUserByEmail(gomock.Any(), reqEmail).
			Return(repoRes, nil)

		passwordMock.EXPECT().CheckPasswordHash(reqPassword, repoRes.Password).
			Return(nil)

		jwtMock.EXPECT().GenerateAuthToken(repoRes.ID.String(), repoRes.Name, util_jwt.USER_ROLE, gomock.Any()).
			Return(resToken, nil)

		token, err := service.Login(context.Background(), reqEmail, reqPassword)

		assert.Nil(t, err)
		assert.Equal(t, resToken, token)
	})

	t.Run("should return error when failed retrieving from db", func(t *testing.T) {
		expectedErr := errors.New("db error")

		repoMock.EXPECT().GetUserByEmail(gomock.Any(), reqEmail).
			Return(user_model.User{}, expectedErr)

		_, err := service.Login(context.Background(), reqEmail, reqPassword)

		require.Error(t, err)
	})

	t.Run("should return unauthorized error when email not found", func(t *testing.T) {
		expectedErr := util_error.NewUnauthorized(sql.ErrNoRows, "Wrong Email or Password")

		repoMock.EXPECT().GetUserByEmail(gomock.Any(), reqEmail).Return(user_model.User{}, sql.ErrNoRows)

		_, err := service.Login(context.Background(), reqEmail, reqPassword)

		require.EqualError(t, err, expectedErr.Error())
	})
}
