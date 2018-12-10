@echo off
@REM 生成 dll 等动态链接文件
rustc --crate-type=cdylib ../share.rs

@REM 运行 python 3
echo "python run:"
python call.py

pause ...