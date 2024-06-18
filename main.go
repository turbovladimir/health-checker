package main

import (
	"fmt"
	"net/http"
)

type Config struct {
	rows []*Service
}

type Service struct {
	Name string
	Host string
}

func main() {
	conf := Config{
		[]*Service{
			{
				Name: "Zaymirubli", Host: "zaymirubli.ru",
			},
		},
	}
	sender := NewSender("7347129568:AAEzHLjzZvWXxZF0FcEY5NNs0Hepm6KI3Co")

	for _, serviceConfig := range conf.rows {

		resp, err := http.Get("https://" + serviceConfig.Host + "/check")

		if err != nil {
			fmt.Println(err)
		}

		if resp.StatusCode != 200 {
			m := fmt.Sprintf("Service #%s# Unexpecred response code %d", serviceConfig.Name, resp.StatusCode)
			err = sender.SendMessage(&Message{ChatID: 107465278, Text: m})
			fmt.Println(err)
		}
	}
}
