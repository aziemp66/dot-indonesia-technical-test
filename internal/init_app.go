package init_app

import (
	task_http "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/http"
	task_http_router "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/http/router"
	task_repository_postgres "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/repository/postgres"
	task_service "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/task/service"
	user_http "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/http"
	user_http_router "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/http/router"
	user_repository_postgres "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/repository/postgres"
	user_service "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/service"
	pkg_config "github.com/aziemp66/dot-indonesia-technical-test/internal/pkg/config"
	util_jwt "github.com/aziemp66/dot-indonesia-technical-test/util/jwt"
	util_password "github.com/aziemp66/dot-indonesia-technical-test/util/password"
	util_redis "github.com/aziemp66/dot-indonesia-technical-test/util/redis"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeApp(router *gin.Engine, cfg *pkg_config.Config, db *gorm.DB) {
	jwtManager := util_jwt.NewjwtManager(cfg.AppJwtSecret)
	passwordManager := util_password.NewPasswordManager(30)
	redisManager := util_redis.NewRedisManager(cfg.RedisHost, cfg.RedisPassword, cfg.RedisDB)

	userRepository := user_repository_postgres.NewUserRepositoryPostgres(db)
	userService := user_service.NewUserService(userRepository, jwtManager, passwordManager, redisManager)
	userHandler := user_http.NewUserHttpHandler(userService, jwtManager)
	user_http_router.BindUserHttpRouter(router.Group("/user"), userHandler, jwtManager)

	taskRepository := task_repository_postgres.NewTaskRepositoryPostgres(db)
	taskService := task_service.NewTaskService(taskRepository, redisManager)
	taskHandler := task_http.NewTaskHttpHandler(taskService, jwtManager)
	task_http_router.BindTaskHttpRouter(router.Group("/task"), taskHandler, jwtManager)
}
