#!/bin/sh -l 

echo "Building the docker to test the system" 
time=$(date)
/app/mainprogram
/app/mainprogram2
/app/test_program
echo "::set-output name=status::building@$time"
