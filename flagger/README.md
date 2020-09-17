# 基于 Flagger 创建应用 A/B Test 原理
---
1. 首先创建出来的 Canary 资源需要有 targetRef 和 ingressRef 以及 service 三个主要配置项
2. targetRef 要填想要托管的服务的 deployment，ingressRef 要写想要托管的服务的 ingress 资源， service 就和要托管的服务的 service 功能一样
3. flagger-runtime 的 pods 会根据 targetRef 找到对应服务的 deployment，然后把本来要启动的 deployment 的 replicas 置为 0，并且创建一个基本上一模一样的 deployment（replicas不为0）
4. 该 deployment 创建出来的 pods 的名字后面会默认加一个 -primary
5. 然后会创建两个对应的 svc，一个 svc 的 selector 也会对应着多加一个 -primary，另一个 svc 的 selector 后面啥也不加，就是和用户创建的 pod 的名字一样。两个 svc，一个叫 xxxx-primary，一个叫 xxxx-canary。
6. 另外还会把最原始也就是用户创建的那个 svc 的 selector 的 label 后面加个 -primary，也就是说 用户自己写的 svc 和 flagger 自己创建的带有 -primary 后缀的 svc 都会把流量打到 flagger 自己启动的带有 -primary 的 pods 中
7. 之后根据 ingressRef 创建一个 xxxx-canary 版本的 ingress 资源，并且该 ingress 规则对应的 svc 是上面第 5 步中的 xxxx-canary。
8. 但是 flagger-runtime 不会修改原始的也就是用户自己写的 ingress 的资源规则，也就是说，当访问用户自定义的 ingress 规则时，会把流量打到最最原始的那个 svc 上，也就是上面第 6 步中那个 svc。由于原始的 svc 的 selector 的 label 后面也加了 -primary，所以流量会打到 flagger 自己启动的 pods 中。

---

↑ 以上是创建一个应用时候 flagger 做的事情
---
---
↓ 当更新一个应用触发 A/B Test 时候
---
---
9. flagger-runtime 检测到原始的 deployment 被更新，然后就会创建一个后缀为 -canary 的 deployment，并且启动不带后缀的 pods
10. 然后由于上面第 7 步中创建的 xxxx-canary 的存在，flagger-runtime 会将 xxxx-canary 的 annotations 中的一些配置进行更新。注意 k8s 的流量分流，A/B Test，金丝雀等功能，基本上都是通过新版本 ingress 的 annotations 中这些配置实现的。
11. 此时 ingress-runtime 一直在监听着 ingress 资源，所以也会检测到 xxxx-canary 的 ingress 资源中的变化，然后动态更新内部的服务代理
12. 所以这个时候如果根据一些规则比如 headers 信息或者 cookie 之类的就可以将流量经由 ingress-runtime 打到 xxxx-canary 的 ingress 的规则
13. xxxx-canary 的 ingress 对应的服务 svc，正好是上面第 5 步中的 xxxx-canary 的 svc，所以流量也就被打倒 xxxx-canary 的 svc 中
14. xxxx-canary 的 svc 中的 selector 对应的 pods，是不带任何后缀的 pods，由于第 9 步中创建的 deployment 并且的 pods 也是不带任何后缀的，所以流量就被打倒了这个更新后的版本了


