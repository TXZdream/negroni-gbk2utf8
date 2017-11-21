# negroi-gbk2utf
Transform GBK to UTF-8 in request and change back in response
## maintainer
TangXuanzhao xuanzhaotang@gmail.com
## Thanks
[phyber](https://github.com/phyber/negroni-gzip)
I learned a lot from him, so the style of code is similar.
## Usage
1. Install
`go get -u github.com/txzdream/negroni-gbk2utf8`
2. Use
First,
`import github.com/txzdream/negroni-gbk2utf8/gbk2utf8`
Then,
```
server := negroni.Classic()
server.Use(gbk2utf8.Transformer())
```
## What you do with this middleware?
- When your request is encoded with gbk and you want to handle it with utf-8
- When you want to sent response back with gbk encoding
## What you can not do with it?
- Learn golang
This demo is quite simple and bad, so I hope you can learn from other places and contribute to this one.
- Use in production
I can not guarance the quality of this code, so please contribute to it when you have better idea.
## TODO
- Test
- Multiple language support
