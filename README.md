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

There are two usage modes:

-	`cf-pancakes export` - returns bash `export` commands to setup flatten variables
-	`cf-pancakes set-env APPNAME` - updates the `APPNAME` with environment variables from that app's `$VCAP_SERVICES`

The former would be run within an application container during startup.

The latter would be run by the developer outside of Cloud Foundry.
