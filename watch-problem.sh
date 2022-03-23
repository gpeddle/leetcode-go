#!/bin/bash

if [[ $# -ne 1 ]] ; then
    echo watch.sh [problem-dir]
    exit 1
fi

problem_dir=$1
if [[ ${problem_dir: -1} == '/' ]] ; then
  problem_dir=${problem_dir:0:-1}
fi

problem=${problem_dir:11}
problem_args="algorithms/${problem}/${problem}.go algorithms/${problem}/${problem}_test.go"

grc gow test -v $problem_args
