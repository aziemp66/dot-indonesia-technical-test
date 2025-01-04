package user_redis

import (
	"context"
	"time"

	user_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/model"
	util_redis "github.com/aziemp66/dot-indonesia-technical-test/util/redis"
)

// Store user data (name & email)
func SetUserData(ctx context.Context, r util_redis.RedisManager, userID string, user user_model.User, expiration time.Duration) error {
	key := getUserKey(userID)

	return r.Set(ctx, key, user, expiration)
}
