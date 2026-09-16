#include <stdio.h>
#include <stdlib.h>
#include <time.h>

void shell(int *items, int count) 
{
    int i, j, gap, k;
    int x, a[5];
    
    a[0] = 9;
    a[1] = 5;
    a[2] = 3;
    a[3] = 2;
    a[4] = 1;
    
    for (k = 0; k < 5; k++) 
    {
        gap = a[k];
        
        for (i = gap; i < count; ++i) 
        {
            x = items[i];
            
            for (j = i - gap; (x < items[j]) && (j >= 0); j = j - gap) 
            {
                items[j + gap] = items[j];
            }
            
            items[j + gap] = x;
        }
    }
}

void qs(int *items, int left, int right) 
{
    int i, j;
    int x, y;
    
    i = left;
    j = right;
    x = items[(left + right) / 2];
    
    do 
    {
        while ((items[i] < x) && (i < right)) 
        {
            i++;
        }
        
        while ((x < items[j]) && (j > left)) 
        {
            j--;
        }
        
        if (i <= j) 
        {
            y = items[i];
            items[i] = items[j];
            items[j] = y;
            i++;
            j--;
        }
    } while (i <= j);
    
    if (left < j) 
    {
        qs(items, left, j);
    }
    
    if (i < right) 
    {
        qs(items, i, right);
    }
}

int compare_int_asc(const void *a, const void *b) 
{
    return (*(int*)a - *(int*)b);
}

int main(void)
{
    setvbuf(stdin, NULL, _IONBF, 0);
    setvbuf(stdout, NULL, _IONBF, 0);

    clock_t start, end;

    int i, randomNumber;

    const int SIZE = 100000;
    int min = -15, max = 20;

    int arr1Shell[SIZE], arr2Shell[SIZE], arr3Shell[SIZE], arr4Shell[SIZE];
    int arr1SQs[SIZE], arr2SQs[SIZE], arr3SQs[SIZE], arr4SQs[SIZE];
    int arr1Qsort[SIZE], arr2Qsort[SIZE], arr3Qsort[SIZE], arr4Qsort[SIZE];

    double shellTime[4];
    double qsTime[4];
    double qsortTime[4];

    srand(time(NULL));

    for (i = 0; i < SIZE; i++) {
        randomNumber = rand() % (max - min + 1) + min;

        arr1Shell[i] = randomNumber;
        arr1SQs[i] = randomNumber;
        arr1Qsort[i] = randomNumber;

        arr2Shell[i] = i;
        arr2SQs[i] = i;
        arr2Qsort[i] = i;

        arr3Shell[i] = SIZE - i;
        arr3SQs[i] = SIZE - i;
        arr3Qsort[i] = SIZE - i;

        if (i < SIZE / 2) {
            arr4Shell[i] = i;
            arr4SQs[i] = i;
            arr4Qsort[i] = i;
        }
        else {
            arr4Shell[i] = SIZE - i;
            arr4SQs[i] = SIZE - i;
            arr4Qsort[i] = SIZE - i;
        }
    }

    start = clock();
    shell(arr1Shell, SIZE);
    end = clock();
    shellTime[0] = (double)(end - start) / CLOCKS_PER_SEC;

    start = clock();
    qs(arr1SQs, 0, SIZE - 1);
    end = clock();
    qsTime[0] = (double)(end - start) / CLOCKS_PER_SEC;

    start = clock();
    qsort(arr1Qsort, SIZE, sizeof(int), compare_int_asc);
    end = clock();
    qsortTime[0] = (double)(end - start) / CLOCKS_PER_SEC;

    start = clock();
    shell(arr2Shell, SIZE);
    end = clock();
    shellTime[1] = (double)(end - start) / CLOCKS_PER_SEC;

    start = clock();
    qs(arr2SQs, 0, SIZE - 1);
    end = clock();
    qsTime[1] = (double)(end - start) / CLOCKS_PER_SEC;

    start = clock();
    qsort(arr2Qsort, SIZE, sizeof(int), compare_int_asc);
    end = clock();
    qsortTime[1] = (double)(end - start) / CLOCKS_PER_SEC;

    start = clock();
    shell(arr3Shell, SIZE);
    end = clock();
    shellTime[2] = (double)(end - start) / CLOCKS_PER_SEC;

    start = clock();
    qs(arr3SQs, 0, SIZE - 1);
    end = clock();
    qsTime[2] = (double)(end - start) / CLOCKS_PER_SEC;

    start = clock();
    qsort(arr3Qsort, SIZE, sizeof(int), compare_int_asc);
    end = clock();
    qsortTime[2] = (double)(end - start) / CLOCKS_PER_SEC;

    start = clock();
    shell(arr4Shell, SIZE);
    end = clock();
    shellTime[3] = (double)(end - start) / CLOCKS_PER_SEC;

    start = clock();
    qs(arr4SQs, 0, SIZE - 1);
    end = clock();
    qsTime[3] = (double)(end - start) / CLOCKS_PER_SEC;

    start = clock();
    qsort(arr4Qsort, SIZE, sizeof(int), compare_int_asc);
    end = clock();
    qsortTime[3] = (double)(end - start) / CLOCKS_PER_SEC;

    printf("\n");
    printf("============================================================================\n");
    printf("            Сравнение алгоритмов сортировки (N = %d)\n", SIZE);
    printf("============================================================================\n");

    printf("%-62s | %8s | %8s | %7s\n",
           "Тип последовательности",
           "Shell",
           "QS",
           "qsort");

    printf("----------------------------------------------------------------------------\n");

    printf("%-58s | %8.6f | %8.6f | %7.6f\n",
           "Случайные значения",
           shellTime[0],
           qsTime[0],
           qsortTime[0]);

    printf("%-71s | %8.6f | %8.6f | %7.6f\n",
           "Возрастающая последовательность",
           shellTime[1],
           qsTime[1],
           qsortTime[1]);

    printf("%-68s | %8.6f | %8.6f | %7.6f\n",
           "Убывающая последовательность",
           shellTime[2],
           qsTime[2],
           qsortTime[2]);

    printf("%-38s | %6.6f | %6.6f | %6.6f\n",
           "Возрастающая-убывающая последовательность",
           shellTime[3],
           qsTime[3],
           qsortTime[3]);

    printf("============================================================================\n");

    return 0;
}