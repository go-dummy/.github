package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	text := flag.String("text", "", "")
	flag.Parse()

	fmt.Println(os.Getenv("TELEGRAM_CHAT_ID"))

	resp, err := http.Get("https://api.telegram.org/bot" + os.Getenv("TELEGRAM_ACCESS_TOKEN") + "/sendMessage?chat_id=" + os.Getenv("TELEGRAM_CHAT_ID") + "&text=" + *text)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
}
