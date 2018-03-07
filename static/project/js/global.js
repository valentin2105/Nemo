PNotify.prototype.options.styling = "fontawesome"
PNotify.prototype.options.delay -= 5000;
function SuccessDelete(name, type) {
  new PNotify({
  title: type + ' deleted.',
  text: name + ' is deleted.',
  type: 'success',
  nonblock: {
      nonblock: true
  }
  });
}
function FailDelete(name, type) {
  new PNotify({
  title: 'Failed to delete.',
  text: type + ' failed to delete',
  icon: 'fas fa-exclamation',
  type: 'error',
  nonblock: {
      nonblock: true
  }
  })
}
