//
// Created by liinda on 05.01.19.
//

#ifndef CORE_API_H
#define CORE_API_H

#include <iostream>
#include <string>
#include "Store.h"

using namespace std;

class Api {
//    void set_store(Store *store);
    Store *store;
    string add_account(string &data);
  public:
    explicit Api(Store *store);
    ~Api(){cout << "~Api" << endl;};
    string dispatch(const string &data);
};

#endif //CORE_API_H
