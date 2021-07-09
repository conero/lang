# Nodejs/Npm 教程

> Joshua Conero
>
> 2018年9月29日 星期六





## NPM

### package.json

> 版本设置含义

- ~ 会匹配最近的小版本依赖包，比如~1.2.3会匹配所有1.2.x版本，但是不包括1.3.0
- ^ 会匹配最新的大版本依赖包，比如^1.2.3会匹配所有1.x.x的包，包括1.3.0，但是不包括2.0.0
- `* `这意味着安装最新版本的依赖包

```json
{
    "devDependencies": {
    "babel-core": "~6.26.3",
    "babel-loader": "~7.1.4",
    "babel-preset-env": "~1.7.0",
    "babel-preset-es2015": "~6.24.1",
    "lodash": "~4.17.10",
    "webpack": "~4.12.1",
    "webpack-cli": "~2.1.5"
  },
  "dependencies": {
    "js-base64": "^2.4.5",
    "pdfmake": "^0.1.37",
    "store": "^2.0.12"
  }
}
```





### node

由于node使用多存在多版本非兼容问题，可使用版本管理工具

- **nvm**    相关连接， [nvm](https://github.com/nvm-sh/nvm)、[nvm-windows](https://github.com/coreybutler/nvm-windows)

