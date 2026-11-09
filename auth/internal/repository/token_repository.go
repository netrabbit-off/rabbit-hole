package repository

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"time"

	"github.com/netrabbit-off/rabbit-hole/pkg/models"
	"github.com/redis/go-redis/v9"
)

type TokenRepository struct {
	Redis *redis.Client
}

func connectRedis(addr string, password string, db int, protocol int) *redis.Client {
	r := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
		Protocol: protocol,
	})
	return r
}

func GetTokenRepository(addr string, password string, db int, protocol int) *TokenRepository {
	return &TokenRepository{Redis: connectRedis(addr, password, db, protocol)}
}

func (r *TokenRepository) CreateToken(ctx context.Context, token string, userID string, isRefresh bool) error {
	hash := sha256.Sum256([]byte(token))
	tokenHash := hex.EncodeToString(hash[:])
	var exp time.Duration
	if isRefresh {
		exp = 24 * time.Hour
	} else {
		exp = 10 * time.Minute
	}

	tokenData := &models.Token{
		Hash:      tokenHash,
		UserID:    userID,
		IsRefresh: isRefresh,
	}
	tokenBytes, _ := json.Marshal(tokenData)
	data := string(tokenBytes)
	if err := r.Redis.Set(ctx, "token:"+tokenHash, data, exp).Err(); err != nil {
		return err
	}

	return nil
}

func (r *TokenRepository) ParseToken(ctx context.Context, token string) (string, error) {
	res, err := r.Redis.Get(ctx, "token:"+token).Result()
	if err != nil {
		return "", err
	}

	tokenData := &models.Token{}
	if err := json.Unmarshal([]byte(res), &tokenData); err != nil {
		return "", err
	}

	if tokenData.IsRefresh {
		return "", errors.New("refresh token cannot be used to request")
	}

	return tokenData.UserID, nil
}
