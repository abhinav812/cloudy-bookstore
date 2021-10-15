#!/bin/bash

protected_branch='master'

## Do not push to master
if read local_ref local_sha remote_ref remote_sha; then
    if [[ "$remote_ref" == *"$protected_branch"* ]]
    then
         echo "Pushing directly to master is disabled."
            exit 1 # push will not execute
    else
        exit 0 # push will execute
    fi
fi