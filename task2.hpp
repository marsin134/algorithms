#include <cstdlib>

void generateRandomArray(int *arr, int size, int minValue, int maxValue) {
    for (int i = 0; i < size; ++i) {
        arr[i] = minValue + rand() % (maxValue - minValue + 1);
    }
}