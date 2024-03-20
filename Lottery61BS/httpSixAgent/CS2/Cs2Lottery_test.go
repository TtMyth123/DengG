package CS2

import (
	"fmt"
	"github.com/TtMyth123/kit/goqueryKit"
	"github.com/opesun/goquery"
	"testing"
)

func Test_GetBetDetail(t *testing.T) {
	str := `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<title></title>
    <link href="css/style.css?v=v190712" rel="stylesheet" type="text/css" />
<link href="skin/dsn2/css/style.css?v=v190712" rel="stylesheet" type="text/css" />
<link href="skin/dsn2/css/balls.css?v=v190712" rel="stylesheet" type="text/css" />
<link href="skin/dsn2/css/listmain.css?v=v190712" rel="stylesheet" type="text/css" />
<link href="skin/dsn2/css/skins.css?v=v190712" rel="stylesheet" type="text/css" />
<script language="javascript" type="text/javascript" src="js/main.js"></script>
<script language="javascript" type="text/javascript" src="js/utils.js"></script>
<script language="javascript" type="text/javascript" src="js/validator.js"></script>
<script type="text/javascript" src="js/jquery.js"></script>
<script type="text/javascript">
jQuery.noConflict();
</script>
</head>
<body class="skins_">
<!--onContextMenu="return false;" onSelectStart="return false"--><div id="design">
		  <div id="form_search">
	      <h1>注单列表 [第2023-050期]>特码 01</h1>
		  <label>注单</label>
		  <select name="is_stock" onchange="showRateList()" id="is_stock">
				  <option value="0">全部</option>
				  <option value="1" >实占货量</option>
				 <option value="2" >手动补货</option>
				  <option value="3" >自动补货</option>
			 </select>
			 <input type="hidden" name="region" value="all" id="region" />
		<!--	<label>会员</label>
			 <select name="region" onchange="showRateList()" id="region">
				 <option  value="all">全部</option>
				 <option  value="A">A盘</option>
				 <option  value="B">B盘</option>
				 <option  value="C">C盘</option>
		    </select>-->
			  <label>排序</label>
			  <select name="ordercol" onchange="showRateList()" id="slordercol" >
				  <option  value="created">下注时间</option>
				  <option  value="ucash">实占金额</option>
				  <option selected="selected" value="cash">下注金额</option>
			  </select>
			  <label>金额</label>
			  <input type="text" size="5" placeholder="最小值" value="100" name="cashmin">-
			  <input type="text" size="5" placeholder="最大值" value="400000" name="cashmax">
			  <input type="hidden" name="lotcode" value="bjamlh_tm" />
			  <input type="hidden" name="item_id" value="64080,64129" />
			  <input type="hidden" name="lottery_id" value="0" />
			  <input name="查询" type="button" onclick="showRateList()"  class="btn" value="查询"/>
		 <div id="goback"><a href="javascript:void(0);" onclick="window.history.back();return false;" target="main">返回上一页</a></div>
		 </div>
		 <div id="main">
		  <!-----------------订单详细页--------------------->
<table width="98%" cellpadding="8" cellspacing="1" class="tab" id="tab1">
    <tr>
      <th>单号</th>
      <th>下注时间</th>
      <th>帐号(名称)</th>
      <th>退水</th>
	  <th>盘口</th>
      <th>内容</th>
      <th>赔率</th>
      <th>金额</th>
	  <th>股份</th>
	  <th>金额</th>
      <!--<th>结果</th>-->
    </tr>
	    <tr>
      <td> r201921080024315921</td>
      <td >02-19 21:08:00</td>
		<td  class="loginwarn0  edusetting231152"><span  >dc3**()</span></td>
      <td>6.5</td>
	 <td>A盘</td>
      <td style="text-align:right; padding-right:15px;">澳门六合彩 2023-050 期  <span class="red">特码</span>  01</td>
      <td><strong class="red">47.5</strong></td>
      <td>222.00</td>
	  <td>		  95
		  %
		  		  <div class="detail-report" pid="/index.php?Y29udHJvbGxlcj1yZXBvcnQmYWN0aW9uPXVzZXJvcmRlcmluZm8mb3JkZXJpZD03OTc1MTE0NDMmc3RpbWU9MTY3NjgxMjA4MA==/&lotcode=bjamlh_tm">
			  [<a href="javascript:;">详细</a>]<div></div>
		  </div>
	  </td>
	  <td>210.90</td>
     <!-- <td >-215.4954</td>-->
    </tr>
	    <tr>
      <td> r201921072798654369</td>
      <td >02-19 21:07:27</td>
		<td  class="loginwarn0  edusetting217751"><span  >m10**()</span></td>
      <td>2.93</td>
	 <td>A盘</td>
      <td style="text-align:right; padding-right:15px;">澳门六合彩 2023-050 期  <span class="red">特码</span>  01</td>
      <td><strong class="red">47.5</strong></td>
      <td>100.00</td>
	  <td>		  5
		  %
		  		  <div class="detail-report" pid="/index.php?Y29udHJvbGxlcj1yZXBvcnQmYWN0aW9uPXVzZXJvcmRlcmluZm8mb3JkZXJpZD03OTc1MDM5OTEmc3RpbWU9MTY3NjgxMjA0Nw==/&lotcode=bjamlh_tm">
			  [<a href="javascript:;">详细</a>]<div></div>
		  </div>
	  </td>
	  <td>5.00</td>
     <!-- <td >-97.07</td>-->
    </tr>
	    <tr>
      <td> r201920441512892408</td>
      <td >02-19 20:44:15</td>
		<td  class="loginwarn0  edusetting217756"><span  >100**()</span></td>
      <td>8.05</td>
	 <td>B盘</td>
      <td style="text-align:right; padding-right:15px;">澳门六合彩 2023-050 期  <span class="red">特码</span>  01</td>
      <td><strong class="red">45</strong></td>
      <td>100.00</td>
	  <td>		  5
		  %
		  		  <div class="detail-report" pid="/index.php?Y29udHJvbGxlcj1yZXBvcnQmYWN0aW9uPXVzZXJvcmRlcmluZm8mb3JkZXJpZD03OTcyNjc1ODMmc3RpbWU9MTY3NjgxMDY1NQ==/&lotcode=bjamlh_tm">
			  [<a href="javascript:;">详细</a>]<div></div>
		  </div>
	  </td>
	  <td>5.00</td>
     <!-- <td >-91.95</td>-->
    </tr>
		 <tr>
	<td colspan="9" style="text-align:right;" bgcolor="#FFF5C5" class="brown"><strong>笔数：3</strong></td>
	<td bgcolor="#FFF5C5" class="brown">220.90</td>
	<!--<td bgcolor="#FFF5C5" class="brown">-404.52</td>-->
  </tr>
	</table>
	 <div class="clearfix">
		 <div id="data_page" style="float:right; clear:none;">
			 <ul id="pagenav">
<li ><a href="#" class="page_txt">共<b>3</b>条</a></li>
<li class="disabled"><a href='#'>&#171; 上一页&nbsp;</a></li>
<li class="current"><a href='#'>1</a></li>
<li class="disabled"><a href='#'>下一页 &#187;</a></li>
<li><a class="page_txt" onclick="window.location='/index.php?Y29udHJvbGxlcj1yZXBvcnQmYWN0aW9uPWRldGFpbGJ5aXRlbSZyZWdpb249QUxMJmlzX3N0b2NrPTAmb3JkZXJieT1kZXNjJm9yZGVyY29sPWNhc2gmY2FzaG1pbj0xMDAmY2FzaG1heD00MDAwMDAmbG90dGVyeV9pZD0wJml0ZW1faWQ9NjQwODAlMkM2NDEyOSZsb3Rjb2RlPWJqYW1saF90bSZvcGVuX2lkPTcwNTg5MjU4JmlzX2hpcz0wJnNob3d0eXBlPWxzdA==/&page='+(document.getElementById('pagenav_gopagenum').value-1);" href="#">跳转到</a> <input type="text" value='1' id="pagenav_gopagenum" name="custompage" size="2" onkeydown="if(event.keyCode==13) {window.location='/index.php?Y29udHJvbGxlcj1yZXBvcnQmYWN0aW9uPWRldGFpbGJ5aXRlbSZyZWdpb249QUxMJmlzX3N0b2NrPTAmb3JkZXJieT1kZXNjJm9yZGVyY29sPWNhc2gmY2FzaG1pbj0xMDAmY2FzaG1heD00MDAwMDAmbG90dGVyeV9pZD0wJml0ZW1faWQ9NjQwODAlMkM2NDEyOSZsb3Rjb2RlPWJqYW1saF90bSZvcGVuX2lkPTcwNTg5MjU4JmlzX2hpcz0wJnNob3d0eXBlPWxzdA==/&page='+(this.value-1); return false;}" />页</li></ul>

		 </div>
	 </div>
</div>
</div>
<!--用户报表统计显示-->
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<!--报表信息弹出层-->
<div id='msgDiv_rpstatbox' style="width: 490px;" class="pop_box pop_liushui">
</div>
<link href="skin/dsn/css/edu.css?v=v190712" rel="stylesheet" type="text/css" />
<script type="text/javascript" src="js/edu.js?v=v190712"></script>
<script type="text/javascript">
onloadEvent(showtable);
function showRateList(){
    var lotcode=jQuery('#form_search').find('input[name="lotcode"]').val();
    var itemid=jQuery('#form_search').find('input[name="item_id"]').val();
    var lotid=jQuery('#form_search').find('input[name="lottery_id"]').val();
    var region=jQuery('#form_search').find('input[name="region"]').val();
    var cashmin=jQuery('#form_search').find('input[name="cashmin"]').val();
    var cashmax=jQuery('#form_search').find('input[name="cashmax"]').val();
    var ordercol=jQuery('#slordercol').val();
    var is_stock=jQuery('#is_stock').val();
   var qstr='&lotcode='+lotcode+'&item_id='+itemid+'&lottery_id='+lotid+'&region='+region+'&is_stock='+is_stock+'&cashmin='+cashmin+'&cashmax='+cashmax+'&ordercol='+ordercol;
   //qstr+="&is_stock="+$("is_stock").value+"&region="+$("region").value+"&ordercol="+$("slordercol").value;
   url='/index.php?Y29udHJvbGxlcj1yZXBvcnQmYWN0aW9uPWRldGFpbGJ5aXRlbQ==/'+qstr;
   gotourl(url);
}
jQuery('document').ready(function(){
    jQuery('.detail-report').bind('click',function () {
        detailReport(jQuery(this));
    });
    jQuery('body').bind('click',function () {
        jQuery('.detail-table').remove();
    });
})
//显示股份列表
function detailReport(obj){
    jQuery('.detail-table').remove();
    var url =obj.attr("pid") + '&is_ajax=1';
    // Ajax.call(url, '', detailReport_ok, 'POST', 'TXT');

    jQuery.ajax({
        type: 'POST',
        url: url,
        data: '',
        success: function (re) {
            // var data = JSON.parse(re);
            var data = re;
            var msg=data.msg;
            if (data.code == 'ok') {
                var divwidth=msg.length * 80 - 1 + 50;
                var html='<div class="detail-table" style="width: '+ divwidth +'px">\n<ul class="tt"><li>账号</li><li>股份</li><li>退水</li></ul>';
                for (var i=0;i<msg.length; i++) {
                    html+='<ul><li>'+msg[i].username+'</li><li>'+msg[i].stk+'</li><li>'+msg[i].rbt+'</li></ul>'
                }
                html+='</div>';
                //console.log(html)
                obj.find('div').html('');
                obj.find('div').append(html);
            } else {
                alert(data.msg);
            }
        },
        dataType: 'json'});
}
</script>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<script type="text/javascript">
//表格颜色
var overcolor='#DFDFDF'; //鼠标经过颜色
var colordef='#fff';  //默认颜色
var colordef2='#fff';  //默认颜色2
var colordown='#ADDFDE';  //按下去的颜色
var oldColor='#F3F3F3';
</Script></body>
</html>`

	dom, e := goquery.ParseString(str)
	tableNodes := dom.Find(".tab")
	arrStrTable := goqueryKit.GetTableNodes2Arr(tableNodes)

	fmt.Println(e, arrStrTable)
}
