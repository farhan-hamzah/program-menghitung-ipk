i = int(input())
hasil2 = 0
komulatif = 0
for _ in range(i):
    ip, sks = input().split()
    sks = int(sks)
    nilai_dict = {"A": 4, "B": 3, "C": 2, "D": 1, "E": 0}
    nilai = nilai_dict.get(ip, -1)
    if sks > 0 and nilai != -1:
        komulatif += nilai
        hasil = nilai * sks
        hasil2 += hasil
if komulatif > 0:
    ips = hasil2 / komulatif
    print(ips)
else:
    print("Error: Tidak ada SKS yang valid")