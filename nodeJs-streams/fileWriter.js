const fs = require("fs");

const writeableStream = fs.createWriteStream("log.txt");
const readableSteam = fs.createReadStream("log.txt");

// process.stdin.pipe(writeableStream);
readableSteam.pipe(process.stdout);