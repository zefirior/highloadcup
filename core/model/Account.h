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

struct Premium {
    int32_t premium_from;
    int32_t premium_to;
};

class Account {
    int32_t id;
    int32_t birth;
    char sex;
    char status;
    string phone;
    string email;
    string fname;
    string sname;
    string country;
    string city;

public:
    Premium premium;
    Like *like;
    Account(
            int32_t id,
            int32_t birth,
            Premium premium,
            char sex,
            char status,
            string phone,
            string email,
            string fname,
            string sname,
            string country,
            string city,
            Like *like
    );
    void repr();
};

#endif //CORE_ACCOUNT_H
