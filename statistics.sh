#!/usr/bin/env bash

echo "Easy $(find ./easy -maxdepth 1 -type d | tail -n +2 | wc -l)"