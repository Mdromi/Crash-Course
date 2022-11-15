/*====================
04.Fundaments Of Node js
CMD: [ node 04.FundamentsOfNodeJs/index.js ]
===================*/
const app = require('./app.js');

// console.log(app.z())

const arr = [2, 4, 5, 7, 9];

let result = arr.filter((item) => {
    // console.log(item)

    return item >= 4
})

console.log(result);