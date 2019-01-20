//
// Created by liinda on 31.12.18.
//

#ifndef CORE_STORE_H
#define CORE_STORE_H

#include <iostream>
#include <vector>
#include <unordered_map>
//#include <map>
#include "Container.h"
#include "model/Account.h"
using namespace std;

typedef unordered_map <string, string*> _store_map;

class Store {
  std::vector <Account*> data;
  Container fname_map;
  Container sname_map;
  Container country_map;
  Container city_map;
  Container interest_map;
  public:
    explicit Store (unsigned long reserve_size);

    ~Store(){cout << "~Store" << endl;};

    void
    add_item(Account* item);

    void
    parse_account(string &data);

    int
    count_account();

    Account*
    get_item(int index);
};

#endif //CORE_STORE_H
