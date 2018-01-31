package myplugins

type Location struct {
	X float64

	Y float64
}

///////////////////2////////////

//location struct

type Userdata struct {
	Uname string

	Uid int64

	Pass string

	Addr string

	Geo Location
}

///////////////////////////3//////////////////

//user struct

type List []struct {
	Name    string
	Address string
}

/////////////////////////4//////////////////
//email group struct
