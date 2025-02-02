#!/bin/bash

# Run R script
Rscript miller-mtpa-chapter-1-program.R

# Run Python script 
python3 miller-mtpa-chapter-1-program.py

# Run Go tests
cd cmd/GoStatsTests
go test -v
cd ../..

# Run Go main program to compare results
cd cmd/GoStatsTests
go run main.go
cd ../..
