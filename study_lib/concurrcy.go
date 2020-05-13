package study_lib

import (
	"fmt"
	"math/rand"
	// "strconv"
	"time"
)

func ProductMsg(name string, c chan string) {

	for {

		msg := fmt.Sprintf("%s message : %d", name, rand.Int())
		fmt.Println(msg)
		c <- msg

		time.Sleep(5 * time.Second)

	}

}

func CustomMsg(c chan string) {

	for {

		// rmsg := "消费：" + <-c
		msg := fmt.Sprintf("消费：%s", <-c)
		fmt.Println(msg)

	}

}
