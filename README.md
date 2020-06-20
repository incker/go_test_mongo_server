# Golang test API server with mongo database

### Installation

Create config

```bash
touch configs/apiserver.toml
```

Config file can be empty, as it has default values, but you can add custom values something like:

```bash
bind_addr = ":8080"
log_level = "debug"
mongodb_url = "mongodb://localhost:27017"
mongodb_name = "testing"
```

Build server

```bash
make
```

Run tests

```bash
make test
```


### Fill in database from `users_go.json` file
 1. Download file provided in test task: [users_go.json](https://drive.google.com/file/d/1tjubsoSwdzPK553ovvmMZs9qQwMjlKh1/view)
 2. Delete some data in `users_go.json` file, so it will start with `[` and end with `]`
 3. Import file data: `mongoimport --db testing --collection user --file users_go.json --type json --jsonArray --batchSize 1`
 4. convert birth_date from string to date
 ```js
db.user.find().forEach((el) => {
    el.birth_date = new Date(el.birth_date);
    db.user.save(el)
});
```
 5. Create index `db.user.createIndex( { "email": 1 }, { unique: true } )`
 