package cache

import "time"

type Cache interface {
	Set(k, v string, expires time.Duration) error
	Get(k string) (string, error)
	Scan(cursor uint64, match string, count int64) (keys []string, newCursor uint64, err error)
}
