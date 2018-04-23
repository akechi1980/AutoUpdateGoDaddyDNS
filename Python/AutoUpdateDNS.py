#!/usr/bin/python
# -*- coding: UTF-8 -*-
# 本地动态获得公网IP并且重置Godaddy域名设定的例子
import os, socket
# pip install requests
import requests

args = {
  "DomainName":"<DomainName>",
  "DomainType":"A",
  "DomainApikey":"<APIKey>",
  "DomainApiSec":"<APISec>",
  "TTL":3600
}


# 处理主函数
def main():
  # 获取当前域名指向IP
  # default_ip = socket.gethostbyname('w1.sz-ming.com')
  default_ip = socket.gethostbyname(args["DomainName"])
  print "当前指向IP : " + default_ip
  # 测试用
  default_ip = "127.0.0.1"
  
  ip = "127.0.0.1"
  # 获取本地的IP地址
  try:
    f = requests.get('http://ipv4.icanhazip.com/')
    ip = f.text.strip()
    print "获得公网IP : " + ip.encode('utf-8') 
  #   args.ip = resp.strip()
  except Exception,err:
    print 1,err
  
  # 区分处理
  if default_ip == ip:
    print "无需处理，直接退出"
    # 直接退出
    os._exit(0)
  else:
    print "正在重新设定域名指向"
  
  
  hostnames = args["DomainName"].split('.')
  
  # 通过API重新设定
  url = 'https://api.godaddy.com/v1/domains/{}/records/{}/{}'.format('.'.join(hostnames[1:]),args["DomainType"],hostnames[0])
  print url
  data ='[{ "data": "' + ip + '", "ttl": 3600, "priority": 0, "weight": 1 }]'
  print data
  headers = {
      'content-type': 'application/json',
      "Accept":"application/json",
      "Authorization": "sso-key {}:{}".format(args["DomainApikey"],args["DomainApiSec"])
  }
  try:
    req = requests.put(url, data=data,headers=headers)
    print "动态设定完成！"
    print req.text
  except Exception,err:
      print 1,err

# 主函数调用
if __name__ == '__main__':
  main()

# 其他API的直接访问例子
# 重新设定域名指向
# curl -H 'Authorization: sso-key dKDGMgeCfUyz_3CBGRJw9pDXcHz5RsPLDK7:3CBJmugDCnV9M4LVsjSm53' -H 'Content-Type: application/json' https://api.godaddy.com/v1/domains/sz-ming.com/records/CNAME/home
# curl -H 'Authorization: sso-key dKDGMgeCfUyz_3CBGRJw9pDXcHz5RsPLDK7:3CBJmugDCnV9M4LVsjSm53' -H 'Content-Type: application/json' https://api.godaddy.com/v1/domains/sz-ming.com/records/A/
# curl -X PUT "https://api.godaddy.com/v1/domains/sz-ming.com/records/A/w1" -H "accept: application/json" -H "Content-Type: application/json" -H "Authorization: sso-key dKDGMgeCfUyz_3CBGRJw9pDXcHz5RsPLDK7:3CBJmugDCnV9M4LVsjSm53" -d "[ { \"data\": \"192.168.0.1\", \"port\": 0, \"priority\": 0, \"protocol\": \"string\", \"service\": \"string\", \"ttl\": 3600, \"weight\": 0 }]"
# curl -X PUT "https://api.godaddy.com/v1/domains/sz-ming.com/records/A/w1" -H "accept: application/json" -H "Content-Type: application/json" -H "Authorization: sso-key dKDGMgeCfUyz_3CBGRJw9pDXcHz5RsPLDK7:3CBJmugDCnV9M4LVsjSm53" -d "[ { \"data\": \"192.168.0.1\", \"priority\": 0,  \"ttl\": 3600, \"weight\": 0 }]" 
