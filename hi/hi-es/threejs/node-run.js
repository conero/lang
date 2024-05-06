// NodeJS 运行
// 2019年1月6日 星期日


console.log('  通过 threeJs 学习 WebGL');
console.log(`  ${Math.random()}`);


const express = require('express')
const app = express()

app.get('/jc/test', function (req, res) {
  res.send('Hello World')
})


// 静态文本
var options = {
    dotfiles: 'ignore',
    etag: false,
    extensions: ['htm', 'html'],
    index: false,
    maxAge: '1d',
    redirect: false,
    setHeaders: function (res, path, stat) {
      res.set('x-timestamp', Date.now())
    }
}
  
app.use(express.static('/public', options))
app.use(express.static(__dirname))

app.listen(3000)