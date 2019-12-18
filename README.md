# Email Service

## Postgres database setup

### Create a postgres database and user

**Prerequisite:**
Postgres installed, Postgres started

Connect to the database via the postgres admin user.
```
sudo -u postgres psql
```

Create a new user
```
CREATE USER stefan
```

Create a new database owned by the user (use your linux user name)
```
CREATE DATABASE email_service OWNER stefan
```

Quit the postgres CLI
```
/q
```

Connect to the database with the new user
```
psql -d email_service
```
