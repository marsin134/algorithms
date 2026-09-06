#include <iostream>

using namespace std;


void sumsLineOrColumn(int **matrix, int rows, int cols, bool sumRows) {
    if (sumRows) {
        for (int i = 0; i < rows; ++i) {
            int sum = 0;
            for (int j = 0; j < cols; ++j) {
                sum += matrix[i][j];
            }
            cout << "Sum of row " << i << ": " << sum << endl;
        }
    } else {
        for (int j = 0; j < cols; ++j) {
            int sum = 0;
            for (int i = 0; i < rows; ++i) {
                sum += matrix[i][j];
            }
            cout << "Sum of column " << j << ": " << sum << endl;
        }
    }
}