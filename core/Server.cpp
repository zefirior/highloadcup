//
// Created by liinda on 01.01.19.
//

#include <unistd.h>
#include <iostream>
//#include <stdio.h>
//#include <stdlib.h>
#include <sys/socket.h>
#include <sys/un.h>
#include <netinet/in.h>
#include "Server.h"

#define SOCKET_CONNECTION_NUM 1
#define BUFFER_SIZE 1024
#define LENGTH_SIZE 9
#define METHOD_SIZE 4

string Api::dispatch(string data) {
  string method = "";

  if (data.size() < METHOD_SIZE) {
    return "format error";
  }

  method.append(data, METHOD_SIZE, 0);
  if (method.compare("PING") == 0){
    return "PING method";
  } else if (method.compare("SYNC") == 0){
    return "PING method";
  }
  return "No method";
}

Server::Server(string unix_sock_file) : unix_sock_file(unix_sock_file) {
  socklen_t addrlen = sizeof(*(this->address));

  // create socket
  this->listen_sockfd = socket(AF_UNIX, SOCK_STREAM, 0);
  if (!this->listen_sockfd) {
    perror("socket fail");
    exit(EXIT_FAILURE);
  }
  // bind socket
  bzero(&(this->address), this->addrlen);
  this->address->sun_family = AF_UNIX;
  this->unix_sock_file.copy(this->address->sun_path, this->unix_sock_file.size(), 0);

  if (bind(this->listen_sockfd, (struct sockaddr *) this->address, this->addrlen) < 0) {
    perror("bind fail");
    exit(EXIT_FAILURE);
  }

  if (listen(this->listen_sockfd, SOCKET_CONNECTION_NUM) < 0) {
    perror("listen fail");
    exit(EXIT_FAILURE);
  }
}

int Server::accept_connection() {
  return accept(this->listen_sockfd, (struct sockaddr *) this->address, &(this->addrlen));
}

string Server::read_massage(int sock) {
  char length[LENGTH_SIZE+1] = "",
    buffer[BUFFER_SIZE+1] = "";
  int ilength, nread;
  string message;

  nread = (int)read(sock, buffer, BUFFER_SIZE);
  if (nread <= LENGTH_SIZE) {
    return "";
  }

  strncpy(length, buffer, LENGTH_SIZE);
  ilength = atoi(length);

  message.append(buffer + LENGTH_SIZE, nread - LENGTH_SIZE);
  ilength -= (nread - LENGTH_SIZE);

  while (ilength > 0){
    nread = (int)read(sock, buffer, BUFFER_SIZE);
    if (nread <= 0) {
      return "";
    }

    message.append(buffer, nread);
    ilength -= nread;
  }

  std::cout << "length: " << length << std::endl;
  std::cout << "ilength: " << ilength << std::endl;
  std::cout << "message: " << message << std::endl;
  return message;
}

void Server::run(Api * api) {
  int connection;
  string request = "", response = "";

  while (true) {
    if ((connection = accept_connection()) < 0) {
      perror("accept fail");
      exit(EXIT_FAILURE);
    }
    std::cout << "accept connection" << std::endl;

    while (true) {
      if ((request = this->read_massage(connection)).empty()) {
        std::cout << "connection is closed" << std::endl;
        close(connection);
        break;
      }

      response = api->dispatch(request);

      
    }
  }
}
