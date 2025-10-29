# i9w

[`vendor/libtai-0.60/`](vendor/libtai-0.60/) contains [libtai](https://cr.yp.to/libtai.html) patched to load data from `/usr/local/etc/leapsecs.dat`.
libtai is in the public domain.

## Build

```shell
# `/usr/local/etc/leapsecs.dat` is required.
(cd vendor/libtai/ || exit 1; make; sudo install -d /usr/local/etc/; sudo install -m 0644 leapsecs.dat /usr/local/etc/)
make
```
