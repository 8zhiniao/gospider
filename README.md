# gospider
go语言开发的爬虫项目
---------

单机版爬虫
----------
http://www.zhenai.com/zhenghun  ---获取--->  Item[city: ,cityUrl:,cityListPares:]

http://www.zhenai.com/zhenghun/  ---获取--->  Item[user: ,userUrl:,cityPares:]

http://www.zhenai.com/zhenghun/  ---获取--->  Item[username: ,userUrl:,userPares:]

----> 用户信息

----------

创建对应的引擎
-----
队列里面
从队列里面拿到一个parserResult进行解析
每个parserResult都包含request,然后从result分别获取如下内容:
request ---> url --- fetch - content -- > 解析器

引擎
----- 


城市用户表
通过每个城市的页面,获取每个城市下的用户的url


获取用户信息
----

----


##### 并发版爬虫
