# gosync
a tool for sync ro download files

# build
windows 
go build -ldflags="-H windowsgui"

# default config
```
[system]
loopsec = 1	#file monitor loop internal
address = 0.0.0.0:3030 #sync address
token = #config html ui token
uiaddr = 0.0.0.0:2020 #html ui address

[httpdownload1]	#httpdown config
dir = .
address = 0.0.0.0:2021

[sync1] #sync config
dir = D:\myproject=>/home/xxx/myproject@192.168.1.111:3030
include = .*\.(cpp|h|go)
exclude = .*
completex = build #ignore file or dir
```
# config
![config](https://github.com/gotask/images/blob/master/gosync.jpg?raw=true)