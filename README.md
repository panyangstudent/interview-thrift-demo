本项目为interview项目下thrift部分的代码示例，包括生成的源码等记录

# mac安装thrift，生成go代码
* 通过brew安装brew
    > brew install thrift

    因为我们只是示例，所以这里没有通过官网安装最新的版本，brew安装的为0.9.3版本
* 查看是否安装成功
  > thrift -version
* 创建git项目拉倒本地，并且书写user.thrift文件,内容如下：
  ```thrift
  namespace go user
  service UserService {
    i32 add(1:i32 num1,2:string str1)
  }
  ```
  
* 生成go代码
  > thrift -r --gen go user.thrift 
