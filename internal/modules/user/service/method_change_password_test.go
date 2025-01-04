package user_service

import (
	"context"
	"errors"
	"testing"

	user_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/model"
	mock_repository "github.com/aziemp66/dot-indonesia-technical-test/mock/repository"
	mock_util "github.com/aziemp66/dot-indonesia-technical-test/mock/util"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestUserServiceChangePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repoMock := mock_repository.NewMockUserRepository(ctrl)
	jwtMock := mock_util.NewMockJWTManager(ctrl)
	passwordMock := mock_util.NewMockPasswordManager(ctrl)

	service := NewUserService(repoMock, jwtMock, passwordMock)

	id := uuid.New()
	oldPassword := "oldPassword123"
	newPassword := "newPassword456"

	t.Run("should successfully change password", func(t *testing.T) {
		repoRes := user_model.User{
			ID:       id,
			Email:    "john@google.com",
			Password: "hashed_old_password",
			Name:     "John",
		}
		hashedPassword := "hashed_new_password"

		repoMock.EXPECT().GetUserByID(gomock.Any(), id).Return(repoRes, nil)

		passwordMock.EXPECT().CheckPasswordHash(oldPassword, repoRes.Password).Return(nil)

		passwordMock.EXPECT().PasswordValidation(newPassword).Return(nil)

		passwordMock.EXPECT().HashPassword(newPassword).Return(hashedPassword, nil)

		repoMock.EXPECT().ChangePassword(gomock.Any(), id, hashedPassword).Return(nil)

		err := service.ChangePassword(context.Background(), id.String(), oldPassword, newPassword)

		require.NoError(t, err)
	})

	t.Run("should return error when failed retrieving db", func(t *testing.T) {
		expectedErr := errors.New("error from db")

		repoMock.EXPECT().GetUserByID(gomock.Any(), "").Return(user_model.User{}, expectedErr)

		err := service.ChangePassword(context.Background(), "", "", "")

		require.EqualError(t, err, expectedErr.Error())
	})
}
