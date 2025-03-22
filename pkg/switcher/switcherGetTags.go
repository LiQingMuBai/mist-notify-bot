package switcher

type SwitcherGetTags struct {
	statuses      []string
	currentStatus map[int64]int
	users         map[int64]string
}

func NewSwitcherGetTags(statuses []string) *SwitcherGetTags {
	return &SwitcherGetTags{
		currentStatus: make(map[int64]int),
		users:         make(map[int64]string),
		statuses:      statuses,
	}
}

func (s *SwitcherGetTags) Next(id int64) {
	val, ok := s.currentStatus[id]
	if val < len(s.statuses)-1 && ok {
		s.currentStatus[id]++
		return
	} else if !ok {
		s.currentStatus[id] = 0
		return
	}
	s.currentStatus[id] = -1
}

func (s *SwitcherGetTags) Current(id int64) string {
	val, ok := s.currentStatus[id]
	if !ok || val == -1 {
		return ""
	}

	return s.statuses[s.currentStatus[id]]
}

func (s *SwitcherGetTags) Previous(id int64) {
	switch s.currentStatus[id] {
	case 0:
		s.currentStatus[id] = -1
		break
	default:
		s.currentStatus[id]--
	}
}

func (s *SwitcherGetTags) IsActive(id int64) bool {
	val, ok := s.currentStatus[id]
	return ok && val >= 0
}
