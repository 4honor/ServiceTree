(function(window, undefined) {
		
    var mlk = window.mlk || {};
    if (mlk.contant) {
        return mlk.contant;
    } else {
        window.mlk = mlk; // 注册didi对象到window下
    }
	
	var contant = function(opts) {
        if (!(this instanceof contant)) {
            _contant = new contant(opts);
            return _contant; 
        } else {
            new contant.fn.init(opts);
        }
    };
	
	contant.fn = contant.prototype = {
		constructor: contant,
		init: function(){
			var myCss = '<link href="../../static/css/milk.css" rel="stylesheet">';
			$('head').append(myCss);
		}	
	}
	
	//{method: 'POST',url: 'www.baidu.com',data: data, contentType:contentType,dataType:dataType, isloading:true callback: func}
	//ajax提交方法
	mlk.ajax = function(opts){
		if(!opts){
			return;	
		}
		if(opts.isloading){
			var loadHtml = $('<div class="mlk-loading-outer"><div class="mlk-loading-inner"><div class="rect1"></div><div class="rect2"></div><div class="rect3"></div><div class="rect4"></div><div class="rect5"></div></div></div>');	
			$('body').append(loadHtml);
			var hideTime = 9000,
				timer = setTimeout(function(){
					loadHtml.remove();	
				},hideTime);
		}
		
		$.ajax({
			method: opts.method,
			url: opts.url,
			data: opts.data,
			contentType: opts.contentType ? opts.contentType : '',
			dataType: opts.dataType ? opts.dataType : '',
			success: function(da){
				if(opts.isloading){
					clearTimeout(timer);
					loadHtml.remove();
				}
				opts.callback(da);
			}
		});
	}
	
	//{flag: true, dom: [a,b,c]}
	//是否可点击状态
	mlk.isClick = function(opts){
		if(!opts){
			return;	
		}
		if(opts.dom && opts.dom.length >= 0){
			if(opts.flag === true){
				for(var i = 0 ; i < opts.dom.length; i++){
					$(opts.dom[i]).removeClass('disabled');	
				}
			}else{
				for(var i = 0 ; i < opts.dom.length; i++){
					$(opts.dom[i]).addClass('disabled');	
				}	
			}
		}	
	}
	
	//{ele: './#a',func:function}
	//编辑
	mlk.edit = function(opts){
		var innerFunction = function(){};
		innerFunction.prototype = {
			constructor: innerFunction,
			init: function(opts){
				if(!opts){
					return;	
				}
				var that = this;//把innerFunction方法传到内部供自己使用
				this.creatBtnEdit(opts,that);
			},
			creatBtnEdit: function(opts,that){//创建编辑按钮
				$(opts.ele).each(function(){//判断有内容才能编辑
					if($(this).text()){
						var editEle = $('<a class="glyphicon glyphicon-pencil edit-btn" title="编辑"></a>');
						$(this).append(editEle);
						$(this).css('position','relative');	
						
						editEle.on('click',function(){
							that.creatBtnOkandNo(opts,that,$(this));	
						});
					}	
				});
					
			},
			creatBtnOkandNo: function(opts,that,obj){//创建输入框及确定取消
				var bntOk = $('<a class="glyphicon glyphicon-ok edit-ok" title="确定"></a>'),
					bntNo = $('<a class="glyphicon glyphicon-remove edit-no" title="取消"></a>'),
					editInput = $('<input type="text" class="form-control edit-input" >');
				obj.parents(opts.ele).append(bntOk);
				obj.parents(opts.ele).append(bntNo);
				obj.parents(opts.ele).append(editInput);
				editInput.val(editInput.parents(opts.ele)[0].childNodes[0].nodeValue);
				editInput.focus();
					
				$('.edit-no').on('click',function(){
					that.removeBtnOkandNo(opts,$(this));	
				});
				$('.edit-ok').on('click',function(){
					opts.func($(this));
					that.removeBtnOkandNo(opts,$(this));	
				});
			},
			removeBtnOkandNo: function(opts,obj){//删除输入框及确定取消
				var parentObj = obj.parents(opts.ele);
				parentObj.find('.edit-ok').remove();
				parentObj.find('.edit-no').remove();
				parentObj.find('.edit-input').remove();	
			}
		};
		var newFunc = new innerFunction();
		newFunc.init(opts);
	}
	
	//{dom: [$(dom1),$(dom2)]}
	//删除
	mlk.removeDom = function(opts){
		if(!opts){
			return;	
		}
		for(var i = 0; i < opts.dom.length; i++){
			opts.dom[i].slideUp(400,function(){
				$(this).remove();	
			});	
		}
	}
	
	//{title: '',html: html,btnfunc: function}
	//alert box
	mlk.alertBox = function(opts){
		if(!opts){
			return;	
		}
		var innerFunction = function(opts){};
		
		innerFunction.prototype = {
			constructor: innerFunction,
			init: function(opts){
				var that = this;
				this.creatBox(opts,that);
			},
			creatBox: function(opts,taht){
				var alertHtml = $('<div class="mlk-alert" style="display:none;"><div class="mlk-alert-box"><h3>' + opts.title + '<span class="glyphicon glyphicon-remove milk-close"></span></h3><div class="mlk-alert-box-inner">' + opts.html + '<div class="mlk-alert-btn-outer clear"><a class="btn btn-primary left milk-ok">确定</a><a class="btn btn-default right milk-close">取消</a></div></div></div></div>');
				$('body').append(alertHtml);
				alertHtml.fadeIn(400);
				$('.milk-close').on('click',function(){
					taht.btnclose($(this));	
				});
				$('.milk-ok').on('click',function(){
					opts.btnfunc($(this));
					taht.btnclose($(this));	
				});	
			},
			btnclose: function(obj){
				obj.parents('.mlk-alert').fadeOut(400,function(){
					$(this).remove();	
				});	
			}
		}
		var newFunc = new innerFunction();
		newFunc.init(opts);
	}
	
	//case
	mlk.eachStatus = function(name){
		var caseName = '';
		switch(name){
			case 1:
				caseName = '<span>使用中</span>';
				break;
			case 2:
				caseName = '<span>已回收</span>';
				break;
			case 3:
				caseName = '<span title="请咨询管理员">新建失败</span>';
				break;
			case 4:
				caseName = '<span>等待审批</span>';
				break;
			case 5:
				caseName = '<span title="大约1小时">制作中</span>';
				break;
			case 6:
				caseName = '<span title="请咨询管理员">审批未通过</span>';
				break;	
		}
		return caseName;	
	}
	
	//{node: node, set: set, movement: ./#div, func: func, style:{width:width,height:height}}
	//tree
	mlk.tree = function(opts){
		if(!opts){
			return;	
		}
		var innerFunction = function(opts){}
		innerFunction.prototype = {
			constructor: innerFunction,
			init: function(opts){
				this.creatTree(opts);
			},
			creatTree: function(opts){
				var treeDiv = $('<div class="tree_menu"></div>'),
					innerTree = $('<ul id="treeDemo" class="ztree"></ul>'),
					btn = $('<a class="btn_tree_open_close btn btn-default" id="tree_btn_falg"><span class="glyphicon glyphicon-th-list"></span></a>'),
					myBody = $('body');
				treeDiv.append(innerTree);
				myBody.append(treeDiv);
				myBody.append(btn);
				var treeWidth = opts.style.width || 300,
					treeHeight = opts.style.height || 400;	
				treeDiv.css({'width':treeWidth,'height':treeHeight});
				$.fn.zTree.init(innerTree, opts.set, opts.node);
				var winWidth = $(window).width();
				$(opts.movement).css({'width':(winWidth - treeWidth - 70),'margin-left':(treeWidth + 20),'height':opts.style.height});
				this.bindClick(opts,btn,treeDiv,treeWidth,treeHeight,winWidth);
			},
			bindClick: function(opts,btn,obj,width,height,winWidth){
				btn.attr('isshow','true');
				btn.bind('click',function(){
					if($(this).attr('isshow') == 'true'){
						$(opts.movement).stop().animate({'width':(winWidth - 90),'margin-left':40},400);
						obj.stop().animate({'width':20,'height':20},400,function(){
							$(this).hide();
							btn.attr('isshow','false');
							if(opts.func){
								opts.func();
							}
						});	
					}
					else{
						obj.show();
						$(opts.movement).stop().animate({'width':(winWidth - width - 70),'margin-left':(width + 20)},400);
						obj.stop().animate({'width':width,'height':height},400,function(){
							btn.attr('isshow','true');
						});	
					}
				});	
			}
		}
		
		var newFunc = new innerFunction();
		newFunc.init(opts);	
	}
	
	//{dom:[a,b,c],isdele:true}
	mlk.editSelect = function(opts){
		var innerFunction = function(){};
		innerFunction.prototype = {
			constructor: innerFunction,
			init: function(opts){
				if(!opts){
					return;	
				}
				this.creatSelect(opts);
			},
			creatSelect: function(opts){
				var that = this;
				for(var i = 0; i < opts.dom.length; i++){
					opts.dom[i].each(function(){
						var selectParent = $(this).parent(),
							oldId = $(this).attr('id'),
							oldName = $(this).attr('name'),
							oldClass = $(this).attr('class'),
							oldOption = $(this).html(),
							inputOuter = $('<div class="mlk-edit-input-outer"></div>'),
							inputMask = $('<div class="mlk-edit-mask"></div>'),
							inputSpan = $('<span class="mlk-edit-down"></span>'),
							inputDele = $('<span class="glyphicon glyphicon-remove mlk-edit-remove" title="删除该条选项"></span>'),
							newInput = $('<input type="text" autocomplete="off" placeholder="可选择，可新增">'),
							newDatalist = $('<datalist style="display:none;"></datalist>');
						newInput.attr({'id':oldId,'name':oldName,'list':oldId+'_list','class':oldClass});
						newDatalist.attr('id',oldId+'_list').html(oldOption);
						$(opts.dom[i]).hide();
						inputOuter.append(newInput);
						inputOuter.append(inputSpan);
						inputOuter.append(inputMask);
						inputOuter.append(inputDele);
						selectParent.append(inputOuter);
						selectParent.append(newDatalist);
						that.bindEdit(newInput);
						var isDele = opts.isdele;
						if(isDele){
							that.bindDele(newInput);
						}
					});
				}
			},
			bindEdit: function(obj){
				obj.bind('keydown',function(event){
					if(event.keyCode == 13){
						var listId = $(this).attr('list');
						if($(this).val()){
							var thisVal = $(this).val(),
								newOption = $('<option value="'+ thisVal +'">' + thisVal + '</option>');
							$('#'+listId).append(newOption);
							var timer = 400;
							$(this).parent().find('.mlk-edit-mask').show();
							$(this).next('.mlk-edit-down').fadeIn(timer,function(){
								$(this).fadeOut(timer,function(){
									$(this).fadeIn(timer,function(){
										$(this).fadeOut(timer,function(){
											$(this).next('.mlk-edit-mask').hide();
										});
									});	
								});	
							});
						}	
					}	
				});	
			},
			bindDele: function(obj){
				obj.bind('change',function(){
					if($(this).val()){
						$(this).parent().find('.mlk-edit-remove').show();
					}
					else{
						$(this).parent().find('.mlk-edit-remove').hide();	
					}	
				});
				obj.parent().find('.mlk-edit-remove').bind('click',function(){
					var thisInput = $(this).parent().find('input'),
						thisInputVal = thisInput.val(),
						thisListId = thisInput.attr('list');
					$('#'+thisListId).children('option').each(function(){
						if($(this).val() == thisInputVal){
							$(this).remove();	
						}	
					});
					thisInput.val('');
					$(this).hide();
				});	
			}
		}
		var newFunc = new innerFunction();
		newFunc.init(opts);	
	}
	
	//{dom:dom}
	mlk.tableInverse = function(opts){
		var selectAll = opts.dom.find('thead input[type=checkbox]').eq(0);
		selectAll.attr('checkFlag','false');
		
		selectAll.bind('click',function(){
			var checkBox = opts.dom.find('tbody input[type=checkbox]');
			if (selectAll.attr('checkFlag')=='false') {
				checkBox.each(function(){
					this.checked=true;
					selectAll.attr('checkFlag','true');
				})
			}else{
				checkBox.each(function(){
					this.checked=false;
					selectAll.attr('checkFlag','false');
				})   
			}
		});
		opts.dom.find('tbody').on('click','input[type=checkbox]',function(){
			var checkBox = opts.dom.find('tbody input[type=checkbox]');
			checkBox.each(function(){
				if (this.checked==false) { 
					selectAll[0].checked=false;
					selectAll.attr('checkFlag','false');
					return false; 
				}
				else{
					selectAll[0].checked=true;
					selectAll.attr('checkFlag','true');
				}    
			});
		}); 
	}
	
	//{ele:./#,btn:./#,ischeckall:true/false}
	mlk.dropdownSearch = function(opts){
		var innerFunction = function(){};
		innerFunction.prototype = {
			constructor: innerFunction,
			init: function(opts){
				if(!opts){
					return;	
				}
				
				this.creatSearch(opts,this);
			},
			creatSearch: function(opts,that){
				$(opts.ele).each(function(){
					if(!$(this).parent().find(opts.btn).attr('ischange')){
						$(this).parent().find(opts.btn).attr('ischange','ok');
						$(opts.ele).css('padding','0px');
						var searchDiv = $('<div class="mlk-dropdown-search"></div>'),
							searchInput = $('<input placeholder="输入搜索" type="text" class="mlk-dropdowninput form-control">'),
							outerDiv = $('<div class="mlk-dropdown-outer"></div>');
						outerDiv.css({'width':'100%','height':($(opts.ele).height()-45),'overflow-y':'auto','padding':'0 3px'});
						var myArray = [];
						$(this).children().each(function(){
							myArray.push($(this).text());	
						});
						$(this).attr('mychild',myArray);
						//是否添加全选按钮
						if(opts.ischeckall){
							var checkAll = $('<input type="checkbox" class="mlk-check-all" title="全选">');	
							searchDiv.append(checkAll);
							that.bindCheckAll(opts,checkAll);
						}
						searchDiv.append(searchInput);
						$(this).append(outerDiv);
						outerDiv.append($(this).children());
						searchDiv.insertBefore(outerDiv);
						that.bindKeyUp(opts,searchInput);
					}
					else{
						return;	
					}
				});
			},
			bindKeyUp: function(opts,obj){
				obj.keyup(function(){
					var thisVal = $(this).val(),
						thisAll = $(this).parents(opts.ele).attr('mychild').split(',');
					if(thisVal){
						var allHtml = '';
						for(var i = 0; i < thisAll.length; i++){
							if(thisAll[i].indexOf(thisVal) >= 0){
								var thisName = $(this).parents(opts.ele).parent().find(opts.btn).text(),
									inputHtml = '<label class="checkbox"><input type="checkbox" name="' + thisName + '" value="' + thisAll[i] + '">' + thisAll[i] + '</label>';
								allHtml += inputHtml;
							}	
						}
						//console.log(allHtml);
						$(this).parents(opts.ele).find('.mlk-dropdown-outer').html(allHtml);
					}
					else{
						var allHtml = '';
						for(var i = 0; i < thisAll.length; i++){
							var thisName = $(this).parents(opts.ele).parent().find(opts.btn).text(),
								inputHtml = '<label class="checkbox"><input type="checkbox" name="' + thisName + '" value="' + thisAll[i] + '">' + thisAll[i] + '</label>';
								allHtml += inputHtml;	
						}
						$(this).parents(opts.ele).find('.mlk-dropdown-outer').html(allHtml);	
					}	
				});	
			},
			bindCheckAll: function(opts,obj){
				obj.change(function(){
					if(this.checked){
						$(this).parents('.mlk-dropdown-search').next('.mlk-dropdown-outer').find('input[type=checkbox]').each(function(){
							this.checked = true;	
						});	
					}
					else{
						$(this).parents('.mlk-dropdown-search').next('.mlk-dropdown-outer').find('input[type=checkbox]').each(function(){
							this.checked = false;	
						});	
						}	
				});	
			}
				
		}
		
		var newFunc = new innerFunction();
		newFunc.init(opts);	
	}
	
	
	window.mlk.contant = contant;

})(window);

