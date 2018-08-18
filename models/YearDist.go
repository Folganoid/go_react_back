package models

type YearDist struct {
	Id     uint64
	Userid uint64
	Year   uint64
	Bike   string
	Dist   float64
}

func (YearDist) TableName() string {
	return "yeardata"
}

func NewYearDist(
	id uint64,
	userid uint64,
	year uint64,
	bike string,
	dist float64,
) *YearDist {

	return &YearDist{id, userid, year, bike, dist}

}
