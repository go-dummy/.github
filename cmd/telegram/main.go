package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	text := flag.String("text", "", "")
	flag.Parse()

	_, err := http.Get("https://api.telegram.org/bot" + os.Getenv("TELEGRAM_ACCESS_TOKEN") + "/sendMessage?chat_id=" + os.Getenv("TELEGRAM_CHAT_ID") + "&text=" + *text)
	if err != nil {
		log.Fatalln(err)
	}
}
