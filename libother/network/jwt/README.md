golang jwt设置过期_一文读懂JWT,JWS,JWE

1.JWT是何物，有哪些常用的场景

JWT(json web token)是设计一种简洁，安全，无状态的token的实现规范rfc7519，通常用于网络请求方和网络接收方之间的网络请求认证。

jwt的常用场景

1.1: restful api接口的无状态认证, 在传统的web应用中,我们通常采用session认证。

session认证流程:

客户端将用户名/密码通过某种加密的方式发送给服务器

服务器接收到客户端请求之后进行验证，验证通过后使用Set-Cookie将用户的唯一sessionid放入到cookie当中，并将生成的sessionid和用户的关联信息存入到内存。

客户端第二次访问后服务器从当前cookie中取出sessionid并从内存中拿到相同的sessionid. 如果不存在，或者没有携带sessionid则说明该用户登陆过期，或者未登陆。

session认证的一些缺点:

由于使用session进行认证的方式必须存储sessionid,当用户量过大时对服务器内存消耗影响巨大. 如果你没有设置session的过期时间(关闭浏览器并不会导致cookie消失)那么对你的服务器来说消耗是致命的。

如果你没有将session存在一个所有服务器都可以获取得到的地方如redis, 那么意味着在本台服务器上面存储的sessionid其他服务器无法获取。用户进行请求时必须请求到这台服务器上面。可扩展性较差。

跨平台性较差，传统的session认证方式在移动端很难行得通。你必须开发二套不同的逻辑对web和移动端进行认证。

jwt认证流程:

客户端将用户名/密码通过某种加密的方式发送给服务器。

服务器接收到客户端请求后进行验证，验证通过服务器生成token返回给客户端。客户端将token存储在本地。

客户端每次请求将token携带在http header头中， 服务器端将token取出进行解密。

jwt认证的优点:

服务器端无需保存token，以加解密的方式代替存储，节省了内存空间。

无状态的token不依赖于服务器保存会话信息，更利于水平扩展。

相比于传统的session认证方式，jwt对移动端的支持更友好。

可以看出jwt认证解决了传统的session认证的一些不足之处。

1.2: 一次性认证：

如需要对某一个应用进行授权使用，此时就可以将jwt进行携带访问。

2.jwt的组成和生成方式

2.1jwt主要由三个部分组成，分别是头部(header)，载荷(payload)，签名(signature)组成。

header:

头部是用来声明此jwt的类型和加密算法，它们通常由alg和typ这二个字段组成。

alg字段通常用于表示加密采用的算法。

typ字段通常用于表示类型。

{"typ":"JWT", "alg":"HS256"} //示范的jwt header头
payload:

载荷就是我们存放公共参数/私有参数的地方.通俗点说该字段就是存放系统中用户的信息和jwt本身的一些信息，rfc文档本身替我们提供了一组字段的声明 (Claims)

iss: 该字段表示jwt的签发者。可以用你的应用唯一标识或者高权限的userid填充此字段。

sub: 该jwt面向的用户。

aud: jwt的接收方。

exp: jwt的过期时间,通常来说是一个时间戳。

iat: jwt的签发时间,常来说是一个时间戳。

jti：此jwt的唯一标识。通常用于解决请求中的重放攻击。该字段在大多数地方没有被提及或使用。因为使用此字段就意味着必须要在服务器维护一张jti表， 当客户端携带jwt访问的时候需要在jti表中查找这个唯一标识是否被使用过。使用这种方式防止重放攻击似乎让jwt有点怪怪的感觉, 毕竟jwt所宣称的优点就是无状态访问。-.-

//一个正儿八经的payload示范{ "iss": "appid_xxxxxx" "sub": "012345122", "exp": "1572246721840", "iat": "1592246721840"}
signature签名流程:

1.首先将header和payload进行base64编码，然后使用"."将header和payload拼接起来。类似于像下面这样:

String base64data = Base64.encode(header)+"."+Base64.encode(payload)
2.在将payload和header进行base64之后进行签名，得到签名后的数据。签名所使用的算法来自于header头中的alg字段。签名过程类似于像下面这样:

final JWSSigner signer = new ECDSASigner(this.privateKey); this.signature = signer.sign(base64data); //第一步， 实例化一个签名对象 //第二步，对base64data 进行签名
3.签名之后将签名的值和base64之后的header和payload用"."号连接起来。此时一个完整的jwt就出来啦。

jwt的最终结构: header.payload.signature

3.使用jwe来使你的jwt更加安全

签名到底在干什么:

在上面我们所谈到的jwt仅仅是签名后的jwt。在这里我们需要明白一个概念，那就是签名并不能保证数据的安全,也就是说如果有人获取到了你的jwt那么他可以通过转码得到你jwt当中所有的信息。那么读者可能会有点想骂人了，你特么上面说了那么多连一个数据安全都不能保证那么我看那么多有啥作用- -. 别急，我们先来了解一下签名的概念。其实对于签名更专业点的来说就是进行了一次哈希散列，对于散列我们首先要保证一下四点概念.

相同的输入将始终产生相同的输出。

多个不同的输入不应产生相同的输出。

从输出到输入应该是不可能的。

