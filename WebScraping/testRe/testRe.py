import re 

s = "I'am python modules test for re modules"
print(re.search("am",s))
print(re.search("am",s).group())
print(re.findall("python",s))