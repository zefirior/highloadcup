//
// Created by liinda on 30.12.18.
//

#ifndef CORE_ACCOUNT_H
#define CORE_ACCOUNT_H

#include <string>
using namespace std;

struct Like {
  int32_t id;
  int32_t ts;
  Like *next;
};

struct Interest {
  string *name;
  Interest *next;
};

struct Premium {
  int32_t from;
  int32_t to;
};

class Account {
  int32_t id;
  int32_t birth;
  int32_t joined;
  char sex;
  char status;
  string phone;
  string email;
  string *fname;
  string *sname;
  string *country;
  string *city;
  Interest *interest;
  Premium *premium;
  Like *like;

public:
  Account(
    int32_t id,
    int32_t birth,
    int32_t joined,
    Premium *premium,
    char sex,
    char status,
    string phone,
    string email,
    string *fname,
    string *sname,
    string *country,
    string *city,
    Like *like,
    Interest *interest
  );
  void repr() const;
};

#endif //CORE_ACCOUNT_H
