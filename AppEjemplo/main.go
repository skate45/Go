package main

import (
	"time"
	"fmt"
	"sync"
	//"strconv"
	"ejercicioApartados/helper"
)

//Package level variables
const conferenceTickets=50
var conferenceName="Go conference"
var remainingTickets uint=50
//var bookings= make([] map[string]string,0)
var bookings= make([] UserData,0)

//var bookings=[50]string {}  //Array
//var bookings [] string //Slice
type UserData struct{
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var waitGroup=sync.WaitGroup{}


func main(){
	

	//fmt.Printf("ConferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n",conferenceTickets,remainingTickets,conferenceName)
	greetUsers()
	
	//for remainingTickets>0 && len(bookings)<50{
		
	firstName,lastName,email,userTickets:=getUserInput()

		//var isValidName bool=len(firstName)>=2 && len(lastName)>=2
		isValidName,isValidEmail,isValidTicketNumber:=helper.ValidateUserInput(firstName,lastName,email,userTickets,remainingTickets)
		//isValidCity:= city=="Singapour" || city=="London"
		//isInvalidCity:= !(city=="Singapour" || city=="London")

		if isValidName && isValidEmail && isValidTicketNumber {
			bookTicket(userTickets, firstName, lastName, email)
			waitGroup.Add(1)
			go sendTicket(userTickets,firstName,lastName,email)

			firstNames:=getFirstNames(bookings)
			fmt.Printf("Thes first names of all the bookings are: %v\n",firstNames)
			
			fmt.Printf("These are all the bookings: %v\n",bookings)
			//var noTicketsRemaining=remainingTickets==0
			
			//noTicketsRemaining:= (remainingTickets==0)
			if remainingTickets==0 {
				//end the program
				fmt.Println("Our conference is book out, come back next year")
				//break
			}
		}else{
			//fmt.Printf("Your input data is invalid. Try again.\n")
			if !isValidName{
				fmt.Println("First name or last name you entered is too short")
			}
			if !isValidEmail{
				fmt.Println("Email address you entered does not contain @ sign")
			}
			if(!isValidTicketNumber){
				fmt.Println("The number of tickets you entered is not valid")
			}
			if(userTickets>remainingTickets){
				fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n",remainingTickets,userTickets)
			}
		}
	//}
	waitGroup.Wait()

}

func greetUsers(){
	fmt.Printf("Welcome to our %v booking application\n",conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available\n",conferenceTickets,remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames(bookings []UserData) []string{
	firstNames:=[]string{}
	for _, bookingItem :=range bookings{
		firstNames=append(firstNames,bookingItem.firstName)
	}
	//fmt.Printf("Thes first names of all the bookings are: %v\n",firstNames)
	return firstNames
}


func getUserInput() (string, string, string, uint ){
	var firstName string
	var lastName string
	var email string
	var userTickets uint
	//ask user for their name

	//fmt.Println(conferenceName)
	//fmt.Println(&conferenceName)

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email address: ")
	fmt.Scan(&email)

	fmt.Println("Enter your number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName,lastName,email,userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string){
	remainingTickets= remainingTickets - userTickets

	/*
	var userData= make(map[string]string)
	userData["firstName"]=firstName
	userData["lastName"]=lastName
	userData["email"]=email
	userData["numberOfTickets"]=strconv.FormatUint(uint64(userTickets),10)
	*/
	var userData= UserData{
		firstName: firstName,
		lastName: lastName,
		email: email,
		numberOfTickets: userTickets,
	}

	bookings=append(bookings,userData)
	/*
	fmt.Printf("The whole slice: %v\n",bookings)
	fmt.Printf("The first value of slice: %v\n",bookings[0])
	fmt.Printf("The type value of slice: %T\n",bookings)
	fmt.Printf("The size of the slice is: %v\n",len(bookings))
	*/
	fmt.Printf("Thank you %v %v for booking %d tickets. You will receive a confirmation email at %v\n",firstName,lastName,userTickets,email)

	fmt.Printf("%v Tickets remaining for %v \n",remainingTickets, conferenceName)

}

func sendTicket(userTickets uint,firstName string, lastName string,email string){
	time.Sleep(10*time.Second)

	var ticket=fmt.Sprintf("%v tickets for %v %v",userTickets,firstName, lastName)
	fmt.Printf("#################\n")
	fmt.Printf("Sending ticket\n %v \nto email address %v\n",ticket,email)
	fmt.Printf("#################\n")
	waitGroup.Done()
}