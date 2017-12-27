# Keepup

[![Code Climate](https://codeclimate.com/github/ustrajunior/keepup/badges/gpa.svg)](https://codeclimate.com/github/ustrajunior/keepup)
[![Build Status](https://travis-ci.org/ustrajunior/keepup.svg?branch=master)](https://travis-ci.org/ustrajunior/keepup)

Keepup is a command line utility to update your records on [Cloudflare](https://www.cloudflare.com/). 

## Instalation

Just run 

```
go get -u github.com/ustrajunior/keepup
```

## Configuration

By default, Keepup will use the file: **$HOME/.keepup.toml(yaml|json)** but you can customize this passing the option --config="/path/to/file.toml". The config file will contain your credentials to your Cloudflare account.

The .keepup.toml has the following format:

```
default = "example"

[example]
  domain = "example.com"
  cfKey = "123456"
  cfEmail = "your@email.com"
  netInter = "en0"

[other]
  domain = "other.com"
  cfKey = "67890"
  cfEmail = "other@email.com"
  netInter = "en0"
```

## Using

To set a new value for your DNS record, you have to pass the **dns** flags when calling the Keepup command on the following formats.

```
keepup update --dns subdomain
keepup update --dns example.com
keepup update --dns subdomain.example.com
```

**dns** is the DNS record you want to update. It could be a subdomain or the main domain.

It will use the ip of the interface you setup on the netInter key on configuration file. If you want set a custom ip, just use the **ip** flag.

```
keepup update --dns subdomain --ip 127.0.0.1
```

To use a separated account, use the **account** flag. 

```
keepup update --dns subdomain --account other
```

## License

Copyright (c) 2016-present [José Carlos Ustra Júnior](https://github.com/ustrajunior)

Licensed under [MIT License](https://github.com/ustrajunior/keepup/blob/master/LICENSE)
