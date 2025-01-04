package user_service

import (
	"context"
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

func TestUserServiceUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mock_repository.NewMockUserRepository(ctrl)
	jwtMock := mock_util.NewMockJWTManager(ctrl)
	passwordMock := mock_util.NewMockPasswordManager(ctrl)

	service := NewUserService(repoMock, jwtMock, passwordMock)

	id := uuid.New()
	name := "John Doe"

	repoRes := user_model.User{
		ID:       id,
		Email:    "johndoe@gmail.com",
		Password: "secure_password",
		Name:     name,
	}

	t.Run("should update user successfully", func(t *testing.T) {
		repoMock.EXPECT().GetUserByID(gomock.Any(), id).Return(repoRes, nil)

		repoMock.EXPECT().UpdateUser(gomock.Any(), id, name).Return(nil)

		err := service.UpdateUser(context.Background(), id.String(), name)

		require.NoError(t, err)
	})

	t.Run("should return error when update user fails", func(t *testing.T) {
		expectedErr := errors.New("failed to update user")

		repoMock.EXPECT().GetUserByID(gomock.Any(), "").Return(user_model.User{}, expectedErr)

		err := service.UpdateUser(context.Background(), "", "")

		require.Error(t, err)
		assert.EqualError(t, err, expectedErr.Error())
	})

	t.Run("should return client error when user is not found", func(t *testing.T) {
		expectedErr := util_error.NewNotFound(fmt.Errorf("user with the id of %s is not found", id), "User not found")
		repoMock.EXPECT().GetUserByID(gomock.Any(), id).Return(user_model.User{}, expectedErr)

		err := service.UpdateUser(context.Background(), id.String(), "")

		require.Error(t, err)
		assert.EqualError(t, err, expectedErr.Error())
	})
}
