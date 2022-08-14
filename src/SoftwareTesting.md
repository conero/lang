# 软件测试

> 2018年10月25日
>
> Joshua Conero





## 概述

*笔者认为： 软件测试中最重要的是对 `所要测试软件` 的业务的了解程度，以及由此编写的 `测试用例` 和测试的结果分析；对于专门的软件测试人员而言。*





## 测试类型



### 渗透测试

> *penetration test*

*渗透测试，是为了证明网络防御按照预期计划正常运行而提供的一种机制。进行这类测试的，都是寻找网络系统安全漏洞的专业人士。*

*渗透测试是通过模拟恶意黑客的攻击方法，来评估计算机网络系统安全的一种评估方法。这个过程包括对系统的任何弱点、技术缺陷或漏洞的主动分析，这个分析是从一个攻击者可能存在的位置来进行的，并且从这个位置有条件主动利用安全漏洞。*



### web 压测

#### 概念

压测的目的是在于找到系统的瓶颈，一定是要确定系统某个方面达到瓶颈了，压力测试才算是基本完成。最理想的系统是低延迟，高吞吐。

- 测试环境（与生成环境尽可能一致）
- 基准值
  - 延迟(latency)
  - 吞吐量(thoughtout)            QPS
  - 峰值并发数                                   用户访问并发数
  - 峰值响应时间



#### 工具

- go 语言版本
  - [vegeta 压测工具](https://github.com/tsenart/vegeta)    可以固定的 QPS 来压测服务。
  - [ddosify 压测工具](https://github.com/ddosify/ddosify)
  - [nakabonne/ali 实时生成图标](https://github.com/nakabonne/ali)
  - [hey](https://github.com/rakyll/hey)   Apache ab 测试替代工具
  
- java版本
  - [jmeter](https://github.com/apache/jmeter) Apache图形化压测工具
  
- ab     Apache http web自带的压测工具。ab是单线程程序，只能利用单一CPU，在给性能好的服务器端应用做压测时，往往跑ab的测试机负荷满了；而服务器应用的性能还绰绰有余。

  

##### jmeter

> 线程属性

- 线程数：模拟真实用户数
- Ramp-up时间(秒): 与上一个线程的间隔时间。
- 循环次数：每个线程重复请求次数。



```shell
# 输出 html 报告文档
# 根据 jmx 文件生成 html 报表文件
jmter -n -t .\聚合报告-local-9310.jmx -l result.jtl -e -o ./rpt
```





##### vegeta

vegeta 的使用，powershell 环境错误，shell 正常。

```shell
echo "GET https://www.gy-imcloud.com" | vegeta attack -rate=10/s > results.gob
```



地址: https://dev.chinadatavalley.net/video/forum/931

```shell
# ddosify 压测工具，配置密码
ddosify -t https://dev.chinadatavalley.net/video/forum/931 -a "dev2022:Dev2022^"


# windows/powershell
echo "GET https://dev.chinadatavalley.net/video/forum/931" | vegeta attack -duration=5s -header="Authorization: Basic ZGVyMjikpEZXYyMDIyXgx" -output chinadatavalley-video.gob

# bash shell
echo "GET https://dev.chinadatavalley.net/video/forum/931" | vegeta attack -duration=5s -header="Authorization: Basic ZGVyMjikpEZXYyMDIyXgx" | tee chinadatavalley-video.gob | vegeta report

echo "GET https://dev.chinadatavalley.net/video/forum/931" | vegeta attack -duration=1s -rate=1000 -header="Authorization: Basic ZGVyMjikpEZXYyMDIyXgx" | tee chinadatavalley-video.gob | vegeta report

echo "GET https://dev.chinadatavalley.net/video/forum/931" | vegeta attack -rate=1/s -header="Authorization: Basic ZGVyMjikpEZXYyMDIyXgx" | tee chinadatavalley-video.gob | vegeta report

# 绘图
cat chinadatavalley-video.gob | vegeta plot --title "expo2022 video forum" > chinadatavalley-video.html
```





### 灰度测试

*灰度测试环境就是生产环境，生产数据，所影响的也是生产环境，只是范围比测试环境更广，更真实。其实就是小范围的生产环境。类似于游戏内测。*