# log.printf参数列表
    printf 参数类型：https://blog.csdn.net/p_function/article/details/84789407
    Go Printf函数副词参数列表
    General
    %v 以默认的方式打印变量的值
    %T 打印变量的类型
    Integer
    %d 整型的十进制表示
    %+d 带符号的整型,fmt.Printf("%+d",255)输出+255
    %q 打印单引号
    %o 不带0的八进制
    %#o 带0的八进制
    %x 小写的十六进制
    %X 大写的十六进制
    %#x 带0x的十六进制
    %U 打印Unicode字符
    %#U 打印带字符的Unicode
    %b 打印整型的二进制
    Integet width
    %5d 表示该整型最大的长度为5
    %-5d 则相反,打印结果会自动左对齐
    %05d 会在数字前面补0
    Fload
    %f(=%.6f) 6位小数点
    %e(=%.6e) 6位小数点(科学计数法)
    %g 用最少的数字来表示
    %.3g 最多3位数字来表示
    %.3f 最多3位小数来表示
    String
    %s 正常输出字符串
    %q 字符串带双引号,字符串中的引号带转义符
    %#q 字符串带反引号,如果字符串内有反引号,就用双引号代替
    %x 将字符串转换为小写的16进制格式
    %X 将字符串转换为大写的16进制格式
    % x 带空格的16进制格式
    String width
    %5s 最小宽度为5
    %-5s 最小宽度为5(左对齐)
    %.5s 最大宽度为5
    %-5.7 最小宽度为5,最大宽度为7(左对齐)
    %5.3 如果宽度大于3则截断
    %05s 如果宽度小于5,就会在字符串前面补0
    Struct
    %v 正常打印.比如:{sam{12345 67890}}
    %+v 带字段名称,比如:{name:sam phone:{mobile:12345 office:67890}}
    %#v 用Go的语法打印,例如:main.People{name:"sam",phone:main.Phone{mobile:"12345",office:"67890"}}
    Boolean
    %t 打印布尔值true或者false
    Pointer
    %p 带0x的指针
    %#p 不带0x的指针