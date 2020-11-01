#include "main.h"
#include <iostream>

static const int N = 10000;

//0-2, 1-4, 2-5, 3-6, 0-4, 6-0, 1-3
int main() {
    int p, q, id[N];

    for (int i = 0; i < N; i++) id[i] = i;

    while (std::cin >> p >> q) {
        int t = id[p];
        if (t == id[q]) continue;

        for (int & i : id) {
            if (i == t) {
                i = id[q];
            }
        }
        std::cout << " " << p << " " << q << std::endl;
    }
}
