FROM busybox:1.27-glibc

# first, download and extract (or build rsc locally)
# https://binaries.rightscale.com/rsbin/rsc/v6/rsc-linux-amd64.tgz
COPY rsc/rsc /usr/local/bin/rsc

WORKDIR /usr/local/bin

CMD ["--help"]

ENTRYPOINT ["rsc"]
