package models

type Year struct {
	Id     uint64
	Userid uint64
	Year   uint64
	Bike   string
	Dist   float64
}

func (Year) TableName() string {
	return "yeardata"
}


func NewYear(
	id uint64,
	userid uint64,
	year uint64,
	bike string,
	dist float64,
) *Year {

	return &Year{
		id,
		userid,
		year,
		bike,
		dist,
	}
}
