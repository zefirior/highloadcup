//
// Created by liinda on 05.01.19.
//

#include "Api.h"
#include "utils.h"
using namespace std;

Api::Api(Store* store) : store(store){};

string Api::dispatch(string* data) {
//  std::cout << "dispatch" << std::endl;
  string method = "";
  size_t pos;

  if ((pos = data->find(" ")) == string::npos){
    return "No method";
  }
  method = data->substr(0, pos);
//  std::cout << "method: " << method << std::endl;

  if (method.compare("AADD") == 0){
    return add_account(data->substr(pos+1));
  } else if (method.compare("SYNC") == 0){
    return "SYNC method";
  }
  return "No method";
}

string Api::add_account(string data) {
//  std::cout << data << std::endl;

  store->parse_account(data);

  return "ok";
}