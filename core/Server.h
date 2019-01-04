//
// Created by liinda on 01.01.19.
//

#ifndef CORE_SERVER_H
#define CORE_SERVER_H

#include <string.h>

using namespace std;

class Api {
  public:
    string dispatch(string data);
};

typedef struct sockaddr_un Address;

class Server {
    string unix_sock_file;
    int listen_sockfd;
    Address* address;
    socklen_t addrlen;
  public:
    Server(string unix_sock_file);
    int accept_connection();
    void run(Api* api);
    string read_massage(int sock);
};


#endif //CORE_SERVER_H
