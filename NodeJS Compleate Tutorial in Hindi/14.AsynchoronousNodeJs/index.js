/*====================
14.Asynchoronous Basics in NodeJs
CMD: [ node 14.AsynchoronousNodeJs/index.js ]
===================*/

let a=20;
let b=0;

setTimeout(()=>{
   b=30;
},2000)


console.log(a+b)