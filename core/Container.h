//
// Created by liinda on 19.01.19.
//

#ifndef CORE_CONTAINER_H
#define CORE_CONTAINER_H

#include <vector>
#include <map>
#include <string>

using namespace std;

class Container {
  public:
//    std::vector <string*> _data;
    std::map <string, string*> _data;
//    size_t _size = 0;
    string* get_ptr(const string &value);
};


#endif //CORE_CONTAINER_H
