package main

import (
    "fmt"
      "github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
)

func main() {
    
    //构建一个阿里云客户端, 用于发起请求。
    //设置调用者（RAM用户或RAM角色）的AccessKey ID和AccessKey Secret。
    client, err := sts.NewClientWithAccessKey("cn-hangzhou", "LTAI5tLYcTEYLuJccbb8uYi5", "suSKAUr09RUmFNBL4M4IAqZ9KxIAqh")
    
    //构建请求对象。
    request := sts.CreateAssumeRoleRequest()
    request.Scheme = "https"
    
    //设置参数。关于参数含义和设置方法，请参见《API参考》。
    request.RoleArn = "acs:ram::1124738755824567:role/djappsmartlock"
    request.RoleSessionName = "chenxiaojun"
    
    //发起请求，并得到响应。
    response, err := client.AssumeRole(request)
    if err != nil {
        fmt.Print(err.Error())
    }
    fmt.Printf("response is %#v\n", response.Credentials)
}