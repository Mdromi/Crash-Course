/*====================
06.MakeBasicServer
CMD: [ node 06.MakeBasicServer/index.js ]
===================*/

const http = require('http');
const { default: test } = require('node:test');

http.createServer((req, resp) => {
    resp.write("<h1>Hello, This is Md Romi</h1>");
    resp.end();
}).listen(4500);
