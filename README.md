# go-uptime

Dirty simple uptime monitor for self hosting enthusiasts

## How to use 

- Clone the repo and update the credentials
- Add the endpoints to be checked inside endpoints.txt (one per line)
- Build

  ```
  go build -o go-uptime main.go
  ```

## Periodically check

Setup a crontab job

```
*/5 * * * * /path/to/go-uptime >> /path/to/logfile.log 2>&1
```


## TO DO

- [ ] Move credentials to environment variables or .env files 
