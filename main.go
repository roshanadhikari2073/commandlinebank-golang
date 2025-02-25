//author: roshanadhikari
//gitusername: roshanadhikari2073
//details :- all codes are raw and self written, -> BANK CLI

package main

import (
	"bufio"
	getLogo "cliapplications/assets"
	sqlconn "cliapplications/dataconfig"
	"cliapplications/src"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"

	"golang.org/x/term"
)

// This leads to the main login page
func main() {
	loginPage("")
}

// the main login page of the application
func loginPage(loginMessage string) {
	clearTheTerminal(src.CLEARTERMINAL)
	if loginMessage != "" {
		fmt.Println(loginMessage)
	}
	redirect, status, alert := false, src.CHECKCREDS, false
	var checkaccount string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(" PRESS --- Y --- IF YOU ARE RETURING USER OR PRESS --- X --- TO CREATE NEW ACCOUNT... ELSE PRESS 9 TO EXIT")
	fmt.Scanf("%s", &checkaccount)
	if strings.ToLower(checkaccount) == "y" {
		for i := 0; i < 5; i++ {
			clearTheTerminal(src.CLEARTERMINAL)
			if i > 0 || alert {
				println(status)
				spacingToTheExit("", 2)
			}
			println("ENTER THE RIGHT CREDENTIALS TO ACCESS THE BANKING APPLICATION")
			spacingToTheExit("", 2)

			fmt.Print("Enter Username: ")
			username, err := reader.ReadString('\n')
			if err != nil {
				loginPage("Something went wrong with the application, reopening the application....")
			}
			verifyTheUserName := src.VerifyTheUserName(strings.TrimSpace(username))
			if verifyTheUserName {
				fmt.Print("Enter Password: ")
				fmt.Println("\033[8m") // Hide input
				bytePassword, err := term.ReadPassword(int(syscall.Stdin))
				fmt.Println("\033[28m")
				if err != nil {
					print("error")
				}

				password := strings.TrimSpace(string(bytePassword))
				pin, _ := strconv.Atoi(password)
				if err != nil {
					// handle error
					fmt.Println(err)
					os.Exit(2)
				}
				successFlag := src.TakeTheUserCreds(strings.TrimSpace(username), pin)
				// after the login gets successful
				if successFlag {
					welcomeloop(strings.TrimSpace(username), true, "", true)
					redirect = true
					break
				} else {
					loginPage("Either your password or username is incorrect please re-type once again else please open a new account")
				}
			} else {
				loginPage("PLEASE ENTER CORRECT USERNAME")
			}

		}
		if !redirect {
			spacingToTheExit(".", 4)
			print(" APPLICATION HAS BEEN LOCKED... EXITING")
			clearTheTerminal("")
		}
	} else if checkaccount == "9" {
		fmt.Print("\033[H\033[2J")
		os.Exit(3)
	} else if strings.ToLower(checkaccount) == "x" {
		saveVal := src.CreateNewAccount()
		if saveVal == "success" {
			clearTheTerminal(src.CLEARTERMINAL)
			status = "* ACCOUNT HAS BEEN SUCCESSFULLY CREATED PLEASE LOGIN TO ENTER *"
			loginPage(status)
		} else if saveVal == "user_exists" {
			for {
				fmt.Print("\033[H\033[2J")
				fmt.Println("USERNAME ALREADY EXISTS.. PLEASE ENTER NEW USERNAME")
				if src.CreateNewAccount() == "success" {
					break
				}
			}
			clearTheTerminal(src.CLEARTERMINAL)
			status = "* ACCOUNT HAS BEEN SUCCESSFULLY CREATED PLEASE LOGIN TO ENTER *"
			loginPage(status)

		} else {
			alert = true
		}
	} else {
		fmt.Print("\033[H\033[2J")
		loginPage(status)
	}

}

