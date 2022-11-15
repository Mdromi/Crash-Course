/*====================
13.CRUD with File System
CMD: [ node 13.CRUDFileSystem/index.js ]
===================*/
const fs = require('fs');
const path = require('path');
const dirPath = path.join(__dirname, `CRUD`);
const filePath = `${dirPath}/apple.txt`

/* CREATE FILE */
// fs.writeFileSync(filePath, 'This is a simple text file');

/* READ FILE */
// fs.readFile(filePath, 'utf8', (err, item) => {
//     console.log(item)
// });

/* UPDATE FILE */
// fs.appendFile(filePath, ' and file name is apple.txt', (err) => {
//     if(!err) console.log('file is update');
// });

/* RENAME FILE */
// fs.rename(filePath, `${dirPath}/fruit.txt`, (err) => {
//     if(!err) console.log('file name is update!');
// });

/* DELTE FILE */
fs.unlinkSync(`${dirPath}/fruit.txt`);