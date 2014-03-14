# !/usr/bin/python

import os

fo = open("foo.txt", "w+")

for root, dirs, files in os.walk(".", topdown=False):
    for name in files:
        fo.write(os.path.join(root, name)+"\n")
    for name in dirs:
        fo.write(os.path.join(root, name)+"\n")



