#!/bin/sh

for i in `seq 14 19`; do
  nohup go run main.go $i &
  sleep 1
done


