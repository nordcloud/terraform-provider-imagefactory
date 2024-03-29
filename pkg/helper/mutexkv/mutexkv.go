// Copyright 2022 Nordcloud Oy or its affiliates. All Rights Reserved.

package mutexkv

import (
	"context"
	"fmt"
	"sync"

	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// MutexKV is a simple key/value store for arbitrary mutexes. It can be used to
// serialize changes across arbitrary collaborators that share knowledge of the
// keys they must serialize on.
type MutexKV struct {
	lock  sync.Mutex
	store map[string]*sync.Mutex
	key   string
}

// Locks the mutex.
func (m *MutexKV) Lock(ctx context.Context, key string) {
	tflog.Debug(ctx, fmt.Sprintf("Locking %s", key))
	m.get().Lock()
	tflog.Debug(ctx, fmt.Sprintf("Locked %s", key))
}

// Unlock the mutex.
func (m *MutexKV) Unlock(ctx context.Context, key string) {
	tflog.Debug(ctx, fmt.Sprintf("Unlocking %s", key))
	m.get().Unlock()
	tflog.Debug(ctx, fmt.Sprintf("Unlocked %s", key))
}

// Returns a mutex for the given key, no guarantee of its lock status.
func (m *MutexKV) get() *sync.Mutex {
	m.lock.Lock()
	defer m.lock.Unlock()
	mutex, ok := m.store[m.key]
	if !ok {
		mutex = &sync.Mutex{}
		m.store[m.key] = mutex
	}
	return mutex
}

// Returns a properly initialized MutexKV.
func NewMutexKV(key string) *MutexKV {
	return &MutexKV{
		store: make(map[string]*sync.Mutex),
		key:   key,
	}
}
