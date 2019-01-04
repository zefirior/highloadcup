//
// Created by liinda on 31.12.18.
//

#ifndef CORE_STORE_H
#define CORE_STORE_H

#include <vector>
#include "model/Account.h"
using namespace std;

class Store {
  std::vector <Account*> data;
  public:
    Store (int reserve_size);
    void add_item(Account* item);
    Account* get_item(int index);
};

#endif //CORE_STORE_H
