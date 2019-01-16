//
// Created by liinda on 31.12.18.
//

#ifndef CORE_STORE_H
#define CORE_STORE_H

#include <iostream>
#include <vector>
#include <map>
#include "model/Account.h"
using namespace std;

class Store {
  std::vector <Account*> data;
  map <string, string*> fname_map;
  map <string, string*> sname_map;
  map <string, string*> country_map;
  map <string, string*> city_map;
  public:
    Store (int reserve_size);
    void add_item(Account* item);
    string* get_ptr_from_map(map <string, string*> *container, string data);
    void parse_account(string data);
    Account* get_item(int index);
};

#endif //CORE_STORE_H
