//
// Created by liinda on 05.01.19.
//

#include "Api.h"
#include "utils.h"
using namespace std;

Api::Api(Store *store) : store(store){
//  cout << "Api::Api store" << store << endl;
//  cout << "Api::Api this->store" << this->store << endl;
};

string Api::dispatch(const string &data) {
  string method;
  string params;
  size_t pos;

  if ((pos = data.find(' ')) == string::npos){
    return "No method";
  }
  method = data.substr(0, pos);
//  std::cout << "method: " << method << std::endl;

  if (method == "AADD"){
    params = data.substr(pos+1);
    return add_account(params);

  } else if (method == "CACC"){ // count account
    return to_string(store->count_account());

  } else if (method == "PLST"){ // print last account
    store->get_item(store->count_account() - 1)->repr();
    return "ok";

  } else if (method == "SYNC"){
    return "SYNC method";
  }
  return "No method";
}

string Api::add_account(string &data) {

  store->parse_account(data);

  return "ok";
}