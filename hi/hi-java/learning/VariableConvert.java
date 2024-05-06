// 2022年3月3日 星期四 <Java11>
// 类型转换
class VariableConvert{
    public static void main(String[] args){
        byte bt1 = 22;
        short s1 = 2000;
        System.out.println("byte -> short: "+ (s1 + bt1));    

        // 强制转换
        bt1 = (byte) s1;
        System.out.println("short -> byte: "+ bt1);        
    }
}