//
// Created by liinda on 31.12.18.
//

#include <iostream>
#include "Account.h"
using namespace std;

Account::Account(
  int id, int birth, Premium *premium, char sex, char status, string phone,
  string email, string *fname, string *sname, string *country, string *city, Like *like, Interest *interest
) :
  id(id), birth(birth), premium(premium), sex(sex), status(status), phone(phone),
  email(email), fname(fname), sname(sname), country(country), city(city), like(like), interest(interest)
  {};

void Account::repr() const {
  std::cout << "<Account: " << endl
            << "  " << "id=" << id << endl
            << "  " << "birth=" << birth << endl
            << "  " << "sex=" << sex << endl
            << "  " << "status=" << status << endl
            << "  " << "phone=" << phone << endl
            << "  " << "email=" << email << endl
            << "  " << "fname=" << *fname << endl
            << "  " << "sname=" << *sname << endl
            << "  " << "country=" << *country << endl
            << "  " << "city=" << *city << endl;

  cout << "  premium->from=" << premium->from << ", premium->to=" << premium->to << endl;

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
