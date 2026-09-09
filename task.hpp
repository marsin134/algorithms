int sumDiagonal(int **matrix, int size) {
    int sum = 0;
    for (int i = 1; i < size; i++) {
        for (int j = 0; j < i; j++) {
            sum += matrix[j][i];
        }
    }
    return sum;
}