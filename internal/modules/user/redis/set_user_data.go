package user_redis

import (
	"context"
	"encoding/json"
	"time"

	user_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/model"
	util_redis "github.com/aziemp66/dot-indonesia-technical-test/util/redis"
)

// Store user data (name & email)
func SetUserData(ctx context.Context, r util_redis.RedisManager, userID string, user user_model.GetUserResponse, expiration time.Duration) error {
	key := getUserKey(userID)

	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return r.Client().Set(ctx, key, data, expiration).Err()
}
