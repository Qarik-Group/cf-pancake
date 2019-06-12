* Any env var that does not start with a letter is now prefixed with an underscore (invalid `3RD_HOSTNAME` becomes valid `_3RD_HOSTNAME`)
* Simple test harness at `test/test_eval.sh`
* Simple demo of each subcommand at `test/demo_fixtures.sh`
* `cf-pancake` CLI built using Go 1.12
* Moved to using Go Modules, instead of `dep`
