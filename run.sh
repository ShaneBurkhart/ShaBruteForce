#!/bin/sh

for i in `seq 5 13`; do
  nohup go run main.go $i &
  sleep 1
done


