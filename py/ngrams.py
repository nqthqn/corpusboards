#!/usr/local/bin/python3
"""
    NOTE: This is being converted to GO! See main.go
    
    Usage: ./bigrams.py text.txt > output.txt

    This takes the content of a file converts it to a list of characters
    and then iterates over the list and builds a dictionary of all the bigrams

    Example

    Input:
      aabbaa

    Output:
      aa 2
      ab 1
      bb 1
      ba 1
"""

import sys

bigrams = {}

with open(sys.argv[1]) as fileobj:

    chars = []

    for line in fileobj:
        chars += [ch for ch in line]

    chlen = len(chars) - 1

    for i in range(chlen):
        if i == chlen:
            break
        key = (chars[i], chars[i + 1])
        if bigrams.get(key):
            bigrams[key] += 1
        else:
            bigrams[key] = 1

for (k, v) in bigrams.items():
    a = k[0].encode('unicode_escape').decode('utf-8')
    b = k[1].encode('unicode_escape').decode('utf-8')
    print("{a}{b} {count}".format(a=a, b=b, count=v))