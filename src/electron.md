# Electron (2017年9月5日 星期二/electron学习笔记/Joshua Conero)
>
    Build cross platform desktop apps with JavaScript, HTML, and CSS 
    Electron + Node + Chromium + V8
    since 2013

    In Electron, Node.js and Chromium share a single V8 instance—usually the version that Chromium is using. Most of the time this just works but sometimes it means patching Node.js

    For instance, Electron uses just the rendering library from Chromium rather than all of Chromium.
    使用 chromium 的渲染库

## Electron    

-    asar   打包源代码为单一文件
-    系统分文： 主进程(Main Process)和渲染进程(Render Process)； 前者负责搭建界面，后者负责渲染视图

### Main Process 主进程

    electron中，由package.json中的main.js运行出来的进程为主进程(Main Process)。主进程用于创建GUI界面以便web页面的展示。
    electron由Chromium负责页面的显示，所以当创建一个页面时，就会对应的创建渲染进程(Renderer Process)。 
    在渲染进程中，直接调用原生的GUI接口是十分危险的。如果你想在渲染进程中使用原生的GUI的功能，需要让渲染进程与主进程进行通信，再由主进程去调用对应接口。

## Node  - 基于 V8 的单事件驱动库  
## Chromium -  
## V8    - javascript 解析器