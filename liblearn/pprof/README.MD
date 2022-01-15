1. 链接：https://www.topgoer.com/%E5%85%B6%E4%BB%96/pprof%E6%80%A7%E8%83%BD%E8%B0%83%E4%BC%98.html
2. pprof性能调优
   在计算机性能调试领域里，profiling 是指对应用程序的画像，画像就是应用程序使用 CPU 和内存的情况。 Go语言是一个对性能特别看重的语言，因此语言中自带了 profiling 的库，这篇文章就要讲解怎么在 golang 中做 profiling。

1.1.1. 前言
写了几吨代码，实现了几百个接口。功能测试也通过了，终于成功的部署上线了

结果，性能不佳，什么鬼？

1.1.2. pprof
想要进行性能优化，首先瞩目在 Go 自身提供的工具链来作为分析依据，本文将带你学习、使用 Go 后花园，涉及如下：

runtime/pprof：采集程序（非 Server）的运行数据进行分析
net/http/pprof：采集 HTTP Server 的运行时数据进行分析
pprof开启后，每隔一段时间（10ms）就会收集下当前的堆栈信息，获取格格函数占用的CPU以及内存资源；最后通过对这些采样数据进行分析，形成一个性能分析报告。

注意，我们只应该在性能测试的时候才在代码中引入pprof。

1.1.3. 是什么
pprof 是用于可视化和分析性能分析数据的工具

pprof 以 profile.proto 读取分析样本的集合，并生成报告以可视化并帮助分析数据（支持文本和图形报告）

profile.proto 是一个 Protocol Buffer v3 的描述文件，它描述了一组 callstack 和 symbolization 信息， 作用是表示统计分析的一组采样的调用栈，是很常见的 stacktrace 配置文件格式

1.1.4. 支持什么使用模式
Report generation：报告生成
Interactive terminal use：交互式终端使用
Web interface：Web 界面
1.1.5. 可以做什么
CPU Profiling：CPU 分析，按照一定的频率采集所监听的应用程序 CPU（含寄存器）的使用情况，可确定应用程序在主动消耗 CPU 周期时花费时间的位置
Memory Profiling：内存分析，在应用程序进行堆分配时记录堆栈跟踪，用于监视当前和历史内存使用情况，以及检查内存泄漏
Block Profiling：阻塞分析，记录 goroutine 阻塞等待同步（包括定时器通道）的位置
Mutex Profiling：互斥锁分析，报告互斥锁的竞争情况
1.1.6. 工具型应用
如果你的应用程序是运行一段时间就结束退出类型。那么最好的办法是在应用退出的时候把 profiling 的报告保存到文件中，进行分析。对于这种情况，可以使用runtime/pprof库。 首先在代码中导入runtime/pprof工具：

    import "runtime/pprof"
1.1.7. CPU性能分析
开启CPU性能分析：

    pprof.StartCPUProfile(w io.Writer)
停止CPU性能分析：

    pprof.StopCPUProfile()
应用执行结束后，就会生成一个文件，保存了我们的 CPU profiling 数据。得到采样数据之后，使用go tool pprof工具进行CPU性能分析。

1.1.8. 内存性能优化
记录程序的堆栈信息

    pprof.WriteHeapProfile(w io.Writer)
得到采样数据之后，使用go tool pprof工具进行内存性能分析。

go tool pprof默认是使用-inuse_space进行统计，还可以使用-inuse-objects查看分配对象的数量

1.1.9. 服务型应用（net/http/pprof）
如果你的应用程序是一直运行的，比如 web 应用，那么可以使用net/http/pprof库，它能够在提供 HTTP 服务进行分析。

如果使用了默认的http.DefaultServeMux（通常是代码直接使用 http.ListenAndServe(“0.0.0.0:8000”, nil)），只需要在你的web server端代码中按如下方式导入net/http/pprof

    import _ "net/http/pprof"
如果你使用自定义的 Mux，则需要手动注册一些路由规则：

    r.HandleFunc("/debug/pprof/", pprof.Index)
    r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
    r.HandleFunc("/debug/pprof/profile", pprof.Profile)
    r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
    r.HandleFunc("/debug/pprof/trace", pprof.Trace)
如果你使用的是gin框架，那么推荐使用"github.com/DeanThompson/ginpprof"。

源码示例

package main

import (
"flag"
"log"
"net/http"
_ "net/http/pprof"
"sync"
"time"
)

func Counter(wg *sync.WaitGroup) {
time.Sleep(time.Second)

    var counter int
    for i := 0; i < 1000000; i++ {
        time.Sleep(time.Millisecond * 200)
        counter++
    }
    wg.Done()
}

