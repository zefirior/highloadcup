#include <iostream>
#include <unistd.h>
#include <stdlib.h>
#include <string.h>
#include "model/Account.h"
#include "Store.h"
#include "Server.h"
#include "utils.h"

using namespace std;

#define ACCOUNT_NUMBER 1300

int main() {

  /*Store acc_store {ACCOUNT_NUMBER};

  std::cout << "Fill store" << std::endl;
  utils::fill_massive(&acc_store, ACCOUNT_NUMBER);

  acc_store.get_item(1000)->repr();*/

  /*
  unlink("/tmp/server-socket.sock");
  std::cout << "Create server" << std::endl;
  Server server = Server("/tmp/server-socket.sock");
  std::cout << "run server" << std::endl;
  server.run(new Api());
   */

  string string1 = "123";

  std::cout << string1.compare("1252343") << std::endl;

  return 0;
}