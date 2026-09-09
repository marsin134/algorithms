#include <iostream>
#include "task1.hpp"
#include "task2.hpp"
#include "task4.hpp"
#include "task5.hpp"
#include "task.hpp"
#include <vector>

using namespace std;

void printArray(int *arr, int size) {
    for (int i = 0; i < size; ++i) {
        cout << arr[i] << " ";
    }
    cout << endl;
}

int main() {
    // Student students[3] = {
    //     {"Smith", "John", "Engineering", 12345},
    //     {"Doe", "Jane", "Science", 67890},
    //     {"Brown", "Mike", "Arts", 54321}
    // };
    // task 3
    // cout << "Enter the rows and cols of the array: ";
    // int rows, cols;
    // cin >> rows >> cols;

    // cout << "Generated matrix: " << endl;
    // int **arr = new int*[rows];
    // for (int i = 0; i < rows; ++i) {
    //     arr[i] = new int[cols];
    //     generateRandomArray(arr[i], cols, -40, 80);
    //     printArray(arr[i], cols);
    // }

    // cout << "Sum of diagonal elements: " << sumDiagonal(arr, rows) << endl;

    // cout << "Maximum value: " << max(arr, size) << endl;
    // cout << "Minimum value: " << min(arr, size) << endl;

    // sumsLineOrColumn(&arr, 1, size, true); 

    // cout << "Enter student details to search (numZach, faculty, family, name): ";
    // int numZach;
    // char faculty[20], family[20], name[20];
    // cin >> numZach >> faculty >> family >> name;

    // Student *foundStudent = getStudent(students, 3, numZach, faculty, family, name);
    // if (foundStudent) {
    //     cout << "Student found: " << foundStudent->family << " " << foundStudent->name << ", Faculty: " << foundStudent->faculty << ", NumZach: " << foundStudent->numZach << endl;
    // } else {
    //     cout << "Student not found." << endl;
    // }

    vector<Student> students;
    string line;
    cout << "Enter student details (family name faculty numZach) or '#' to finish:" << endl;
    while (getline(cin, line)) {
        if (line == "#") break;
        Student student;
        sscanf(line.c_str(), "%s %s %s %d", student.family, student.name, student.faculty, &student.numZach);
        students.push_back(student);
    }
    for (const auto& student : students) {
        cout << "Student: " << student.family << " " << student.name << ", Faculty: " << student.faculty << ", NumZach: " << student.numZach << endl;
    }
    cout << students.size() << " students entered." << endl;
    Student student;
    cout << "Enter student details to search (family, name, faculty, numZach): ";
    cin >> student.family >> student.name >> student.faculty >> student.numZach;
    getStudent(students, students.size(), student.numZach, student.faculty, student.family, student.name);
    //  for (int i = 0; i < rows; ++i) {
    //     delete[] arr[i];
    // }
    // delete[] arr;
    return 0;
}