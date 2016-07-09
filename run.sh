#!/bin/sh
mkdir /etc/letsencrypt
cd /etc/letsencrypt
s3 get s3://pinionsse/PinionSSE.zip
unzip PinionSSE.zip
mv /etc/letsencrypt/live/sse.pinion.site/fullchain.pem /etc/letsencrypt/live/sse.pinion.site/fullchain.pem.old
ln -s /etc/letsencrypt/archive/sse.pinion.site/fullchain1.pem /etc/letsencrypt/live/sse.pinion.site/fullchain.pem
mv /etc/letsencrypt/live/sse.pinion.site/cert.pem /etc/letsencrypt/live/sse.pinion.site/cert.pem.old
ln -s /etc/letsencrypt/archive/sse.pinion.site/cert1.pem /etc/letsencrypt/live/sse.pinion.site/cert.pem
mv /etc/letsencrypt/live/sse.pinion.site/chain.pem /etc/letsencrypt/live/sse.pinion.site/chain.pem.old
ln -s /etc/letsencrypt/archive/sse.pinion.site/chain1.pem /etc/letsencrypt/live/sse.pinion.site/chain.pem
mv /etc/letsencrypt/live/sse.pinion.site/privkey.pem /etc/letsencrypt/live/sse.pinion.site/privkey.pem.old
ln -s /etc/letsencrypt/archive/sse.pinion.site/privkey1.pem /etc/letsencrypt/live/sse.pinion.site/privkey.pem
certbot renew
cd /go
/go/pinion_sse
