#define CATCH_CONFIG_MAIN

#include <vector>
#include "catch2/catch.hpp"

class BST {
public:
    int value;
    BST *left;
    BST *right;

    BST(int val);

    BST &insert(int val);
};

BST::BST(int val) {
    value = val;
    left = nullptr;
    right = nullptr;
}

void findClosestValueInBstHelper(BST *tree, int target, int *closest) {
    if (tree == nullptr) {
        return;
    }

    if (tree->value == target) {
        return;
    }


    if (std::abs(target - *closest) > std::abs(target - tree->value)) {
        *closest = tree->value;
    }

    if (target < tree->value) {
        return findClosestValueInBstHelper(tree->left, target, closest);
    } else if (target > tree->value) {
        return findClosestValueInBstHelper(tree->right, target, closest);
    }
}


int findClosestValueInBst(BST *tree, int target) {
    // find in root
    if (tree->value == target) {
        return target;
    }

    int closest = INFINITY;
    findClosestValueInBstHelper(tree, target, &closest);

    return closest;
}


TEST_CASE("findClosestValueInBst", "[findClosestValueInBst]") {
    BST *root = new BST(10);
    root->left = new BST(5);
    root->left->left = new BST(2);
    root->left->left->left = new BST(1);
    root->left->right = new BST(5);
    root->right = new BST(15);
    root->right->left = new BST(13);
    root->right->left->right = new BST(14);
    root->right->right = new BST(22);
    int expected = 13;
    int actual = findClosestValueInBst(root, 12);
    REQUIRE(actual == expected);
}