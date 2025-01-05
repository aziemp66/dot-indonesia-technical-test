package task_redis

import "fmt"

func getTaskKey(userID string) string {
	return fmt.Sprintf("task:%s", userID)
}