func main() {
flag.Parse()

    //远程获取pprof数据
    go func() {
        log.Println(http.ListenAndServe("localhost:8080", nil))
    }()

    var wg sync.WaitGroup
    wg.Add(10)
    for i := 0; i < 10; i++ {
        go Counter(&wg)
    }
    wg.Wait()

    // sleep 10mins, 在程序退出之前可以查看性能参数.
    time.Sleep(60 * time.Second)
}
编译运行之后在浏览器访问 http://localhost:8080/debug/pprof/



这个路径下还有几个子页面：

/debug/pprof/profile：访问这个链接会自动进行 CPU profiling，持续 30s，并生成一个文件供下载
/debug/pprof/heap： Memory Profiling 的路径，访问这个链接会得到一个内存 Profiling 结果的文件
/debug/pprof/block：block Profiling 的路径
/debug/pprof/goroutines：运行的 goroutines 列表，以及调用关系
1.1.10. 通过交互式终端使用
不管是工具型应用还是服务型应用，我们使用相应的pprof库获取数据之后，下一步的都要对这些数据进行分析，我们可以使用go tool pprof命令行工具。

go tool pprof最简单的使用方式为:

    go tool pprof [binary] [source]
其中：

binary 是应用的二进制文件，用来解析各种符号；
source 表示 profile 数据的来源，可以是本地的文件，也可以是 http 地址。
注意事项： 获取的 Profiling 数据是动态的，要想获得有效的数据，请保证应用处于较大的负载（比如正在生成中运行的服务，或者通过其他工具模拟访问压力）。否则如果应用处于空闲状态，得到的结果可能没有任何意义。

示例代码：

// runtime_pprof/main.go
package main

import (
"flag"
"fmt"
"os"
"runtime/pprof"
"time"
)

// 一段有问题的代码
func logicCode() {
var c chan int // nil
for {
select {
case v := <-c: // 阻塞
fmt.Printf("recv from chan, value:%v\n", v)
default:
time.Sleep(time.Millisecond * 500)
}
}
}

func main() {
var isCPUPprof bool // 是否开启CPUprofile的标志位
var isMemPprof bool // 是否开启内存profile的标志位

    flag.BoolVar(&isCPUPprof, "cpu", false, "turn cpu pprof on")
    flag.BoolVar(&isMemPprof, "mem", false, "turn mem pprof on")
    flag.Parse()

    if isCPUPprof {
        f1, err := os.Create("./cpu.pprof") // 在当前路径下创建一个cpu.pprof文件
        if err != nil {
            fmt.Printf("create cpu pprof failed, err:%v\n", err)
            return
        }
        pprof.StartCPUProfile(f1) // 往文件中记录CPU profile信息
        defer func() {
            pprof.StopCPUProfile()
            f1.Close()
        }()
    }
    for i := 0; i < 6; i++ {
        go logicCode()
    }
    time.Sleep(20 * time.Second)

    if isMemPprof {
        f2, err := os.Create("./mem.pprof")
        if err != nil {
            fmt.Printf("create mem pprof failed, err:%v\n", err)
            return
        }
        pprof.WriteHeapProfile(f2)
        f2.Close()
    }
}
执行

    go run main.go -cpu
等一会就出在同级目录下生成一个cpu.pprof文件

1.1.11. 通过交互式终端使用
我们使用go工具链里的pprof来分析一下。

    go tool pprof cpu.pprof
执行上面的代码会进入交互界面如下：

Type: cpu
Time: Nov 14, 2019 at 11:21am (CST)
Duration: 20s, Total samples = 50ms ( 0.25%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof)
我们可以在交互界面输入top3来查看程序中占用CPU前3位的函数：

    Type: cpu
    Time: Nov 14, 2019 at 11:21am (CST)
    Duration: 20s, Total samples = 50ms ( 0.25%)
    Entering interactive mode (type "help" for commands, "o" for options)
    (pprof) top3
    Showing nodes accounting for 40ms, 80.00% of 50ms total
    Showing top 3 nodes out of 19
          flat  flat%   sum%        cum   cum%
          20ms 40.00% 40.00%       20ms 40.00%  runtime.stdcall1
          10ms 20.00% 60.00%       10ms 20.00%  runtime.casgstatus
          10ms 20.00% 80.00%       10ms 20.00%  runtime.findrunnable
结束后将默认进入 pprof 的交互式命令模式，可以对分析的结果进行查看或导出。具体可执行 pprof help 查看命令说明

flat：给定函数上运行耗时
flat%：同上的 CPU 运行耗时总比例
sum%：给定函数累积使用 CPU 总比例
cum：当前函数加上它之上的调用运行总耗时
cum%：同上的 CPU 运行耗时总比例
最后一列为函数名称，在大多数的情况下，我们可以通过这五列得出一个应用程序的运行情况，加以优化

我们还可以使用list 函数名命令查看具体的函数分析，例如执行list logicCode查看我们编写的函数的详细分析。

    (pprof) list logicCode
    Total: 50ms
