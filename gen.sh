#!/bin/bash

#protoc --go_out=plugins=grpc:.  greet/greetpb/greet.proto
protoc --go_out=plugins=grpc:.  calculator/calcpb/calc.proto