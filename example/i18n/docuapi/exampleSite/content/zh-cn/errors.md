---
weight: 20
title: 错误

---

# 错误

<aside class ="notice">此误差部分被存储在一个单独的文件，errors.md。 DocuAPI可以让你尽可能多的文件，需要拆分单页文档。文件包含在默认雨果页面顺序。将页排序的一种方式是通过在前面的问题设置页面`weight`。具有较低体重页将被首先列出。</aside>

该Kittn API使用下面的错误代码: 


错误代码|含义
---------- | -------
400 |错误的请求 - 要求很烂
401 |未经授权 - 您的API密钥是错误的
403 | - 禁止小猫请求的是隐藏的管理员只
404 |未找到 - 指定的小猫找不到
405 |不允许的方法 - 您试图访问的小猫无效的方法
406 |无法接受 - 您要求的形式，是不是JSON
410 |飘 - 请求的小猫已经从我们的服务器中删除
418 |我是一个茶壶
429 |太多的要求 - 你要求太多的小猫！慢一点！
500 |内部服务器错误 - 我们必须与我们的服务器有问题。稍后再试。
503 |服务不可用 - 我们temporarially离线状态下进行maintanance。请稍后再试。
