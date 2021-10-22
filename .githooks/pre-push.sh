#!/bin/bash

protected_branch='master'

## Do not push to master
if read -r local_ref local_sha remote_ref remote_sha; then
    if [[ "$remote_ref" == *"$protected_branch"* ]]
    then
         echo "Pushing directly to master is disabled."
            exit 1 # push will not execute
    else
        exit 0 # push will execute
    fi
fi

## Do not push if trunk is ahead of current branch
trunk="origin/develop"
local_rev=$(git rev-parse @)
echo "Local revision: ${local_rev}"
trunk_rev=$(git rev-parse "$trunk")
echo "Trunk revision: ${trunk_rev}"
base_rev=$(git merge-base @ "$trunk")
echo "Base revision: ${base_rev}"

if [ "$local_rev" = "$trunk_rev" ]; then
  echo "Pre-push git hook check finished. Up-to-date. No need to push."
elif [ "$local_rev" = "$base_rev" ]; then
  echo "Pre-push git hook check finished. Trunk is ahead of local branch. Please take a pull from trunk first. Failed to Push."
  exit 1
elif [ "$trunk_rev" = "$base_rev" ]; then
  echo "Pre-push git hook check finished. OK to push."
else
  echo "Pre-push git hook check finished. Diverged with trunk. Please take a pull from trunk first. Failed to Push."
    exit 1
fi