#include <iostream>
#include "task1.hpp"
#include "task2.hpp"
#include "task4.hpp"
#include "task5.hpp"

using namespace std;

void printArray(int *arr, int size) {
    for (int i = 0; i < size; ++i) {
        cout << arr[i] << " ";
    }
    cout << endl;
}

int main() {
    Student students[3] = {
        {"Smith", "John", "Engineering", 12345},
        {"Doe", "Jane", "Science", 67890},
        {"Brown", "Mike", "Arts", 54321}
    };
    // task 3
    cout << "Enter the size of the array: ";
    int size;
    cin >> size;

    int *arr = new int[size];

    generateRandomArray(arr, size, 1, 100);
    cout << "Generated array: ";
    printArray(arr, size);

    cout << "Maximum value: " << max(arr, size) << endl;
    cout << "Minimum value: " << min(arr, size) << endl;

    sumsLineOrColumn(&arr, 1, size, true); 

    cout << "Enter student details to search (numZach, faculty, family, name): ";
    int numZach;
    char faculty[20], family[20], name[20];
    cin >> numZach >> faculty >> family >> name;

    Student *foundStudent = getStudent(students, 3, numZach, faculty, family, name);
    if (foundStudent) {
        cout << "Student found: " << foundStudent->family << " " << foundStudent->name << ", Faculty: " << foundStudent->faculty << ", NumZach: " << foundStudent->numZach << endl;
    } else {
        cout << "Student not found." << endl;
    }

    delete[] arr;
    return 0;
}