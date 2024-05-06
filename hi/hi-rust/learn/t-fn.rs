/**
    generics-fn(T-fn)
    trait fn 函数
**/

// 统计总数
fn num_count<T>(nums: &[T]) -> T {
    let mut all_num = &nums[0];
    let mut ci = 0;
    for &t in nums.iter(){
        ci += 1;
        if ci == 1{
            continue;
        }
        all_num = all_num + &t;
        // 有错误
    }
    *all_num
}

fn main(){
    let intVet = vec![1, 78, 41, 27, 63, 10, 12, 15, 16, 18, 19];
    println!("{}", num_count(&intVet));
}