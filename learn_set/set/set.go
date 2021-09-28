package set

var void = struct{}{}

type Set struct {
	m map[interface{}]struct{}
}

/**
	实例化
 */
func New(items ...interface{}) *Set {
	// 实例化set
	s := &Set{
		m: make(map[interface{}]struct{}),
	}
	// 实例化成员
	s.Add(items...)
	return s
}

/**
	添加
 */
func (s *Set) Add(items ...interface{}) error {
	for _, item := range items {
		s.m[item] = void
	}
	return nil
}

/**
	包含
 */
func (s *Set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}

/**
	删除
 */
func (s *Set) Remove(item interface{}) {
	delete(s.m, item)
}

/**
	长度
 */
func (s *Set) Size() int {
	return len(s.m)
}

/**
	清除
 */
func (s *Set) Clear() {
	s.m = make(map[interface{}]struct{})
}

/**
	相等
 */
func (s *Set) Equal(other *Set) bool {
	// 如果两个set size不相等，直接返回false
	if s.Size() != other.Size() {
		return false
	}
	// 遍历，只要有一个不存在就返回false
	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

/**
	子集
 */
func (s *Set) IsSubset(other *Set) bool {
	if s.Size() > other.Size() {
		return false
	}
	// 遍历
	for key := range s.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}
