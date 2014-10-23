define(["jquery", "milk", "datepicker", "allfunction", "highchart"], function($, milk, datepicker, allfunction, highchart) {
	var o = {};
	  
  	o.init = function(){
		
		/*点击下拉*/
		$(document).on('click','.dropdown-toggle',function(ev){
			$('.dropdown-menu').hide();
			$(this).next('.dropdown-menu').show();	
		});
		$(document).on('click','.btn-group',function(ev){
			ev.stopPropagation();
		});
		$(document).click(function(ev){
			$('.dropdown-menu').hide();
			//$('.select_div').hide();
		});
		/*点击下拉end*/
		
		/*绑定拖拽*/
		/*$( "#outer" ).sortable({
			revert: true,
			delay: 300 ,
			opacity: 0.6,
		});*/
		 //$( ".dropdown-menu" ).disableSelection();
		/*绑定拖拽end*/
		
		/*删除*/
		$(document).on('click','.close',function(){
			$(this).parents('.drag_cont').slideUp(500,function(){
				var thisId = $(this).attr('id');
				$('.' + thisId)[0].checked = false;
				$(this).remove();	
			});	
		});
		/*删除end*/
		
		//刷新
		$(document).on('click','.btn_sx',function(){
			var loading = $(this).parents('.drag_cont').find('.progress');
			loadingBar(loading,200,'70%');
			var inputDte = timeSearch(3600);
			$(this).parents('.form-group').find('.time-search:eq(1)').val(inputDte.nowTime);
			var that = $(this);
			seachNext(that);
			
		});
	
		/*点击checkbox查询状态*/
		$('.monitor_inner').on('click','input[type=checkbox]',function(){
			if(this.checked){//选中状态
			
				var serVal = $(this).attr('name'),
					thisClass = $(this).attr('class'),
					hostVal = '';
					
				var thisHelp = $(this).parents('p').find('.for_help'),
					helpName = thisHelp.attr('name'),
					helpType = thisHelp.attr('type'),
					helpComment = thisHelp.attr('comment');
					
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
				
				var preTime = new Date((dateNow/1000-3600)*1000),
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
				
				var	timeStringBegin = preHours + ':' + preMin + ':' + preSec,//起始时间
					timeStringEnd = dateHours + ':' + dateMin + ':' + dateSec,//结束时间
					timeYearBeagin = preYear + '/' + preMon + '/' + preDay,//起始年月
					timeYearEnd = dateYear + '/' + dateMon + '/' + dateDay;//结束年月
				var dateArry = [timeYearBeagin,timeYearEnd,timeStringBegin,timeStringEnd],
					dataBegin = preYear + ',' + preMon + ',' + preDay + ',' + preHours + ',' + preMin + ',' + preSec;
				if(hostVal){
					var valFlag = true;	
				}else{
					var valFlag = false;	
				}
				var aggregate = 'num';
				if(serVal){
					var thisNewDiv = '';//保留创建的那个div
					$('.use_input').val('');
					$.ajax({
						type: 'GET',
						//url: 'http://10.231.146.171/api/ts?metric=' + serVal + '@host=' + hostVal + '&stime=' + timeYearBeagin + '-' + timeStringBegin + '&etime=' + timeYearEnd + '-' + timeStringEnd +'&aggregate=sum',
						url: 'http://10.231.146.171/api/ts2?metric=' + serVal + '@host=' + hostVal + '&stime=2014/10/19-18:39:07&etime=2014/10/22-18:39:07&aggregate=sum',
						data: null,
						dataType: "jsonp",
						success: function(data){
							if(data){
								/*var xTime = [],
									yNum = [],
									serviceName = '',
									serviceUnit = '';
								serviceName = data[0].name;
								serviceUnit = data[0].unit;
								for(var i = 0; i < data[0].data.length; i++){
									var j = data[0].data[i]
									for(var attr in j){
										//xTime.push(getLocalTime(parseInt(attr)));//处理坐标
										xTime.push(attr);
										yNum.push(j[attr]);//处理每个点的值
									}	
								}
								xTime = xTime[0]
								//console.log(xTime);
								creatDiv(dataBegin,serVal,hostVal,thisClass,(timeYearBeagin + '-' + timeStringBegin),(timeYearEnd + '-' + timeStringEnd),xTime,yNum,serviceName,serviceUnit,dateArry,valFlag,aggregate,helpType,helpComment,helpName);
								$('.mask_srch').find('.sumOrAvg option:eq(0)').attr('selected','true');
								var ns = $('#machine_drop').html().split(",");
								ns = ns.slice(0,ns.length);
								ns.push("resource:machine");*/
								var serviceName = '',
									serviceUnit = data[0].unit;
								
								creatDiv(dataBegin,serVal,hostVal,thisClass,(timeYearBeagin + '-' + timeStringBegin),(timeYearEnd + '-' + timeStringEnd),data,serviceName,serviceUnit,dateArry,valFlag,aggregate,helpType,helpComment,helpName);
								$('.mask_srch').find('.sumOrAvg option:eq(0)').attr('selected','true');
								var ns = $('#machine_drop').html().split(",");
								ns = ns.slice(0,ns.length);
								ns.push("resource:machine");
									
								//当第一步完成后进行第二步ajax请求
								$.ajax({
									type: 'GET',
									url: '/api/tagkv?metric=' + serVal + '@host=' + hostVal + '&ns='+ns.join(),
									data: null,
									dataType: "json",
									success: function(data){
										if(data){
											creatInput(data);
											mlk.dropdownSearch({
												ele: '.dropdown-menu',
												btn: '.btn',
												ischeckall: true
											});
										}
									}
								});
							
								}
							}
						});
					
						//
				}	
			
			//if end		
			}
			else{//取消选中状态
				var thisClass = $(this).attr('class');
				$('#' + thisClass).slideDown(500,function(){
					$(this).remove();	
				});	
			}
		});
		
		//$('.time-search').bind('blur')
		
		/*点击查询状态end*/
		
		/*全选*/
		$(document).on('change','.mlk-check-all',function(){
			  var obj = $(this).parents('.mlk-dropdown-search').next('.mlk-dropdown-outer').find('input[type=checkbox]').eq(0);
			  var loading = $(this).parents('.drag_cont').find('.progress');
			  loadingBar(loading,200,'70%');
			  seachNext(obj);
		});
		
		/*点选下拉*/
		$(document).on('change','.dropdown-menu input[type=checkbox]',function(){
			if(!$(this).hasClass('mlk-check-all')){//排除全选按钮
				var obj = $(this),
					that = this;
				var loading = $(this).parents('.drag_cont').find('.progress');
				loadingBar(loading,200,'70%');
				seachNext(obj);
				clcikHostData(obj,that);
			}
		});
		
		/*end点选下拉*/
		
		/*查询时间*/
		$(document).on('click','.time_search',function(){
			var that = $(this);
			var loading = $(this).parents('.drag_cont').find('.progress');
			loadingBar(loading,200,'70%');
			seachNext(that);
		});
		/*end查询时间*/
		
		//弹出层
		$('#btn_add').click(function(){
			$('.mask_srch').fadeIn(500,function(){
				$('.input_frst').focus();	
			});
		});
		
		$('#close_alert').click(function(){
			$('.mask_srch').fadeOut(500);	
		});
		
		//给同比环比绑定
		$(document).on('change','.more_checked input',function(){
			var obj = $(this),
				that = this;
			clickMoreData(obj,that);
			seachNext(obj);
		});
		
		//给聚合计算绑定
		$(document).on('change','.aggregate',function(){
			var obj = $(this);
			var loading = $(this).parents('.drag_cont').find('.progress');
			loadingBar(loading,200,'70%');
			seachNext(obj);	
		});
		
		//给时间选择绑定
		$(document).on('click','.other_btn_outer a',function(){
			var thisTime = parseFloat($(this).attr('time')),
				time = timeSearch(thisTime),
				nowTime = time.nowTime,
				preTime = time.preTime,
				that = $(this);
			var loading = $(this).parents('.drag_cont').find('.progress');
			loadingBar(loading,200,'70%');
			timeValueInput(that,nowTime,preTime);
			seachNext(that);
		});
	  
	}
	
	return o;

});
