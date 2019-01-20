//
// Created by liinda on 19.01.19.
//

#include "Container.h"

///*
string*
Container::get_ptr(const string &value) {
  string* str_ptr = new string(value);
  if (_data.count(*str_ptr)) {
    return _data.find(*str_ptr)->second;
  }
  _data[*str_ptr] = str_ptr;
  return str_ptr;
};
// */

/*
string*
Container::get_ptr(const string &value) {
  auto str_ptr = new string(value);
  unsigned long i, size;
//  map <string, string*> test_map;

  size=_data.size();
  for (i=0; i<size; i++){
    if (*_data[i] == value) {
      return _data[i];
    }
  }
  _data.push_back(str_ptr);
  return str_ptr;
}
 */
