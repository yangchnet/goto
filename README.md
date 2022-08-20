# goto

管理ssh服务器

## 使用方法：

- 连接服务器
```sh
$ goto $server_name
```

- 增加服务器
```sh
goto server -a $server_name:$user_name:$server_ip [-s $secret_key_file] [-p $passwd]
```

- 列出当前所有服务器
```sh
goto server -l
```

- 删除某个服务器
```sh
goto server -d $server_name
```

- 添加密钥
```sh
goto key -a $key_name
```

- 删除密钥
```sh
goto key -d $key_name
```

- list keys
```sh
goto key -l
```
