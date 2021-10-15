#!/bin/sh

STAGED_GO_FILES=$(git diff --cached --name-only | grep ".go$")

# shellcheck disable=SC2039
if [[ "$STAGED_GO_FILES" == "" ]]; then
  exit 0
fi

PASS=true

for FILE in $STAGED_GO_FILES; do
  goimports -w "$FILE"

  golint "-set_exit_status" "$FILE"
  # shellcheck disable=SC2039
  if [[ $? == 1 ]]; then
    printf "**** golint FAILED for %s\n", "$FILE"
    PASS=false
  fi

  go vet "$FILE"
  # shellcheck disable=SC2039
  # shellcheck disable=SC2181
  if [[ $? != 0 ]]; then
    printf "**** go vet FAILED for %s\n", "$FILE"
    PASS=false
  fi


  go test -v "$FILE"
  # shellcheck disable=SC2039
  # shellcheck disable=SC2181
  if [[ $? != 0 ]]; then
    printf "**** go test FAILED for %s\n", "$FILE"
    PASS=false
  fi
done

if ! $PASS; then
  printf "COMMIT FAILED\n"
  exit 1
else
  printf "COMMIT SUCCEEDED\n"
fi

exit 0
