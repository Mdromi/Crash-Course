/*====================
12.Show File List
CMD: [ node 12.ShowFileList/index.js ]
===================*/

const { dir } = require('console');
const fs = require('fs');
// fs.writeSync('apple.txt', 'this is a aplle file');

const path = require('path');
const dirPath = path.join(__dirname, `files`);
// console.warn(dirPath)

// for(let i = 0; i < 5; i++){
//     fs.writeFileSync(dirPath+`/hello${i}.txt`, 'a simple txt file')
// }

fs.readdir(dirPath, (err, files) => {
    files.forEach((items) => {
        console.log("Foles Name is ", items)
    })
})