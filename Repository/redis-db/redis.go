package redis_db

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/greenbahar/manage-order-process/config/env"
	"github.com/greenbahar/manage-order-process/order"
	"log"
	"time"
)

const (
	maxIdleConnections   = 10
	idleTimeout          = 20 * time.Second
	maxActiveConnections = 20
)

type RedisRepo struct {
	Pool *redis.Pool
}

func NewRepo() *RedisRepo {
	redisConf := env.GetRedisConfig()

	pool := &redis.Pool{
		MaxIdle:     maxIdleConnections,
		MaxActive:   maxActiveConnections,
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", redisConf.Host, redisConf.Port))
			if err != nil {
				panic(err)
			}
			return conn, err
		},
	}

	repo := &RedisRepo{pool}
	ping(repo)

	return repo
}

func ping(r *RedisRepo) {
	con := r.Pool.Get()
	defer con.Close()

	_, err := redis.String(con.Do("PING"))
	if err != nil {
		panic(err)
	}
}

func (r *RedisRepo) SetOrder(order *order.Order) error {
	// get conn and put back when exit from method
	conn := r.Pool.Get()
	defer conn.Close()

	key := order.OrderId
	val := order
	_, err := conn.Do("SET", key, val)
	if err != nil {
		log.Printf("ERROR: fail set key %s, val %s, error %s", key, val, err.Error())
		return err
	}

	return nil
}

//func (r *RedisRepo) GetOrder(key string) (string, error) {
//	// get conn and put back when exit from method
//	conn := r.Pool.Get()
//	defer conn.Close()
//
//	s, err := redis.String(conn.Do("GET", key))
//	if err != nil {
//		log.Printf("ERROR: fail get key %s, error %s", key, err.Error())
//		return "", err
//	}
//
//	return s, nil
//}
