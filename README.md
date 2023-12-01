# The Anatomy of a Linux Container using Go
[Basic Anatomy of a Linux Container](https://www.youtube.com/watch?v=jeTKgAEyhsA&t=1179s) taught by Liz Rice, Aqua Security.

Source: https://github.com/alpinelinux/alpine-make-rootfs  <br />

## Download alpine-make-rootfs and set the executable.
```console
wget https://raw.githubusercontent.com/alpinelinux/alpine-make-rootfs/v0.7.0/alpine-make-rootfs \
  && echo 'e09b623054d06ea389f3a901fd85e64aa154ab3a  alpine-make-rootfs' | sha1sum -c \
  || exit 1
```
```console
chmod +x alpine-make-rootfs
```

## Command used to create an alpine linux file system
```console
sudo ./alpine-make-rootfs --branch v3.18 --packages 'python3 ruby' --timezone 'America/New_York' --script-chroot alpinefs-$(date +%Y%m%d).tar.gz - <<'SHELL'
apk add --no-progress -t .make build-base
apk del --no-progress .make
SHELL
```
```console
mkdir alpinefs && mv alpinefs-20231201.tar.gz alpinefs/
```
```console
cd alpinefs && tar zxvf alpinefs-20231201.tar.gz
```
You now have an alpine filesystem (version 3.18.5) that you can use to chroot to!  <br />
This is required in order to follow Liz's video through its entirity.

Additional Resources:
https://github.com/rootless-containers/rootlesskit
