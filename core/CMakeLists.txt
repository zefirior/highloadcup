cmake_minimum_required(VERSION 3.13)
project(core)

set(CMAKE_CXX_STANDARD 14)

add_executable(core main.cpp model/Account.h model/Account.cpp
        Store.cpp Store.h utils.cpp utils.h Server.cpp Server.h Api.cpp Api.h)