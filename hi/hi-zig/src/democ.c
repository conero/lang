// 测试脚本
// 2025年4月11日

#include <stdio.h>
void jc_greet(){
    printf("demo use include a other file!\n");
    printf("\n");
    printf("中文脚本： jc_greet \n");

}

void print_cli(int argc, char *argv[]){
    for (int i = 0; i < argc; i++){
        printf("%d: %s\n", i+1, argv[i]);
    }
}