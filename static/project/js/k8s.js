function DeletePod(ns, name) {
  jQuery.ajax({
    type: 'DELETE',
    url: '/delete/' + ns + '/pod/' + name,
    success: function(data){
      SuccessDelete(name, "Pod");
      setTimeout(function() {
        window.location.replace("/pods");
      }, 1000);
    },
    error: function(data){
          FailDelete(name, "Pod");
    }
  });
}
