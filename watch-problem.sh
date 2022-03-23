#!/bin/bash

if [[ $# -ne 1 ]] ; then
    echo watch.sh [problem-dir]
    exit 1
fi

fullname=$1
problem=${fullname:11}
problem_args="algorithms/${problem}/${problem}.go algorithms/${problem}/${problem}_test.go"

grc gow test -v $problem_args
