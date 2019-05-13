//cgo 测试
#include <stdio.h>
#include "base.h"

int count = 1;

/**
 * int
**/
int add(){
    count += 1;
    return count;
}