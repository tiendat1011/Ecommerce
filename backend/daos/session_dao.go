package daos

import (
	"context"
	"ecommerce-project/databases"
	"time"
)

type SessionDAO interface {
	SaveSession(token, userID string, ttl time.Duration) error
	GetSession(token string) (string, error)
	DeleteSession(token string) error
}

type sessionDAO struct{}

func (d *sessionDAO) SaveSession(token, userID string, ttl time.Duration) error {
	return databases.Rdb.Set(context.TODO(), "session:"+token, userID, ttl).Err()
}

func (d *sessionDAO) GetSession(token string) (string, error) {
	return databases.Rdb.Get(context.TODO(), "session:"+token).Result()
}

func (d *sessionDAO) DeleteSession(token string) error {
	return databases.Rdb.Del(context.TODO(), "session:"+token).Err()
}
