#!/bin/bash

export failed=0
for fixture in $(ls fixtures/*.json); do
  echo "Testing $fixture"
  rm -f missingvar
  envvars=$(VCAP_APPLICATION={} VCAP_SERVICES=$(cat $fixture) go run main.go envvars)
  (
    eval "$(VCAP_APPLICATION={} VCAP_SERVICES=$(cat $fixture) go run main.go exports)"
    for envvar in $envvars; do
      [[ "${!envvar:-MISSING}" == "MISSING" ]] && { touch missingvar; }
    done
  )
  [[ -f missingvar ]] && { failed=1; rm -f missingvar; echo "Failed."; }
done

[[ "$failed" == "0" ]] && { echo "Success!"; } || { echo "FAILED!"; }
exit $failed
