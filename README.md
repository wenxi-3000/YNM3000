# 简介
要你命三千，集多种渗透工具于一身的终极武器霸王。<br>可以配置任意命令行工具的调度引擎。

<img src="./libs/images/YNM3000.jpeg" width=390 height=250/>

# 源码安装
```
git clone https://github.com/ShadowFl0w/YNM3000.git
cd YNM3000
go build
```

下载如下二进制到binaries目录
anew https://github.com/tomnomnom/anew
ksubdomain https://github.com/boy-hack/ksubdomain
cleansub  https://github.com/j3ssie/go-auxs/tree/master/cleansub
finddomain https://github.com/ShadowFl0w/mytools/tree/main/finddomain
httpx https://github.com/projectdiscovery/httpx
nuclei https://github.com/projectdiscovery/nuclei
subfinder https://github.com/projectdiscovery/subfinder


### 使用
1. 将需要的二进制文件放入到binaries文件夹
2. 在workflow里面配置模板。

```
YNM3000 scan -f ./target.txt --org HackedBycsdzds
```

# 模板
模板有数个变量，可以进行任意组合配置



