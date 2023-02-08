const { Readable, Writable } = require("stream");

const writableStream = new Writable({ objectMode: true, write(s) {
    // console.log('Writing - ', s.toString());
    console.log('Writing - ', s);
} });
// writableStream.write('Hello World')


const readableStream = new Readable({ highWaterMark: 6, objectMode: true ,read() {} });
readableStream.on("data", (chunk) => {
//   console.log("Data Coming - ", chunk.toString());
  console.log("Data - ", chunk);
  writableStream.write(chunk);
});

// readableStream.push("Hello from readable stream");
readableStream.push({
  name: 'MD Romi',
  age: 19,
  job: "student"
});

