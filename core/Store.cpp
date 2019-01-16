//
// Created by liinda on 31.12.18.
//

#include "Store.h"
#include "utils.h"
#include "model/Account.h"
using namespace std;

Store::Store(int reserve_size) {
  this->data.reserve(reserve_size);
}

void Store::add_item(Account* item) {
  this->data.push_back(item);
}

string* Store::get_ptr_from_map(map<string, string *> *container, string data) {
  if (!container->count(data)){
    (*container)[data] = new string(data);
  }
  return container->find(data)->second;
}

void Store::parse_account(string data) {
  size_t left=0, right=0;
  int id, birth;
  char sex, status;
  string
    marker,
    phone = "",
    email,
    *fname = nullptr,
    *sname = nullptr,
    *country = nullptr,
    *city = nullptr;
  Premium *premium = nullptr;
  Like *like_root = nullptr, *like_new = nullptr;

  while (right < string::npos) {
    marker = utils::next_block(data, left, right);
    if (marker.compare("id") == 0){
      id = utils::int_from_string(utils::next_block(data, left, right));

    } else if (marker.compare("b") == 0) {
      birth = utils::int_from_string(utils::next_block(data, left, right));

    } else if (marker.compare("p") == 0) {
      premium = new Premium;
      premium->premium_from = utils::int_from_string(utils::next_block(data, left, right));
      premium->premium_to = utils::int_from_string(utils::next_block(data, left, right));

    } else if (marker.compare("s") == 0) {
      sex = utils::next_block(data, left, right)[0];

    } else if (marker.compare("st") == 0) {
      status = utils::next_block(data, left, right)[0];

    } else if (marker.compare("ph") == 0) {
      phone = utils::next_block(data, left, right);

    } else if (marker.compare("e") == 0) {
      email = utils::next_block(data, left, right);

    } else if (marker.compare("fn") == 0) {
      fname = get_ptr_from_map(&fname_map, utils::next_block(data, left, right));

    } else if (marker.compare("sn") == 0) {
      sname = get_ptr_from_map(&sname_map, utils::next_block(data, left, right));

    } else if (marker.compare("co") == 0) {
      country = get_ptr_from_map(&country_map, utils::next_block(data, left, right));

    } else if (marker.compare("ci") == 0) {
      city = get_ptr_from_map(&city_map, utils::next_block(data, left, right));

    } else if (marker == "l") {
      like_new = new Like;
      like_new->id = utils::int_from_string(utils::next_block(data, left, right));
      like_new->ts = utils::int_from_string(utils::next_block(data, left, right));
      like_new->next = like_root;
      like_root = like_new;

    } else {
      cout << "unexpected marker" << marker << endl;
      perror("unexpected marker");
    }
  }

  auto acc = new Account(
    id, birth, premium, sex, status, phone, email,
    fname, sname, country, city, like_root
  );
  add_item(acc);

}

Account* Store::get_item(int index){
  return this->data[index];
}
