package structures

type Account struct {
	Id        int
	Birth     int
	Joined    int
	Country   int
	City      int
	Email     string
	Fname     int
	Sname     int
	Phone     string
	Sex       int
	Status    int
	Interests []int
	Likes     []Likes
	Premium   Premium
}

func (a1 *Account) Equal(a2 Account) bool {
	if
		a1.Id == a2.Id &&
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

func SexIndex(sex string) int {
	switch sex {
	case "f":
		return 0
	case "m":
		return 1

	default:
		panic("Unexpected sex " + sex)
	}
}

func SexValue(i int) string {
	sex := [2]string{"f", "m"}
	return sex[i]
}

func StatusIndex(status string) int {
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

func StatusValue(i int) string {
	sex := [3]string{"свободны", "заняты", "всё сложно"}
	return sex[i]
}
