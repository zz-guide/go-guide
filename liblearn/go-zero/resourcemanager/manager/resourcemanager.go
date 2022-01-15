package manager

import (
	"io"
	"sync"
)

// ResourceManager 通用的资源池管理对象，需实现Close()
type ResourceManager struct {
	resources map[string]io.Closer
	lock      sync.RWMutex
}

func NewResourceManager() *ResourceManager {
	return &ResourceManager{
		resources: make(map[string]io.Closer),
	}
}
