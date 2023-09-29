# ParamReplace

## Install

```
go install github.com/MantisSTS/ParamReplace@latest
```

## Usage

```
cat wayback.txt | ParamReplace -param 'testparam' -value '"><script>alert("mantis")</script>' | httpx -ms 'alert("mantis")'
```

