# DeepLearning

> 2018年8月2日 星期四
>
> Joshua Conero



机器学习（machine learning，ML）是一类强大的可以从经验中学习的技术。  通常采用观测数据或与环境交互的形式，机器学习算法会积累更多的经验，其性能也会逐步提高。 使用数据训练模型（数据驱动）。

深度学习是机器学习的一个主要分支。



当前主流的深度学习框架有

- pytorch   [pytorch/pytorch](https://github.com/pytorch/pytorch)，[官网-pytorch.org](https://pytorch.org/)。由Facebook人工智能学院提供支持服务的
- TensorFlow  [tensorflow/tensorflow](https://github.com/tensorflow/tensorflow)，[官网-tensorflow.org](https://tensorflow.org/)。由Google智能机器研究部门研发
- paddle 百度飞将学习平台，[paddlepaddle/paddle](https://github.com/paddlepaddle/paddle)，[官网-www.paddlepaddle.org](http://www.paddlepaddle.org/)
- keras    [keras-team/keras](https://github.com/keras-team/keras)，[官网-keras.io](http://keras.io/)
- scikit-learn  [scikit-learn/scikit-learn](https://github.com/scikit-learn/scikit-learn)，[官网-scikit-learn.org](https://scikit-learn.org/)
- transformers  自然语言处理库 [huggingface/transformers](https://github.com/huggingface/transformers)





学习资源

- [动手学深度学习 v2版本](https://zh.d2l.ai/)



> 用途方向及领域

- 计算机视觉
- 自动语音识别
- 强化学习
- 统计建模
- 医疗保健和基因组学



> 所需数学知识

- 集合代数
- 线性代数
- 微积分
- 概率与信息论



## 基础

机器学习组件要素

- data：用来学习的**数据**
- model：如何转换数据的**模型**
- object function：一个用来量化模型的有效性的**目标函数**
- algorithm：调整模型参数以优化目标函数的**算法**



**数据**

每个数据集由一个个样本（example, sample）组成，样本有时也叫做数据点（data point）或者数据实例（data instance），通常每个样本由一组称为特征（features，或协变量（covariates））的属性组成。 机器学习模型会根据这些属性进行预测。在监督学习问题中，要预测的是一个特殊的属性，它被称为标签（label，或目标（target））。

数据集由一些为训练而收集的样本组成，称为训练数据集（training dataset，或称为训练集（training set））；“新数据集”通常称为测试数据集（test dataset，或称为测试集（test set））。因此，可用数据集通常可以分成两部分：训练数据集用于拟合模型参数，测试数据集用于评估拟合的模型。



**目标函数**

在机器学习中，我们需要定义模型的优劣程度的度量，这个度量在大多数情况是“可优化”的，这被称之为目标函数（objective function）。 我们通常定义一个目标函数，并希望优化它到最低点。 因为越低越好，所以这些函数有时被称为损失函数（loss function，或cost function）



监督学习处理方法

- 回归
- 分类
- 标记的问题
- 搜索
- 推荐系统：个性化
- 序列学习



### 主流框架比较

**pytorch**

以动态计算图为基础，使得模型的设计和调试更加直观。

简单易用，适合研究人员和初学者，也广泛应用于学术界



**TensorFlow**

采用静态计算图，使其更适用于生产环境和大规模部署。

丰富的生态系统，支持广泛的应用领域，从移动设备到云端。



**paddle**

注重产业实践，提供全面的深度学习平台，包括框架、工具和服务。

具有开放、易扩展的特性，适用于产业界的多样需求。



## TensorFlow



### 简介

TensorFlow™ 是一个开放源代码软件库，用于进行高性能数值计算。借助其灵活的架构，用户可以轻松地将计算工作部署到多种平台（CPU、GPU、TPU）和设备（桌面设备、服务器集群、移动设备、边缘设备等）。TensorFlow™ 最初是由 Google Brain 团队（隶属于 Google 的 AI 部门）中的研究人员和工程师开发的，可为机器学习和深度学习提供强力支持，并且其灵活的数值计算核心广泛应用于许多其他科学领域。 



_Tensor的意思是张量，代表N维数组；Flow的意思是流，代表基于数据流图的计算。把N维数字从流图的一端流动到另一端的过程，就是人工智能神经网络进行分析和处理的过程。_



//@TODO  [《使用 Python 进行深度学习》(中文)](https://github.com/cnbeining/deep-learning-with-python-cn)

//@TODO 阅读 [机器学习速成课程
使用 TensorFlow API](https://developers.google.cn/machine-learning/crash-course/)



#### 教程

```powershell
# 安装 tensorflow
pip install tensorflow
```





## 机器学习

*机器学习教计算机执行人和动物与生俱来的活动：<span style="color: red; font-size: 1.2em;">从经验中学习</span>。机器学习算法使用计算方法直接从数据中“学习”信息，而不依赖于预定方程模型。当可用于学习的样本数量增加时，这些算法可自适应提高性能。*

*机器学习算法可从能够带来洞察力的数据中发现自然模式，帮助您更好地制定决策和做出预测。*



*当您遇到涉及大量数据和许多变量的复杂任务或问题，但没有现成的处理公式或方程式时，可以考虑使用机器学习。*



*机器学习采用两种类型的技术：监督式学习和无监督学习。监督式学习根据已知的输入和输出训练模型，让模型能够预测未来输出；无监督学习从输入数据中找出隐藏模式或内在结构。*

![](../image/AI/dl/ml-classify-AB.png)



### 监督式学习

在“给定输入特征”的情况下预测标签

*监督式机器学习旨在构建能够根据存在不确定性的证据做出预测的模型。监督式学习算法接受已知的输入数据集和对数据的已知响应（输出），然后训练模型，让模型能够为新输入数据的响应生成合理的预测。*

*监督式学习采用分类和回归技术开发预测模型。*

- 分类技术可预测离散的响应
- 回归技术可预测连续的响应



### 无监督学习

*无监督学习可发现数据中隐藏的模式或内在结构。这种技术可根据包含未标记响应的输入数据的数据集执行推理。*

*聚类是一种最常用的无监督学习技术。这种技术可通过探索性数据分析发现数据中隐藏的模式或分组。*

*聚类的应用包括基因序列分析、市场调查和对象识别。*



//@TODO  ~conero\docs-era\matlab之于人工智能 **93116v00_machine_learning_section3_ebook_v05**



> *机器学习的挑战*

- 数据会以各种形式和大小出现(如离散之类)
- 预处理数据需要掌握专业的知识和工具
- 找到拟合数据的最佳模型需要时间



> *反复尝试和出错才是机器学习的核心*



## Keras

- [中文文档](https://keras.io/zh/)



## TensorFlow.js

> javascript 版 TensorFlow 库，可用于浏览器(Browser)或NodeJs中

- 网站： [https://js.tensorflow.org/](https://js.tensorflow.org/)
- API 文档: https://js.tensorflow.org/api/latest/



*TensorFlow.js 与 Tensors (张量)、Layers (图层)、Optimizers (优化器) 和损失函数等概念兼容，TensorFlow.js 为 JavaScript 中神经网络编程提供了灵活的构建块。*



**环境**

当一个用TensorFlow.js开发的程序运行时，所有的配置被统称为环境。它包含一个全局的backend，以及一些可以精确控制TensorFlow.js特性的标记。

```javascript
// 打印当前所使用的环境
console.log(tf.getBackend());
```





### 起始

1. `npm install @tensorflow/tfjs`   *安装纯 js 版本TensorFlow* (<font style="color:red;">NodeJs/Browser</font>)

2. `npm install @tensorflow/tfjs-node`  *安装原生 C++ 绑定的TensorFlow* (<font style="color:red;">NodeJs</font>)



### 张量(Tensors)和操作(operations)

*TebsorFlow.js是一个在JavaScript中使用张量来定义并运行计算的框架。张量是向量和矩阵向更高维度的推广。*



*tf.Tensor是TensorFlow.js中的最重要的数据单元，它是一个形状为一维或多维数组组成的数值的集合。*

一个 `tf.Tensor` 包含如下属性（`tf.tensor (values, shape?, dtype?)`）：

- `rank`: 张量的维度
- `shape`: 每个维度的数据大小
- `dtype`: 张量中的数据类型

`tf.Tensor` 可进行进本的数值运算，如取平方 tensor.square, 相加 tensor.add



## LLM

large language model，大语言模型。

模型基于 Transformer 架构，需要非常大的 GPU 数据中心。





## 参考

- [https://tensorflow.google.cn/](https://tensorflow.google.cn/)   谷歌开发者中心
- [Tensorflow 中文网](http://www.tensorfly.cn/)
- [深度学习框架之争：PyTorch、TensorFlow和PaddlePaddle的全面对比](https://blog.csdn.net/BetrayFree/article/details/135522180)
- [目前学习深度学习框架哪个比较好（paddlepaddle、tensorflow、pytorch）？](https://www.zhihu.com/question/421753332/answer/3123241291)
- llm
  - [周末推荐一篇LLM的必读好文](https://mp.weixin.qq.com/s/plok-Bh20A8pkeOXgFX4Ig)


