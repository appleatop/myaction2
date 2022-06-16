#!/bin/sh -l 

echo "Building the docker to test the system" 
time=$(date)
ls -al /app/mainprogram
/app/mainprogram
/app/mainprogram2
go test -v ./test2
echo "::set-output name=status::building@$time"
