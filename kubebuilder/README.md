# Kubebuilder笔记

-----

## kubebuilder 作用
- 提供脚手架工具初始化 CRDs 工程，自动生成 boilerplate 代码和配置
- 提供代码库封装底层的 K8s go-client

## kubebuilder整体流程
![kubebuilder-process](./kubebuilder.png)

1. 用户自定义crd，将自定义的crd注册到scheme中，这样通过GVK能找到对应的go的struct，也能通过go的struct找对对应的GVK
2. Cache监听Scheme中的GVK，同时和Api Server简历list-watch的连接
3. 当发现某个controller需要的GVK资源发生状态的改变就reconcile调度对应的controller
4. 对应的controller中是用户自己定义的逻辑，用来保证该类型的crd和k8s集群中声明的该资源的yaml/json中的字段一致
5. 保证一致主要是通过调用Clients，然后Clients可以和Api Server交互


## 使用方法
1. 初始化脚手架
这不创建了一个项目模板并引入一些依赖
```bash
kubebuilder init --domain ding.test.com
```
2. 创建api
会在根目录下创建一个api目录
里头有:
resourcename_types.go、groupversion_info.go、
zz_generated.deepcopy.go
![kubebuilder-process](./kubebuilder-catalog.png)
```bash
kubebuilder create api --group ding.shin.com --version v1alpha666 --kind DingShinType
```
groupversion_info.go: 会声明 GroupVersion、SchemeBuilder、AddToScheme三个东西(也可以手动给改到别的文件引入)
![kubebuilder-catalog](./kubebuilder-groupversion.png)
每当再执行一遍craete api 如:
```bash
kubebuilder create api --group ding.shin.com --version v1alpha666 --kind DingShin888Type
```
就会多生成一个xxxxx_type.go
![kubebuilder-catalog](./kubebuilder-app-new-api.png)
同时在 zz_generated.deepcopy.go 文件中也会对新的api创建对应的 deepcopy 相关的函数(对runtime.Object的interface的实现)</br>
另外 config 目录下的 crd 目录下的 yaml 们是对要创建的crd资源的描述</br>
controllers 目录下是 自定义的 Kind 的 controller



