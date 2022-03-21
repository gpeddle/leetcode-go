#!/bin/bash

if [[ $# -ne 1 ]] ; then
    echo init-problem.sh [problem-dir]
    exit 1
fi

mkdir algorithms/$1
touch algorithms/$1/$1.go
touch algorithms/$1/problem.md

