#define CATCH_CONFIG_MAIN

#include <cstdlib>
#include <vector>
#include "catch2/catch.hpp"

class BinaryTree {
public:
    int value;
    BinaryTree *left;
    BinaryTree *right;

    BinaryTree(int value) {
        this->value = value;
        left = nullptr;
        right = nullptr;
    }
};

class TestBinaryTree : public BinaryTree {
public:
    TestBinaryTree(int value) : BinaryTree(value) {};

    BinaryTree *insert(std::vector<int> values, int i = 0) {
        if (i >= values.size())
            return nullptr;
        std::vector<BinaryTree *> queue = {this};
        while (!queue.empty()) {
            BinaryTree *current = queue[0];
            queue.erase(queue.begin());
            if (current->left == nullptr) {
                current->left = new BinaryTree(values[i]);
                break;
            }
            queue.push_back(current->left);
            if (current->right == nullptr) {
                current->right = new BinaryTree(values[i]);
                break;
            }
            queue.push_back(current->right);
        }
        insert(values, i + 1);
        return this;
    }
};

std::vector<int> branchSums(BinaryTree *root) {
    // Write your code here.
    return {};
}

TEST_CASE("branchSums", "[branchSums]") {
    auto *tree = new TestBinaryTree(1);
    tree->insert({2, 3, 4, 5, 6, 7, 8, 9, 10});
    std::vector<int> expected = {15, 16, 18, 10, 11};
    REQUIRE(branchSums(tree) == expected);
}