# AutoUpdateGoDaddyDNS

最近网络环境趋紧，流行的几个DDNS要么要实名认证，要么直接就不能免费使用，很烦。一个很简单的东西，天朝居然都要搞得这么复杂，最近无聊顺手些了一个，
用Python以及GoLang实现动态域名解析。

大概的流程

   1，通过调用一个远程网站，获得本地的公网IP

   2，确认域名的指向与1获得的IP是否一致，如果一致就不继续，如果不一致前进到3

   3，通过调用GoDaddy的API重置特定域名的IP指向。

每隔30分钟，循环上述流程，实现DDNS

使用方式，
  嘿嘿目前不支持配置文件及参数，直接把开头的变量改一下就行，
  比如GoLang的
    const TIME_INTERVAL int = 15*60     //循环间隔
    const DOMAIN_NAME string = "<DomainName>"     //域名
    const DOMAIN_SUBNAME string = "<SubDomainName>"     //你想要动态设定的子域名
    const DOMAIN_TYPE string = "A"            
    const DOMAIN_TTL int = 900                    //域名的TTL
    const GoDaddy_Key string = "<APIKey>"         //GoDaddy的API
    const GoDaddy_Sec string = "<APISec>"
 通过build.sh编译一下，Linux，windows一次编译生成执行文件。。。。舒爽   
 
 Python，直接修改前面的设定之后 python AutoUpdateDNS.py执行
 
 以上。
