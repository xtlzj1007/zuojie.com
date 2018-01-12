import csv

dic = {
    "a":"123",
    "b":"456",
    "c":"789",
}

f = open("t.txt","w")
dname = ["a","b","c"]
wr = csv.DictWriter(f,fieldnames=dname)
wr.writerow(dic)
f.close()
