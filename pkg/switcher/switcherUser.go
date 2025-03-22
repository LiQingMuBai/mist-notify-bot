package switcher

type SwitcherUser struct {
	statuses      []string
	currentStatus map[int64]int
	users         map[int64]string
}

func NewSwitcherUser(statuses []string) *SwitcherUser {
	return &SwitcherUser{
		statuses:      statuses,
		users:         make(map[int64]string),
		currentStatus: make(map[int64]int),
	}
}

func (s *SwitcherUser) Next(id int64) {
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

func (s *SwitcherUser) Current(id int64) string {
	val, ok := s.currentStatus[id]
	if !ok || val == -1 {
		return ""
	}

	return s.statuses[s.currentStatus[id]]
}

func (s *SwitcherUser) Previous(id int64) {
	switch s.currentStatus[id] {
	case 0:
		s.currentStatus[id] = -1
		break
	default:
		s.currentStatus[id]--
	}
}

func (s *SwitcherUser) IsActive(id int64) bool {
	val, ok := s.currentStatus[id]
	return ok && val >= 0
}
