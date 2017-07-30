#!/usr/bin/python
import markovify
import string

with open("corpus.txt") as f:
	text = f.read()

model = markovify.NewlineText(text)

with open("fakedeaths.txt",'w') as outfile:
    for i in range(1000):
		outfile.writelines(str(model.make_sentence())+"\n")