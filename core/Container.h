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
    std::vector <string> _data;
    int16_t cur_index = 0;
//    std::map <string, string*> _data;
  public:
//    size_t _size = 0;
    int get_idx(const string &value);
};


#endif //CORE_CONTAINER_H
