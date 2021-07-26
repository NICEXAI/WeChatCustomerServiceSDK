package cache

import "time"

type Cache interface {
	Set(k, v string, expires time.Duration) error
	Get(k string) (string, error)
}
