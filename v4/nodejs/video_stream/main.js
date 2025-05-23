const utils = require('../lib/utils')

const url = 'http://api-videostream-sh.fengkongcloud.com/videostream/v4'
const access_key = '{ACCESS_KEY}'
const bt_id = '{BT_ID}'
const video_url = 'https://jsonplaceholder.typicode.com/posts/'

const payload = {
    'accessKey': access_key,
    'appId': 'default',
    'eventId': 'video',
    'imgType': 'POLITY_QRCODE_ADVERT',
    'audioType': 'POLITY_EROTIC',
    'callback': 'https://jsonplaceholder.typicode.com/posts/',
    'data': {
        'url': video_url,
        'streamType': 'NORMAL',
        'btId': bt_id,
    },
}

utils.sendRequest(url, 'POST', payload, data => {
    // 通过 code 和 riskLevel 判断返回，具体请参考接口文档
    console.log(data)
})