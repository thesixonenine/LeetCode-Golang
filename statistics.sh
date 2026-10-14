#!/usr/bin/env bash
e=$(find ./easy -maxdepth 1 -type d | tail -n +2 | wc -l)
m=$(find ./medium -maxdepth 1 -type d | tail -n +2 | wc -l)
echo "Easy   ${e}"
echo "Medium ${m}"
echo -e "\ntotal  $((e+m))"