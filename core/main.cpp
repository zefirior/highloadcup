#define DEBUG_MODE

#include <iostream>
#include <map>
#include <unistd.h>
#include <stdlib.h>
#include <string.h>
#include "model/Account.h"
#include "Store.h"
#include "Server.h"
#include "utils.h"
#define ACCOUNT_NUMBER 1300000
#define SOCK_FILE "/tmp/server-socket.sock"

using namespace std;

int main() {


  /*
  Store acc_store (ACCOUNT_NUMBER);
  Api api = Api(acc_store);
  unlink(SOCK_FILE);
  std::cout << "Create server" << std::endl;
  Server server = Server(SOCK_FILE);
  std::cout << "run server" << std::endl;
  server.run(api);
   */

//  /*
  string data = "AADD fn asdgkajdew id 123 b 2345345 p 42 235 s 0 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 2345 678683 st 1 ph 8(912)6290012 e alegfbhhh@afdh.ru sn qweasd co country ci city";
  Store acc_store (ACCOUNT_NUMBER);
  auto api = Api(acc_store);
  for (int i = 0; i < ACCOUNT_NUMBER; i++){
    auto resp = api.dispatch(data);
  }
  acc_store.get_item(100)->repr();
//  */

  string src, dst;
  size_t hash_res;
  src = "Source row";
  dst = src.substr(2, 4);
  hash_res = hash<string>()(src);
  cout << hash_res << endl;

  cout << dst << endl;

  return 0;
}