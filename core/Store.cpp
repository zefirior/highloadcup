//
// Created by liinda on 31.12.18.
//

#include "Store.h"
#include "model/Account.h"
using namespace std;

Store::Store(int reserve_size) {
  this->data.reserve(reserve_size);
}

void Store::add_item(Account* item) {
  this->data.push_back(item);
}

Account* Store::get_item(int index){
  return this->data[index];
}
