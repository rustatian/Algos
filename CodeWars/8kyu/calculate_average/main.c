//
// Created by Valery Piashchynski
//
#include "main.h"
#include "stdio.h"


double find_average(double *array, int length) {
    double sum = 0;
    for (int i = 0; i < length; ++i) {
        sum += array[i];
    }

    return sum/length;
}

int main() {
    double array[] = {17, 16, 16, 16, 16, 15, 17, 17, 15, 5, 17, 17, 16};

    double a = 0;

    a = find_average(array, sizeof(array) / sizeof(array[0]));
    printf("after %lf\n", a);

}