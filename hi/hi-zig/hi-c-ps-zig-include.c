// 2025年4月10日
// 通过 zig 来编译 c 语言代码
// 
// 编译代码：  
// zig cc hi-c-ps-zig-include.c src/democ.c -Iinclude -o out/hi-c-ps-zig-include.exe
// 
// 使用 -Os/-Oz 减少文件编译大小。-Os  优化以减小生成文件的大小；-Oz   进一步优化以最小化文件大小（Clang 特有的选项）。
// zig cc -Os hi-c-ps-zig-include.c src/democ.c -Iinclude -o out/hi-c-ps-zig-include.exe
// 
// Windows 中文乱码: [Console]::OutputEncoding = [System.Text.Encoding]::UTF8
#include <stdio.h>
#include "cinclude/democ.h"

int main (int argc, char *argv[]) {
    printf("hello world\n");
    printf(" 💗 贵阳 中国.\n");
    printf(" 💗 使用用 zig 编译的c语言代码\n");
    
    // 文件调用
    jc_greet();
    print_cli(argc, argv);
    return 0;
}