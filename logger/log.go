package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

// AddLog will track and log when u use app
func AddLog(text string) {
	file, err := os.OpenFile(".log", os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		log.Fatal(err)
		return
	}

	// use defer for closing file !
	defer file.Close()

	_, err = file.WriteString(fmt.Sprintf("[%v] %v \n", time.Now(), text))
	if err != nil {
		log.Fatal(err)
		return
	}
}
