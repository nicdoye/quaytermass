# Quaytermass - a quay repo deleter

## Intro

Deletes repos from Quay. You need to set the environment variable `API_ACCESS_TOKEN`
matching someone with `repo:admin` permissions.

I fought with this for ages, and the only authentication method I got to work was by
using [Quay's API browser](https://docs.quay.io/api/swagger/#!/repository/deleteRepository),
and grabbing the bearer token from Chrome's Developer Tools. There must be a better way.

## Running

```bash
go run main.go namespace/repo1 namespace/repo2
```

Or if you have a list of them stored in a file.

```bash
go run main.go $(cat /tmp/repos| xargs)
```

## Author

Nic Doye - _2018.06.22_