#!/usr/bin/python
#usage python ppastats.py PPATEAM (ex: webupd8team) PPA (ex: gthumb) DIST (Ubuntu version eg maverick) ARCH (ubuntu arch eg i386 or amd64)

#example - highest downloaded file: python ppastats.py webupd8team y-ppa-manager maverick amd64 | tr '\t' ',' | cut -d ',' -f3 | sort -gr

import sys
from launchpadlib.launchpad import Launchpad
import sqlite3

PPAOWNER = "elementary-os"
PPA = "stable"
OSVERSION = "trusty"
ARCH = "amd64"

class DB:
       def __init__(self):
              db_loc = "prod.db"
              self.conn = sqlite3.connect(db_loc)
              self.rev = None

       def start_transaction(self):
              cursor = self.conn.cursor()
              try:
                     cursor.execute("SELECT MAX(rev) FROM dates")
                     rev = cursor.fetchone()[0]
                     if rev is None:
                            rev = 0
              except:
                     rev = 0

              self.rev = rev + 1

              cursor.execute('INSERT INTO dates(date, rev) VALUES(strftime("%Y%m%d"), ?)', (self.rev,))

       def store(self, name, download_val):
              cursor = self.conn.cursor()
              
              if self.rev is None:
                     self.start_transaction()

              cursor.execute("INSERT INTO data(rev, name, downloads) VALUES(?, ?, ?)", (self.rev, name, download_val, ))
              cursor.fetchone()
              self.conn.commit()

class Updater:
       def __init__(self, db_man):
              self.db_man = db_man

       def run(self):
              distro_uri = 'https://api.launchpad.net/devel/ubuntu/trusty/amd64'
              cachedir = "~/.launchpadlib/cache/"
              
              lp_ = Launchpad.login_anonymously('elementary-ppa-stats', 'edge', cachedir, version='devel')
              
              owner = lp_.people['elementary-os']
              archive = owner.getPPAByName(name='daily')

              self.db_man.start_transaction()

              for individualarchive in archive.getPublishedBinaries(status='Published', distro_arch_series=distro_uri):
                     x = individualarchive.getDownloadCount()
                     
                     if x > 0:
                            self.db_man.store(individualarchive.binary_package_name,  int(individualarchive.getDownloadCount()))
                     elif x < 1:
                            self.db_man.store(individualarchive.binary_package_name,  0)


def main():
       db = DB()
       updater = Updater(db)
       updater.run()

if __name__ == "__main__":
       main()
