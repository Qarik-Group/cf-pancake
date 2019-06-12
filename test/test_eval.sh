#!/bin/bash

failed=

for fixture in $(ls fixtures/*.json); do
  echo "Testing $fixture"
  (eval "$(VCAP_APPLICATION={} VCAP_SERVICES=$(cat $fixture) go run main.go exports)")
  [[ $? != 0 ]] && { failed=1; }
done

[[ $failed == 0 ]] && { echo "Success!"; } || { echo "FAILED!"; }
exit $failed
