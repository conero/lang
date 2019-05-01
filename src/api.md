# API

> Joshua Conero
>
> 2018年10月11日 星期四



*Application Programming Interface,应用程序编程接口*





## 应用

### WebService



**WebService是一种跨编程语言和跨操作系统平台的远程调用技术**



*XML+XSD,SOAP和WSDL就是构成WebService平台的三大技术。*



> 组成部分

- **XML**

- **SOAP** 

  *`Simple Object Access Protocol` 简单对象访问协议。 **SOAP协议 = HTTP协议 + XML数据格式***

- **WSDL**

  *Web Services Description Language*

  *WebService务器端首先要通过一个WSDL文件来说明自己家里有啥服务可以对外调用，服务是什么（服务中有哪些方法，方法接受 的参数是什么，返回值是什么），服务的网络地址用哪个url地址表示，服务通过什么方式来调用。*





> 包含 *服务端 (Service)* 和 *客服端 (client)*





#### WSDL

*definitions是WSDL文档的根元素*

| 元素        | 定义                       |
| ----------- | -------------------------- |
| <portType>  | web service 执行的操作     |
| <message>   | web service 使用的消息     |
| <types>     | web service 使用的数据类型 |
| <binding>   | web service 使用的通信协议 |
| <part>      | 消息参数                   |
| <operation> | 操作                       |



```xml
<!--参数： xml string-->
<part name="xml" type="xsd:string"/>
```





## RESTful API

