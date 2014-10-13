requirejs.config({
    //baseUrl: "/fist/lib",
    paths: {
        'jquery': "jquery",
		'milk': "milk",
		'datatables': "datatables"
    },
    shim: {
        'handlebars': {
            exports: 'Handlebars'
        },
        "highcharts": {
            deps: ["jquery"],
            exports: 'jQuery.fn.highcharts'
        },
		"datatables": {
			deps: ["jquery"]	
		}
    }
});