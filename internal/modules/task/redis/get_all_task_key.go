package task_redis

import "fmt"

func getAllTaskKey(userID string) string {
	return fmt.Sprintf("task:%s:all", userID)
}
