#include <string>
#include <algorithm>

//Given a string S, remove the vowels 'a', 'e', 'i', 'o', and 'u' from it, and return the new string.
class Solution {
public:
    std::string removeVowels(const std::string& S) {
        std::string newString;
        newString.reserve(S.size());
        for (char i : S) {
            if ((i == 'a') || (i == 'e') || (i == 'i') || (i == 'o') || (i == 'u')) {
                continue;
            } else {
                newString += i;
            }

        }
        return newString;
    }
};