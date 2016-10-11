(function(window){
  let user = JSON.parse(sessionStorage.getItem('user'));
  if (user === null || typeof user === undefined) {
    window.location.href = '/';
  }
})(window);
