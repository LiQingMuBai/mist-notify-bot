package switcher

type ISwitcher interface {
	Current(id int64) string
	Next(id int64)
	Previous(id int64)
	IsActive(id int64) bool
}

type ISwitcherAdd interface {
	ISwitcher
}

type ISwitcherUpdate interface {
	ISwitcher
}

type ISwitcherGetTags interface {
	ISwitcher
}

type ISwitcherUser interface {
	ISwitcher
}
