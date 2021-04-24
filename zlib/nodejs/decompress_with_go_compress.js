const { createInflate } = require('zlib');
const { pipeline } = require('stream');
const {
  createReadStream,
  createWriteStream
} = require('fs');


const gunzip = createInflate();
// console.log(gunzip)
const sourceUnzip = createReadStream('../go/data.zip');
const destinationUnzip =  createWriteStream('go_data.txt');
pipeline(sourceUnzip, gunzip, destinationUnzip, (err) => {
  if (err) {
    console.error('发生错误:', err);
    process.exitCode = 1;
  }
});