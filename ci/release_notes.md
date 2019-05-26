* env vars geneated with the service instance name as the prefix (upper cased, with underscores instead of hyphens etc)

    For example, `cf create-service p-mysql 10mb mydb` will generate `MYDB_` prefixes:

    ```plain
    export MYDB_USERNAME='rR8QDBUvwvdkZAwJ'
    export MYDB_HOSTNAME='10.144.0.17'
    export MYDB_PASSWORD='LhggDEV5ObpQz8PH'
    ```