package main

type StringIntMap struct {
	stringIntMap map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{stringIntMap: make(map[string]int)}
}

func (m *StringIntMap) Add(key string, value int) {
	m.stringIntMap[key] = value

}

func (m *StringIntMap) Remove(key string) {
	delete(m.stringIntMap, key)

}

func (m *StringIntMap) Copy() *StringIntMap {
	CopyStringIntMap := NewStringIntMap()
	for key, value := range m.stringIntMap {
		CopyStringIntMap.Add(key, value)
	}
	return CopyStringIntMap
}

func (m *StringIntMap) Exist(key string) (int, bool) {

	value, exists := m.stringIntMap[key]
	return value, exists
}

func (m *StringIntMap) Get(key string) (int, bool) {

	value, exists := m.stringIntMap[key]
	return value, exists
}

func main() {

}
