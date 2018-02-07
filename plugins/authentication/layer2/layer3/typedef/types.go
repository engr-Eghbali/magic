package magic

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Location struct {
	X float64

	Y float64
}

///////////////////2////////////

//location struct

type Userdata struct {
	ID bson.ObjectId `json:"id" bson:"_id,omitempty"`

	Uname string `json:"uname"`

	Uid int64 `json:"uid"`

	Pass string `json:"pass"`

	Addr string `json:"addr"`

	Geo Location `json:"geo"`
}

///////////////////////////3//////////////////

//user struct

type List []struct {
	Name    string
	Address string
}

/////////////////////////4//////////////////
//email group struct
type FileInfo struct {
	Path    string
	Name    string    // base name of the file
	Size    int64     // length in bytes for regular files; system-dependent for others
	ModTime time.Time // modification time
	IsDir   bool      // abbreviation for Mode().IsDir()
}
