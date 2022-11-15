/*====================
16.How Node JS Works
CMD: [ node 16.HowNodeJSWorks/index.js ]
===================*/

console.log("Starting up");

setTimeout(()=>{
    console.log("2 second log")
 },2000);

 setTimeout(()=>{
    console.log("0 second log")
 },0);

 console.log("finishing up")