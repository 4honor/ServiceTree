define(["jquery", "WdatePicker", "milk"], function($, datePicker, milk) {
	var o = {};
	
	o.init = function(){
		//启动milk插件css
		mlk.contant();
		
		//datePicker
		$(document).on('focus','.datePickker .edit-input',function(){
			var thisVal = $(this).val();
			WdatePicker({
				startDate:thisVal,//'%y-%M-01 00:00:00'
				dateFmt:'yyyy-MM-dd HH:mm:ss',
				alwaysUseStartDate:true
			});	
		});
		
		//点击按钮
		$(document).on('click','#table_inner .btn',function(){
			if(!$(this).hasClass('btn-approve')){//操作按钮
				if($(this).hasClass('clear-back')){
					if(window.confirm('你确定要删除虚拟机吗？')){}
					else{
						return false;	
					}	
				}
				var thisName = $(this).attr('name'),
					thisListId = $(this).parents('tr').attr('id'),
					thisBtnData = {
						id: thisListId,
						op: thisName	
					},
					that = $(this);
				mlk.ajax({
					method: 'POST',
					url: '/operation/',
					data: thisBtnData,
					callback: function(da){
						var myData = JSON.parse(da);
						if(myData.errno == 0){
							if(that.hasClass('close-ma')){
								var allBtns = that.parents('tr').find('.btn:not(.open-ma,.clear-back)');
								mlk.isClick({
									flag: false,
									dom: allBtns	
								});	
							}
							else if(that.hasClass('open-ma')){
								var allBtns = that.parents('tr').find('.btn');
									mlk.isClick({
										flag: true,
										dom: allBtns	
									});		
							}
							else{}
							alert('操作成功！')	
						}
						else{
							alert(myData.errmsg);	
						}
					}	
				});	
				
			}
			else{//审批按钮
				if($(this).hasClass('btn-pass')){
					var yeOrNo = 'y';
				}
				else if($(this).hasClass('btn-fail')){
					var yeOrNo = 'n';
				}
				else{}
				
				var approData = {
					id: $(this).parents('tr').attr('id'),
					yorn: yeOrNo	
				};
				var btnThis = $(this);
				mlk.ajax({
					method: 'POST',
					url: '/approvalaction/',
					data: approData,
					callback: function(da){
						var myData = JSON.parse(da);
						if(myData.errno == 0){
							console.log(btnThis.parents('tr'));
							mlk.removeDom({//删除
								dom: [btnThis.parents('tr').find('td')]	
							});
						}
						else{
							alert(myData.errmsg);	
						}
					}
				});
				
			}
		});
		
		
	}
	
	return o;	
});


