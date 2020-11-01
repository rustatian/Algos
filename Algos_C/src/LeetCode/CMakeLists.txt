cmake_minimum_required(VERSION 3.16)
project(LeetCode)

set(CMAKE_CXX_STANDARD 17)

add_subdirectory(remove_vovels_from_string)
add_subdirectory(two_sum)

add_executable(add_two_numbers add_two_numbers.cpp)