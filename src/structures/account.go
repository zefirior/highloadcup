package structures

import "fmt"

type JsonAcc struct {
	Accounts []Account // `json:"accounts"`
}

type Account struct {
	Id        int32
	Birth     int32
	Joined    int32
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
	Premium   Premium
}

type Premium struct {
	Start  int32
	Finish int32
}

type Likes struct {
	Ts int32
	id int32
}

func (a *Account) Print() {
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
