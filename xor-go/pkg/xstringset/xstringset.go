package xstringset

type Set map[string]struct{}

func (s Set) Add(item string) {
	s[item] = struct{}{}
}

func (s Set) AddItems(items []string) {
	for _, item := range items {
		s[item] = struct{}{}
	}
}

func (s Set) Remove(item string) {
	delete(s, item)
}

func (s Set) Contains(item string) bool {
	_, found := s[item]
	return found
}

func (s Set) Size() int {
	return len(s)
}

func (s Set) Items() []string {
	items := make([]string, 0, len(s))
	for item := range s {
		items = append(items, item)
	}
	return items
}
