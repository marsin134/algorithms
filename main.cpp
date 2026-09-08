#include <iostream>
#include <cstdlib>
#include <ctime>


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

int main()
{
    std::ios_base::sync_with_stdio(false);
    std::cin.tie(nullptr);
    
    clock_t start, end;
    int i=0, j=0, r;
    const int SIZE = 1000;
    static int a[SIZE][SIZE], b[SIZE][SIZE], c[SIZE][SIZE], elem_c;
    
    srand(time(NULL));
    while(i<SIZE)
    {
        j = 0;
        while(j<SIZE)
        {
            a[i][j]=rand()% 100 + 1;
            j++;
        }
        i++;
    }
    
    srand(time(NULL));
    i=0; j=0;
    while(i<SIZE)
    {
        j = 0;
        while(j<SIZE)
        {
            b[i][j]=rand()% 100 + 1;
            j++;
        }
        i++;
    }
    start = clock();
    for(i=0;i<SIZE;i++)
    {
        for(j=0;j<SIZE;j++)
        {
            elem_c=0;
            for(r=0;r<SIZE;r++)
            {
                elem_c=elem_c+a[i][r]*b[r][j];
                c[i][j]=elem_c;
            }
        }
    }
    end = clock();
    std::cout << "Size: " << SIZE << " x " << SIZE << std::endl;
    std::cout << "Time taken for matrix multiplication: " << static_cast<double>(end - start) / CLOCKS_PER_SEC << " seconds\n" << std::endl;
    return 0;
}