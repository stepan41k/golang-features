#bin/sh

find /var/log -type f -size +100M -mtime +7

# find /var/log -type f -size +100M -mtime +7 -delete