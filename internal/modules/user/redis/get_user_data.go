package user_redis

import (
	"context"

	user_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/model"
	util_redis "github.com/aziemp66/dot-indonesia-technical-test/util/redis"
	"github.com/redis/go-redis/v9"
)

// Retrieve user data (name & email)
func GetUserData(ctx context.Context, r util_redis.RedisManager, userID string) (user_model.User, error) {
	key := getUserKey(userID)

	var user user_model.User

	err := r.Get(ctx, key, &user)
	if err == redis.Nil {
		return user_model.User{}, nil
	}
	if err != nil {
		return user_model.User{}, err
	}

	return user, nil
}
