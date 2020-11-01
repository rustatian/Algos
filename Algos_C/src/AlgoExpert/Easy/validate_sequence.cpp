#define CATCH_CONFIG_MAIN

#include <cstdlib>
#include <catch2/catch.hpp>
#include <boost/container/vector.hpp>

bool isValidSubsequence(boost::container::vector<int> array, boost::container::vector<int> sequence) {
    int sIndx = 0;
    for (auto &num: array) {
        if (sIndx == sequence.size())
            break;
        if (sequence[sIndx] == num)
            sIndx++;
    }
    return sIndx == sequence.size();
}

TEST_CASE("validate sequence", "[validate sequence]") {
    auto a = boost::container::vector<int>{5, 1, 22, 25, 6, -1, 8, 10};
    auto b = boost::container::vector<int>{1, 6, -1, 10};
    REQUIRE(isValidSubsequence(a, b) == true);
}