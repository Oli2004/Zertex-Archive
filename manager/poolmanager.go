package manager

import (
	"errors"
	"sync"
)

var (
	PoolStarted = errors.New("pool already started, please stop it")
	PoolStopped = errors.New("pool has not been started")
)

type PoolManager struct {
	workers chan interface{}
	control chan interface{}
	wait sync.WaitGroup
	size int
	workFunc func(*sync.WaitGroup, *chan interface{})
}

func NewPoolManager(size int, workFunc func(*sync.WaitGroup, *chan interface{})) *PoolManager {
	return &PoolManager{
		size: size,
		workFunc: workFunc,
	}
}

func (pm *PoolManager) Start() error {
	if len(pm.workers) > 0 {
		return PoolStarted
	}
	pm.workers = make(chan interface{},pm.size)
	pm.control = make(chan interface{})

	pm.wait.Add(pm.size)
	for i:=0; i<pm.size; i++ {
		pm.workers <- 0
		go pm.workFunc(&pm.wait, &pm.control)
	}
	return nil
}

func (pm *PoolManager) Stop(waitOnExit bool) error {
	if len(pm.workers) == 0 {
		return PoolStopped
	}
	for i:=0; i < pm.size; i++ {
		pm.control <- 0
	}
	if waitOnExit {
		pm.wait.Wait()
	}
	return nil
}

func (pm *PoolManager) Resize(newSize int) error {
	if pm.size < newSize { // bigger new
		pm.workers = make(chan interface{}, newSize)
		pm.size = newSize
		for i:=0; i < newSize; i++ {
			pm.workers <- 0
			go pm.workFunc(&pm.wait, &pm.control)
		}
	} else {
		diff := pm.size - newSize
		for i:=0; i < diff; i++ {
			pm.control <- 0
		}
	}
	return nil
}