//
// Created by liinda on 31.12.18.
//

#include <iostream>
#include "Account.h"
using namespace std;

Account::Account(
  int id, int birth, int joined, Premium *premium, char sex, char status, string phone,
  string email, string *fname, string *sname, string *country, string *city, Like *like, Interest *interest
) :
  id(id), birth(birth), joined(joined), premium(premium), sex(sex), status(status), phone(phone),
  email(email), fname(fname), sname(sname), country(country), city(city), like(like), interest(interest)
  {};

void Account::repr() const {
  std::cout << "<Account: " << endl;
  std::cout << "  " << "id=" << id << endl;
  std::cout << "  " << "birth=" << birth << endl;
  std::cout << "  " << "joined=" << joined << endl;
  std::cout << "  " << "sex=" << sex << endl;
  std::cout << "  " << "status=" << status << endl;
  std::cout << "  " << "phone=" << phone << endl;
  std::cout << "  " << "email=" << email << endl;
  if (fname){std::cout << "  " << "fname=" << *fname << endl;}
  if (sname){std::cout << "  " << "sname=" << *sname << endl;}
  if (country){std::cout << "  " << "country=" << *country << endl;}
  if (city){std::cout << "  " << "city=" << *city << endl;}

  if (premium){cout << "  premium->from=" << premium->from << ", premium->to=" << premium->to << endl;}

  Like *_like = like;
  cout << "  likes:" << endl;
  while (_like != nullptr) {
    cout << "    like->id=" << _like->id << ", like->ts=" << _like->ts << endl;
    _like = _like->next;
  }

  Interest *_interest = interest;
  cout << "  interests:" << endl;
  while (_interest != nullptr) {
    cout << "    interest->name=" << *_interest->name << endl;
    _interest = _interest->next;
  }

  cout << ">" << std::endl;
};
