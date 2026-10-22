package main

import (
	"fmt"
	"sync"
)

type Dictionary struct {
	elements map[string]string
	lock     sync.RWMutex
}

func (dict *Dictionary) Put(key, value string) {
	dict.lock.Lock()
	defer dict.lock.Unlock()

	if dict.elements == nil {
		dict.elements = make(map[string]string)
	}

	dict.elements[key] = value
}

func (dict *Dictionary) Remove(key string) bool {
	dict.lock.Lock()
	defer dict.lock.Unlock()

	_, exists := dict.elements[key]
	if exists {
		delete(dict.elements, key)
	}

	return exists
}

func (dict *Dictionary) Contains(key string) bool {
	dict.lock.Lock()
	defer dict.lock.Unlock()

	_, exists := dict.elements[key]
	return exists

}

func (dict *Dictionary) Find(key string) string {
	dict.lock.RLock()
	defer dict.lock.RUnlock()

	return dict.elements[key]
}

func (dict *Dictionary) Reset() {
	dict.lock.Lock()
	defer dict.lock.Unlock()

	dict.elements = make(map[string]string)
}

func (dict *Dictionary) NumberOfElements() int {
	dict.lock.RLock()
	defer dict.lock.RUnlock()

	return len(dict.elements)

}

func (dict *Dictionary) GetKeys() []string {
	dict.lock.RLock()
	defer dict.lock.RUnlock()

	dictKeys := []string{}
	for key := range dict.elements {
		dictKeys = append(dictKeys, key)
	}
	return dictKeys
}

func (dict *Dictionary) GetValues() []string {
	dict.lock.RLock()
	defer dict.lock.RUnlock()

	dictValues := []string{}

	for key := range dict.elements {
		dictValues = append(dictValues, dict.elements[key])
	}
	return dictValues
}

func main() {
	dict := &Dictionary{}
	dict.Put("1", "1")
	dict.Put("2", "2")
	fmt.Println("Put : ", dict.elements)
}
