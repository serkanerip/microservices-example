package basket

import (
	"context"

	"github.com/vmihailenco/msgpack/v5"

	"github.com/go-redis/redis/v8"
)

type Repository interface {
	GetBasket(ctx context.Context, username string) (*ShoppingCart, error)
	UpdateBasket(ctx context.Context, basket ShoppingCart) (*ShoppingCart, error)
	DeleteBasket(ctx context.Context, username string) error
}

type RedisRepository struct {
	c          *redis.Client
	compressor ZlibCompressor
}

func NewRedisRepository(c *redis.Client) Repository {
	return &RedisRepository{c: c, compressor: ZlibCompressor{}}
}

func (r *RedisRepository) GetBasket(ctx context.Context, username string) (*ShoppingCart, error) {
	cmd := r.c.Get(ctx, username)
	if err := cmd.Err(); err != nil {
		return nil, err
	}
	b, err := cmd.Bytes()
	if err != nil {
		return nil, err
	}
	dBytes, err := r.compressor.Decompress(b)
	if err != nil {
		return nil, err
	}
	var sc ShoppingCart
	err = msgpack.Unmarshal(dBytes, &sc)
	if err != nil {
		return nil, err
	}
	return &sc, nil
}

func (r *RedisRepository) UpdateBasket(ctx context.Context, basket ShoppingCart) (*ShoppingCart, error) {
	bytes, err := msgpack.Marshal(basket)
	if err != nil {
		return nil, err
	}
	cBytes := r.compressor.Compress(bytes)

	r.c.Set(ctx, basket.Username, cBytes, 0)
	return r.GetBasket(ctx, basket.Username)
}

func (r *RedisRepository) DeleteBasket(ctx context.Context, username string) error {
	cmd := r.c.Del(ctx, username)
	return cmd.Err()
}
