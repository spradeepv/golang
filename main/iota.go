package main

import (
	"fmt"
	"github.com/spradeepv/golang/constants"
)

func main() {
	fmt.Println("1 KB is", constants.KB, "bytes!")
	fmt.Println("1 MB is", constants.MB, "bytes!")
	fmt.Println("1 GB is", constants.GB, "bytes!")
	fmt.Println("1 TB is", constants.TB, "bytes!")
	fmt.Println("1 PB is", constants.PB, "bytes!")
	fmt.Println("1 EB is", constants.EB, "bytes!")
	fmt.Println("1 ZB is", constants.ZB, "bytes!")
	fmt.Println("1 YB is", constants.YB, "bytes!")

	fmt.Print("SUNDAY is the ", constants.SUNDAY, "st day of the week\n")
	fmt.Print("MONDAY is the ", constants.MONDAY, "nd day of the week\n")
	fmt.Print("TUESDAY is the ", constants.TUESDAY, "rd day of the week\n")
	fmt.Print("WEDNESDAY is the ", constants.WEDNESDAY, "th day of the week\n")
	fmt.Print("THURSDAY is the ", constants.THURSDAY, "th day of the week\n")
	fmt.Print("FRIDAY is the ", constants.FRIDAY, "th day of the week\n")
	fmt.Print("SATURDAY is the ", constants.SATURDAY, "th day of the week")
}
