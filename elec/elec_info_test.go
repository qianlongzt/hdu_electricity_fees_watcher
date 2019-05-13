package elec

import (
	"testing"
)

var rawData = `
<html class="no-js"><head>
    <meta charset="utf-8">
    <title>电费充值</title>
    <meta name="description" content="">
    <meta name="HandheldFriendly" content="True">
    <meta name="MobileOptimized" content="320">
    <meta name="viewport" content="width=device-width, initial-scale=1, minimal-ui, user-scalable=no">
    <meta http-equiv="cleartype" content="on">
    <meta name="format-detection" content="telephone=no">
    <link rel="stylesheet" href="http://static.xiaotui.so/stable/GC2/assets/base/electricity/css/index.css?ver=">
	<link rel="stylesheet" href="http://static.xiaotui.so/stable/GC2/assets/base/ecs/css/index.css">
	<link rel="stylesheet" type="text/css" href="http://static.xiaotui.so/stable/GC2/assets/bower_components/sweetalert/lib/sweet-alert.css">
</head>
<body>
<div class="wraper">
    <div class="top">
        <a href="/base/electricity_hd/query/ele_id/7" class="iconfont icon-zuo pull-left"></a>
        <span class="pull-left">电量查询</span>
        <a href="/base/electricity_hd/index/ele_id/7" class="iconfont icon-shouye pull-right text-right"></a>
    </div>  <div class="field">截至 <font style="color:#2d9fd3"><b>2019-05-06 17:26:06</b></font></div>
    <div class="info-box margin-top">
        <p>实际所剩金额：<span class="price" style="color:#2d9fd3">62.10元</span></p>
        <p>园区：下沙本部</p>
        <p>楼幢：12号楼南</p>
        <p>楼层：3层</p>
        <p>寝室号：321</p>
    </div>
    <div class="btn-box">
        <a href="/base/electricity_hd/index/ele_id/7" class="btn">确认</a>
    </div>
</div>
<script src="http://static.xiaotui.so/stable/GC2/assets/base/electricity/js/jquery-2.1.0.min.js"></script>
<script src="http://static.xiaotui.so/stable/GC2/assets/base/electricity/js/jq.utils.js"></script>
<script src="http://static.xiaotui.so/stable/GC2/assets/bower_components/underscore/underscore-min.js"></script>
<script src="http://static.xiaotui.so/stable/GC2/assets/bower_components/sweetalert/lib/sweet-alert.min.js"></script>
<script src="http://static.xiaotui.so/stable/GC2/assets/public/layermobile/layer.js"></script>
<link rel="stylesheet" href="http://static.xiaotui.so/stable/GC2/assets/bower_components/sweetalert/lib/sweet-alert.css?v=">


</body>
</html>`

func TestGetElecInfoFromString(t *testing.T) {
	r, e := getElectInfoFromString(rawData)
	if e != nil {
		t.Fatal(r, e)
	} else {
		t.Log(r)
	}
}
