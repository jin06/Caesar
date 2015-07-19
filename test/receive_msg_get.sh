#!/bin/bash
for((i=1;i<=1000000;i++))
do
curl -i http://127.0.0.1:1991/receive_msg/32002
done

