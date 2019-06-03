//
// Created by Valery Piashchynski on 2019-06-02.
//

#include <stdbool.h>
#include "main.h"
#include "stdio.h"
#include "mm_malloc.h"
#include "string.h"
//Trolls are attacking your comment section!
//A common way to deal with this situation is to remove all of the vowels from the trolls' comments, neutralizing the threat.
//Your task is to write a function that takes a string and return a new string with all vowels removed.
//For example, the string "This website is for losers LOL!" would become "Ths wbst s fr lsrs LL!".
//Note: for this kata y isn't considered a vowel.

bool is_vowel(char c) {
    c |= 'A' ^ 'a';
    return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u';
}


//solution must allocate all required memory
//and return a free-able buffer to the caller.
char *disemvowel(const char *str) {
    // strlen don't encounter a 0-character
    const size_t len = strlen(str);

    // to have a properly size we have to add 1 (include 0 character) to res
    char *const res = malloc(len + 1);
    size_t i_dst = 0;

    for (size_t i_src = 0; i_src < len; i_src++)
        if (!is_vowel(str[i_src]))
            res[i_dst++] = str[i_src];

    res[i_dst] = '\0';
    return res;
}


int main() {
    // "Ths wbst s fr lsrs LL!"
    printf("result is: %s", disemvowel("This website is for losers LOL!"));
}

