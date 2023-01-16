#!/bin/bash

function read_dir(){
    hasTestFile=0
    for file in `ls $1`
    do
        if [ -d $1"/"$file ]
        then
            read_dir $1"/"$file
        else
            if [[ $file =~ .*_test.go$ ]]; then
                echo $1"/"$file
                cd $1"/"
                go test
                cd -
            fi
        fi
    done
}

read_dir $1
