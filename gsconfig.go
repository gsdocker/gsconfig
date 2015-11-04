package gsconfig

import (
	"errors"
	"sync"
	"time"
)

//gsconfig package's public errors
var (
	ErrConvert = errors.New("config object convert error")
)

// Watcher config value changed watcher
type Watcher func(key, value interface{})

// Provider the gsconfig provider
type Provider struct {
	sync.RWMutex                        // mixin read/write locker
	properties   map[string]interface{} // config kvs
	watchers     map[string][]Watcher   // register watchers
	events       chan func()            // config change events
}

// NewProvider create new config provider with eventQ size
func NewProvider(eventQ int) *Provider {
	provider := &Provider{
		properties: make(map[string]interface{}),
		watchers:   make(map[string][]Watcher),
		events:     make(chan func(), eventQ),
	}

	go func() {
		for f := range provider.events {
			f()
		}
	}()

	return provider
}

// Save update config kvs
func (provider *Provider) Save(kvs map[string]interface{}) {
	provider.Lock()
	defer provider.Unlock()

	for k, v := range kvs {
		provider.properties[k] = v

		if watchers, ok := provider.watchers[k]; ok && watchers != nil {
			provider.addEvent(func() {
				for _, watcher := range watchers {
					watcher(k, v)
				}
			})
		}
	}
}

// Update update config
func (provider *Provider) Update(key string, val interface{}) {
	provider.Lock()
	defer provider.Unlock()

	provider.properties[key] = val

	if watchers, ok := provider.watchers[key]; ok && watchers != nil {
		provider.addEvent(func() {
			for _, watcher := range watchers {
				watcher(key, val)
			}
		})
	}
}

func (provider *Provider) addEvent(f func()) {
	select {
	case provider.events <- f:
	default:
	}
}

// Watch add config key's watcher
func (provider *Provider) Watch(key string, watcher Watcher) {
	provider.Lock()
	defer provider.Unlock()

	provider.watchers[key] = append(provider.watchers[key], watcher)

	if val, ok := provider.properties[key]; ok {
		provider.addEvent(func() {
			watcher(key, val)
		})
	}
}

// Get get config value
func (provider *Provider) Get(key string) (val interface{}, ok bool) {
	provider.RLock()
	defer provider.RUnlock()

	val, ok = provider.properties[key]

	return
}

var globalProvider = NewProvider(1024)

//Save save as global config
func Save(kvs map[string]interface{}) {
	globalProvider.Save(kvs)
}

// Update set new config item
func Update(k string, v interface{}) {
	globalProvider.Update(k, v)
}

//Get .
func Get(key string) (val interface{}, ok bool) {

	return globalProvider.Get(key)
}

//String .
func String(key string, defaultval string) string {
	if val, ok := globalProvider.Get(key); ok {
		if vall, ok := val.(string); ok {
			return vall
		}
	}

	return defaultval
}

//Int64 .
func Int64(key string, defaultval int64) int64 {
	if val, ok := globalProvider.Get(key); ok {
		if vall, ok := val.(int64); ok {
			return vall
		}
	}

	return defaultval
}

//Int32 .
func Int32(key string, defaultval int32) int32 {
	return int32(Int64(key, int64(defaultval)))
}

//Int16 .
func Int16(key string, defaultval int16) int16 {
	return int16(Int64(key, int64(defaultval)))
}

//Int .
func Int(key string, defaultval int) int {
	return int(Int64(key, int64(defaultval)))
}

//Uint64 .
func Uint64(key string, defaultval uint64) uint64 {
	if val, ok := globalProvider.Get(key); ok {
		if vall, ok := val.(uint64); ok {
			return vall
		}
	}

	return defaultval
}

//Uint32 .
func Uint32(key string, defaultval uint32) uint32 {
	return uint32(Uint64(key, uint64(defaultval)))
}

//Uint16 .
func Uint16(key string, defaultval uint16) uint16 {
	return uint16(Uint64(key, uint64(defaultval)))
}

//Uint .
func Uint(key string, defaultval uint) uint {
	return uint(Uint64(key, uint64(defaultval)))
}

//Bool .
func Bool(key string, defaultval bool) bool {
	if val, ok := globalProvider.Get(key); ok {
		if vall, ok := val.(bool); ok {
			return vall
		}
	}

	return defaultval
}

//Float64 .
func Float64(key string, defaultval float64) float64 {
	if val, ok := globalProvider.Get(key); ok {
		if vall, ok := val.(float64); ok {
			return vall
		}
	}

	return defaultval
}

//Float32 .
func Float32(key string, defaultval float32) float32 {
	return float32(Float64(key, float64(defaultval)))
}

//Seconds .
func Seconds(key string, defaultval int64) time.Duration {
	return time.Duration(Int64(key, defaultval)) * time.Second
}

//Milliseconds .
func Milliseconds(key string, defaultval int64) time.Duration {
	return time.Duration(Int64(key, defaultval)) * time.Millisecond
}
