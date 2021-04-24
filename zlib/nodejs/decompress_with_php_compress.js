const { createInflate } = require('zlib');
const { pipeline } = require('stream');
const {
  createReadStream,
  createWriteStream
} = require('fs');


const gunzip = createInflate();
const sourceUnzip = createReadStream('../php/data.zip');
const destinationUnzip =  createWriteStream('php_data.txt');

pipeline(sourceUnzip, gunzip, destinationUnzip, (err) => {
  if (err) {
    console.error('发生错误:', err);
    process.exitCode = 1;
  }
});