//
// Created by Valery Piashchynski on 2019-06-16.
//

//Write a function called repeatStr which repeats the given string string exactly n times.
//
//repeatStr(6, "I") // "IIIIII"
//repeatStr(5, "Hello") // "HelloHelloHelloHelloHello"

#include "main.h"
#include "stdlib.h"
#include "string.h"
#include "stdio.h"

char *repeat_str(size_t count, const char *src) {
    int length = strlen(src);
    char *dest = malloc(count * length * sizeof(char));
    for (int i = 0; i < count; i++) {
        strcpy(dest + i * length, src);
    }
    return dest;
}

int main() {
    size_t count = 10;
    const char *str = "hello ";
    char *result = repeat_str(count, str);
    printf("result is: %s", result);
    free(result);
}