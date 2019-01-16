//
// Created by liinda on 01.01.19.
//

#include <iostream>
#include "utils.h"
#include "model/Account.h"

int utils::int_from_string(string data){
  return atoi(data.c_str());
}

string utils::next_block(string const &data, size_t &left, size_t &right){
  if (right != 0){
    left = right + 1;
  }
  right = data.find(' ', left);
  return data.substr(left, right - left);

}

/*void utils::fill_massive(Store *acc_store, int acc_num) {
  Like *like_1, *like_2;
  Premium *premium;
  Account *account;
  for (int32_t i = 0; i < acc_num; i++) {
    like_1 = new Like();
    like_1->id = 123423;
    like_1->ts = 2431234;
    like_1->next = nullptr;

    like_2 = new Like();
    like_2->id = 451234;
    like_2->ts = 24313523;
    like_2->next = like_1;
    premium = new Premium();
    premium->premium_from = 1324;
    premium->premium_to = 132523;

    account = new Account(
      123, 23452543, premium, '2', '3', "89126290012", "wizard.liinda@gmail.com",
      "liinda", "wizard", "рус", "ekat", like_2
    );
    acc_store->add_item(account);
  }
}*/
