# Gen RSA key
>
>  telegram客户端使用的是pkcs1格式公钥证书
>
>  参考
>
>  [rsa秘钥介绍及openssl生成命令](https://medium.com/@oyrxx/rsa%E7%A7%98%E9%92%A5%E4%BB%8B%E7%BB%8D%E5%8F%8Aopenssl%E7%94%9F%E6%88%90%E5%91%BD%E4%BB%A4-d3fcc689513f)
>

## Genkey

```
    openssl genrsa -out server.key 2048

    convert pcks8:
    openssl pkcs8 -topk8 -inform PEM -in server.key -outform pem -nocrypt -out server8.key
    openssl rsa -in server.key -pubout > public_pkcs8.pub

    convert pcks1:
    openssl rsa -in server.key -outform PEM -RSAPublicKey_out -out public_pkcs1.key
```


## gen fingerprint
./fingerprint key
