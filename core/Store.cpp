//
// Created by liinda on 31.12.18.
//

#include "Store.h"
#include "utils.h"
#include "model/Account.h"
using namespace std;

Store::Store(unsigned long reserve_size) {
  this->data.reserve(reserve_size);
}

void Store::add_item(Account* item) {
  this->data.push_back(item);
}

void
Store::parse_account(string &data) {
  size_t left=0, right=0;
  int id, birth;
  char sex, status;
  string
    block,
    marker,
    phone,
    email,
    *fname = nullptr,
    *sname = nullptr,
    *country = nullptr,
    *city = nullptr;
  Premium *premium = nullptr;
  Like *like_root = nullptr, *like_new = nullptr;
  Interest *interest_root = nullptr, *interest_new = nullptr;

  while (right < string::npos) {
    marker = utils::next_block(data, left, right);
    if (marker == "id"){
      id = utils::int_from_string(utils::next_block(data, left, right));

    } else if (marker == "b") {
      birth = utils::int_from_string(utils::next_block(data, left, right));

    } else if (marker == "p") {
      premium = new Premium;
      premium->from = utils::int_from_string(utils::next_block(data, left, right));
      premium->to = utils::int_from_string(utils::next_block(data, left, right));

    } else if (marker == "s") {
      sex = utils::next_block(data, left, right)[0];

    } else if (marker == "st") {
      status = utils::next_block(data, left, right)[0];

    } else if (marker == "ph") {
      phone = utils::next_block(data, left, right);

    } else if (marker == "e") {
      email = utils::next_block(data, left, right);

    } else if (marker == "fn") {
      block = utils::next_block(data, left, right);
      fname = fname_map.get_ptr(block);

    } else if (marker == "sn") {
      block = utils::next_block(data, left, right);
      sname = sname_map.get_ptr(block);

    } else if (marker == "co") {
      country = country_map.get_ptr(utils::next_block(data, left, right));

    } else if (marker == "ci") {
      city = city_map.get_ptr(utils::next_block(data, left, right));

    } else if (marker == "l") {
      like_new = new Like;
      like_new->id = utils::int_from_string(utils::next_block(data, left, right));
      like_new->ts = utils::int_from_string(utils::next_block(data, left, right));
      like_new->next = like_root;
      like_root = like_new;

    } else if (marker == "in") {
      sname = sname_map.get_ptr(utils::next_block(data, left, right));
      interest_new = new Interest;
      interest_new->name = interest_map.get_ptr(utils::next_block(data, left, right));
      interest_new->next = interest_root;
      interest_root = interest_new;

    } else {
      cout << "unexpected marker" << marker << endl;
      perror("unexpected marker");
    }
  }

  auto acc = new Account(
    id, birth, premium, sex, status, phone, email,
    fname, sname, country, city, like_root, interest_root
  );
  add_item(acc);
//  acc->repr();

}

Account* Store::get_item(int index){
  return this->data[index];
}

int
Store::count_account()
{ return (int)data.size(); }