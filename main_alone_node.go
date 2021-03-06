package main

import "D7024E/dht"

//import "D7024E/dht"
import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	//fmt.Print("Enter id: ")
	//Id, _ := reader.ReadString('\n')
	//Id := ""
	//fmt.Println(text)

	//fmt.Println("Enter Ip: ")
	//Ip, _ := reader.ReadString('\n')
	//Ip := "localhost"
	//fmt.Scanln(Ip)

	//fmt.Println("Enter port: ")
	//Port, _ := reader.ReadString('\n')
	//fmt.Scanln(Port)

	//id0 := "00"
	//fmt.Println(Id)
	//fmt.Println(Ip)
	//fmt.Println(Port)

	//id := strings.TrimSpace(Id)
	//ip := strings.TrimSpace(Ip)
	//port := strings.TrimSpace(Port)
	//id_static := 1
	id_static := "1"
	ip_static := "localhost"
	port_static := "1111"

	n := dht.MakeDHTNode(&id_static, ip_static, port_static)
	//n.JoinRing("localhost:1112")
	go func() {

		http.HandleFunc("/chord/", dht.Chord)

		http.HandleFunc("/chord/post/", func(w http.ResponseWriter, r *http.Request) {
			dht.Post(w, r, n)
		})

		http.HandleFunc("/chord/get/", func(w http.ResponseWriter, r *http.Request) {
			dht.Get(w, r, n)
		})

		http.HandleFunc("/chord/put/", func(w http.ResponseWriter, r *http.Request) {
			dht.Put(w, r, n)
		})

		http.HandleFunc("/chord/delete/", func(w http.ResponseWriter, r *http.Request) {
			dht.Del(w, r, n)
		})

		http.HandleFunc("/chord/list/", func(w http.ResponseWriter, r *http.Request) {
			dht.List(w, r, n)
		})

		http.ListenAndServe(":"+port_static, nil)

		fmt.Println("The page is rolling")

	}()

	go func() {
		c := time.Tick(3 * time.Second)
		for {
			select {
			case <-c:
				n.AutoFingers()
				//node.autoFingers()
			}
		}
	}()

	for {
		fmt.Println("Enter command: ")
		Input, _ := reader.ReadString('\n')
		fmt.Scanln(Input)
		input := strings.TrimSpace(Input)

		switch input {
		//		case "join":
		//			go n.Join(input)

		case "joinRing":
			go n.JoinRing("localhost:1000")

			//		case "changePredecessor":
			//			go n.changePredecessor(input)
		case "fingers":
			go n.FingerPrint()

		case "id":
			go n.IdPrint()

		case "preid":
			go n.PreID()

		case "sucid":
			go n.SucID()

		case "ping":
			fmt.Println("hate")
			go n.Ping()

		}

	}

}
