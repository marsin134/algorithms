#include <stdio.h>
#include <stdlib.h>
#include <time.h>

void shell(int *items, int count)
{
    int i, j, gap, k;
    int x, a[5];
    
    a[0]=9; a[1]=5; a[2]=3; a[3]=2; a[4]=1;
    
    for(k=0; k < 5; k++) {
        gap = a[k];
        for(i=gap; i < count; ++i) {
            x = items[i];
            for(j=i-gap; (x < items[j]) && (j >= 0); j=j-gap)
                items[j+gap] = items[j];
            items[j+gap] = x;
        }
    }
}

void qs(int *items, int left, int right)
{
    int i, j;
    int x, y;
    
    i = left; j = right;
    x = items[(left+right)/2];
    
    do {
        while((items[i] < x) && (i < right)) i++;
        while((x < items[j]) && (j > left)) j--;
        
        if(i <= j) {
            y = items[i];
            items[i] = items[j];
            items[j] = y;
            i++; j--;
        }
    } while(i <= j);
    
    if(left < j) qs(items, left, j);
    if(i < right) qs(items, i, right);
}

int compare_int_asc(const void *a, const void *b) {
    return (*(int*)a - *(int*)b);
}

int main(void)
{
    setvbuf(stdin, NULL, _IONBF, 0);
    setvbuf(stdout, NULL, _IONBF, 0);
    
    clock_t start, end;
    int i=0, j=0, r;
    const int SIZE = 1000;
    int a[SIZE*2], b[SIZE/2];
    
    srand(time(NULL));
    while(i<SIZE)
    {
        a[i]=rand()% 100 + 1;
        b[i]=rand()% 100 + 1;
        i++;
    }
    shell(b, SIZE/2);
    for(i=0; i<SIZE/2; i++) {
        int tmp = a[i];
        a[i] = a[SIZE-i-1];
        a[SIZE-i-1] = tmp;
    }
    

    start = clock();
    qsort(a, SIZE, sizeof(int), compare_int_asc);
    end = clock();
    printf("Random | Size: %d\n", SIZE);
    printf("Time shell: %f seconds\n", (double)(end - start) / CLOCKS_PER_SEC);

    shell(a, SIZE);
    start = clock();
    qsort(a, SIZE, sizeof(int), compare_int_asc);
    end = clock();
    printf("Increasing sequence| Size: %d\n", SIZE);
    printf("Time qs: %f seconds\n", (double)(end - start) / CLOCKS_PER_SEC);\

    for (i=0; i<SIZE/2; i++) {
        int tmp = a[i];
        a[i] = a[SIZE-i-1];
        a[SIZE-i-1] = tmp;
    }
    start = clock();
    qsort(a, SIZE, sizeof(int), compare_int_asc);
    end = clock();
    printf("In descending order | Size: %d\n", SIZE);
    printf("Time qs: %f seconds\n", (double)(end - start) / CLOCKS_PER_SEC);

    shell(b, SIZE/2);
    for (i=0; i<SIZE/2; i++) {
        a[SIZE+i] = b[i];
    }
    start = clock();
    qsort(a, SIZE, sizeof(int), compare_int_asc);
    end = clock();
    printf("Mixed | Size: %d\n", SIZE);
    printf("Time qs: %f seconds\n", (double)(end - start) / CLOCKS_PER_SEC);
    return(0);
}