package main

import (
	"log"
	"net/http"

	"github.com/gabrielcerdam-olxautos/goreddit/mysql"
	"github.com/gabrielcerdam-olxautos/goreddit/web"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	store, err := mysql.NewStore("root:@(127.0.0.1:3306)/godb?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	h := web.NewHandler(store)
	http.ListenAndServe(":3000", h)
}
