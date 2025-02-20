package main

import "fmt"

func main() {
        var sks, i, nilai, hasil, komulatif, ips, hasil2 int
        var ip string
        fmt.Scan(&i)
        hasil2 = 0
        komulatif = 0
        for i > 0{
                fmt.Scan(&ip, &sks)
                if ip == "A"{
                        nilai = 4
                }else if ip == "B"{
                        nilai = 3
                }else if ip == "C"{
                        nilai = 2
                }else if ip == "D"{
                        nilai = 1
                }else if ip == "E"{
                        nilai = 0
                }
                if sks > 0 && (ip == "A" || ip == "B" || ip == "C" || ip == "D" || ip == "E"){
                        komulatif = komulatif + nilai
                        hasil = nilai*sks
                        hasil2 = hasil2 + hasil
                        i = i - 1
                }
                
        }
        ips = hasil2/komulatif
        fmt.Print(ips)
}