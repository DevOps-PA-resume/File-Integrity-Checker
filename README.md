# File-Integrity-Checker

https://roadmap.sh/projects/file-integrity-checker


## How to use

```bash
go mod init integrity-check
go mod tidy
go build -o integrity-check ./src
```

```bash
> ./integrity-check init /var/log  # Initializes and stores hashes of all log files in the directory
> Hashes stored successfully.

> ./integrity-check check /var/log/syslog
> Status: Modified (Hash mismatch)
# Optionally report the files where hashes mismatched

> ./integrity-check -check /var/log/auth.log
> Status: Unmodified

> ./integrity-check update /var/log/syslog
> Hash updated successfully.
```