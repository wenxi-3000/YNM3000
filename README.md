# 简介
要你命三千，集多种渗透工具于一身的终极武器霸王。<br>
通过yaml文件配置任意命令行工具的调度引擎。

<img src="./libs/images/YNM3000.jpeg" width=390 height=250/>

# 源码安装
```
git clone https://github.com/ShadowFl0w/YNM3000.git
cd YNM3000
go build
```

下载如下二进制到binaries目录<br>
anew https://github.com/tomnomnom/anew<br>
ksubdomain https://github.com/boy-hack/ksubdomain<br>
cleansub  https://github.com/j3ssie/go-auxs/tree/master/cleansub<br>
finddomain https://github.com/ShadowFl0w/mytools/tree/main/finddomain<br>
httpx https://github.com/projectdiscovery/httpx<br>
nuclei https://github.com/projectdiscovery/nuclei<br>
subfinder https://github.com/projectdiscovery/subfinder<br>


### 使用
1. 将需要的二进制文件放入到binaries文件夹
2. 在workflow里面配置模板。

```
YNM3000 scan -f ./target.txt --org HackedBycsdzds
```

# 模板
workflow目录用于存放模板，可以进行任意组合配置，通过yaml文件定义工具调用



