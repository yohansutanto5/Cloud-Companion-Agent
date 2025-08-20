package model

type Connect struct {
	ID int `gorm:"primaryKey;autoIncrement"`
	Host string `gorm:""`
	Port int `gorm:"notNull"`

}

// DTO input and func to populate it
type ConnectInput struct {

}
func (m *Connect) PopulateFromDTOInput(input ConnectInput) {

}

// DTO out and func to populate it
type ConnectOutput struct {

}

func (m *Connect) PopulateDTOOutput() (output ConnectOutput){

  return
}