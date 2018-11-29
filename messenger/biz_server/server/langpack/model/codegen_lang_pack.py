#!/usr/bin/python
#-*- coding: utf-8 -*-
#encoding=utf-8

'''
// Copyright (c) 2018-present,  NebulaChat Studio (https://nebula.chat).
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Author: Benqi (wubenqi@gmail.com)
'''

import glob, re, binascii, os, sys

'''
{ langPackString
key: "lng_about_done" [STRING],
value: "Done" [STRING],
},
'''

if (len(sys.argv) !=2):
    print('Input file required.')
    sys.exit(1)


def ToCamelName(name):
    ss = name.split("_")
    for i in range(len(ss)):
        if (i == 0):
            continue
        s = ss[i]
        ss[i] = s[0:1].upper() + s[1:]

    return ''.join(ss)

langPackStrings = []
langPackStringPluralizeds = []

with open(sys.argv[1]) as f:
    lastName = ''
    for line in f:
        line=line.strip()
        line=line.strip(',')
        if (line == '{ langPackString'):
            langPackStrings.append({'key':'', 'value':''})
            lastName = 'langPackString'
            continue

        if (line == '{ langPackStringPluralized'):
            langPackStringPluralizeds.append({'key':'', 'zero_value':'', 'one_value':'', 'two_value':'', 'few_value':'', 'many_value':'', 'other_value':''})
            lastName = 'langPackStringPluralized'
            continue

        if (line.find('flags:') >=0 ):
            continue

        if (line == '},'):
            continue

        idx = line.rfind('[')
        line = line[0:idx-1]

        if (line == ''):
            continue

        idx = line.find(':')
        s = []
        if (idx > 0):
            s = [line[0:idx], line[idx+2:]]
        else:
            continue

        if (len(s) == 0):
            continue

        s[1] = s[1].strip('"')

        #print s

        if (lastName == 'langPackString'):
            langPackStrings[-1][s[0]] = s[1]

        if (lastName == 'langPackStringPluralized'):
            langPackStringPluralizeds[-1][s[0]] = s[1]

    #print langPackStrings
    #print langPackStringPluralizeds
    #print len(langPackStrings) + len(langPackStringPluralizeds)

print'''# lang_pack_en.toml

langCode = "en"
version = 77
'''


for d in langPackStrings:
    print '[[Strings]]'
    for k, v in d.items():
        print('%s = "%s"' % (ToCamelName(k), v))

    print('')

for d in langPackStringPluralizeds:
    print '[[StringPluralizeds]]'
    for k, v in d.items():
        print('%s = "%s"' % (ToCamelName(k), v))

    print('')
