// 2022年3月3日 星期四 <Java11>
// 常量
class VariableConst{
    public static final double PI = 3.1415926535898;	// 静态类常量
    final int Index = 220303; // 类成员常量
    
    public static void main(String[] args){
        // 局部常量
        final int V16 = 0xff;   // 0x/0X  16 进制
        final byte V8 = 013;     // 8进制
        final long vLong = 12L;

        // 自动推断
        var f16 = 3.51f;

        // 数据输出
        System.out.println("PI: "+ PI);
        System.out.println("V8: "+ V8);
        System.out.println("vLong: "+ vLong);
        System.out.println("f16: "+ f16);
    }
}