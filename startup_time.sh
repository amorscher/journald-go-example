#!/bin/bash

sum=0
n=$2
echo Running ${1} for ${2} times.

for ((i=1;i<=$2;i++)); 
do 
    timestamp=$(date +%s%N | cut -b1-13)
    # call the programm and capture the time: programm has to print time: ${time}
    captured_time=$($1  | sed -n 's/Time: \(.*\)/\1/p')
    # calculate the needed time
    ellapsedTime="$((captured_time-timestamp))"
    #used time for the programm execution
    echo ellapsedTime : ${ellapsedTime} Milliseconds
    sum="$((sum+ellapsedTime))"

done

# calculate the average of the run
average="$((sum/n))"
echo Average : ${average} Milliseconds
