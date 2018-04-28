package session

import "sync"

type Manager struct {
	cookieName string
	lock sync.Mutex
	lifetime int
}