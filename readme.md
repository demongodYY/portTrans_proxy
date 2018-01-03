# 端口轉發
## 使用方法：
- 服務器模式： tr.exe -s lServerPort rServerPort
- 客戶端模式（需要與服務器模式配合使用）： tr.exe -c rServerIp:port targetIp:port
- 通過http代理客戶端模式（需要與服務器模式配合使用）： tr.exe -pc rServerIp:port targetIp:port proxyIp:port
- 傳輸轉發模式：tr.exe -t lPort targetIp:port