package switcher

type Switcher struct {
	ISwitcherAdd
	ISwitcherUpdate
	ISwitcherGetTags
	ISwitcherUser
}

func (s *Switcher) Next(id int64) {
	if s.ISwitcherAdd.IsActive(id) {
		s.ISwitcherAdd.Next(id)
	} else if s.ISwitcherUpdate.IsActive(id) {
		s.ISwitcherUpdate.Next(id)
	} else if s.ISwitcherGetTags.IsActive(id) {
		s.ISwitcherGetTags.Next(id)
	} else if s.ISwitcherUser.IsActive(id) {
		s.ISwitcherUser.Next(id)
	}
}

func NewSwitcher(statusesAdd []string, statusesUpdate []string, statusesGetTags []string, statusesUser []string) *Switcher {
	return &Switcher{
		ISwitcherAdd:     NewSwitcherAdd(statusesAdd),
		ISwitcherUpdate:  NewSwitcherUpdate(statusesUpdate),
		ISwitcherGetTags: NewSwitcherGetTags(statusesGetTags),
		ISwitcherUser:    NewSwitcherUser(statusesUser),
	}
}
