# protocel-example

This repository is an example of [protocel](https://github.com/Neakxs/protocel) tools usage.

## Requirements

For testing this project, you will need :

- go, for building the server
- git, for cloning repositories
- [grpcurl](https://github.com/fullstorydev/grpcurl) (or another tool) for interacting with the gRPC API

After cloning the repository, you can simply launch or build the server located in `cmd/server`.

## Examples

> The following examples uses the `gcurl='grpcurl -import-path submodules/protocel -import-path submodules/googleapis/ -import-path . -proto proto/library/v1/library_service.proto -plaintext'` alias

- CreateAuthor with wrong credentials

```bash
$ gcurl -H 'x-api-key: wrongsecret' -d '{}' 127.0.0.1:8765 library.v1.LibraryService.CreateAuthor
ERROR:
  Code: PermissionDenied
  Message: permission denied on "/library.v1.LibraryService/CreateAuthor"
```

- CreateAuthor without payload
  
```bash
gcurl -H 'x-api-key: mysecret' -d '{}' 127.0.0.1:8765 library.v1.LibraryService.CreateAuthor
ERROR:
  Code: Unknown
  Message: validation failed on library.v1.CreateAuthorRequest.author
```

- CreateAuthor with empty author

```bash
$ gcurl -H 'x-api-key: mysecret' -d '{"author": {}}' 127.0.0.1:8765 library.v1.LibraryService.CreateAuthor
ERROR:
  Code: Unknown
  Message: validation failed on library.v1.Author.display_name
```

- CreateAuthor without birth_date

```bash
$ gcurl -H 'x-api-key: mysecret' -d '{"author": {"display_name": "Victor Hugo"}}' 127.0.0.1:8765 library.v1.LibraryService.CreateAuthorERROR:
  Code: Unknown
  Message: validation failed on library.v1.Author.birth_date
```

- CreateAuthor with all fields completed

```bash
$ gcurl -H 'x-api-key: mysecret' -d '{"author": {"display_name": "Victor Hugo", "birth_date": "1802-02-26T00:00:00Z", "death_date": "1885-05-22T00:00:00Z"}}' 127.0.0.1:8765 library.v1.LibraryService.CreateAuthor
{
  "name": "authors/5c543b25-ee5d-433d-99b9-241be242906e",
  "displayName": "Victor Hugo",
  "birthDate": "1802-02-26T00:00:00Z",
  "deathDate": "1885-05-22T00:00:00Z"
}
```

- ListAuthors

```bash
$ gcurl -d '{}' 127.0.0.1:8765 library.v1.LibraryService.ListAuthors
{
  "authors": [
    {
      "name": "authors/5c543b25-ee5d-433d-99b9-241be242906e",
      "displayName": "Victor Hugo",
      "birthDate": "1802-02-26T00:00:00Z",
      "deathDate": "1885-05-22T00:00:00Z"
    }
  ],
  "nextPageToken": "14"
}
```