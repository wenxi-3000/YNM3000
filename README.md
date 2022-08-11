## 1. 简介
要你命三千，集多种渗透工具于一身的终极武器霸王。<br>
通过yaml文件配置任意命令行工具的调度引擎。

<img src="./code/libs/images/YNM3000.jpeg" width=390 height=250/>

<br>

## 2. 源码安装(MacOS/Linux)
```
git clone https://github.com/ShadowFl0w/YNM3000.git
cd YNM3000
go build
```
安装基础工具(确保安装有golang,并配置好了$GOPATH): 
```
./install.sh
```
install.sh会安装如下二进制文件到binaries目录
```
anew    https://github.com/tomnomnom/anew
ksubdomain  https://github.com/boy-hack/ksubdomain
cleansub    https://github.com/j3ssie/go-auxs/tree/master/cleansub
finddomain  https://github.com/ShadowFl0w/mytools/tree/main/finddomain
httpx   https://github.com/projectdiscovery/httpx
nuclei  https://github.com/projectdiscovery/nuclei
subfinder   https://github.com/projectdiscovery/subfinder
```

## 3. 使用
1. 需要的二进制文件已放入binaries目录
2. 在workflow里面配置模板。

```
YNM3000 scan -f ./target.txt --org test
```

## 4. 结果保存
结果保存在results目录<br>
通过--org参数指定results目录下的子目录<br>
如果不适用--org参数则保存在./results/no-org目录


## 4. 模板
workflow目录用于存放模板，可以进行任意组合配置（通过yaml配置）



