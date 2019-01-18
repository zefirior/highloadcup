//
// Created by liinda on 01.01.19.
//

#include <unistd.h>
#include <iostream>
//#include <stdio.h>
//#include <stdlib.h>
#include "Server.h"

#define SOCKET_CONNECTION_NUM 1
#define BUFFER_SIZE 1024

//#define DEBUG_MODE


Server::Server(string unix_sock_file) : unix_sock_file(unix_sock_file) {
  sockaddr_un address;
  socklen_t addrlen = sizeof(address);
  this->address = &address;
  this->addrlen = addrlen;

  // create socket
  this->listen_sockfd = socket(AF_UNIX, SOCK_STREAM, 0);
  if (!this->listen_sockfd) {
    perror("socket fail");
    exit(EXIT_FAILURE);
  }
  // bind socket
  bzero(this->address, this->addrlen);
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

string* Server::read_massage(int sock) {
  char buffer[BUFFER_SIZE+1] = "";
  int ilength, nread, pos;
  string* message = nullptr;
  string* length = nullptr;

  nread = (int)read(sock, buffer, BUFFER_SIZE);
  if (nread <= 0){
    return nullptr;
  }

  pos = strpbrk(buffer, " ") - buffer;
  length = new string(buffer, pos);
  ilength = atoi(length->c_str());

  if (ilength == 0){
    return nullptr;
  }

  message = new string(buffer + pos + 1, nread - pos - 1); //append(buffer + pos + 1, nread - pos - 1);
  ilength -= (nread - (pos + 1));

  while (ilength > 0){
    nread = (int)read(sock, buffer, BUFFER_SIZE);
    if (nread <= 0) {
      return nullptr;
    }

    message->append(buffer, nread);
    ilength -= nread;
#ifdef DEBUG_MODE
    std::cout << "message: " << *message << std::endl;
    std::cout << "ilength: " << ilength << std::endl;
#endif
  }

#ifdef DEBUG_MODE
  std::cout << "length: " << *length << std::endl;
  std::cout << "ilength: " << ilength << std::endl;
  std::cout << "message: " << *message << std::endl;
#endif
  delete length;
  return message;
}

bool Server::send_massage(int sock, string message) {
  string chlength = "";
  int nsend, length = message.length();

#ifdef DEBUG_MODE
  std::cout << "message: " << message << std::endl;
#endif
  chlength = to_string(length);
  chlength.push_back(' ');
  nsend = send(sock, chlength.c_str(), chlength.length(), 0);

  if (nsend < chlength.length()){
    perror("data send error");
    return false;
  }

  const char * char_message = message.c_str();

  while (length){
    nsend = send(sock, char_message, length, 0);
    if (nsend <= 0){
      perror("data send error");
      return false;
    }
    length -= nsend;
    char_message += nsend;
  }
  return true;

}

void Server::run(Api &api) {
  int connection;
  string const *request;
  string response;

  while (true) {
    if ((connection = accept_connection()) < 0) {
      perror("accept fail");
      exit(EXIT_FAILURE);
    }
    std::cout << "accept connection" << std::endl;

    while (true) {
      if ((request = this->read_massage(connection)) == nullptr) {
        std::cout << "connection is closed" << std::endl;
        close(connection);
        delete request;
        break;
      }
      std::cout << "read " << *request << std::endl;
      response = api.dispatch(*request);
      delete request;
      std::cout << "dispatch " << response << std::endl;

      if (!send_massage(connection, response)){
        std::cout << "connection is closed" << std::endl;
        close(connection);
        break;
      }
    }
  }
}