给定输入的任何修改都将导致哈希值发生巨大变化。

从上面4点大家看明白了吗？哈希与身份验证结合使用，可以产生强有力的证据来证明给定的消息尚未被修改。也就是说这个jwt从我这里签发以后无法改变，从而保证了数据来源的可靠性。

当客户端携带jwt进行请求时服务器在执行一遍签名步骤。如果签名的值一样就表示这个jwt是可靠的。此时我们在客户端和服务器之间就建立一种可信任的token机制。

应该怎样保证数据安全:

对于如何保证jwt本身的数据安全很多文章或文档都可以提及,我们可以把上面所生成的jwt成为jws(JSON Web Signed). 他本身的数据并没有进行加密。

此时如果我们想保证数据的安全就需要使用jwe(JSON Web Encryption)对jwt进行加密。jwe加密的秘文如下所示

// jwe相对于jws来说多了二个组成部分。eyJhbGciOiJSU0EtT0FFUCIsImVuYyI6IkEyNTZHQ00ifQ. OKOawDo13gRp2ojaHV7LFpZcgV7T6DVZKTyKOMTYUmKoTCVJRgckCL9kiMT03JGe ipsEdY3mx_etLbbWSrFr05kLzcSr4qKAq7YN7e9jwQRb23nfa6c9d-StnImGyFDb Sv04uVuxIp5Zms1gNxKKK2Da14B8S4rzVRltdYwam_lDp5XnZAYpQdb76FdIKLaV mqgfwX7XWRxv2322i-vDxRfqNzo_tETKzpVLzfiwQyeyPGLBIO56YJ7eObdv0je8 1860ppamavo35UgoRdbYaBcoh9QcfylQr66oc6vFWXRcZ_ZT2LawVCWTIy3brGPi 6UklfCpIMfIjf7iGdXKHzg. 48V1_ALb6US04U3b. 5eym8TW_c8SuK0ltJ3rpYIzOeDQz7TALvtu6UG9oMo4vpzs9tX_EFShS8iB7j6ji SdiwkIr3ajwQzaBtQD_A. XFBoMYUZodetZdvTiFvSkQ
jwe的5个组成部分:

JWE header: 描述用于创建jwe加密密钥和jwe密文的加密操作，类似于jws中的header。参数不一一描述，详情请见jwe header参数

JWE Encrypted Key：用来加密文本内容所采用的算法。

JWE initialization vector: 加密明文时使用的初始化向量值，有些加密方式需要额外的或者随机的数据。这个参数是可选的。

JWE Ciphertext:明文加密后产生的密文值。

JWE Authentication Tag：数字认证标签。

//一个完整的jwe json结构 { "protected":"jwe受保护的header头", "unprotected":"JWE Shared Unprotected Header数据", "header":"", "encrypted_key":"密钥加密后数据 ", "aad":"额外的认证数据", "iv":"同上的 JWE initialization vector", "ciphertext":"同上的JWE Ciphertext", "tag":"同上的JWE Authentication Tag" }
jwe创建流程:

根据头部 alg 的声明，将header头进行编码

随机生成密钥

加密密钥

生成iv如果不需要，此步骤可以省略.

加密原始报文

生成认证算法得到Authentication Tag.

如果明文有声明zip压缩，那么压缩明文

Base64.encode(header)+"."+Base64.encode(encrypted_key)+","+Base64.encode(iv)+"."+Base64.encode(ciphertext)+"."Base64.encode(tag)
4. jwt的缺点以及常见的理解误区

缺点：

无法主动的过期token. 常见的场景为后台踢出用户或封禁用户,此时若是token还在生效
时间范围内,那么意味着该用户在被踢出系统后还可以在这个时间范围内进行访问。又或者用户修改了密码，此时原token依然可以进行加解密也就代表着用户仍然可以继续访问。更为致命的是大多数客户端会将token存放在Local Storage或者vuex中，这意味着除非用户点击退出登陆。否则就算关闭软件在重新打开也会造成原密码可以登陆的假象。

jwt对比于传统的session认证方案并不会提高运行效率，因为本质上jwt做的是一个以时间换空间的动作。频繁的加解密会带来不小的性能开销。

常见的理解误区:

对jwt进行存储,这个是最常见的理解误区。事实上很多人都觉得将jwt进行存储可以完美的解决主动过期的问题, 然而这是一种赔了夫人又损兵的做法。既无法节省空间也无法节省时间。如果你的应用必须要主动过期的功能，那么我推荐你使用传统的session，事实上传统的session也有不少成熟的解决方案。如spring-session等。

jwt被盗用会导致数据泄漏不安全, 事实上使用jwe加密的jwt是不存在数据不安全的问题的。第二，jwt的数据一般都是非敏感数据，由于签名机制的存在所以你盗用了jwt做不了任何事情。

不存储如何实现退出功能。直接在客户端清除token就行了。因为服务器不存储所以只需要在客户端清除token就可以做到用户退出。

将token的时期设计的非常长。这是一个会造成很多隐性问题的设计，比如上述的密码修改等问题。这个时候可以参见oauth2的做法，将token的过期时间设置为一小时，但是增加一个refreshToken。客户端在没有触发退出或者修改密码等操作时通过refreshToken来刷新token。
