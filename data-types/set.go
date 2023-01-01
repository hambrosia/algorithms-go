package datatypes

// inspired by https://gist.github.com/fatih/6206844, Python sets API, and a conversation in the Gophers Slack

import "sync"

// Set data type implementing commonly used read and write methods
type Set struct {
	m  map[string]struct{}
	mu sync.RWMutex
}

// Constructor: returns a new empty set struct
func Constructor() Set {
	return Set{
		m:  make(map[string]struct{}),
		mu: sync.RWMutex{},
	}
}

// Adds a member to the set
func (s *Set) Add(member string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m[member] = struct{}{}
}

// Removes a member to the set
func (s *Set) Remove(member string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.m, member)
}

// Removes all members of the set
func (s *Set) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.m = make(map[string]struct{})
}

// Returns true if the set contains the provided value, otherwise returns false
func (s *Set) Contains(member string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	_, ok := s.m[member]
	return ok
}

// Returns the length of the set, the number of members in the set
func (s *Set) Len() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return len(s.m)
}

// Returns true if the set is empty, otherwise returns false
func (s *Set) IsEmpty() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.Len() < 1
}

// Returns a slice of strings containing the members of the set
func (s *Set) List() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	ret := make([]string, s.Len())
	i := 0
	for member := range s.m {
		ret[i] = member
		i++
	}
	return ret
}

// Returns s1 - s2, the members of the set that are not in the second set
func (s1 *Set) Difference(s2 *Set) *Set {
	s1.mu.RLock()
	s2.mu.RLock()
	defer s1.mu.RUnlock()
	defer s2.mu.RUnlock()

	diff := Constructor()

	for member := range s1.m {
		if !s2.Contains(member) {
			diff.Add(member)
		}
	}

	return &diff
}

// Returns s1 Δ s2, the union of the members of s1 that are not in s2 and the members of s2 that are not in s1
func (s1 *Set) SymmetricDifference(s2 *Set) *Set {
	s1.mu.RLock()
	s2.mu.RLock()
	defer s1.mu.RUnlock()
	defer s2.mu.RUnlock()

	diff := Constructor()

	for member := range s1.m {
		if !s2.Contains(member) {
			diff.Add(member)
		}
	}
	for member := range s2.m {
		if !s1.Contains(member) {
			diff.Add(member)
		}
	}

	return &diff
}

// Returns s1 ∩ s2, the members of s1 that are also in s2
func (s1 *Set) Intersection(s2 *Set) *Set {
	s1.mu.RLock()
	s2.mu.RLock()
	defer s1.mu.RUnlock()
	defer s2.mu.RUnlock()

	intersect := Constructor()

	for member := range s1.m {
		if s2.Contains(member) {
			intersect.Add(member)
		}
	}
	return &intersect
}

// Returns s1 ∪ s2, the deduplicated unique members of both s1 and s2
func (s1 *Set) Union(s2 *Set) *Set {
	s1.mu.RLock()
	s2.mu.RLock()
	defer s1.mu.RUnlock()
	defer s2.mu.RUnlock()

	union := Constructor()

	for member := range s1.m {
		union.Add(member)
	}
	for member := range s2.m {
		union.Add(member)
	}
	return &union
}

// Returns true if s1 is a superset of s2
func (s1 *Set) IsSuperset(s2 *Set) bool {

	s1.mu.RLock()
	s2.mu.RLock()
	defer s1.mu.RUnlock()
	defer s2.mu.RUnlock()

	for member := range s2.m {
		if !s1.Contains(member) {
			return false
		}
	}
	return true
}

// Returns true if s1 is a subset of s2
func (s1 *Set) IsSubset(s2 *Set) bool {
	s1.mu.RLock()
	s2.mu.RLock()
	defer s1.mu.RUnlock()
	defer s2.mu.RUnlock()

	for member := range s1.m {
		if !s2.Contains(member) {
			return false
		}
	}
	return true
}

// Updates s1 to contain only members that are not also in s2
func (s1 *Set) DifferenceUpdate(s2 *Set) {
	s1.mu.Lock()
	defer s1.mu.Unlock()
	s2.mu.RLock()
	defer s2.mu.RUnlock()

	s1.m = s1.Difference(s2).m
}

// Updates s1 to contain only members which are in s1 and s2
func (s1 *Set) IntersectionUpdate(s2 *Set) {
	s1.mu.Lock()
	s2.mu.RLock()
	defer s1.mu.Unlock()
	defer s2.mu.RUnlock()

	s1.m = s1.Intersection(s2).m
}

// Updates s1 to contain only members which are in s1 only or s2 only, not in both s1 and s2
func (s1 *Set) SymmetricDifferenceUpdate(s2 *Set) {
	s1.mu.Lock()
	s2.mu.RLock()
	defer s1.mu.Unlock()
	defer s2.mu.RUnlock()

	s1.m = s1.SymmetricDifference(s2).m
}
