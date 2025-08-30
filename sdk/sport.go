package sdk

// In the context of this package, a sport is simply a group of conferences.
type sport struct {
	conferences map[uint8]Conference
}

func (s *sport) add(c Conference) {

	if s.conferences == nil {
		s.conferences = make(map[uint8]Conference)
	}

	s.conferences[c.ID()] = c
}

func (s *sport) LeagueByID(id uint8) Conference {
	return s.conferences[id]
}
