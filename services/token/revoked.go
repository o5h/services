package token

import (
	"sync"
	"time"
)

type revokedTokens struct {
	mutex sync.RWMutex
	list  map[string]struct{}
}

var revoked = revokedTokens{
	list: make(map[string]struct{})}

func (l *revokedTokens) Revoke(s string, timeout time.Duration) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	time.AfterFunc(timeout, func() { l.remove(s) })
	l.list[s] = struct{}{}
}

func (l *revokedTokens) IsRevoked(s string) bool {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	_, exists := l.list[s]
	return !exists

}

func (l *revokedTokens) remove(s string) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	delete(l.list, s)
}
