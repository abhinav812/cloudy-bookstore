#!/bin/sh

STAGED_GO_FILES=$(git diff --cached --name-only | grep ".go$")

# shellcheck disable=SC2039
if [[ "$STAGED_GO_FILES" == "" ]]; then
  exit 0
fi

PASS=true

for FILE in $STAGED_GO_FILES; do
  printf "RUNNING GOIMPORTS\n"
  goimports -w "$FILE"

  printf "RUNNING GO LINT\n"
  golint "-set_exit_status" "$FILE"
  # shellcheck disable=SC2039
  if [[ $? == 1 ]]; then
    PASS=false
  fi

  printf "RUNNING GO VET\n"
  go vet "$FILE"
  # shellcheck disable=SC2039
  # shellcheck disable=SC2181
  if [[ $? != 0 ]]; then
    PASS=false
  fi

  printf "RUNNING GO TEST\n"
  go test -v "$FILE"
  # shellcheck disable=SC2039
  # shellcheck disable=SC2181
  if [[ $? != 0 ]]; then
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
