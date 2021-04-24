<?php
// Perform GZIP compression:
$compressed = file_get_contents("data.json");
$deflateContext = deflate_init(ZLIB_ENCODING_DEFLATE);
$compressed = deflate_add($deflateContext, $compressed, ZLIB_NO_FLUSH);
$compressed .= deflate_add($deflateContext, NULL, ZLIB_FINISH);

file_put_contents("data.zip",$compressed);

// Perform GZIP decompression:
$inflateContext = inflate_init(ZLIB_ENCODING_DEFLATE);
$uncompressed = inflate_add($inflateContext, $compressed, ZLIB_NO_FLUSH);
$uncompressed .= inflate_add($inflateContext, NULL, ZLIB_FINISH);
echo $uncompressed;
?>