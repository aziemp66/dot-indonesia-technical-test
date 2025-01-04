package user_http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	user_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/model"
	mock_service "github.com/aziemp66/dot-indonesia-technical-test/mock/service"
	mock_util "github.com/aziemp66/dot-indonesia-technical-test/mock/util"
	util_error "github.com/aziemp66/dot-indonesia-technical-test/util/error"
	util_http "github.com/aziemp66/dot-indonesia-technical-test/util/http"
	util_http_middleware "github.com/aziemp66/dot-indonesia-technical-test/util/http/middleware"
	util_jwt "github.com/aziemp66/dot-indonesia-technical-test/util/jwt"
	util_logger "github.com/aziemp66/dot-indonesia-technical-test/util/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserHttpHandlerChangePassword(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	serviceMock := mock_service.NewMockUserService(ctrl)
	jwtMock := mock_util.NewMockJWTManager(ctrl)
	handler := NewUserHttpHandler(serviceMock, jwtMock)

	url := "/user/change-password"

	util_logger.InitLogger(gin.TestMode, "test", "./user_http_handler.log")

	app := util_http.NewHTTPServer(gin.TestMode)
	app.Use(
		util_http_middleware.TraceIdAssignmentMiddleware(),
		util_http_middleware.LogHandlerMiddleware(),
		util_http_middleware.ErrorHandlerMiddleware(),
		util_http_middleware.JWTAuthentication(jwtMock),
		util_http_middleware.JWTAuthorization(util_jwt.USER_ROLE),
	)
	app.POST(url, handler.ChangePassword)

	reqBody := user_model.ChangePasswordRequest{
		OldPassword: "old_password123",
		NewPassword: "new_password123",
	}
	reqBodyBytes, _ := json.Marshal(reqBody)

	authToken := "user-token"
	tokenBearer := fmt.Sprintf("BEARER %s", authToken)

	claims := util_jwt.AuthClaims{
		ID:   uuid.NewString(),
		Name: "John",
		Role: util_jwt.USER_ROLE,
	}

	t.Run("should successfully change password", func(t *testing.T) {
		jwtMock.EXPECT().VerifyAuthToken(authToken).Return(&claims, nil)

		serviceMock.EXPECT().ChangePassword(gomock.Any(), claims.ID, reqBody.OldPassword, reqBody.NewPassword).Return(nil)

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(reqBodyBytes))
		req.Header.Add(util_http.HEADER_AUTH, tokenBearer)
		w := httptest.NewRecorder()

		app.ServeHTTP(w, req)

		resBody := util_http.Response{}
		err := json.NewDecoder(w.Body).Decode(&resBody)

		assert.Equal(t, http.StatusOK, w.Result().StatusCode)
		assert.Nil(t, err)
		assert.NotEqual(t, util_http.Response{}, resBody)
	})

	t.Run("should return error when service error", func(t *testing.T) {
		jwtMock.EXPECT().VerifyAuthToken(authToken).Return(&claims, nil)

		serviceMock.EXPECT().ChangePassword(gomock.Any(), claims.ID, reqBody.OldPassword, reqBody.NewPassword).
			Return(util_error.NewBadRequest(nil, "Password dont match"))

		req := httptest.NewRequest(http.MethodPost, url, bytes.NewReader(reqBodyBytes))
		req.Header.Add(util_http.HEADER_AUTH, tokenBearer)
		w := httptest.NewRecorder()

		app.ServeHTTP(w, req)

		resBody := util_http.Response{}
		err := json.NewDecoder(w.Body).Decode(&resBody)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusBadRequest, w.Result().StatusCode)
	})

}
