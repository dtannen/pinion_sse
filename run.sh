#!/bin/sh
mkdir /etc/letsencrypt
cd /etc/letsencrypt
s3 get s3://pinionsse/PinionSSE.zip
unzip PinionSSE.zip
cd /go
/go/pinion_sse
