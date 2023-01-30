# 代码位置
代码放在了utils/jwt.go

# token生成
* 使用"github.com/dgrijalva/jwt-go"包生成。
* token字符串中主要包含几个信息：用户ID， 签发人，过期时间。

涉及到几个参数
```
secret: "secret" // 密钥。暂时使用secret
insuer: "Frogs"  // 签发人。暂时使用Frogs， 这个参数是我为后续做token校验时候做准备
TokenExpireDuration: "" // token有效时间。暂时使用两个小时
```

# token解析
调用解析方法即可将token字符串解析成对象指针