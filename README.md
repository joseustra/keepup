# Keepup

Keepup is a command line utility to update your records on [Cloudflare](https://www.cloudflare.com/). It will keep the last IP you updated, so will not be necessary update the record every single time.

## Instalation

Just run 

```
go get -u github.com/ustrajunior/keepup
```

## Configuration

By default, Keepup will use the file: **$HOME/.keepup.yaml** but you can customize this passing the option --config="/path/to/file.yaml". The config file will contain your credentials to your Cloudflare account.

The .keepup.yaml has the following format:

```
cfKey: "your cloudflare key"
cfEmail: "your cloudflare email"
```

## Using

To set a new value for your DNS record, you have to pass the **zone** and **dns** flags when calling the Keepup command.

```
keepup update --zone="domain.com" --dns="sub.domain.com"
```

the **zone** is the domain you want to update, if you have more than 1 domain on your account, choose the one you want to work on.

**dns** is the DNS record you want to update. It could be a subdomain or the main domain.

Passing only this two flags, Keepup will update your DNS record with your current public IP. To set a specify IP to update, you have to pass the **ip** flag, like this:

```
keepup update --zone="domain.com" --dns="sub.domain.com" --ip="127.0.0.1"
```

All your domains with IPs will be stored in the file **$HOME/.keepup/keepup.db**. The IP on this file will be checked before the update occurred on your Cloudflare account, and if the IP you are passing is the same on the file, the update will not happen.

To force the update, you need to pass the **force** flag.

```
keepup update --zone="domain.com" --dns="sub.domain.com" --force
```

This way, the IP will be updated and the new IP will be saved on the local file.

## License

Copyright (c) 2016-present [José Carlos Ustra Júnior](https://github.com/ustrajunior)

Licensed under [MIT License](https://github.com/ustrajunior/keepup/blob/master/LICENSE)
