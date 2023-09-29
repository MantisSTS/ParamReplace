# ParamReplace

## Install

```
go install github.com/MantisSTS/ParamReplace@latest
```

## Usage

```
cat wayback.txt | ParamReplace -param 'testparam' -value '"><script>alert("mantis")</script>' | httpx -ms 'alert("mantis")'
```

## Help
```
 -add
        Add the parameter if it doesn't exist
  -append
        Append the value to the parameter if it exists
  -param string
        Parameter to replace
  -value string
        Replacement value
  -verbose
        Verbose error messages

```
