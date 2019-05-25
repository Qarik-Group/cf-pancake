cf-pancake
==========

Flatten `$VCAP_SERVICES` into regular environment variables.

Consider a `$VCAP_SERVICES` with a Postgresql binding:

```json
{
    "postgresql93": [
      {
        "credentials": {
          "dbname": "6jmpbpi4wovsk44l",
          "hostname": "10.10.2.7",
          "password": "yinredrg1va6xihy",
          "port": "49165",
          "uri": "postgres://2ibio9h9data939m:yinredrg1va6xihy@10.10.2.7:49165/6jmpbpi4wovsk44l",
          "username": "2ibio9h9data939m"
        },
        "label": "postgresql93",
        "name": "atk-pg",
        "plan": "free",
        "tags": [
          "postgresql93",
          "postgresql",
          "relational"
        ]
      }
    ]
  }
```

There are two usage modes - within the app container before startup (`cf-pancake exports`); or from your local machine (`cf-pancake set-env`).

If you bundle the linux64 version of `cf-pancake` into your application, you can run it during startup to setup environment variables.

`cf-pancake exports` returns bash `export` commands to setup flatten variables.

The output would look like:

```plain
$ cf-pancake e
export POSTGRESQL93_URI='postgres://ldo5vforrkoplrfb:2xmfoewk1ggtrm8y@10.10.2.7:49169/wybjfszhcd9xmbp1'
export POSTGRESQL93_HOSTNAME='10.10.2.7'
export POSTGRESQL93_PORT='49169'
export POSTGRESQL93_USERNAME='ldo5vforrkoplrfb'
export POSTGRESQL93_PASSWORD='2xmfoewk1ggtrm8y'
export POSTGRESQL93_DBNAME='wybjfszhcd9xmbp1'
```

To source the environment variables:

```plain
$(cf-pancake e)
```

Save the output to a script and then `source` that script to setup the variables.

Alternately, you can setup the environment variables from your local machine and store them within Cloud Foundry environment variables (as seen by `cf env`).

`cf-pancake set-env APPNAME` - updates the `APPNAME` with environment variables from that app's `$VCAP_SERVICES`

The output would look like:

```plain
$ cf-pancake set-env myapp
Setting env variable 'POSTGRESQL93_URI' to 'postgres://ldo5vforrkoplrfb:2xmfoewk1ggtrm8y@10.10.2.7:49169/wybjfszhcd9xmbp1' for app myapp in org intel / space myapp as admin...
OK
TIP: Use 'cf restage' to ensure your env variable changes take effect

Setting env variable 'POSTGRESQL93_USERNAME' to 'ldo5vforrkoplrfb' for app myapp in org intel / space myapp as admin...
OK
TIP: Use 'cf restage' to ensure your env variable changes take effect

Setting env variable 'POSTGRESQL93_DBNAME' to 'wybjfszhcd9xmbp1' for app myapp in org intel / space myapp as admin...
OK
TIP: Use 'cf restage' to ensure your env variable changes take effect

Setting env variable 'POSTGRESQL93_HOSTNAME' to '10.10.2.7' for app myapp in org intel / space myapp as admin...
OK
TIP: Use 'cf restage' to ensure your env variable changes take effect

Setting env variable 'POSTGRESQL93_PASSWORD' to '2xmfoewk1ggtrm8y' for app myapp in org intel / space myapp as admin...
OK
TIP: Use 'cf restage' to ensure your env variable changes take effect

Setting env variable 'POSTGRESQL93_PORT' to '49169' for app myapp in org intel / space myapp as admin...
OK
TIP: Use 'cf restage' to ensure your env variable changes take effect
```

The former would be run within an application container during startup.

The latter would be run by the developer outside of Cloud Foundry.

Installation
------------

If you want to use `cf-pancake export` within your application on Cloud Foundry, then download the `cf-pancake_linux_amd64` release from the [releases](https://github.com/cloudfoundry-community/cf-pancake/releases) page and add it to the project being uploaded. You will then need to create a custom startup script that uses it to create environment variables.

If you want to use `cf-pancake set-env APPNAME` locally, then you can either:

- download from the [releases](https://github.com/cloudfoundry-community/cf-pancake/releases)
- install via `go get`

```plain
go get -u https://github.com/cloudfoundry-community/cf-pancake
```

Local development
-----------------

As `cf-pancake exports` is designed to be run within an application container, if you try to run it locally then `$VCAP_SERVICES` will be missing. You can setup a local `$VCAP_SERVICES` (and `$VCAP_APPLICATION` is also required) for a Cloud Foundry application.

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
