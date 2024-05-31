#!/bin/bash

# Prompt user to enter option 1) Build 2) Run 3) Test 4) Migrate
echo "Please select an option :"
echo "1) Build"
echo "2) Run"
echo "3) Test"

# Read user input
read -p "Enter option: " option

# Check user input
if [ $option -eq 1 ]
then
    echo "Building..."
    go build -o bin/maestrore-archive.exe main.go
    cp -r public bin/
elif [ $option -eq 2 ]
then
    echo "Running..."
    go build -o bin/maestrore-archive.exe main.go
    cp -r public bin/

    ./bin/maestrore-archive.exe
elif [ $option -eq 3 ]
then
    echo "Testing..."
    go test ./...
else
    echo "Invalid option"
fi