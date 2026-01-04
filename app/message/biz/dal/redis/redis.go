package redis

import (
	"common/pkg/errno"
	"context"
	"gen/kitex_gen/base"
	"github.com/redis/go-redis/v9"
	"time"

	"github.com/bytedance/sonic"
)

type Dict struct {
	RedisClient *redis.Client
}

func NewDict(client *redis.Client) *Dict {
	return &Dict{RedisClient: client}
}

func (r *Dict) GetDict(c context.Context, id string) (*base.Dict, error) {
	p, err := r.RedisClient.Get(c, id).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, errno.NotFound
		}
		return nil, err
	}
	var pv base.Dict
	if err = sonic.UnmarshalString(p, &pv); err != nil {
		return nil, err
	}
	return &pv, nil
}

func (r *Dict) SetDict(c context.Context, id string, p *base.Dict) error {
	_, err := r.RedisClient.Get(c, id).Result()
	if err != redis.Nil {
		if err == nil {
			return errno.AlreadyExist
		} else {
			return err
		}
	}
	pv, err := sonic.Marshal(p)
	if err != nil {
		return err
	}
	if err = r.RedisClient.Set(c, id, pv, 168*time.Hour).Err(); err != nil {
		return err
	}
	return nil
}

func (r *Dict) RemoveProfile(c context.Context, id string) error {
	if err := r.RedisClient.Del(c, id).Err(); err != nil {
		return err
	}
	return nil
}
