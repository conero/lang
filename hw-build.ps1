# 2018年11月13日 星期二
# 静态语言编辑处理

$tardir = './test-hw'
rmdir -r $tardir
mkdir $tardir
cd $tardir


# gcc -> c
gcc ../hi-c-cpp/hw.c -o hwc


# g++ -> cpp
g++ ../hi-c-cpp/hw.cpp -o hwcpp


# go
go build ../hi-golang/test/hw.go
# rename ./hw.exe ./hwgo.exe
$goexe = 'hw.exe'
copy $goexe hwgo.exe
unlink $goexe


# rust
rustc ../hi-rust/hw.rs -o hwrust.exe

# kotlin-native
kotlinc ../hi-kotlin/hw.kt -o hwkt
