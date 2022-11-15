/*====================
15.Handle Asynchoronous Data
CMD: [ node 15.HandleAsynchoronousData/index.js ]
===================*/
let a=20;
let b=0;



let wattingData = new Promise((resolve, reject) => {
    setTimeout(()=>{
        resolve(30)
    },2000);
})

wattingData.then((data) => {
    b = data
    console.log(a+b)
})