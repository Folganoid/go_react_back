package models

type Marker struct {
	Id    	uint64
	Userid	uint64
	X		float64
	Y		float64
	Name	string
	Subname	string
	Link  	string
	Color	string
}

func NewMarker(
	id    	uint64,
	userid	uint64,
	x		float64,
	y		float64,
	name	string,
	subname	string,
	link  	string,
	color	string) *Marker {

	return &Marker{id, userid, x, y, name, subname, link, color}
}