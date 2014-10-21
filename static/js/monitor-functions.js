//all functions
	//时间挫转换时间
	function getLocalTime(nS) {     
	   return new Date(parseInt(nS) * 1000).toLocaleString().replace(/年|月/g, "-").replace(/日/g, " ");      
	}  
	
	//创建div
	function creatDiv(ser,host,stime,etime,x,num,name,uni,arry,flag,aggr){
		var newDiv = $('#outer .drag_cont:eq(0)').clone();
		$('#outer').append(newDiv);
		newDiv.find('h3 span').text(ser);
		newDiv.find('h3 span').attr('host',host);
		newDiv.find('input.time-search:eq(0)').val(stime);
		newDiv.find('input.time-search:eq(1)').val(etime);
		newDiv.slideDown(500,function(){
			var newDivTop = newDiv.offset().top;
			$('html,body').animate({'scrollTop':newDivTop},500);
		});
		thisNewDiv = newDiv;
		var dateInput1 = newDiv.find('.time-search:eq(0)'),
			dateInput2 = newDiv.find('.time-search:eq(1)');
		newDiv.find('.aggregate option').each(function(){//求和均值初始化
			if($(this).val() == aggr){
				$(this).attr('selected','true');	
			}	
		});
		datePicker(dateInput1,arry[0],arry[2]);
		datePicker(dateInput2,arry[1],arry[3]);
		var oChart = newDiv.find('.cont_box');
		creatHighcharts(oChart,ser,x,num,name,uni,flag);
	}
	
	//时间选择
	function datePicker(obj,year,time){
		$(obj).bind('focus',function(){
			var thisVal = $(this).val();
			var date = thisVal ? thisVal.replace('-',' ') : (year + ' ' + time);
			WdatePicker({
				startDate:date,//'%y-%M-01 00:00:00'
				dateFmt:'yyyy/MM/dd-HH:mm:ss',
				alwaysUseStartDate:true
			});
		});	
	}
	
	//创建每个input（host）
	function creatInput(da){
		var allHtml = '';
		for(var i = 0 ; i < da.length; i++){
			var //inputCheck = da[i].checked == 'true' ? 'checked' : '',
				inputName = da[i].name,
				inputValue = da[i].value,
				allInputHtml = '';
			var inputCheck = da[i].checked,
				checked = (inputCheck == true ? 'checked' : '');
			
			for(var j = 0; j < inputValue.length; j++){//遍历input
				var inputHtml = '<label class="checkbox"><input type="checkbox" name="' + inputName + '"' + checked + ' value="' + inputValue[j] + '">' + inputValue[j] + '</label>';
				allInputHtml += inputHtml;	
			}
			
			var divHtml = '<div class="btn-group" style="margin-right:5px;"><button type="button" style="width:77px;" class="btn btn-default dropdown-toggle" data-toggle="dropdown">' + inputName + '<span class="caret" style="margin-left:5px;"></span></button><div class="dropdown-menu" role="menu">' + allInputHtml + '</div></div>';
			allHtml += divHtml;
		}
		
		thisNewDiv.find('.select_outer').append(allHtml);
	}
	
	//创建Highcharts
	function creatHighcharts(obj,attr,x,num,name,uni,flag){//初始化时创建数据表
		if(!flag){
			name = name.replace('@','');
		}
		$(obj).attr('valflag',flag);
		var highOptions = {
				chart: { 
					zoomType: 'xy'
				},
				title: {
					text: name,
					x: -20 //center
				},
				subtitle: {
					text: '',
					x: -20
				},
				xAxis: {
					categories: x,
					lineWidth :0,
					tickWidth:0,
					tickInterval: 3600 * 1000,
					labels:{
						enabled:false
					}
				},
				yAxis: {
					title: {
						text: '单位 (' + uni + ')'
					},
					plotLines: [{
						value: 0,
						width: 1,
						color: '#808080'
					}]
				},
				tooltip: {
					shared: true,
          			crosshairs: true
				},
				legend: {
					layout: 'vertical',
					align: 'right',
					verticalAlign: 'middle',
					borderWidth: 0
				},
				exporting: {
					enabled: false	
				},
				series: [
				{
					name: attr,
					data: num
				}]
			}
		obj.attr('highchart',JSON.stringify(highOptions));//把属于自己的曲线属性保存下来，方便以后的操作
		obj.highcharts(highOptions);
	}
	
	//重新上传参数拿数据（核心方法）
	function seachNext(that){
			var thisBeginTime = that.parents('.cont-mess').find('.time-search').eq(0).val(),
				thisEndTime = 	that.parents('.cont-mess').find('.time-search').eq(1).val(),
				thisMetric = that.parents('.cont-mess').find('h3 span').text();
			//var thatHost = that.parents('.drag_cont').find('h3 span').attr('host')
            var tags = '';
			that.parents('.cont-mess').find('.dropdown-menu input[type=checkbox]').each(function(){
				if(this.checked){
					var thisName = $(this).attr('name'),
						thisVal = $(this).val();
                    if (tags.indexOf(thisName + '=') == -1 ) {
                        tags += ':' + thisName + '=';
                    }
                    tags +=  thisVal + ',';
				}
			});
            tags = tags.substr(1,tags.length);
			tags = tags.substr(0,tags.length-1);
			tags = tags.replace(',:',':');
			var thisChart = that.parents('.drag_cont').find('.cont_box');
			var objHighChart = JSON.parse(thisChart.attr('highchart'));//取出属于自己的曲线属性
			var tags_condition = ''//处理host
			if (tags != '') {
				tags_condition = '@' + tags
			}else {
				tags_condition = '@host='
			}
			objHighChart.series = [];
			var compare ='';//处理同比和环比
			that.parents('.cont-mess').find('.more_checked input:checked').each(function(){
				if(this.checked){
					var eachCompare	= $(this).val();
					compare += eachCompare + ',';
				}	
			});
			compare = compare.substr(0,compare.length-1);
			var aggregate = that.parents('.cont-mess').find('.aggregate').val();
			$.ajax({
				type: 'GET',
				url: 'http://10.231.146.171/api/ts?metric=' + thisMetric + tags_condition + '&stime=' + thisBeginTime + '&etime=' + thisEndTime + '&compare=' + compare +'&aggregate=' + aggregate,
				data: null,
				dataType: "jsonp",
				success: function(da){
					if(da){
						var loading = that.parents('.drag_cont').find('.progress');
						endLoading(loading);
						var xTime = [];//x轴坐标解析
						for(var l = 0; l < da[0].data.length; l++){
							var n = da[0].data[l];
							for(var i in n){
								xTime.push(getLocalTime(parseInt(i)));//处理坐标
							}
						}
						for(var n = 0; n < da.length; n++){
							var yNum = [];
								for(var i = 0; i < da[n].data.length; i++){
									var j = da[n].data[i];
									for(var attr in j){
										yNum.push(j[attr]);//处理每个点的值
									}
								}
								var thisDivHighChar = that.parents('.drag_cont').find('cont_box').attr('valflag');
								if(!thisDivHighChar){
									var lineName = da[n].name.split('@');
								}
								var	newLine = {
											name: lineName[1] == '' ? thisMetric : lineName[1],
											data: yNum
										};
								//alert(xTime);
								objHighChart.series.push(newLine);
							}
							objHighChart.xAxis.categories = xTime;
							newHighCharts(thisChart,objHighChart);
					}
				}
			});
	}
	
	//重新给HighCharts参数
	function newHighCharts(obj,hop){
		 //alert(hop);
		 obj.highcharts(hop);
		 //hop.series.splice(1,hop.series.length-1);
		 //var jsonStr = JSON.stringify(hop);
		 ///obj.attr('highchart',jsonStr);//把属于自己的曲线属性保存下来，方便以后的操作
	}

	//启动left-btn动画
	function animateBtnAdd(){
		if($('#btn_add').attr('aniflag') != 'false'){
			$('#btn_add').animate({'width':50,'height':50,'left':'50%','top':300,'line-height':'45px','font-size':'30px'},1000,function(){
				$(this).animate({'top':300,'left':0},500,function(){
					$(this).removeClass('add_block').addClass('add_left_block');
					window.setTimeout(function(){
						$('#btn_add').stop().animate({'left':-25},800);
						$('#btn_add').attr('aniflag','false');
						$('#btn_add').hover(
							function(){
								$('#btn_add').stop().animate({'left':0},300);	
							},
							function(){
								$('#btn_add').stop().animate({'left':-25},300);	
							}
						);	
					},500);	
				});	
			});	
		}
	}
	
	//点击同比环比限制
	var limitCount = 20;
	function clickMoreData(obj,that){
		var clickLength = obj.parents('.drag_cont').find('.dropdown-menu').children('input:checked').length,
			sibInputs = obj.parent('label').siblings('label').children('input');
		if(that.checked){
			if(clickLength > limitCount){
				sibInputs.attr('disabled','disabled');	
			}
			else{
				sibInputs.removeAttr('disabled');	
			}	
		}
		else{
			sibInputs.removeAttr('disabled');	
		}
		
	}
	
	//点击host限制
	function clcikHostData(obj,that){
		var clickLength = obj.parents('.drag_cont').find('.more_checked input:checked').length,
			sibInputs = obj.parent('label').siblings('label').children('input'),
			sibCheckLeng = obj.parents('.dropdown-menu').find('input:checked').length;
		if(that.checked){
			if(clickLength > limitCount){
				sibInputs.attr('disabled','disabled');	
			}
			else{
				sibInputs.removeAttr('disabled');	
			}	
		}
		else{
			sibInputs.removeAttr('disabled');	
		}
		
		if(sibCheckLeng > (limitCount + 1)){
			obj.parents('.drag_cont').find('.more_checked input').removeAttr('checked').attr('disabled','disabled');
			sibInputs.removeAttr('disabled');	
		}
		else{
			obj.parents('.drag_cont').find('.more_checked input').removeAttr('disabled');
		}
	}
	
	//time按钮
	function timeSearch(time){
		var dateNow = new Date(),
			dateHours = dateNow.getHours(),
			dateMin = dateNow.getMinutes(),
			dateYear = dateNow.getFullYear(),
			dateMon = dateNow.getMonth() + 1,
			dateDay = dateNow.getDate(),
			dateSec = dateNow.getSeconds();
		dateHours = dateHours.toString().length != 1 ? dateHours : ('0' + dateHours);
		dateMin = dateMin.toString().length != 1 ? dateMin : ('0' + dateMin);
		dateSec = dateSec.toString().length != 1 ? dateSec : ('0' + dateSec);
		dateMon = dateMon.toString().length != 1 ? dateMon : ('0' + dateMon);
		dateDay = dateDay.toString().length != 1 ? dateDay : ('0' + dateDay);
		var inputDate = dateYear + '/' + dateMon + '/' + dateDay + '-' + dateHours + ':' + dateMin + ':' + dateSec;
		
		var preTime = new Date((dateNow/1000-time)*1000),
			preHours = preTime.getHours(),
			preMin = preTime.getMinutes(),
			preYear = preTime.getFullYear(),
			preMon = preTime.getMonth() + 1,
			preDay = preTime.getDate(),
			preSec = preTime.getSeconds();
		preHours = preHours.toString().length != 1 ? preHours : ('0' + preHours);
		preMin = preMin.toString().length != 1 ? preMin : ('0' + preMin);
		preSec = preSec.toString().length != 1 ? preSec : ('0' + preSec);
		preMon = preMon.toString().length != 1 ? preMon : ('0' + preMon);
		preDay = preDay.toString().length != 1 ? preDay : ('0' + preDay);
		var inputPreDate = preYear + '/' + preMon + '/' + preDay + '-' + preHours + ':' + preMin + ':' + preSec;
		var timeObj = {};
		timeObj.nowTime = inputDate;
		timeObj.preTime = inputPreDate;
		return timeObj;
	}
	
	//给时间input重新赋值
	function timeValueInput(obj,nowTime,preTime){
		var thisBeginTime = obj.parents('.cont-mess').find('.time-search').eq(0),
			thisEndTime = 	obj.parents('.cont-mess').find('.time-search').eq(1);
		thisBeginTime.val(preTime);
		thisEndTime.val(nowTime);
	}
	
	//进度条初始化
	function beginLoading(obj){
		window.setTimeout(function(){
			obj.hide();
			obj.find('.progress-bar').css({'width':0});	
		},1000)
	}
	
	//进度条动画
	function loadingBar(obj,time,endWidth,callback){
		obj.show();
		obj.find('.progress-bar').stop().animate({'width':endWidth},time);
	}	
	
	function endLoading(loading){
		loadingBar(loading,200,'100%');//进度条结尾动画
		beginLoading(loading);
	}
