#!/bin/bash

if [[ $# -ne 1 ]] ; then
    echo init-tests.sh [problem-dir]
    exit 1
fi

# create tests

gotests -all -parallel -w $1