# 2019年1月11日 星期五
# version 3.6
# TensorFlow hello world
# tf-hello-world.py



import tensorflow as tf

# link+ https://blog.csdn.net/hq86937375/article/details/79696023
# TIPS: I tensorflow/core/platform/cpu_feature_guard.cc:141] Your CPU supports instructions that this TensorFlow binary was not compiled to use: AVX2

# method 1
import os
os.environ['TF_CPP_MIN_LOG_LEVEL'] = '2'

tf.enable_eager_execution()

print(tf.add(100, 3))

hello = tf.constant('Hello, TensorFlow!')
print(hello.numpy())

print(' TF-V' + tf.VERSION)