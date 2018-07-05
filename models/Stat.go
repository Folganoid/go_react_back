package models

type Stat struct {
	Id      uint64
	Dist    float64
	Time    uint64
	Bike    string
	Avgspd  float64
	Maxspd  float64
	Avgpls  uint64
	Maxpls  uint64
	Tires   string
	Date    uint64
	Surfasf uint8
	Surftvp uint8
	Surfgrn uint8
	Srfbzd  uint8
	Prim    string
	Teh     string
	Temp    string
	Wind    string
	Userid  uint64
}

func NewStat(
	id uint64,
	dist float64,
	time uint64,
	bike string,
	avgspd float64,
	maxspd float64,
	avgpls uint64,
	maxpls uint64,
	tires string,
	date uint64,
	surfasf uint8,
	surftvp uint8,
	surfgrn uint8,
	srfbzd uint8,
	prim string,
	teh string,
	temp string,
	wind string,
	userid uint64,
) *Stat {

	return &Stat{
		id,
		dist,
		time,
		bike,
		avgspd,
		maxspd,
		avgpls,
		maxpls,
		tires,
		date,
		surfasf,
		surftvp,
		surfgrn,
		srfbzd,
		prim,
		teh,
		temp,
		wind,
		userid,
	}
}
