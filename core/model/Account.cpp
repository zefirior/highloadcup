//
// Created by liinda on 31.12.18.
//

#include <iostream>
#include "Account.h"
using namespace std;

Account::Account(
  int id, int birth, Premium *premium, char sex, char status, string phone,
  string email, string *fname, string *sname, string *country, string *city, Like *like
) :
  id(id), birth(birth), premium(premium), sex(sex), status(status), phone(phone),
  email(email), fname(fname), sname(sname), country(country), city(city), like(like)
  {};

void Account::repr() {
  std::cout << "<Account: "
            << "id=" << this->id
            << ", birth=" << this->birth
            << ", sex=" << this->sex
            << ", status=" << this->status
            << ", phone=" << this->phone
            << ", email=" << this->email
            << ", fname=" << *this->fname
            << ", sname=" << *this->sname
            << ", country=" << *this->country
            << ", city=" << *this->city
            << ">" << std::endl;
};
