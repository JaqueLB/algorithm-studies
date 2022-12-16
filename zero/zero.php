<?php
while (FALSE !== ($line = fgets(STDIN))) {
    $data = explode(" ", $line);
	if ($data[0] == $data[1] && $data[1] == $data[2]) {
        echo "*";
    } elseif ($data[0] == $data[1]) {
        echo "C";
    } else if ($data[1] == $data[2]) {
        echo "A";
    } else {
        echo "B";
    }
    exit();
}
