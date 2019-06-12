#!/bin/bash

# USAGE:
#   ./test/demo_fixtures.sh
#   ./test/demo_fixtures.sh exports
#   ./test/demo_fixtures.sh envvars

CMD=${1:-exports}
for fixture in $(ls fixtures/*.json); do
  echo "# Testing $fixture"
  VCAP_APPLICATION={} VCAP_SERVICES=$(cat $fixture) go run main.go $CMD
done