结合代码可以找到需要优化的代码行数

1.1.12. 图形化
或者可以直接输入web，通过svg图的方式查看程序中详细的CPU占用情况。 想要查看图形化的界面首先需要安装graphviz图形化工具。

Mac：

    brew install graphviz
Windows: 下载graphviz 将graphviz安装目录下的bin文件夹添加到Path环境变量中。 在终端输入dot -version查看是否安装成功。



关于图形的说明： 每个框代表一个函数，理论上框的越大表示占用的CPU资源越多。 方框之间的线条代表函数之间的调用关系。 线条上的数字表示函数调用的次数。 方框中的第一行数字表示当前函数占用CPU的百分比，第二行数字表示当前函数累计占用CPU的百分比。

1.1.13. go-torch和火焰图
火焰图（Flame Graph）是 Bredan Gregg 创建的一种性能分析图表，因为它的样子近似 🔥而得名。上面的 profiling 结果也转换成火焰图，如果对火焰图比较了解可以手动来操作，不过这里我们要介绍一个工具：go-torch。这是 uber 开源的一个工具，可以直接读取 golang profiling 数据，并生成一个火焰图的 svg 文件。

安装go-touch

      go get -v github.com/uber/go-torch
火焰图 svg 文件可以通过浏览器打开，它对于调用图的最优点是它是动态的：可以通过点击每个方块来 zoom in 分析它上面的内容。

火焰图的调用顺序从下到上，每个方块代表一个函数，它上面一层表示这个函数会调用哪些函数，方块的大小代表了占用 CPU 使用的长短。火焰图的配色并没有特殊的意义，默认的红、黄配色是为了更像火焰而已。

go-torch 工具的使用非常简单，没有任何参数的话，它会尝试从http://localhost:8080/debug/pprof/profile获取 profiling 数据。它有三个常用的参数可以调整：

-u –url：要访问的 URL，这里只是主机和端口部分
-s –suffix：pprof profile 的路径，默认为 /debug/pprof/profile
–seconds：要执行 profiling 的时间长度，默认为 30s
安装 FlameGraph 要生成火焰图，需要事先安装 FlameGraph工具，这个工具的安装很简单（需要perl环境支持），只要把对应的可执行文件加入到环境变量中即可。

1.下载安装perl：https://www.perl.org/get.html

2.下载FlameGraph：git clone https://github.com/brendangregg/FlameGraph.git

3.将FlameGraph目录加入到操作系统的环境变量中。

4.Windows平台的同学，需要把go-torch/render/flamegraph.go文件中的GenerateFlameGraph按如下方式修改，然后在go-torch目录下执行go install即可。

// GenerateFlameGraph runs the flamegraph script to generate a flame graph SVG. func GenerateFlameGraph(graphInput []byte, args ...string) ([]byte, error) {
flameGraph := findInPath(flameGraphScripts)
if flameGraph == "" {
return nil, errNoPerlScript
}
if runtime.GOOS == "windows" {
return runScript("perl", append([]string{flameGraph}, args...), graphInput)
}
return runScript(flameGraph, args, graphInput)
}
1.1.14. 压测工具wrk
推荐使用 https://github.com/wg/wrk 或 https://github.com/adjust/go-wrk

1.1.15. 使用go-torch
使用wrk进行压测:go-wrk -n 50000 http://127.0.0.1:8080/book/list 在上面压测进行的同时，打开另一个终端执行go-torch -u http://127.0.0.1:8080 -t 30，30秒之后终端会初夏如下提示：Writing svg to torch.svg

然后我们使用浏览器打开torch.svg就能看到火焰图了。

火焰图的y轴表示cpu调用方法的先后，x轴表示在每个采样调用时间内，方法所占的时间百分比，越宽代表占据cpu时间越多。通过火焰图我们就可以更清楚的找出耗时长的函数调用，然后不断的修正代码，重新采样，不断优化。

1.1.16. pprof与性能测试结合
go test命令有两个参数和 pprof 相关，它们分别指定生成的 CPU 和 Memory profiling 保存的文件：

-cpuprofile：cpu profiling 数据要保存的文件地址
-memprofile：memory profiling 数据要报文的文件地址
我们还可以选择将pprof与性能测试相结合，比如：

比如下面执行测试的同时，也会执行 CPU profiling，并把结果保存在 cpu.prof 文件中：

    go test -bench . -cpuprofile=cpu.prof
比如下面执行测试的同时，也会执行 Mem profiling，并把结果保存在 cpu.prof 文件中：

    go test -bench . -memprofile=./mem.prof
需要注意的是，Profiling 一般和性能测试一起使用，这个原因在前文也提到过，只有应用在负载高的情况下 Profiling 才有意义。