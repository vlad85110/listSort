package main

import (
	"strconv"
	"os"
	"fmt"
	"time"
	l "listSort/list"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: listSort <number of thread>")
		return
	}

	param := args[1]
	num, err := strconv.Atoi(param)
	if err != nil {
		fmt.Println("Invalid number:", param)
		return
	}

	list := l.NewList()
	
	for i := 0; i < num; i++ {
		go sortRoutine(list)
	}

	for true {
		var str string
		fmt.Scanln(&str)

		if len(str) == 0 {
			fmt.Println(list)
			continue
		}

		if str == "q" {
			break;
		}

		if len(str) >= 80 {
			for i := 0; i < len(str); i+=40 {
				s := str[i : i+40]	
				list.AddToBegin(s)
			}
		} else {
			list.AddToBegin(str)
		}
	}
}

func sortRoutine(list *l.List) {
	for true {
		time.Sleep(1)
		list.Sort()
	}
}
