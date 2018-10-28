package main

import (
        "fmt"
        "net/http"
        "net"
        "time"
        "os"
        "io/ioutil"
        "strings"
)


func main() {
    const TIME_INTERVAL int = 15*60
    const DOMAIN_NAME string = "<DomainName>"
    const DOMAIN_SUBNAME string = "<SubDomainName>"
    const DOMAIN_TYPE string = "A"
    const DOMAIN_TTL int = 900
    const GoDaddy_Key string = "<APIKey>"
    const GoDaddy_Sec string = "<APISec>"

    fmt.Println("Starting the application...")
    for true  {
        ip1 :=   getPublicIP()
        ip2 :=   getDNSRecordIP(DOMAIN_SUBNAME + "." + DOMAIN_NAME)
        // ip2 = "0.0.0.0"
        formatMsg("PublicIP : " + ip1)
        formatMsg("DNSRecord : " + DOMAIN_SUBNAME + "." + DOMAIN_NAME + ":" + ip2)
        
        if len([]rune(ip1)) > 15 {
        
            formatMsg("Get PublicIP Error ")
            
        } else if ip2 == "0.0.0.0" {
            formatMsg("Get DNSRecord Error ")

        } else if ip1 == ip2 {
            formatMsg("No Need For Update Domian Record ")

        }else{

            formatMsg("Ready To Update Domian Record ...")
        	setDNSRecordIP(ip1,DOMAIN_TTL,DOMAIN_NAME,DOMAIN_TYPE,DOMAIN_SUBNAME,GoDaddy_Key,GoDaddy_Sec)
        }
        

        time.Sleep(time.Duration(TIME_INTERVAL) * time.Second)

    }
}

func getPublicIP() string {
    var ip = "0.0.0.0"
    response, err := http.Get("http://ipv4.icanhazip.com/")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        ip = string(data)
       
    }
    return strings.TrimSpace(ip)
}

func getDNSRecordIP(domainName string) string {
    var ip = "0.0.0.0"
	domain := domainName
	ipAddr, err := net.ResolveIPAddr("ip", domain)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return ip
	}
	ip = ipAddr.IP.String()
// 	fmt.Println(ip)
    return strings.TrimSpace(ip)
}

func setDNSRecordIP(ip string, ttl int, domain string, domainType string, name string, key string, sec string) {
    url := fmt.Sprintf("https://api.godaddy.com/v1/domains/%s/records/%s/%s",domain,domainType,name)
    data :=fmt.Sprintf("[{ \"data\": \"%s\", \"ttl\": %v, \"priority\": 0, \"weight\": 1 }]",ip,ttl)
    //生成client 参数为默认
    client := &http.Client{}
    //提交请求
    reqest, err := http.NewRequest("PUT", url, strings.NewReader(data))
    if err != nil {
        //panic(err)
        fmt.Printf("The HTTP request failed with error %s\n", err)
		return 
    }
    reqest.Header.Add("content-type", "application/json")
    reqest.Header.Add("Accept", "application/json")
    ssokey :=fmt.Sprintf("sso-key %s:%s",key,sec)
    reqest.Header.Add("Authorization", ssokey)
    //处理返回结果
    response, err := client.Do(reqest)
   //返回的状态码
    if err != nil {
        //panic(err)
        fmt.Printf("The HTTP request failed with error %s\n", err)
		return 
    }
    status := response.StatusCode
    if status == 200 {
        formatMsg("Done!")
    }else{
        formatMsg("Error Hapened!")
    }
    
}

func formatMsg(msg string){
	formatTimeStr:=time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(formatTimeStr + " : " +  msg)
}

