package model

type Chaos struct {
	Type   string `gorm:""`
	Dose   int    `gorm:"notNull"`
	Thread int    `gorm:"notNull"`
}

// DTO input and func to populate it
type ChaosInput struct {
	Type   string `json:"type" binding:"required"`
	Dose   int    `json:"dose" binding:"required"`
	Thread int    `json:"thread" binding:"required"`
}

func (m *Chaos) PopulateFromDTOInput(input ChaosInput) {
	m.Type = input.Type
	m.Dose = input.Dose
	m.Thread = input.Thread
}

// DTO out and func to populate it
type ChaosOutput struct {
}

func (m *Chaos) PopulateDTOOutput() (output ChaosOutput) {

	return
}
