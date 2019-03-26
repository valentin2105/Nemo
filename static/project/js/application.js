$(document).ready(function() {
	// get current URL path and assign 'active' class for Navbar
	var pathname = window.location.pathname;
	$('.nav > li > a[href="'+pathname+'"]').parent().addClass('active');

	// Filter Table
	var jobCount = $('.results tbody tr').length;
	$('.counter').text('(' + jobCount + ' items)');

	$(".search").keyup(function () {
	  var searchTerm = $(".search").val();
	  var listItem = $('.results tbody').children('tr');
	  var searchSplit = searchTerm.replace(/ /g, "'):containsi('")

	$.extend($.expr[':'], {'containsi': function(elem, i, match, array){
	      return (elem.textContent || elem.innerText || '').toLowerCase().indexOf((match[3] || "").toLowerCase()) >= 0;
	  }
	});

	$(".results tbody tr").not(":containsi('" + searchSplit + "')").each(function(e){
	  $(this).attr('visible','false');
	});

	$(".results tbody tr:containsi('" + searchSplit + "')").each(function(e){
	  $(this).attr('visible','true');
	});

	var jobCount = $('.results tbody tr[visible="true"]').length;
	$('.counter').text('(' + jobCount + ' items)');

	if(jobCount == '0') {$('.no-result').show();}
	else {$('.no-result').hide();}
	});
})



function onReady(callback) {
    var intervalID = window.setInterval(checkReady, 1000);

    function checkReady() {
        if (document.getElementsByTagName('body')[0] !== undefined) {
            window.clearInterval(intervalID);
            callback.call(this);
        }
    }
}

function show(id, value) {
    document.getElementById(id).style.display = value ? 'block' : 'none';
}

onReady(function () {
    show('page', true);
    show('loading', false);
});
