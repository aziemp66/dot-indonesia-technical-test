package user_redis

import (
	"context"
	"encoding/json"

	user_model "github.com/aziemp66/dot-indonesia-technical-test/internal/modules/user/model"
	util_redis "github.com/aziemp66/dot-indonesia-technical-test/util/redis"
)

// Retrieve user data (name & email)
func GetUserData(ctx context.Context, r util_redis.RedisManager, userID string) (user_model.GetUserResponse, error) {
	key := getUserKey(userID)

	var user user_model.GetUserResponse

	data, err := r.Client().Get(ctx, key).Result()
	if err != nil {
		return user_model.GetUserResponse{}, err
	}

	err = json.Unmarshal([]byte(data), &user)
	if err != nil {
		return user_model.GetUserResponse{}, err
	}

	return user, nil
}
