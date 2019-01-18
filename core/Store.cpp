//
// Created by liinda on 31.12.18.
//

#include "Store.h"
#include "utils.h"
#include "model/Account.h"
using namespace std;

Store::Store(unsigned long reserve_size) {
  this->data.reserve(reserve_size);
  this->city_map =    new map <size_t, string*>;
  this->country_map = new map <size_t, string*>;
  this->fname_map =   new map <size_t, string*>;
  this->sname_map =   new map <size_t, string*>;
}

void Store::add_item(Account* item) {
  this->data.push_back(item);
}

string* Store::get_ptr_from_map(map<size_t, string *> *container, string const &data) {
  size_t const block_hash = hash<string>()(data);
  if (!container->count(block_hash)){
    (*container)[block_hash] = new string(data);
  }
  return container->find(block_hash)->second;
}

void Store::parse_account(string data) {
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

  while (right < string::npos) {
    marker = utils::next_block(data, left, right);
    if (marker == "id"){
      id = utils::int_from_string(utils::next_block(data, left, right));

    } else if (marker == "b") {
      birth = utils::int_from_string(utils::next_block(data, left, right));

    } else if (marker == "p") {
      premium = new Premium;
      premium->premium_from = utils::int_from_string(utils::next_block(data, left, right));
      premium->premium_to = utils::int_from_string(utils::next_block(data, left, right));

    } else if (marker == "s") {
      sex = utils::next_block(data, left, right)[0];

    } else if (marker == "st") {
      status = utils::next_block(data, left, right)[0];

    } else if (marker == "ph") {
      phone = utils::next_block(data, left, right);

    } else if (marker == "e") {
      email = utils::next_block(data, left, right);

    } else if (marker == "fn") {
//      block = utils::next_block(data, left, right);
//
//      if (!fname_map->count(hash<string>()(block))){
//        (*fname_map)[hash<string>()(block)] = new string(block);
//      }
//      fname = fname_map->find(hash<string>()(block))->second;
        block = utils::next_block(data, left, right);

      fname = get_ptr_from_map(fname_map, block);

    } else if (marker == "sn") {
//      block = utils::next_block(data, left, right);
//
//      if (!sname_map->count(block)){
//        (*sname_map)[block] = new string(block);
//      }
//      sname = sname_map->find(block)->second;

      sname = get_ptr_from_map(sname_map, utils::next_block(data, left, right));

    } else if (marker == "co") {
//      block = utils::next_block(data, left, right);
//
//      if (!country_map->count(block)){
//        (*country_map)[block] = new string(block);
//      }
//      country = country_map->find(block)->second;
//
      country = get_ptr_from_map(country_map, utils::next_block(data, left, right));

    } else if (marker == "ci") {
//      block = utils::next_block(data, left, right);
//
//      if (!fname_map->count(block)){
//        (*fname_map)[block] = new string(block);
//      }
//      city = fname_map->find(block)->second;
//
      city = get_ptr_from_map(city_map, utils::next_block(data, left, right));

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
