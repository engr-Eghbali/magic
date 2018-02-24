package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	magic_struct "./pkg"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func upload(c echo.Context) error {

	// Read form fields

	U := &magic_struct.UpDT{ID: c.FormValue("ID"), Path: c.FormValue("path")}

	//mongo connect
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	////////
	result := magic_struct.Userdata{}
	cc := session.DB("magic").C("userInfo")

	log.Print(U.ID)
	err = cc.FindId(bson.ObjectIdHex(U.ID)).One(&result)
	if err != nil {
		//redirect here
		log.Fatal(err)
	}
	if result.Stat {

		log.Print("%%%%%%%%%%uplogin done %%%%%%%%%%")
	}

	//------------
	// Read files
	//------------

	// Multipart form
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	files := form.File["files"]

	for _, file := range files {
		// Source
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Destination
		dst, err := os.Create(file.Filename)
		if err != nil {
			return err
		}
		defer dst.Close()

		reader := bufio.NewReader(src)
		writer := bufio.NewWriter(dst)
		buffer := make([]byte, 1024)
		for {
			n, err := reader.Read(buffer)
			if err != nil && err != io.EOF {
				panic(err)
			}
			if n == 0 {
				break
			}

			if _, err := writer.Write(buffer[:n]); err != nil {
				panic(err)
			}

		}
		if err = writer.Flush(); err != nil {
			panic(err)
		}

	}

	return c.HTML(http.StatusOK, fmt.Sprintf("<p>Uploaded successfully %d files with fields name=%s and email=%s.</p>", len(files), U.ID, U.Path))
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.POST("/upload", upload)
	e.Logger.Fatal(e.Start(":1323"))
}
