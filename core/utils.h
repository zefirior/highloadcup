//
// Created by liinda on 01.01.19.
//

#ifndef CORE_UTILS_H
#define CORE_UTILS_H

#include "Store.h"

namespace utils {
    int int_from_string(string data);
    string next_block(string const &data, size_t &left, size_t &right);
//    void parse_account(string data);
    void fill_massive(Store *acc_store, int acc_num);
}

#endif //CORE_UTILS_H
