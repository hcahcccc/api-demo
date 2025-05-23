<?php

require 'lib/utils.php';

$url = 'http://api-video-bj.fengkongcloud.com/video/v4';
$access_key = '{ACCESS_KEY}';
$bt_id = '{BT_ID}';
$tokenId = '{TOKENID}';
$video_url = 'https://jsonplaceholder.typicode.com/posts/';

$payload = array(
    'accessKey' => $access_key,
    'appId' => 'default',
    'eventId' => 'video',
    'imgType' => 'POLITY_QRCODE_ADVERT',
    'audioType' => 'POLITY_EROTIC',
    'callback' => 'https://jsonplaceholder.typicode.com/posts/',
    'data' => array(
        'url' => $video_url,
        'btId' => $bt_id,
        'tokenId' => $tokenId,
    ),
);

$res = request_post($url, json_encode($payload));
$data = json_decode($res);
// 通过 code 和 riskLevel 判断返回，具体请参考接口文档
var_dump($data);