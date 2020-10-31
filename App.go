package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/robfig/cron"
)

func main() {
	host := os.Args[1]
	log.Println("report host --> ", host)
	c := cron.New()
	c.AddFunc("*/5 * * * * *", func() {
		defer func() {
			e := recover()
			if e != nil {
				log.Println(e)
			}
		}()
		log.Println("report ip")
		resp, err := http.Get(host)
		if err != nil {
			log.Println(err.Error())
		}
		body, err := ioutil.ReadAll(resp.Body)
		log.Println(string(body))
	})
	c.Start()
	select {}
}
