local json = require('cjson')
local hc = require('lib.httpclient').new()

local url = 'http://api-videostream-sh.fengkongcloud.com/videostream/v4'
local accessKey = '{ACCESS_KEY}'

local payload = json.encode({
  accessKey = accessKey,
  appId = 'default',
  eventId = 'video',
  imgType = 'POLITY_QRCODE_ADVERT',
  audioType = 'POLITY_EROTIC',
  callback = 'https://jsonplaceholder.typicode.com/posts/',
  data = {
    url = 'https://jsonplaceholder.typicode.com/posts/',
    streamType = 'NORMAL',
    btId = '{BT_ID}',
  }
})

local res = hc:post(url, payload)
local data = json.decode(res['body'])
-- 通过 code 和 riskLevel 判断返回，具体请参考接口文档
print(res['body'])
