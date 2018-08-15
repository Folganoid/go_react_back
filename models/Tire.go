package models

type Tire struct {
	Id     uint64
	Name   string
	Userid uint64
}

func (Tire) TableName() string {
	return "tire"
}

func NewTire(
	id uint64,
	name string,
	userid uint64,
) *Tire {

	return &Tire{id, name, userid}

}
