#!/bin/bash

protoc --go_out=plugins=grpc:.  calcpb/calc.proto
cmd /k