package config

import (
	"fmt"
	"testing"
)

func TestRedis(t *testing.T) {

	RedisConfig = (&redisConfig{}).Load("../../config/app.ini").Init()
	fmt.Println(RedisConfig)
	if RedisConfig == nil {
		t.Fail()
	}

}
