package query

type Query struct {
	Sex string

	StatusEq  string
	StatusNeq string

	EmailLt  string
	EmailGt  string
	EmailDom string

	FnameEq   string
	FnameAny  string
	FnameNull string

	SnameEq     string
	SnameStarts string
	SnameNull   string

	PhoneNull string

	CountryEq   string
	CountryNull string

	CityEq   string
	CityNull string

	BirthLt string
	BirthGt string

	PremiumNow  string
	PremiumNull string
}
