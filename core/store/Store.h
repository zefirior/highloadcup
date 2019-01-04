//
// Created by liinda on 31.12.18.
//

#ifndef CORE_STORE_H
#define CORE_STORE_H

#include <vector>
using namespace std;

template <typename T>
class Store {
  std::vector <T> data;
  public:
    void add_item(T item);
    T get_item(int index);
};

#endif //CORE_STORE_H
