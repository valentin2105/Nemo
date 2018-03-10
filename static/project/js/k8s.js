function DeleteComponent(what, ns, name) {
  jQuery.ajax({
    type: 'DELETE',
    url: '/delete/' + ns + '/' + what + '/' + name,
    success: function(data){
      SuccessDelete(name, what);
      setTimeout(function() {
        window.location.replace('/' + what + 's');
      }, 1000);
    },
    error: function(data){
          FailDelete(name, what);
    }
  });
}
