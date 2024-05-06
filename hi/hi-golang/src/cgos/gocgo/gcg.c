#include <stdio.h>
#include "_cgo_export.h"
int c2Go(){
    printf("C code Inner.");
    // C 中调用 go 函数
    DateString4Go();
    return 0;
}