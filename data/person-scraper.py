#!/usr/bin/python
import json
import os
with open('person-original.json') as data_file:    
    data = json.load(data_file)
d = {}
i = []
j = 0
print len(data)
for person in data:
	if 'year' in person['inquest']:
		death_year = int(person['inquest']['year'])
	else:
		continue
	if 'birth' in person and 'year' in person['birth']:
		start_year = int(person['birth']['year'])
	elif 'convict' in person and 'year' in person['convict']:
		start_year = int(person['convict']['year'])
	elif 'immigration' in person:
		start_year = int(person['immigration']['year'])
	else:
		continue	
	if death_year < start_year:
		continue
	if 'marriage' in person and (start_year > int(person['marriage']['year']) or int(person['marriage']['year']) > death_year):
		continue
	if 'bankruptcy' in person and (start_year > int(person['bankruptcy']['year']) or int(person['bankruptcy']['year']) > death_year):
		continue
	if 'census' in person and (start_year > int(person['census']['year']) or int(person['census']['year']) > death_year):
		continue
		if 'health-welfare' in person and (start_year > int(person['health-welfare']['year']) or int(person['health-welfare']['year']) > death_year):
			continue
	j = j+1
	i.append(person)
print j

with open('person.json', 'w') as outfile:
    json.dump(i, outfile, indent=4)		
