bounceback
----------

connect to private spaces heroku postgres databases from peered vpcs


[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/heroku/bounceback)


Deploying this app will create a single `bounceback` dyno in your space that exposes a pgbouncer/stunnel server on your dyno private network.
This pgbouncer instance is accessible from VPCs peered to your space.

By setting the BOUNCEBACK_URLS config var, you may configure a comma separated list of endpoints that reside inside your peered VPC that accept a json post anytime the IP address or credentials
for the pgbouncer change, which is every time the web dyno running pgbouncer is restarted. 

*NOTE* please assure the endpoints you configured are secure, by either having them reside in your peered VPC or by using TLS + basic auth.

The endpoints will receive a payload with the correct _PGBOUNCER url for each database attached to the app, for example:

```
{
"DATABASE_URL_PGBOUNCER":"postgres://user:pass@private.ip:6000/databaseName",
"OTHER_DATABASE_URL_PGBOUNCER":"postgres://user:pass@private.ip:6000/otherDatabaseName"
}
```