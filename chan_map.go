package chan_map

// channel map , simply and efficient
// A "thread" safe map of type Anything:Anything.
// Can be upgraded with reduced lock granularity
type chanMap struct {
	cMap map[interface{}]interface{}
	c    chan struct{}
}

// NewChanMap new a safe map
func NewChanMap() *chanMap {
	return &chanMap{
		cMap: make(map[interface{}]interface{}),
		c:    make(chan struct{}, 1),
	}
}

// Get get a value of key
func (m *chanMap) Get(key interface{}) interface{} {
	m.lock()
	defer m.unlock()
	return m.cMap[key]
}

// Set set a value of key
func (m *chanMap) Set(key interface{}, value interface{}) {
	m.lock()
	defer m.unlock()
	m.cMap[key] = value
}

// Delete delete the key
func (m *chanMap) Delete(key interface{}) {
	m.lock()
	defer m.unlock()
	delete(m.cMap, key)
}

// Size get the map size
func (m *chanMap) Size() int {
	m.lock()
	defer m.unlock()
	return len(m.cMap)
}

// Range calls f sequentially for each key and value present in the map.
// If f returns false, range stops the iteration.

func (m *chanMap) Range(fn func(k, v interface{}) bool) {
	m.lock()
	defer m.unlock()
	for k, v := range m.cMap {
		if !fn(k, v) {
			break
		}
	}
}

func (m *chanMap) lock() {
	m.c <- struct{}{}
}

func (m *chanMap) unlock() {
	<-m.c
}
