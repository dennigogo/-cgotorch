#include "utils.h"

#include <stdio.h>
#include <string.h>

const char *exception_str(const char *e) {
    auto len = strlen(e);
    auto r = new char[len + 1];
    snprintf(r, len + 1, "%s", e);
    return r;
}

void FreeString(const char *s) {
    delete[] s;
}