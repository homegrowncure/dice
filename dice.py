#!/usr/bin/env python

# Usage: dice.py [numwords=6]

import sys

WORDLIST = 'dicewords'

def roll():
    with open('/dev/urandom') as rand:
        return ord(rand.read(1)) % 6 + 1

def word():
    combo = []
    for x in range(5):
        combo.append(roll())

    with open(WORDLIST) as words:
        lines = words.readlines()
        for l in lines:
            if l[:5] == ''.join(str(x) for x in combo):
                return l[6:len(l) - 1]

for x in range(int(sys.argv[1]) if len(sys.argv) > 1 else 6):
    print word(),
