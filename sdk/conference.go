package sdk

/* Conference is the major subdivision of a league. For some reason MLB will
 * confusingly refer to this as a league.
 */
type Conference interface {
	ID() uint8
	add(*Division)
	DivisionByID(uint8) *Division
	DivisionByName(string) *Division
}

/*
 * basicConference is a building-block type. It represents a conference (league)
 * in a sport, helping perform the basic tasks.
 */
type basicConference struct {
	// This is an ID as defined by upstream API being used.
	id uint8
	// The most common, human readable name. i.e. "National League"
	Name string
	// The common, shortned version of the name. i.e. "NL"
	NameShort string
	// The common was for people to type a unique name for the conference. i.e. "nl"
	Slug      string
	divisions map[uint8]*Division
}

func (bc *basicConference) add(d *Division) {

	if bc.divisions == nil {
		bc.divisions = make(map[uint8]*Division)
	}

	bc.divisions[d.ID] = d
}

func (bc *basicConference) DivisionByID(id uint8) *Division {
	return bc.divisions[id]
}

func (bc *basicConference) DivisionByName(name string) *Division {

	for _, d := range bc.divisions {
		if d.Name == name {
			return d
		}
	}

	return nil
}

func (bc *basicConference) DivisionBySlug(slug string) *Division {

	for _, d := range bc.divisions {
		if d.Slug == slug {
			return d
		}
	}

	return nil
}

func (bc *basicConference) ID() uint8 {
	return bc.id
}
