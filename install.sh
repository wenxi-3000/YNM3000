#! /bin/sh

uNames=`uname -s`

echo "当前系统: " $uNames

osName=${uNames: 0: 4}


if [[ "$osName" == "Darw" || "$osName" == "Linu" ]] # Darwin/Linux
then
	# echo "当前操作系统: Mac OS X"


    # 下载anew
    if [ -f ./binaries/anew ]
    then
        echo "anew已经存在"

    else
        echo "开始下载anew: "
        go install -v github.com/tomnomnom/anew@latest
        cp $GOPATH/bin/anew ./binaries
    fi

    #下载cleansub
    if [ -f ./binaries/cleansub ]
    then 
        echo "cleansub已经存在"
    else 
        echo "开始下载cleansub:"
        curl -o cleansub.go https://raw.githubusercontent.com/j3ssie/go-auxs/master/cleansub/main.go
        go build cleansub.go
        mv ./cleansub ./binaries
        rm cleansub.go
    fi


    #下载finddomain
    if [ -f ./binaries/finddomain ]
    then 
        echo "finddomain已经存在"
    else 
        echo "开始下载finddomain:"
        curl -o finddomain.go https://raw.githubusercontent.com/ShadowFl0w/mytools/main/finddomain/finddomain.go
        go build finddomain.go
        mv ./finddomain ./binaries
        rm finddomain.go
    fi


    #下载httpx
    if [ -f ./binaries/httpx ]
    then 
        echo "httpx已经存在"
    else 
        echo "开始下载httpx:"
        go install -v github.com/projectdiscovery/httpx/cmd/httpx@latest
        cp $GOPATH/bin/httpx ./binaries
    fi   


    #下载ksubdomain
    if [ -f ./binaries/ksubdomain ]
    then 
        echo "ksubdomain已经存在"
    else 
        echo "开始下载ksubdomain:"
        go install -v github.com/boy-hack/ksubdomain/cmd/ksubdomain@latest
        cp $GOPATH/bin/ksubdomain ./binaries
    fi

    #下载naabu
    if [ -f ./binaries/naabu ]
    then 
        echo "naabu已经存在"
    else 
        echo "开始下载naabu:"
        go install -v github.com/projectdiscovery/naabu/v2/cmd/naabu@latest
        cp $GOPATH/bin/naabu ./binaries
    fi


    #下载nuclei
    if [ -f ./binaries/nuclei ]
    then 
        echo "nuclei已经存在"
    else 
        echo "开始下载nuclei:"
        go install -v github.com/projectdiscovery/nuclei/v2/cmd/nuclei@latest
        cp $GOPATH/bin/nuclei ./binaries
    fi

    #下载subfinder
    if [ -f ./binaries/subfinder ]
    then 
        echo "subfinder已经存在"
    else 
        echo "开始下载subfinder:"
        go install -v github.com/projectdiscovery/subfinder/v2/cmd/subfinder@latest
        cp $GOPATH/bin/subfinder ./binaries
    fi


# elif [ "$osName" == "Linu" ] # Linux
# then
# 	echo "GNU/Linux"
elif [ "$osName" == "MING" ] # MINGW, windows, git-bash
then 
	echo "Windows, git-bash"
    exit 
else
	echo "unknown os"
    exit
fi