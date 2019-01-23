#define DEBUG_MODE

#include <iostream>
#include <bitset>
#include <unistd.h>
#include <stdlib.h>
#include <string.h>
#include "model/Account.h"
#include "Store.h"
#include "Server.h"
#include "utils.h"
#define ACCOUNT_NUMBER 130000
#define SOCK_FILE "/tmp/server-socket.sock"

using namespace std;

int main() {


//  /*
  Store acc_store (ACCOUNT_NUMBER);
  Api api = Api(&acc_store);
  unlink(SOCK_FILE);
  std::cout << "Create server" << std::endl;
  Server server = Server(SOCK_FILE);
  std::cout << "run server" << std::endl;
  server.set_api(&api);
  server.run();
//   */

  /*
  string data = "AADD fn asdgkajdew id 123 b 2345345 p 42 235 s 0 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 134 65473 l 2345 678683 st 1 ph 8(912)6290012 e alegfbhhh@afdh.ru sn qweasd co country ci city";
  Store acc_store (ACCOUNT_NUMBER);
  Api api = Api(&acc_store);
  for (int i = 0; i < ACCOUNT_NUMBER; i++){
    auto resp = api.dispatch(data);
  }
  acc_store.get_item(100)->repr();
  */

  vector<Like> cont;
  cont.reserve(5);
  cont[1].id = 12;
  cont[1].ts = 1242;

  cout << 1 << endl;
  for (Like l : cont){
    cout << "id = " << l.id << ", ts = " << l.ts << endl;
  }

  return 0;
}