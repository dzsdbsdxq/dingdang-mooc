import Fly from 'flyio/dist/npm/fly';
const fly = new Fly();
const APP_API = process.env.NODE_ENV === 'development' ? 'http://localhost:50288/' : '/' ;
// 设置请求基地址
fly.config.baseURL = `${APP_API}`;

// 添加请求拦截器
fly.interceptors.request.use((request) => {
  request.timeout = 30000;
  //将参数对象进行字典顺序排序
  request.body['_t'] = new Date().getTime();
  //let bodyStr = utils.jsonSort(request.body);
  //request.body['sign'] = Base64.encode(md5(utils.enCrypt(bodyStr)))

  request.headers = {
    "content-type": "application/x-www-form-urlencoded"
  }
  return request;
})

// 添加响应拦截器，响应拦截器会在then/catch处理之前执行
fly.interceptors.response.use((response) => {
    //let responseData = JSON.parse(response.data);
    return Promise.resolve(response.data)
  },
  (err) => {
    // 发生网络错误后会走到这里
    return Promise.reject(err.response.data);
  }
)
export { fly }
