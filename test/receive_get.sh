#!/bin/bash
for((i=1;i<=100;i++))
do
curl -i http://127.0.0.1:1991/receive/32002
done
