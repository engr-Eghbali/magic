package myplugins

import (
	"database/sql"
	"net"
	"net/http"
	"time"
	"strconv"
	"strings"
)

func CheckErrr(err error) {

	if err != nil {

		panic(err)

	}

}


//!! move this on reddist on memory DB...its so quick.

func Checkbrute(id int64, r *http.Request) (flg bool) {

	db, err := sql.Open("mysql", "root:toor@/test?charset=utf8")

	defer db.Close()

	CheckErrr(err)

	start := time.Now()

	now, err := strconv.ParseInt(start.Format("20060102150405"), 10, 64)

	CheckErrr(err)

	stmt, err := db.Query("SELECT `Time`,`no` FROM `brute`	WHERE `ID`=" + strconv.FormatInt(id, 10))

	CheckErrr(err)

	var ttime int64

	var no int

	stmt.Scan(&ttime, &no)

	if ttime-now < 5 {

		if no >= 3 {

			IP := IPfinder(r)

			stmt4, err := db.Prepare("UPDATE `brute` SET `IP`=" + IP + ",`no`=no+1 WHERE `ID`=?")

			CheckErrr(err)

			_, err4 := stmt4.Exec(id)

			CheckErrr(err4)

			return true

		} else {

			stmt2, err := db.Prepare("UPDATE `brute` SET `Time`=" + strconv.FormatInt(now, 10) + ",`no`=no+1 WHERE `ID`=?")

			CheckErrr(err)

			_, err2 := stmt2.Exec(id)

			CheckErrr(err2)

			return false

		}

	} else {

		stmt3, err := db.Prepare("UPDATE `brute` SET `Time`=" + strconv.FormatInt(now, 10) + " WHERE `ID`=?")

		CheckErrr(err)

		_, err3 := stmt3.Exec(id)

		CheckErrr(err3)

		return false

	}

}

///////////////////////////////////////////////////////////////10/////////////////////////////////////////////

//check bruteforce(in absence of cfs)

func IPfinder(r *http.Request) (ip string) {

	addr := getRealAddr(r)

	return addr

}

/////////////////////////////////////////////////////////////////8///////////////////////////

//finde IP

func getRealAddr(r *http.Request) string {

	remoteIP := ""

	// the default is the originating ip. but we try to find better options because this is almost

	// never the right IP

	if parts := strings.Split(r.RemoteAddr, ":"); len(parts) == 2 {

		remoteIP = parts[0]

	}

	// If we have a forwarded-for header, take the address from there

	if xff := strings.Trim(r.Header.Get("X-Forwarded-For"), ","); len(xff) > 0 {

		addrs := strings.Split(xff, ",")

		lastFwd := addrs[len(addrs)-1]

		if ip := net.ParseIP(lastFwd); ip != nil {

			remoteIP = ip.String()

		}

		// parse X-Real-Ip header

	} else if xri := r.Header.Get("X-Real-Ip"); len(xri) > 0 {

		if ip := net.ParseIP(xri); ip != nil {

			remoteIP = ip.String()

		}

	}

	return remoteIP

}

///////////////////////////////////////////////////////////////////////////////////////9///////////////////////////

////////real addr find
