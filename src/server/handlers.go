package server

import (
	"fmt"
	"highload/query"
	s "highload/structures"
	"net/http"
	"net/url"
)

func init() {
	fillQueryParams()

	http.HandleFunc("/accounts/filter/", filterHandler)
}

var filterQueryArgs = make(map[string]argSpec)

type argSpec struct {
	AllowValues []string
}

func (a argSpec) InValues(value string) bool {
	fmt.Println("len argSpec", len(a.AllowValues))
	if len(a.AllowValues) == 0 {
		return true
	}
	for _, v := range a.AllowValues {
		if v == value {
			return true
		}
	}
	return false
}

func fillQueryParams() {
	filterQueryArgs["sex_eq"] = argSpec{[]string{s.SexF, s.SexM}}

	filterQueryArgs["status_eq"] = argSpec{[]string{s.StatFree, s.StatBusy, s.StatComplex}}
	filterQueryArgs["status_neq"] = argSpec{[]string{s.StatFree, s.StatBusy, s.StatComplex}}

	filterQueryArgs["email_domain"] = argSpec{}
	filterQueryArgs["email_lt"] = argSpec{}
	filterQueryArgs["email_gt"] = argSpec{}

	filterQueryArgs["fname_eq"] = argSpec{}
	filterQueryArgs["fname_any"] = argSpec{}
	filterQueryArgs["fname_null"] = argSpec{}

	filterQueryArgs["sname_eq"] = argSpec{}
	filterQueryArgs["sname_starts"] = argSpec{}
	filterQueryArgs["sname_null"] = argSpec{}

	filterQueryArgs["phone_code"] = argSpec{}
	filterQueryArgs["phone_null"] = argSpec{}

	filterQueryArgs["country_eq"] = argSpec{}
	filterQueryArgs["country_null"] = argSpec{}

	filterQueryArgs["city_eq"] = argSpec{}
	filterQueryArgs["city_any"] = argSpec{}
	filterQueryArgs["city_null"] = argSpec{}

	filterQueryArgs["birth_lt"] = argSpec{}
	filterQueryArgs["birth_gt"] = argSpec{}
	filterQueryArgs["birth_year"] = argSpec{}

	filterQueryArgs["interests_contains"] = argSpec{}
	filterQueryArgs["interests_any"] = argSpec{}

	filterQueryArgs["likes_contains"] = argSpec{}

	filterQueryArgs["premium_now"] = argSpec{}
	filterQueryArgs["premium_null"] = argSpec{}

	filterQueryArgs["limit"] = argSpec{}
	filterQueryArgs["query_id"] = argSpec{}
}

func filterHandler(writer http.ResponseWriter, r *http.Request) {
	//fmt.Println(filterQueryArgs)
	//fmt.Println(r.URL.Query())
	//fmt.Println(validateQuery(r.URL.Query()))
}

func validateQuery(values url.Values) (*query.Query, bool) {
	for k, v := range values {
		spec, ok := filterQueryArgs[k]
		if !ok {
			return nil, false
		}
		if ok := spec.InValues(v[0]); !ok {
			return nil, false
		}
	}
	return nil, true
}
