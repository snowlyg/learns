const { createDeflate } = require('zlib');
const { pipeline } = require('stream');
const {
  createReadStream,
  createWriteStream
} = require('fs');

const gzip = createDeflate();
const source = createReadStream('data.json');
const destination = createWriteStream('data.zip');

pipeline(source, gzip, destination, (err) => {
  if (err) {
    console.error('发生错误:', err);
    process.exitCode = 1;
  }
});
