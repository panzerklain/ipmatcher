# ipmatcher
Simple tool for matching ip address with list of ipnets/ipaddresses

## usage
```
./ipmatcher ip_for file_with_list_for_match
```

## example
Create a file with a list of ipaddress/ipnets which the entered ip address will be compared.
For example, whitelist.txt:
```
192.168.0.0/16
172.30.0.1
```

Run ipmatcher:
```
./ipmatcher 172.30.0.1 whitelist.txt
```
Return code is 0 if matched, else 214.
