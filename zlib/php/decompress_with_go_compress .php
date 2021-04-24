<?php
// Perform GZIP compression:
$compressed = file_get_contents("../go/data.zip");
$inflateContext = inflate_init(ZLIB_ENCODING_DEFLATE);
$uncompressed = inflate_add($inflateContext, $compressed, ZLIB_NO_FLUSH);
$uncompressed .= inflate_add($inflateContext, NULL, ZLIB_FINISH);
file_put_contents("go_data.json",$uncompressed);
echo $uncompressed;
?>

