/*====================
05.CoreModule
CMD: [ node 05.CoreModule/index.js ]
===================*/
// const fs = require('fs');
fs = require('fs').writevSync;

console.log("Code By Step Step");

fs.writeFileSync("hello.txt","Code By Step Step");

console.log("--> ",__dirname);
console.log("--> ",__filename);