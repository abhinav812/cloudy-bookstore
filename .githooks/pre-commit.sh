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
done

# go vet/test sometime gives warning on individual files, ignoring dependent files in the same package.
go test -v ./...
# shellcheck disable=SC2039
# shellcheck disable=SC2181
if [[ $? != 0 ]]; then
  printf "**** go test FAILED"
  PASS=false
fi

go vet ./...
# shellcheck disable=SC2039
# shellcheck disable=SC2181
if [[ $? != 0 ]]; then
  printf "**** go vet FAILED "
  PASS=false
fi

if ! $PASS; then
  printf "COMMIT FAILED\n"
  exit 1
else
  printf "COMMIT SUCCEEDED\n"
fi

exit 0
