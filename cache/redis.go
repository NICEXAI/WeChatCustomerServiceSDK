package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"sync"
	"time"
)

const GlobalEvent = "global_event"

type Redis struct {
	//订阅服务器实例
	Point *redis.Client
	//订阅列表
	PbFns sync.Map
	//读写锁
	lock sync.Mutex
}

type RedisOptions struct {
	Addr     string
	Password string
	DB int
}

func NewRedis(options RedisOptions) *Redis {
	ctx := context.TODO()
	instance := Redis{}
	//实例化连接池，解决每次重新连接效率低的问题
	instance.Point = redis.NewClient(&redis.Options{
		Addr:     options.Addr,
		Password: options.Password,
		DB:       options.DB,
	})

	instance.PbFns = sync.Map{}
	go func() {
		pubSub := instance.Point.Subscribe(ctx, "__keyevent@0__:expired")
		for {
			msg, err := pubSub.ReceiveMessage(ctx)
			if err != nil {
				fmt.Println(err)
				return
			}
			if msg.Channel == "__keyevent@0__:expired" {
				pbFnList, _ := instance.PbFns.Load(msg.Payload)
				if pbFnList != nil {
					cbList, ok := pbFnList.([]func(message string))
					if ok {
						for _, cb := range cbList{
							cb(msg.Payload)
						}
					}
				}
				//处理全局订阅回调
				globalFnList, _ := instance.PbFns.Load(GlobalEvent)
				if globalFnList != nil {
					cbList, ok := globalFnList.([]func(message string))
					if ok {
						for _, cb := range cbList{
							cb(msg.Payload)
						}
					}
				}
			}
		}
	}()

	return &instance
}

// GetOriginPoint 获取原始redis实例
func (r *Redis) GetOriginPoint() *redis.Client {
	return r.Point
}

// Subscribe 订阅指定键过期时间，需要redis开启键空间消息通知：config set notify-keyspace-events Ex
func (r *Redis) Subscribe(k string, pb func(message string)) {
	var cbList []func(message string)
	pbFnList, ok := r.PbFns.Load(k)
	if ok {
		cbList, ok = pbFnList.([]func(message string))
		if ok {
			r.lock.Lock()
			cbList = append(cbList, pb)
			r.lock.Unlock()
		}
	} else {
		cbList = []func(message string){pb}
	}

	r.PbFns.Store(k, cbList)
}

// SubscribeAllEvents 订阅所有键过期事件
func (r* Redis) SubscribeAllEvents(pb func(message string))  {
	var cbList []func(message string)
	pbFnList, ok := r.PbFns.Load(GlobalEvent)
	if ok {
		cbList, ok = pbFnList.([]func(message string))
		if ok {
			r.lock.Lock()
			cbList = append(cbList, pb)
			r.lock.Unlock()
		}
	} else {
		cbList = []func(message string){pb}
	}

	r.PbFns.Store(GlobalEvent, cbList)
}

func (r *Redis) Set(k, v string, expires time.Duration) error {
	return r.Point.Set(context.TODO(), k, v, expires * time.Second).Err()
}

func (r *Redis) Get(k string) (string, error) {
	con, err := r.Point.Get(context.TODO(), k).Result()
	if err != nil && err != redis.Nil {
		return "", err
	}
	return con, nil
}

func (r *Redis) Scan(cursor uint64, match string, count int64) (keys []string, newCursor uint64, err error) {
	return r.Point.Scan(context.TODO(), cursor, match, count).Result()
}