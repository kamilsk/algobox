<?php
$handle = fopen('php://stdin', 'r');
echo 'Hello, World.', PHP_EOL, fgets($handle), PHP_EOL;
fclose($handle);
