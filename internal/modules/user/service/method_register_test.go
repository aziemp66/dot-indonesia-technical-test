package user_service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"testing"

	user_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/model"
	mock_repository "github.com/aziemp66/dot-indonesia-technical-test/mock/repository"
	mock_util "github.com/aziemp66/dot-indonesia-technical-test/mock/util"
	util_error "github.com/aziemp66/dot-indonesia-technical-test/util/error"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUserServiceRegister(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mock_repository.NewMockUserRepository(ctrl)
	jwtMock := mock_util.NewMockJWTManager(ctrl)
	passwordMock := mock_util.NewMockPasswordManager(ctrl)
	redisMock := mock_util.NewMockRedisManager(ctrl)

	service := NewUserService(repoMock, jwtMock, passwordMock, redisMock)

	reqEmail := "test@example.com"
	reqPassword := "password123"
	reqName := "Test User"

	hashedPassword := "secured_password"

	resID := uuid.New()

	t.Run("should register a new user", func(t *testing.T) {
		repoMock.EXPECT().GetUserByEmail(gomock.Any(), reqEmail).Return(user_model.User{}, sql.ErrNoRows)

		passwordMock.EXPECT().PasswordValidation(reqPassword).Return(nil)

		passwordMock.EXPECT().HashPassword(reqPassword).Return(hashedPassword, nil)

		repoMock.EXPECT().CreateUser(gomock.Any(), reqEmail, hashedPassword, reqName).
			Return(resID.String(), nil)

		id, err := service.Register(context.Background(), reqEmail, reqPassword, reqName)

		require.NoError(t, err)
		assert.Equal(t, resID.String(), id)
	})

	t.Run("should return error when failed retrieving from db", func(t *testing.T) {
		expectedErr := errors.New("db error")

		repoMock.EXPECT().GetUserByEmail(gomock.Any(), reqEmail).Return(user_model.User{}, expectedErr)

		id, err := service.Register(context.Background(), reqEmail, reqPassword, reqName)

		assert.Error(t, err)
		assert.Empty(t, id)
		assert.EqualError(t, err, expectedErr.Error())
	})

	t.Run("should return client error when email is already used", func(t *testing.T) {
		repoRes := user_model.User{
			ID:       resID,
			Email:    reqEmail,
			Password: reqPassword,
			Name:     reqName,
		}
		repoMock.EXPECT().GetUserByEmail(gomock.Any(), reqEmail).Return(repoRes, nil)

		id, err := service.Register(context.Background(), reqEmail, reqPassword, reqName)

		expectedErr := util_error.NewBadRequest(fmt.Errorf("%s is already registered", reqEmail), fmt.Sprintf("Email %s is already used", reqEmail))

		assert.Error(t, err)
		assert.Empty(t, id)
		assert.EqualError(t, err, expectedErr.Error())
	})
}
