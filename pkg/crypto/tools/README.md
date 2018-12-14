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
> fingerprint.go

```
./fingerprint test.key 

N:  bca2c43964f3b7d1e7dfff4a769fd174770487399df315de2d2a47208cda5d32c90f0f01849cb58d1fe2a9e1bc25ee72aed55a6ea312900ea5b48a60ca51fffff1688ccb17d411eee043d8397420074a8e8ba92bd3c8976481fdfe238f40e583b0bf8bb7c8031b4c41cbeb0f7bfd991ddcca3235fa3bd078b0eb318c5ae4e6a0e8583ae2a09a2b009ede1407cfa4e05fdb0ef7a215ee752ac913495b43ca4258da4c63c701f62f2bf96062b5cbe8b8b0c0be6b674d7eda921a03ce62a0a49058962018e2a03bdefeeee5421ea44f10815d2308e8712423ee6cff1d83efcf94b2d52b2c54e4276242d663d84332e2cf7194d2b35fc5decc4d0c1c46ba6d0a6717
E:  010001
fingerprint, decimal: 12240908862933197005, hexadecimal： a9e071c1771060cd

```

