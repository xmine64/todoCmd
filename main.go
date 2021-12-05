package main

import (
	"TodoCmd/cmd"
	"fmt"
)

func main() {
	fmt.Println(start())

	cmd.CheckFiles()
}

func start() string {
	start :=
		`
  _______          _                           
 |__   __|        | |                          
    | | ___     __| | ___     __ _ _ __  _ __  
    | |/ _ \   / _'' |/ _ \   / _'' | '_ \| '_ \ 
    | | (_) | | (_| | (_) | | (_| | |_) | |_) |
    |_|\___/   \__,_|\___/   \__,_| .__/| .__/ 
                                  | |   | |    
                                  |_|   |_|

`

	return start
}
