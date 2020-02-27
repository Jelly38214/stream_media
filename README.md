API 设计原则

- 以 URL(统一资源定位符) 风格设计 API
- 通过不同的 METHOD（GET, POST, PUT, DELETE）来区分对资源对 CRUD
- 返回码（Status Code）符合 HTTP 资源描述的规定

用户 (发表)-> 评论 (从属关系)-> 视频
||> 上传，观看，下载，删除

- 创建（注册）用户:

  - URL: /user
  - Method: POST
  - Status Code: 201(创建成功), 400, 500

- 用户登录:
  - URL: /user/:username
  - Method: POST
  - Status Code: 200, 400

* 获取用户基本信息：

  - URL: /user/:username
  - Method: GET
  - Status Code: 200, 400(Not Found), 401(Bad Request), 403(forbid), 500

* 用户注销:
  - URL: /user/:username
  - Method: DELETE
  - Status Code: 204(Not Return Content), 400, 401, 403, 500
