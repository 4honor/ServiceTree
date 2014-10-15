define(["jquery", "milk", "zTree"], function($, milk, zTree) {
	var o = {};
	
	o.init = function(){
		//启动milk插件css
		mlk.contant();
		
		var setting = {
			data: {
				simpleData: {
					enable: true
				}
			},
			callback: {
				onClick: function(ev){
					var treeObj = $.fn.zTree.getZTreeObj("treeDemo");
					var sNodes = treeObj.getSelectedNodes();
					if (sNodes.length > 0) {
						var my = sNodes[0];
						var node = sNodes[0].getParentNode();
						var nodeType = sNodes[0].getParentNode().ename;
						//var parent = sNodes[0].getParentNode().getParentNode();
						var level = sNodes[0].level;
						console.log(node);
					}	
				} 	
			}
		};

		var zNodes =[
			{ id:1, pId:0, name:"父节点1 - 展开", open:false, iconSkin:"pIcon01"},
			{ id:11, pId:1, name:"父节点11 - 折叠", iconSkin:"pIcon01", ename:"nimei", A:"abc"},
			{ id:111, pId:11, name:"叶子节点111", iconSkin:"icon01"},
			{ id:112, pId:11, name:"叶子节点112", iconSkin:"icon01"},
			{ id:113, pId:11, name:"叶子节点113", iconSkin:"icon01"},
			{ id:114, pId:11, name:"叶子节点114", iconSkin:"icon01"},
			{ id:12, pId:1, name:"父节点12 - 折叠", iconSkin:"pIcon01"},
			{ id:121, pId:12, name:"叶子节点121", iconSkin:"icon01"},
			{ id:122, pId:12, name:"叶子节点122", iconSkin:"icon01"},
			{ id:123, pId:12, name:"叶子节点123", iconSkin:"icon01"},
			{ id:124, pId:12, name:"叶子节点124", iconSkin:"icon01"},
			{ id:13, pId:1, name:"父节点13 - 没有子节点", isParent:true, iconSkin:"pIcon01"},
			{ id:2, pId:0, name:"父节点2 - 折叠", iconSkin:"pIcon01"},
			{ id:21, pId:2, name:"父节点21 - 展开", iconSkin:"pIcon01"},
			{ id:211, pId:21, name:"叶子节点211", iconSkin:"icon01"},
			{ id:212, pId:21, name:"叶子节点212", iconSkin:"icon01"},
			{ id:213, pId:21, name:"叶子节点213", iconSkin:"icon01"},
			{ id:214, pId:21, name:"叶子节点214", iconSkin:"icon01"},
			{ id:22, pId:2, name:"父节点22 - 折叠", iconSkin:"pIcon01"},
			{ id:221, pId:22, name:"叶子节点221", iconSkin:"icon01"},
			{ id:222, pId:22, name:"叶子节点222", iconSkin:"icon01"},
			{ id:223, pId:22, name:"叶子节点223", iconSkin:"icon01"},
			{ id:224, pId:22, name:"叶子节点224", iconSkin:"icon01"},
			{ id:23, pId:2, name:"父节点23 - 折叠", iconSkin:"pIcon01"},
			{ id:231, pId:23, name:"叶子节点231", iconSkin:"icon01"},
			{ id:232, pId:23, name:"叶子节点232", iconSkin:"icon01"},
			{ id:233, pId:23, name:"叶子节点233", iconSkin:"icon01"},
			{ id:234, pId:23, name:"叶子节点234", iconSkin:"icon01"},
			{ id:3, pId:0, name:"父节点3 - 没有子节点", isParent:true, iconSkin:"pIcon01"}
		];
		
		var windowHeight = $(window).height();
		mlk.tree({//启动服务树
			node: zNodes, 
			set: setting, 
			movement: '.tree_cont', 
			func: function(){}, 
			style:{
				width: 300,
				height: windowHeight-190
			}	
		});
		
		mlk.edit({//启动编辑
			ele: '#tag_show',
			func: function(obj){}	
		});
		
		
		$('#check_tab li').bind('click',function(){
			var thisIndex = $(this).index();
			$(this).addClass('active').siblings().removeClass('active');
			$('.odin_check_outer .odin_check_inner').eq(thisIndex).show()
			.attr('isshow','yes').siblings()
			.hide().attr('isshow','no');	
		});
		
		//批量删除table tr
		function removeTableTr( url ){
			var tableEachTr = $('.table tbody tr input[type=checkbox]:checked'),
				list = '';
			if(tableEachTr.length > 0){//确保优选tag
				tableEachTr.each(function(){
					var thisId = $(this).parents('tr').attr('id');
					list += thisId + ',';//拼接id	
				});
				list = list.substring(0,list.length-1);
				mlk.ajax({
					method: 'DELETE',
					url: url + list,
					data: null,
					isloading: true,
					callback: function(da){
						if(da.Success){
							var domArray = [];
							tableEachTr.each(function(){
								var thisTr = $(this).parents('tr');
								domArray.push(thisTr);	
							});
							mlk.removeDom({//删除dom
								dom: domArray	
							});	
						}	
					}
				});
			}
			else{
				alert('请选择需要删除的机器！')	
			}	
		}
		
		//tag 页批量删除
		$(document).on('click','#remove_tag',function(){
			removeTableTr('/v1/tag_meta/');
		});
		
		//machine 页批量删除
		$(document).on('click','#bulk_production',function(){
			removeTableTr('/v1/machine/');
		});
			
		
	}
	
	return o;	
});


