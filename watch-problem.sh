#!/bin/bash

if [[ $# -ne 1 ]] ; then
    echo watch.sh [problem-dir]
    exit 1
fi

problem_args="algorithms/$1/$1.go algorithms/$1/$1_test.go"

grc gow test -v $problem_args
