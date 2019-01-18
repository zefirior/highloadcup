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
public:
  std::vector <Account*> data;
  map <size_t, string*> *fname_map;
  map <size_t, string*> *sname_map;
  map <size_t, string*> *country_map;
  map <size_t, string*> *city_map;
    explicit Store (unsigned long reserve_size);
    void add_item(Account* item);
    string* get_ptr_from_map(map <size_t, string*> *container, string const &data);
    void parse_account(string data);
    Account* get_item(int index);
};

#endif //CORE_STORE_H
