package models

type Bike struct {
	Id     uint64
	Name   string
	Userid uint64
}

func (Bike) TableName() string {
	return "ts"
}

func NewBike(
	id uint64,
	name string,
	userid uint64,
) *Bike {

	return &Bike{id, name, userid}

}
