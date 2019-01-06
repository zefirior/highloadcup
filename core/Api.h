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
    Store* store;
    string add_account(string data);
  public:
    Api(Store* store);
    string dispatch(string* data);
};

#endif //CORE_API_H
