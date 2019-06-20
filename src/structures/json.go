package structures

import "fmt"

type Json struct {
	Accounts []JAccount
}

type JAccount struct {
	Id        int
	Birth     int
	Joined    int
	Country   string
	City      string
	Email     string
	Fname     string
	Sname     string
	Phone     string
	Sex       string
	Status    string
	Interests []string
	Likes     []Likes
	Premium   *Premium
}

type Premium struct {
	Start  int32
	Finish int32
}

type Likes struct {
	Ts int32
	Id int32
}

func (a *JAccount) Print() {
	fmt.Println("Id: ", a.Id)
	fmt.Println("Birth: ", a.Birth)
	fmt.Println("Joined: ", a.Joined)
	fmt.Println("Country: ", a.Country)
	fmt.Println("City: ", a.City)
	fmt.Println("Email: ", a.Email)
	fmt.Println("Fname: ", a.Fname)
	fmt.Println("Sname: ", a.Sname)
	fmt.Println("Phone: ", a.Phone)
	fmt.Println("Sex: ", a.Sex)
	fmt.Println("Status: ", a.Status)
	fmt.Println("Interests: ", a.Interests)
	fmt.Println("Likes: ", a.Likes)
	fmt.Println("Premium: ", a.Premium)
}
