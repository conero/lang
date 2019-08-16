# New-ther

> *IT 新兴技术汇总*
>
> 2019年3月25日 星期一





## VR 全景

*全景(英文名称是 Panorama)是把相机环360 度拍摄的一组或多组照片拼接成一个全景图像，通过计算机技术实现全方位互动式观看的真实场景还原展示方式。VR全景最早起源于谷歌地图。*





## oAuth2

*一个关于授权（authorization）的开放网络标准，发布标准 [RFC 6749](http://www.rfcreader.com/#rfc6749)。*



*OAuth2.0授权框架允许第三方应用获取对HTTP服务的有限的访问权限，既可以以资源所有者名义在资源所有者和HTTP服务之间进行允许的交互，也可以允许第三方应用以自己的名义进行访问。*



**访问令牌/token**  *一个代表特定作用域、生命周期以及其他访问权限属性的字符串。*



**角色**

​		resource owner：  		用户，资源所有者

​		resource server：  		服务器，资源服务器

​		client：							客户端

​		authorization server：  授权服务器，认证服务

*请求流程：*

```
 +--------+                               +---------------+
 |        |--(A)- Authorization Request ->|   Resource    |
 |        |                               |     Owner     |
 |        |<-(B)-- Authorization Grant ---|               |
 |        |                               +---------------+
 |        |
 |        |                               +---------------+
 |        |--(C)-- Authorization Grant -->| Authorization |
 | Client |                               |     Server    |
 |        |<-(D)----- Access Token -------|               |
 |        |                               +---------------+
 |        |
 |        |                               +---------------+
 |        |--(E)----- Access Token ------>|    Resource   |
 |        |                               |     Server    |
 |        |<-(F)--- Protected Resource ---|               |
 +--------+                               +---------------+
```



授权码（*Authorization Code*） 资源所有者密码凭据（即用户名和密码），可以直接作为获取访问令牌的授权许可。

访问令牌(access token)  令牌代表了访问权限的由资源所有者许可并由资源服务器和授权服务器实施的具体范围和期限。令牌字符串由数据和签名组成。