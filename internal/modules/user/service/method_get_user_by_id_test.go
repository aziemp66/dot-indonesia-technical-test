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

func TestUserServiceGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mock_repository.NewMockUserRepository(ctrl)
	jwtMock := mock_util.NewMockJWTManager(ctrl)
	passwordMock := mock_util.NewMockPasswordManager(ctrl)
	redisMock := mock_util.NewMockRedisManager(ctrl)

	service := NewUserService(repoMock, jwtMock, passwordMock, redisMock)

	idReq := uuid.New()

	userRes := user_model.GetUserResponse{
		ID:    idReq.String(),
		Name:  "John Smith",
		Email: "johnsmith123@gmail.com",
	}

	t.Run("should get user by id", func(t *testing.T) {
		repoRes := user_model.User{
			ID:       idReq,
			Email:    userRes.Email,
			Password: "secured_password",
			Name:     userRes.Name,
		}

		repoMock.EXPECT().GetUserByID(gomock.Any(), idReq).
			Return(repoRes, nil)

		res, err := service.GetUserByID(context.Background(), idReq.String())

		require.Nil(t, err)
		assert.Equal(t, userRes, res)
	})

	t.Run("should return client error when theres no user equal to id requirement", func(t *testing.T) {
		expectedErr := util_error.NewNotFound(sql.ErrNoRows, fmt.Sprintf("User with the id of %s is not found", idReq))

		repoMock.EXPECT().GetUserByID(gomock.Any(), idReq).
			Return(user_model.User{}, sql.ErrNoRows)

		_, err := service.GetUserByID(context.Background(), idReq.String())

		require.Error(t, err)
		assert.EqualError(t, expectedErr, err.Error())
	})

	t.Run("should return error when failed retrieving from db", func(t *testing.T) {
		expectedErr := errors.New("error from db")

		repoMock.EXPECT().GetUserByID(gomock.Any(), idReq).
			Return(user_model.User{}, expectedErr)

		_, err := service.GetUserByID(context.Background(), idReq.String())
		require.Error(t, err)
		assert.EqualError(t, expectedErr, err.Error())
	})
}
