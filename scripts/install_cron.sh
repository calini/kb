#!/bin/bash

# example: ./install_cron.sh e-c07kilf16{01..52}.it.manchester.ac.uk

for hostname in "$@"
do
   echo "* * * * * bash -c \"w -h > ~/kb/\$HOSTNAME.log\"" | ssh -o "StrictHostKeyChecking no" $hostname "crontab -"
done
