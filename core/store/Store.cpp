//
// Created by liinda on 31.12.18.
//

#include "Store.h"
using namespace std;

template <typename T>
void Store<T>::add_item(T item) {
  this->data.push_back(item);
}

template <typename T>
T Store<T>::get_item(int index){
  return this->data[index];
}
