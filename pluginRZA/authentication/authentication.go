package myplugins

import (
	"database/sql"
	f "fmt"
	"log"

	"net/http"

	"regexp"

	"strconv"

	"strings"

	davinchi "./layer2/layer3"
	typeRZA "./layer2/layer3/typedef"
	_ "github.com/go-sql-driver/mysql"
)

func checkErr(err error) {

	if err != nil {

		panic(err)

	}

}

/////////////////////////////////////////////////////////////7///////////////////////

//make panic on  error
//"id=8736&name=day&pass=me2&add=24#56#"
func Initialsubmit(r *http.Request, in *typeRZA.Userdata) {
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

func Initiallogin(r *http.Request, in *typeRZA.Userdata) {

	r.ParseForm()

	var tid string = r.Form["id"][0]

	in.Geo.Y, in.Geo.X, in.Addr, in.Uname = 22.0000, 22.0000, "login", "login"

	in.Uid, _ = strconv.ParseInt(tid, 10, 0)

	in.Pass = r.Form["pass"][0]

}

func Login(w http.ResponseWriter, r *http.Request, input typeRZA.Userdata) (flg bool) {

	db, err := sql.Open("mysql", "root:toor@/test?charset=utf8")
	defer db.Close()
	var name, addr, x, y string
	var pass []byte
	checkErr(err)
	stmt, err2 := db.Prepare("SELECT `pass`,`uname`,`address`,`X`,`Y` FROM `userinfo` WHERE `ID`=?")
	checkErr(err2)

	res, err3 := stmt.Query(input.Uid)

	if err3 != nil {

		checkErr(err3)
		return false
	} else {
		defer res.Close()
		for res.Next() {
			xerr := res.Scan(&pass, &name, &addr, &x, &y)
			f.Print(name)
			if xerr != nil {
				log.Fatal(xerr)
			}

		}
		if input.Pass == davinchi.Decipher(pass) {
			return true
		} else {
			return false
		}
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

func Validation(in *typeRZA.Userdata, w http.ResponseWriter, r *http.Request) (flg bool) {

	var empty_flg bool = false

	if len(in.Pass) == 0 || in.Uid == 0 {

		empty_flg = true

		f.Fprintf(w, "fill the form") //empty form error

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

func Submit(w http.ResponseWriter, r *http.Request, input typeRZA.Userdata) (flg bool) {

	db, err := sql.Open("mysql", "root:toor@/test?charset=utf8")
	defer db.Close()
	checkErr(err)

	stmt, err2 := db.Prepare("INSERT INTO `userinfo` (`ID`,`uname`,`pass`,`address`,`X`,`Y`) VALUES (?,?,?,?,?,?)")
	checkErr(err2)

	_, err3 := stmt.Exec(strconv.FormatInt(input.Uid, 10), input.Uname, davinchi.Cipher(input.Pass), input.Addr, strconv.FormatFloat(input.Geo.X, 'f', 6, 64), strconv.FormatFloat(input.Geo.Y, 'f', 6, 64))

	if err3 != nil {

		if strings.Contains(err3.Error(), "Duplicate") {

			f.Fprintf(w, "user duplicate")

			return false

		}

		checkErr(err3)

		return false

	} else {

		return true

	}

}
