bounceback
----------

connect to private spaces heroku postgres databases from peered vpcs


[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/heroku/bounceback)


Deploying this app will create a single web dyno in your space that also exposes a pgbouncer/stunnes server on your dyno private network
this pgbouncer instance is accessible from VPCs peered to your space.

You may configure a comma seperated list of endpoints that reside inside your peered VPC that accept a json post anytime the IP address or credentials
for the pgbouncer change, which is every time the web dyno running pgbouncer is restarted.

The endpoints will receive a payload with the correct _PGBOUNCER url for each database attached to the app.

```
{
"DATABASE_URL_PGBOUNCER":"postgres://user:pass@private.ip:6000/databaseName",
"OTHER_DATABASE_URL_PGBOUNCER":"postgres://user:pass@private.ip:6000/otherDatabaseName"
}
```