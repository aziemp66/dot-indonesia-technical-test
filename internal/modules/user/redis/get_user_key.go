package user_redis

import "fmt"

func getUserKey(userID string) string {
	return fmt.Sprintf("user:%s:data", userID)
}
