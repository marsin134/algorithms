#include <cstring>
#include <iostream>
#include <vector>

using namespace std;

struct Student {
    char family[20];
    char name[20];
    char faculty[20];
    int numZach;
};

void getStudent(const vector<Student>& students, int size, const int numZach, const char *faculty, const char *family, const char *name) {
    for (int i = 0; i < size; i++) {
        if (students[i].numZach == numZach) {
            cout << "Student for numZach " << numZach << ": " << students[i].family << " " << students[i].name << ", Faculty: " << students[i].faculty << ", NumZach: " << students[i].numZach << endl;
        }
        if (strcmp(students[i].faculty, faculty) == 0) {
            cout << "Student for faculty " << faculty << ": " << students[i].family << " " << students[i].name << ", Faculty: " << students[i].faculty << ", NumZach: " << students[i].numZach << endl;
        }
        if (strcmp(students[i].family, family) == 0) {
            cout << "Student for family " << family << ": " << students[i].family << " " << students[i].name << ", Faculty: " << students[i].faculty << ", NumZach: " << students[i].numZach << endl;
        }
        if (strcmp(students[i].name, name) == 0) {
            cout << "Student for name " << name << ": " << students[i].family << " " << students[i].name << ", Faculty: " << students[i].faculty << ", NumZach: " << students[i].numZach << endl;
        }
        }

}