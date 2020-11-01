cmake_minimum_required(VERSION 3.17 FATAL_ERROR)
project(easy)
set(CMAKE_CXX_STANDARD 20)

find_program(CCACHE_PROGRAM ccache)
if(CCACHE_PROGRAM)
    set_property(GLOBAL PROPERTY RULE_LAUNCH_COMPILE "${CCACHE_PROGRAM}")
endif()

# TESTS
find_package(Catch2 CONFIG REQUIRED)
find_package(Boost "1.73" REQUIRED)

add_executable(tns two_number_sum.cpp)
target_link_libraries(tns Catch2::Catch2)

add_executable(valSeq validate_sequence.cpp)
target_link_libraries(valSeq Catch2::Catch2 ${Boost_LIBRARIES})

add_executable(fcvINbst find_closest_value_in_BST.cpp)
target_link_libraries(fcvINbst PRIVATE Catch2::Catch2)

add_executable(branch_sums branch_sums.cpp)
target_link_libraries(branch_sums PRIVATE Catch2::Catch2)