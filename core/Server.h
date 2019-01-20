//
// Created by liinda on 01.01.19.
//

#ifndef CORE_SERVER_H
#define CORE_SERVER_H

#include <sys/socket.h>
#include <netinet/in.h>
#include <sys/un.h>
#include <string.h>
#include "Api.h"

using namespace std;

class Server {
    string unix_sock_file;
    int listen_sockfd;
    sockaddr_un* address;
    socklen_t addrlen;
    Api* _api;
  public:
    explicit Server(string unix_sock_file);
    int accept_connection();
    void set_api(Api *api);
    void run();
    string* read_massage(int sock);
    bool send_massage(int sock, string message);
};

#endif //CORE_SERVER_H
