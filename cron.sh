#!/bin/sh

cd /home/voldyman/code/e-stats/
python ppastats.py

echo "PPA stats updated" >> /tmp/email-data
echo  "New Records: " `sqlite3 prod.db 'SELECT COUNT(*) FROM data WHERE rev=(SELECT MAX(rev) FROM dates)'` >> /tmp/email-data

echo "Most Downloaded: " `sqlite3 prod.db 'SELECT name, MAX(downloads) FROM data WHERE rev=(SELECT MAX(rev) FROM dates)'` >> /tmp/email-data

cat /tmp/email-data | mail -s "daily update" akshay.is.gr8@gmail.com

rm /tmp/email-data
