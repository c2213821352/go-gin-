1. 微信登录  ，把code 发送到微信的服务器 ，返回的openid和session——key 存在服务器内部 ，然后 返回一个登录状态到小程序，比如说登录成功， 使用wx.request 携带登录状态请求，通过自定义登录状态，查询openid和session ，返回数据

2. 携带上自己的钱数量，登录状态，返回增加的钱

   ​

3. 接口数量：

- 一个登录的接口，
- 一个加钱的接口，（另外算，完成看视频之后，发送一个请求，加钱）
- 一个提现的接口，需要openid ，
- 一个签到的接口（另外算）
- 一个小游戏最高记录的接口（另外算）