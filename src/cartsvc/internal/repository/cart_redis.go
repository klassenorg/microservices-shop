package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type CartRepo struct {
	db *redis.Client
}

func NewCartRepo(db *redis.Client) *CartRepo {
	return &CartRepo{db: db}
}

func (c *CartRepo) GetByID(ctx context.Context, id string) (map[string]string, error) {
	return c.db.HGetAll(ctx, id).Result()
}

func (c *CartRepo) AddByID(ctx context.Context, id string, product string, count int64) error {
	return c.db.HIncrBy(ctx, id, product, count).Err()
}

func (c *CartRepo) RemoveByID(ctx context.Context, id string, product string, count int64) error {
	res, err := c.db.HGet(ctx, id, product).Result()
	if err != nil {
		return err
	}

	if res == "" {
		return nil
	}

	current, err := strconv.ParseInt(res, 10, 64)
	if err != nil {
		return err
	}

	if current <= count {
		return c.db.HDel(ctx, id, product).Err()
	}

	return c.db.HIncrBy(ctx, id, product, count*-1).Err()
}

func (c *CartRepo) RemoveAllByID(ctx context.Context, id string) error {
	all, err := c.db.HGetAll(ctx, id).Result()
	if err != nil {
		return err
	}

	if len(all) == 0 {
		return nil
	}

	keys := make([]string, 0, len(all))
	for k := range all {
		keys = append(keys, k)
	}

	return c.db.HDel(ctx, id, keys...).Err()
}
