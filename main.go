package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

func main() {
	color.Set(color.FgHiWhite, color.BgBlack)
	var periodName int
	openingText := "Hello Mr.Walz\n Enter The Class Period -> "
	boldWhite := color.New(color.FgWhite, color.Bold)
	boldWhite.Print(openingText)
	fmt.Scan(&periodName)
	for {
		switch periodName {
		case 1, 2, 3, 4, 5, 6:
			fmt.Print("\nThe Period Is ", periodName)
			time.Sleep(2 * time.Second)
			keyboard.Open()
			defer keyboard.Close()
			getStudents(periodName)
		default:
			fmt.Println("Invalid Period Name")
			return
		}
	}
}

func getStudents(periodName int) {
	period1 := []string{"Apple", "Bob", "Charlie", "David", "Eva"}
	period2 := []string{"Dany", "Ella", "Finn", "George", "Hannah"}
	period3 := []string{"Jane", "Ivan", "Jack", "Kara", "Liam"}
	period4 := []string{"Mark", "Nina", "Olivia", "Paul", "Quinn"}
	period5 := []string{"Johann", "Rachel", "Sam", "Tom", "Uma"}
	period6 := []string{"Isis", "Victor", "Will", "Xena", "Yara"}

	var studentList []string
	var selectedStudents []string

	switch periodName {
	case 1:
		studentList = period1
	case 2:
		studentList = period2
	case 3:
		studentList = period3
	case 4:
		studentList = period4
	case 5:
		studentList = period5
	case 6:
		studentList = period6
	default:
		fmt.Println("Invalid period number.")
		return
	}

	rand.Seed(time.Now().UnixNano())

	for len(studentList) > 0 {
		fmt.Println("\nPress 'A' to select a student...")

		for {
			key, _, err := keyboard.GetKey()
			if err != nil {
				fmt.Println("Error reading keyboard input:", err)
				return
			}
			if key == 'a' {
				randomIndex := rand.Intn(len(studentList))
				selectedStudent := studentList[randomIndex]
				boldWhite := color.New(color.FgWhite, color.Bold)
				boldWhite.Println(selectedStudent)
				selectedStudents = append(selectedStudents, selectedStudent)
				studentList = append(studentList[:randomIndex], studentList[randomIndex+1:]...)
				if len(studentList) == 0 {
					fmt.Println("\nAll students have been selected.")
					fmt.Print("\nExiting now in 3 seconds")
					time.Sleep(3 * time.Second)
					os.Exit(0)
				}
				break
			}
		}
	}
}
