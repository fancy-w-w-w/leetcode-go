package pkg

import (
	"context"

	"github.com/go-redis/redis/v8"
)

func SetRedis() string {
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "sh-crs-ix7qfv5m.sql.tencentcdb.com:21139",
		Password: "!WC2971218", // no password set
		DB:       0,            // use default DB
	})
	// log.Debugf(ctx, "redis conn success")
	err := rdb.Set(ctx, "firstMessage", "niubi", 0).Err()
	if err != nil {
		// log.Errorf(ctx, err.Error())
		panic(err)
	}
	defer rdb.Close()
	return "success"
}

func GetValue() string {
	return "ss"
}
