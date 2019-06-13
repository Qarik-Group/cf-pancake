# Pancake

Flatten `$VCAP_SERVICES` into simple, regular environment variables.

Instead of `$VCAP_SERVICES` with a PostgreSQL binding:

```json
{
  "elephantsql":[{
  "label": "elephantsql",
  "plan": "turtle",
  "name": "myapp-db",
  "tags": ["postgresql"],
  "instance_name": "myapp-db",
  "credentials": {
    "uri": "postgres://user:password@raja.db.elephantsql.com:5432/dbname",
    "max_conns": "5"
  }
}]}
```

You will get environment variables based on the service name/label `elephantsql`:

```plain
export ELEPHANTSQL_URI=postgres://user:password@raja.db.elephantsql.com:5432/dbname
export ELEPHANTSQL_MAX_CONNS=5
```

And you'll get env vars based on the service instance named you choose `myapp-db`:

```plain
export MYAPP_DB_URI=postgres://user:password@raja.db.elephantsql.com:5432/dbname
export MYAPP_DB_MAX_CONNS=5
```

And you'll get env vars based on each tag within the service instance:

```plain
export POSTGRESQL_URI=postgres://user:password@raja.db.elephantsql.com:5432/dbname
export POSTGRESQL_MAX_CONNS=5
```

## Install

You don't typically install or use `cf-pancake` CLI directly.

### Buildpack Usage

Commonly you will use the [pancake-buildpack Cloud Foundry buildpack](https://github.com/starkandwayne/pancake-buildpack).

If your platform operator has installed the buildpack as a system buildpack:

```yaml
applications:
- name: myapp
  services:
  - myapp-db
  buildpacks:
  - pancake_buildpack
  - php_buildpack
```

If you do not see `pancake_buildpack` within your `cf buildpacks` list, then you have two options:

Y can use the buildpack with its Git URL:

```yaml
applications:
- name: myapp
  services:
  - myapp-db
  buildpacks:
  - https://github.com/starkandwayne/pancake-buildpack
  - php_buildpack
```

Or you can use a [pre-built buildpack release](https://github.com/starkandwayne/pancake-buildpack/releases), for example:

```yaml
applications:
- name: myapp
  services:
  - myapp-db
  buildpacks:
  - https://github.com/starkandwayne/pancake-buildpack/releases/download/v1.0.0/pancake_buildpack-cached-cflinuxfs3-v1.0.0.zip
  - php_buildpack
```

### Cloud Native Buildpack

If you are dabbling with Cloud Native Builpacks then check out [pancake-cnb](https://github.com/starkandwayne/pancake-cnb).

### Legacy Cloud Foundry

If your Cloud Foundry is old and does not support multi-buildpacks, or does not allow "supply buildpacks" (you get errors with the buildpack complaining about a missing `bin/compile` file), then you have one more option.

Download the latest `cf-pancake` [release](https://github.com/starkandwayne/cf-pancake/releases) for linux-amd64 and store within your application source code prior to `cf push`:

```plain
mkdir -p vendor
curl -o vendor/cf-pancake -L $(curl -sSL https://api.github.com/repos/starkandwayne/cf-pancake/releases/latest | jq -r '.assets[] | select(.name == "cf-pancake-linux-amd64") | .browser_download_url')
```

And create a `.profile` script to use `cf-pancake`:

```bash
#!/bin/bash

chmod +x vendor/cf-pancake
eval "$(./vendor/cf-pancake exports)"
```

## Usage

Once you have installed `cf-pancake` using one of the methods above, your application and your `cf ssh` shell should now see the generated environment variables.

When using `cf ssh`, remember to run `/tmp/lifecycle/shell` in order to load your application's runtime environment (which includes `cf-pancake` behaviour):

```plain
# cf ssh phpapp
$ /tmp/lifecycle/shell
$ env
```

If you used the `.profile` (non-buildpack) approach above, then you need to source the `.profile` explicitly during `cf ssh`:

```plain
# cf ssh phpapp
$ /tmp/lifecycle/shell
$ source .profile
$ env
```

## Local development

As `cf-pancake exports` is designed to be run within an application container, if you try to run it locally then `$VCAP_SERVICES` will be missing.

There are some `fixtures/` JSON you can use:

```plain
(VCAP_APPLICATION={} VCAP_SERVICES=$(cat fixtures/elephantsql.json) go run main.go exports)
(VCAP_APPLICATION={} VCAP_SERVICES=$(cat fixtures/cleardb.json) go run main.go exports)
(VCAP_APPLICATION={} VCAP_SERVICES=$(cat fixtures/p-mysql.json) go run main.go exports)
(VCAP_APPLICATION={} VCAP_SERVICES=$(cat fixtures/two-services.json) go run main.go exports)
(VCAP_APPLICATION={} VCAP_SERVICES=$(cat fixtures/empty.json) go run main.go exports)
```

To check for errors, `eval` the results:

```plain
./test/test_eval.sh
```

You can also setup a local `$VCAP_SERVICES` (and `$VCAP_APPLICATION` is also required) for a Cloud Foundry application.

The [jq](http://stedolan.github.io/jq/) CLI is required for the commands below (assuming your example app has a unique name):

```plain
export NAME=myapp-name

export VCAP_APPLICATION="{}"
export VCAP_SERVICES=$(cf curl $(cf curl "/v2/apps?q=name:$NAME" | jq ".resources[0].metadata.url" | xargs echo)/env | jq -c -M .system_env_json.VCAP_SERVICES)
```

Confirm they are setup:

```plain
env | grep VCAP
```

Now run the `exports` command:

```plain
go run main.go exports
```
