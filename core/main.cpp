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

using namespace std;

#define ACCOUNT_NUMBER 30000

int main() {

  /*Store acc_store {ACCOUNT_NUMBER};

  std::cout << "Fill store" << std::endl;
  utils::fill_massive(&acc_store, ACCOUNT_NUMBER);

  acc_store.get_item(1000)->repr();*/

  /*
  Store acc_store {ACCOUNT_NUMBER};
  auto api = Api(acc_store);
  unlink("/tmp/server-socket.sock");
  std::cout << "Create server" << std::endl;
  Server server = Server("/tmp/server-socket.sock");
  std::cout << "run server" << std::endl;
  server.run(&api);
   */

//  /*
  string data = "AADD fn asdgkajdew i 123 b 2345345 p 42 235 s 0 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 2345 678683 st 1 ph 8(912)6290012 e alegfbhhh@afdh.ru sn qweasd co country ci city";
  Store acc_store {ACCOUNT_NUMBER};
  auto api = Api(&acc_store);
  for (int i = 0; i < 1300000; i++){
    auto resp = api.dispatch(&data);
  }
  acc_store.get_item(100)->repr();
//  */

  return 0;
}