#!/bin/bash

while [ -n "$1" ];do
i=$1
j=$2
k=$3
shift
shift
shift
if [ "$j" == '+' ];
        then
        result=`expr $i + $k`
elif [ "$j" == '-' ];
        then
        result=`expr $i - $k`
elif [ "$j" == '/' ];
then
        result=`expr $i / $k`
else
        result=`expr $i \* $k`
fi
echo $result
done
