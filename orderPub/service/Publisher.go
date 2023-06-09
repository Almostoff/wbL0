package service

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"os"
	"strings"
)

type StanClient struct {
	sc stan.Conn
}

func CreateStan() *StanClient {
	stan := StanClient{}
	return &stan
}

func (sCl *StanClient) Connect(clusterID string, clientID string, url string) error {
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(url))
	if err != nil {
		return err
	}
	sCl.sc = sc
	return err
}

func (sCl *StanClient) Close() {
	if sCl != nil {
		sCl.sc.Close()
	}
}

func (sCl *StanClient) PublishFromFile(ch string, filepath string) error {
	text, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	sCl.sc.Publish(ch, text)
	return err
}

func (sCl *StanClient) PublishFromStdinCycle(ch string) error {
	var filepath string
	var err error
	for {
		var text []byte
		fmt.Print("Enter filepath: ")
		fmt.Fscan(os.Stdin, &filepath)
		filepath = strings.TrimSuffix(filepath, "\r\n")
		if filepath == "exit" {
			return nil
		}
		text, err = os.ReadFile(filepath)
		if err != nil {
			log.Println("Error reading file")
			return err
		}
		sCl.sc.Publish(ch, text)
	}
	return err
}
