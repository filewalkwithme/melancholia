melancholia
===========

Melancholia enshrines all triumph.

Prepare your database
---------------------

Melancholia is using PostgreSQL.

Create your database

```
createdb melancholia
```

Import the *data.sql* file

```
psql melancholia -f data.sql
```

If you feel like testing
------------------------

```
createdb melancholia_test
```

Also import the *data.sql* file

```
psql melancholia_test -f data.sql
```

And go inside the models (the only one tested so far) folder and run:

```
go test
```

Done ;)