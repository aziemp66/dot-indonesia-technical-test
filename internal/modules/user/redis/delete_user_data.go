package user_redis

import (
	"context"

	util_redis "github.com/aziemp66/dot-indonesia-technical-test/util/redis"
)

// Delete user data
func DeleteUserData(ctx context.Context, r util_redis.RedisManager, userID string) error {
	key := getUserKey(userID)
	return r.Client().Del(ctx, key).Err()
}
