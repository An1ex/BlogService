# BlogService
一个完整的博客网站
## 技术栈
前端：

后端：gin + gorm + JWT + mysql + redis + Jaeger

文档：swaggo

## TODO
- [x] 实现标签和标签接口的去重判断
- [x] 实现文章与标签的一对多关系
- [x] 实现文章与标签的数据库事务、回滚
- [ ] 实现多图片的上传接口
- [ ] 实现分布式的限流器(Redis实现)