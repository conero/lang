#include <stdio.h>
#include "_cgo_export.h"
int helloFromC() {
    printf("Hi from C\n");
    //call Go function
    HelloFromGo();
    return 0;
}