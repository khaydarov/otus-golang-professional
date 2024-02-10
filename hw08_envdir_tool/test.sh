#!/usr/bin/env bash
set -xeuo pipefail

go build -o go-envdir

export HELLO="hello"
export FOO="foo"
export UNSET=""
export ADDED="from original env"
export EMPTY=""
export CUSTOM="custom"

result=$(./go-envdir "$(pwd)/testdata/env" "/bin/bash" "$(pwd)/testdata/echo.sh" arg1=1 arg2=2)
expected='HELLO is ("hello")
BAR is (bar)
FOO is (   foo
with new line)
UNSET is ()
ADDED is (from original env)
EMPTY is ()
CUSTOM is (custom)
arguments are arg1=1 arg2=2'

[ "${result}" = "${expected}" ] || (echo -e "invalid output: ${result}, expected: ${expected}" && exit 1)

rm -f go-envdir
echo "PASS"
