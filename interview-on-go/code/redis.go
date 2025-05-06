package main

import "sync"

type Storage interface {
	Get(key string) (string, error)
	Put(key string, value string) error
}

type Redis struct {
	mp map[string]string
	rw sync.RWMutex
}

func NewRedis() *Redis {
	return &Redis{
		mp: make(map[string]string),
		rw: sync.RWMutex{},
	}
}

func (r *Redis) Get(key string) (string, error) {
	r.rw.RLock()
	defer r.rw.RUnlock()
	return r.mp[key], nil
}

func (r *Redis) Put(key string, value string) error {
	r.rw.Lock()
	defer r.rw.Unlock()
	r.mp[key] = value
	return nil
}
