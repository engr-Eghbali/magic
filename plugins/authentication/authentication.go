package magic

import (
	"fmt"
	f "fmt"
	"log"

	"net/http"

	"regexp"

	"strconv"

	"strings"

	magic_gcm "./layer2/layer3"
	magic_struct "./layer2/layer3/typedef"
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func checkErr(err error) {

	if err != nil {

		log.Fatal(err)
		//panic(err)

	}

}

/////////////////////////////////////////////////////////////7///////////////////////

//make panic on  error
//"id=8736&name=day&pass=me2&add=24#56#"
func Initialsubmit(r *http.Request, in *magic_struct.Userdata) {

	var err error = nil
	r.ParseForm()

	var tid string = r.Form["id"][0]
	in.Uname = r.Form["name"][0]
	in.Uid, err = strconv.ParseInt(tid, 10, 0)
	checkErr(err)
	in.Pass = r.Form["pass"][0]
	in.Addr = r.Form["add"][0]

	in.Geo.X, err = strconv.ParseFloat(strings.Split(in.Addr, "#")[1], 64)
	checkErr(err)
	in.Geo.Y, err = strconv.ParseFloat(strings.Split(in.Addr, "#")[2], 64)
	//checkErr(err)
}

/////////////////////////////////////4////////////////////

//initial variables for submit form

func Initiallogin(r *http.Request, in *magic_struct.Userdata) {

	r.ParseForm()

	var tid string = r.Form["id"][0]
	var err error
	in.Geo.X, err = strconv.ParseFloat(strings.Split(in.Addr, "#")[1], 64)
	checkErr(err)
	in.Geo.Y, err = strconv.ParseFloat(strings.Split(in.Addr, "#")[2], 64)
	//checkErr(err)

	in.Addr, in.Uname = r.Form["add"][0], r.Form["name"][0]

	in.Uid, _ = strconv.ParseInt(tid, 10, 0)

	in.Pass = r.Form["pass"][0]

}

func Login(w http.ResponseWriter, r *http.Request, input magic_struct.Userdata) (flg bool) {

	var pass []byte

	//var name, addr, x, y string
	//mysql method
	//	db, err := sql.Open("mysql", "root:toor@/test?charset=utf8")
	//	defer db.Close()
	//	checkErr(err)
	//	stmt, err2 := db.Prepare("SELECT `pass`,`uname`,`address`,`X`,`Y` FROM `userinfo` WHERE `ID`=?")
	//	checkErr(err2)
	//	res, err3 := stmt.Query(input.Uid)
	//	if err3 != nil {
	//
	//		checkErr(err3)
	//		return false
	//	} else {
	//		defer res.Close()
	//		for res.Next() {
	//			xerr := res.Scan(&pass, &name, &addr, &x, &y)
	//			f.Print(name)
	//			if xerr != nil {
	//				log.Fatal(xerr)
	//			}
	//
	//		}

	// MongoDB

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	////////
	c := session.DB("magic").C("userInfo")

	// Index
	//	index := mgo.Index{
	//		Key:        []string{"Uname", "Uid"},
	//		Unique:     true,
	//		DropDups:   true,
	//		Background: true,
	//		Sparse:     true,
	//	}

	//	err = c.EnsureIndex(index)
	//	if err != nil {
	//		panic(err)
	//	}

	// Query One
	result := magic_struct.Userdata{}
	log.Print(input.Uname)
	log.Print(input.Pass)
	err = c.Find(bson.M{"uname": input.Uname}).Select(bson.M{"pass": magic_gcm.Cipher(input.Pass)}).One(&result)
	if err != nil {
		panic(err)
	}
	log.Print("**\n")
	f.Println("pass", result)
	pass = []byte(result.Pass)

	//////

	if input.Pass == magic_gcm.Decipher(pass) {

		// create session here
		fmt.Fprint(w, result.ID)

		return true

	} else {

		fmt.Fprint(w, "login failed")
		return false
	}
}

func stripchars(str, chr string) string {

	return strings.Map(func(r rune) rune {

		if strings.IndexRune(chr, r) < 0 {

			return r

		}

		return -1

	}, str)

}

//////////////////////////////////////////////5///////////////////////

//scape_special_char

func Validation(in *magic_struct.Userdata, w http.ResponseWriter, r *http.Request) (flg bool) {

	var empty_flg bool = false

	if len(in.Pass) == 0 || in.Uid == 0 {

		empty_flg = true

		f.Fprintf(w, "fill the forms") //empty form error

		return empty_flg

	} else {

		var valid_flg bool = true

		f1, _ := regexp.MatchString("^[a-zA-Z0-9].$.@", string(in.Pass))

		f2, _ := regexp.MatchString("[0-9]", strconv.FormatInt(in.Uid, 10))

		in.Uname = stripchars(in.Uname, "<>/'=`")

		in.Addr = stripchars(in.Addr, "<>/'=`")

		valid_flg = f1 && f2

		return !valid_flg

	}

}

////////////////////////////////////////////////6/////////////////////////////

//check validation of inputs

func Submit(w http.ResponseWriter, r *http.Request, input magic_struct.Userdata) (flg bool) {

	//	db, err := sql.Open("mysql", "root:toor@/test?charset=utf8")
	//	defer db.Close()
	//	checkErr(err)
	//
	//	stmt, err2 := db.Prepare("INSERT INTO `userinfo` (`ID`,`uname`,`pass`,`address`,`X`,`Y`) VALUES (?,?,?,?,?,?)")
	//	checkErr(err2)
	//
	//	_, err3 := stmt.Exec(strconv.FormatInt(input.Uid, 10), input.Uname, magic_gcm.Cipher(input.Pass), input.Addr, strconv.FormatFloat(input.Geo.X, 'f', 6, 64), strconv.FormatFloat(input.Geo.Y, 'f', 6, 64))
	//
	//	if err3 != nil {
	//
	//		if strings.Contains(err3.Error(), "Duplicate") {
	//
	//			f.Fprintf(w, "user duplicate")
	//
	//			return false
	//
	//		}
	//
	//		checkErr(err3)
	//
	//		return false
	//
	//	} else {
	//
	//		return true
	//
	//	}
	//

	////// MONGO DB
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
		return false
	}

	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	////////
	c := session.DB("magic").C("userInfo")

	// Insert Datas
	err = c.Insert(&magic_struct.Userdata{Uname: input.Uname, Uid: input.Uid, Pass: string(magic_gcm.Cipher(input.Pass)), Addr: input.Addr, Geo: input.Geo})

	if err != nil {
		fmt.Fprintf(w, "invalid or duplicated submit data")
		return false
	} else {
		fmt.Fprintf(w, "submited")
		return true

	}

}
