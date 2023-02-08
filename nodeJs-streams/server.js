const http = require("http");
const fs = require("fs");
const { Transform, pipeline } = require("stream");
const replaceWordProcessing = require("./replaceWordProcessor");
const uppercaseWordProcessing = require("./uppercaseWordProcessor");

const server = http.createServer((req, res) => {
  if (req.url != "/") return res.end();

  // Downloading big file using good way
  // const readableStream = fs.createReadStream('sample.txt');
  // readableStream.pipe(res);

  // copy file good way
  // const readStream = fs.createReadStream('sample.txt');
  // const writeStream = fs.createWriteStream('output.txt');
  // readStream.on('data', (chunk) => {
  //     console.log('Chunk', chunk);
  //     writeStream.write(chunk);
  // });

  // String processing
  const sampleFileStream = fs.createReadStream("sample.txt");
  const outputWriteableStream = fs.createWriteStream("output.txt");

  pipeline(
    sampleFileStream,
    replaceWordProcessing,
    uppercaseWordProcessing,
    outputWriteableStream,
    (err) => {
      if(err) console.log('ERROR HANDLING: ',err);
    }
  );

  // sampleFileStream.pipe(replaceWordProcessing).pipe(uppercaseWordProcessing).pipe(outputWriteableStream);

  // sampleFileStream.on('data', chunk => {
  //     console.log('Data Received', chunk.toString());

  //     // Process
  //     const finalString = chunk.toString().replaceAll(/ipsum/gi, 'Romi').toUpperCase();

  //     // Writeable stream write
  //     outputWriteableStream.write(finalString);
  // });

  res.end();
});

const PORT = process.env.PORT || 4000;
server.listen(PORT, console.log(`Listening on port ${PORT}`));
