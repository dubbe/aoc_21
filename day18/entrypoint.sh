#!/bin/sh
ts=$(date +%s%N)
$@
te=$(date +%s%N)
echo $((($te - $ts)/1000000))ms