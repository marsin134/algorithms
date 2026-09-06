#include <cstring>

struct Student {
    char family[20];
    char name[20];
    char faculty[20];
    int numZach;
};

Student* getStudent(Student *students, int size, const int numZach, const char *faculty, const char *family, const char *name) {
    for (int i = 0; i < size; i++) {
        if (students[i].numZach == numZach && strcmp(students[i].faculty, faculty) == 0 &&
            strcmp(students[i].family, family) == 0 && strcmp(students[i].name, name) == 0) {
            return &students[i];
        }
    }
    return nullptr;
}