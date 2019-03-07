//cgo 测试
#include <stdio.h>
#include "base.h"

int count = 1;

int add(){
    count += 1;
    return count;
}