// the main page
func welcomeloop(usertrue string, cont bool, status string, updateTheTable bool, params ...map[string]string) {
	clearTheTerminal(src.CLEARTERMINAL)
	var customerGlobalScope = make(map[string]string)
	if updateTheTable {
		customerGlobalScope = sqlconn.UserInfo(usertrue)
	} else {
		customerGlobalScope = params[0]
	}
	if cont {
		// add the ending parameters
		fmt.Println(status)
		fmt.Println(getLogo.BankLogo())
		spacingToTheExit("", 3)
		fmt.Printf("-------   WELCOME TO THE BANKING APPLICATIONS Mr. %s    -------    ", strings.ToUpper(customerGlobalScope["fullname"]))
		fmt.Printf("-  %s  -    ", time.Now().Format(time.RFC850))
		spacingToTheExit("", 3)
		fmt.Println("HINT -> TYPE NUMBERS ASSOCIATED WITH THE MODULES MENTIONED BELOW")
		spacingToTheExit("", 2)
		fmt.Println("|-----------------------------------------------|")
		fmt.Printf(" Customer Name  |        %s          \n|", customerGlobalScope["fullname"])
		fmt.Println("-----------------------------------------------|")
		fmt.Printf(" Address        |        %s          \n|", customerGlobalScope["address"])
		fmt.Println("-----------------------------------------------|")
		fmt.Printf(" Phone Number   |        %s          \n|", customerGlobalScope["phone"])
		fmt.Println("-----------------------------------------------|")
		spacingToTheExit("", 2)
		fmt.Println("[ 1 ]  -> |      CHECK BALANCE              |")
		fmt.Println("")
		fmt.Println("[ 2 ]  -> |      TAKE LOAN                  |")
		fmt.Println("")
		fmt.Println("[ 3 ]  -> |      TOP UP BALANCE             |")
		fmt.Println("")
		fmt.Println("[ 4 ]  -> |      CHECK YOUR BANK STATEMENT  |")
		fmt.Println("")
		if true {
			fmt.Println("[ 5 ]  -> |      REPAY THE LOAN             |")
		}
		spacingToTheExit("", 3)
		fmt.Println("******  press 9 to exit  ******")
		_, int := takeTheUserInput("int")
		switch int {
		case 1, 2, 3, 4:
			bankingModules(usertrue, int, "", customerGlobalScope)
		case 5:
			if true {
				bankingModules(usertrue, int, "")
			} else {
				welcomeloop(usertrue, true, src.CHECKCREDS, false)
			}

		case 9:
			clearTheTerminal(src.CLEARTERMINAL)
			spacingToTheExit(".", 4)
			println(src.GOODBYENOTE)
			spacingToTheExit(".", 4)
			cont = false
		default:
			welcomeloop(usertrue, true, src.CHECKCREDS, false, customerGlobalScope)
		}
	} else {
		// add a ending paramter
		spacingToTheExit(".", 4)
		println(src.GOODBYENOTE)
		spacingToTheExit(".", 4)
	}

}

// modules for the bank application
func bankingModules(usertrue string, head int, blockStat string, custinf ...map[string]string) {
	updateTheTable := false
	if head == 2 {
		updateTheTable = true
	}
	clearTheTerminal(src.CLEARTERMINAL)
	if blockStat == "blocked" {
		fmt.Println(src.CHECKCREDS)
		spacingToTheExit("", 4)
	}
	header := [6]string{"", "MAIN BALANCE", "TAKE THE LOAN", "TOP UP BALANCE", "CHECK EXPENDITURE"}
	// create if there is loan to be paid
	if true {
		header[5] = "REPAY THE LOAN"
	}
	println(header[head])
	spacingToTheExit("", 4)
	if head == 1 {
		// to check the main balance of the user
		src.CheckBalance(custinf[0])

	} else if head == 2 {
		// take the loan
	} else if head == 3 {
		// top up the balance

	} else if head == 4 {
		// check expenditure
		src.PrintBankStatement(custinf[0])

	} else if head == 5 {
		println("5")
		//repay the loan
	}
	checkStat, Status := exitTextSignal(usertrue, head)
	if Status == "" {
		welcomeloop(usertrue, checkStat, Status, updateTheTable, custinf...)
	}
}

// this function gives the exiting text
func exitTextSignal(usertrue string, currentInt int) (bool, string) {
	spacingToTheExit("", 4)
	println(src.EXITAPP)
	var reader string
	fmt.Scanf("%s", &reader)
	if reader == "" {
		return true, ""
	} else if reader == "9" {
		return false, ""
	} else {
		bankingModules(usertrue, currentInt, "blocked")
		return false, ""
	}
}

// TODO: Implement Interface here and learn more about it
func takeTheUserInput(dataType string) (string, int) {
	username, password := "", 0
	if dataType == "str" {
		fmt.Scanf("%s", &username)
		return username, password
	} else if dataType == "int" {
		fmt.Scanf("%d", &password)
		return username, password
	} else {
		panic("error while parsing the correct datatype")
	}
}

//this function gives spacing
func spacingToTheExit(char string, totalspace int) {
	j := 0
	for {
		fmt.Println(char)
		j++
		if j >= totalspace {
			break
		}
	}
}

// This function clears the terminal and prints the designated text
func clearTheTerminal(s string) bool {
	forceClear := func() {
		fmt.Print("\033[H\033[2J")
	}
	if s == src.CLEARTERMINAL {
		forceClear()
	} else {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println(s)
		text, _ := reader.ReadString('\n')
		if text == "\n" || text != "\n" {
			forceClear()
		}
	}
	return false
}
