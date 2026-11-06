CSH Plug
========

Plug is the advertisement platform of the future.

Development
-----------

## Running 
It is likely easier to use containers like `podman` or `docker` or the corresponding compose file

With podman, I have been using 

```sh
podman compose up --watch
```

If you want, you can run without auto rebuild using
```sh
podman compose up --force-recreate --build
```
Which can be restarted every time changes are made.

## Config
To easily configure your workspace, copy the `.env.example` into a new `.env` file.

### Add OIDC Config
Reach out to an RTP to get OIDC credentials that will allow you to develop locally behind OIDC auth

You can also fetch dev credentials from the wiki (https://wiki.csh.rit.edu/wiki/SSO).

The `CSH_AUTH_STATE` and `CSH_AUTH_CLIENT_SECRET` are each a unqiue sequence of 16+ random characters, I recommend using the openssl command line utilities `openssl rand -hex 16`

```bash
CSH_AUTH_CLIENT_ID=
CSH_AUTH_CLIENT_SECRET=
CSH_AUTH_STATE=
CSH_AUTH_JWT_SECRET=
CSH_AUTH_SERVER_HOST=sso.csh.rit.edu
CSH_AUTH_REDIRECT_URI=http://localhost:8080/auth/redir
```

### Add LDAP Config
Authenticate with LDAP using your CSH account credentials by following the instructions on the wiki (https://wiki.csh.rit.edu/wiki/LDAP).

Set the `LDAP_WRITE_MODE` environment variable to `false` to prevent attempted writes to LDAP, it will default to false if unset for safety.

```bash
LDAP_BIND_DN=uid=<username>,cn=users,cn=accounts,dc=csh,dc=rit,dc=edu
LDAP_BIND_PW=(Your CSH account password)
```

### Add S3 Config
An S3 bucket is used to store images that users upload. In order to have this work properly, you need to provide some credentials to the app.

There are 2 ways that you can get the needed credentials. 
1. Reach out to an RTP for creds to the dev bucket
2. Create your own bucket using [DEaDASS](https://deadass.csh.rit.edu/), and the site will give you the credentials you need. 

```bash
S3_ACCESS_ID=
S3_SECRET_KEY=
S3_BUCKET_NAME=
S3_HOST=s3.csh.rit.edu
```

### Database
You can either develop using the dev database, or use the local database provided in the docker compose file

Using the local database is detailed below, but both options will require the dev database password, so you will have to ask an RTP for this too

## Testing Locally

### Admin Dashboard Access
Modify the `CheckIfAdmin` function in `ldap.go` to always return true which will grant you admin access.

### Drink Credits
Modify the `DecrementCredits` function in `ldap.go` to give yourself credits if your account doesn't have them already.
