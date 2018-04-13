# ISSUE explanation to "Jet Template Engine for Go" project

When I apply filter `url` to `struct` variable I get the following stack trace:

```
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title>Test</title>
</head>
<body>
    <a href="someurl?id=panic: reflect: reflect.Value.call using value obtained using unexported field [recovered]
        panic: reflect: reflect.Value.call using value obtained using unexported field

goroutine 1 [running]:
github.com/user/jetfilters/vendor/github.com/CloudyKit/jet.(*Runtime).recover(0xc4200d2440, 0xc42004dee8)
        /home/maxix/Workspace/go/src/github.com/user/jetfilters/vendor/github.com/CloudyKit/jet/eval.go:231 +0x103
panic(0x5220c0, 0xc42008c1b0)
        /usr/local/go/src/runtime/panic.go:505 +0x229
reflect.flag.mustBeExported(0xb8)
        /usr/local/go/src/reflect/value.go:218 +0x145
reflect.Value.call(0x529a20, 0x563c88, 0x13, 0x55696b, 0x4, 0xc42004d9e8, 0x1, 0xa, 0x5220c0, 0x5220c0, ...)
        /usr/local/go/src/reflect/value.go:424 +0x780
reflect.Value.Call(0x529a20, 0x563c88, 0x13, 0xc42004d9e8, 0x1, 0xa, 0x0, 0x99, 0xc42004d9c8)
        /usr/local/go/src/reflect/value.go:308 +0xa4
github.com/user/jetfilters/vendor/github.com/CloudyKit/jet.reflect_Call(0xc42004d9e8, 0x1, 0xa, 0xc4200d2440, 0x529a20, 0x563c88, 0x13, 0x0, 0x0, 0x0, ...)
        /home/maxix/Workspace/go/src/github.com/user/jetfilters/vendor/github.com/CloudyKit/jet/eval.go:1257 +0x6ff
github.com/user/jetfilters/vendor/github.com/CloudyKit/jet.reflect_Call10(0x1, 0xc4200d2440, 0x529a20, 0x563c88, 0x13, 0x0, 0x0, 0x0, 0xc420098100, 0x1, ...)
        /home/maxix/Workspace/go/src/github.com/user/jetfilters/vendor/github.com/CloudyKit/jet/eval.go:1262 +0x107
github.com/user/jetfilters/vendor/github.com/CloudyKit/jet.(*Runtime).evalCallExpression(0xc4200d2440, 0x529a20, 0x563c88, 0x13, 0x0, 0x0, 0x0, 0xc420098100, 0x1, 0x1, ...)
        /home/maxix/Workspace/go/src/github.com/user/jetfilters/vendor/github.com/CloudyKit/jet/eval.go:1134 +0x12c
github.com/user/jetfilters/vendor/github.com/CloudyKit/jet.(*Runtime).evalCommandPipeExpression(0xc4200d2440, 0xc4200ce1e0, 0x5220c0, 0xc4200980a0, 0xb8, 0x0, 0x0, 0x0, 0x7e)
        /home/maxix/Workspace/go/src/github.com/user/jetfilters/vendor/github.com/CloudyKit/jet/eval.go:1190 +0x175
github.com/user/jetfilters/vendor/github.com/CloudyKit/jet.(*Runtime).evalPipelineExpression(0xc4200d2440, 0xc4200d2200, 0x7e, 0x80, 0x7e, 0x0)
        /home/maxix/Workspace/go/src/github.com/user/jetfilters/vendor/github.com/CloudyKit/jet/eval.go:1203 +0xa9
github.com/user/jetfilters/vendor/github.com/CloudyKit/jet.(*Runtime).executeList(0xc4200d2440, 0xc4200d2100)
        /home/maxix/Workspace/go/src/github.com/user/jetfilters/vendor/github.com/CloudyKit/jet/eval.go:398 +0x109
github.com/user/jetfilters/vendor/github.com/CloudyKit/jet.(*Template).ExecuteI18N(0xc4200dc000, 0x0, 0x0, 0x575fa0, 0xc420086008, 0xc42008e3f0, 0x0, 0x0, 0x0, 0x0)
        /home/maxix/Workspace/go/src/github.com/user/jetfilters/vendor/github.com/CloudyKit/jet/template.go:401 +0x166
github.com/user/jetfilters/vendor/github.com/CloudyKit/jet.(*Template).Execute(0xc4200dc000, 0x575fa0, 0xc420086008, 0xc42008e3f0, 0x0, 0x0, 0x0, 0xc42008e210)
        /home/maxix/Workspace/go/src/github.com/user/jetfilters/vendor/github.com/CloudyKit/jet/template.go:374 +0x6f
main.main()
        /home/maxix/Workspace/go/src/github.com/user/jetfilters/main.go:24 +0x174
exit status 2

```

Is it bug? Or I did something wrong?