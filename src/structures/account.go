package structures

import "fmt"

type Account struct {
	Sex     int8
	Status  int8
	Country int16
	City    int16
	Fname   int16
	Sname   int16
	Domain  int16

	Id     int
	Birth  int
	Joined int

	Email string
	Phone string

	Interests []int16
	Likes     []Likes
	Premium   *Premium
}

func (a1 *Account) Equal(a2 Account) bool {
	if a1.Id == a2.Id &&
		a1.Birth == a2.Birth &&
		a1.Joined == a2.Joined &&
		a1.Country == a2.Country &&
		a1.City == a2.City &&
		a1.Email == a2.Email &&
		a1.Fname == a2.Fname &&
		a1.Sname == a2.Sname &&
		a1.Phone == a2.Phone &&
		a1.Sex == a2.Sex &&
		a1.Status == a2.Status {
		return true
	}
	return false
}

func (acc *Account) Print() {
	fmt.Println("    Id: ", acc.Id)
	fmt.Println("    Sex: ", acc.Sex)
	fmt.Println("    Status: ", acc.Status)
	fmt.Println("    Birth: ", acc.Birth)
	fmt.Println("    Country: ", acc.Country)
	fmt.Println("    City: ", acc.City)
	fmt.Println("    Fname: ", acc.Fname)
	fmt.Println("    Sname: ", acc.Sname)
	fmt.Println("    Email: ", acc.Email)
	fmt.Println("    Domain: ", acc.Domain)
	fmt.Println("    Phone: ", acc.Phone)
	fmt.Println("    Joined: ", acc.Joined)
	fmt.Println("    Interests: ", acc.Interests)
	fmt.Println("    Likes: ", acc.Likes)
	fmt.Println("    Premium: ", acc.Premium)
}

func SexIndex(sex string) int8 {
	switch sex {
	case "f":
		return 0
	case "m":
		return 1

	default:
		panic("Unexpected sex " + sex)
	}
}

func SexValue(i int8) string {
	sex := [2]string{"f", "m"}
	return sex[i]
}

func StatusIndex(status string) int8 {
	switch status {
	case "свободны":
		return 0
	case "заняты":
		return 1
	case "всё сложно":
		return 2

	default:
		panic("Unexpected status " + status)
	}
}

func StatusValue(i int8) string {
	sex := [3]string{"свободны", "заняты", "всё сложно"}
	return sex[i]
}
