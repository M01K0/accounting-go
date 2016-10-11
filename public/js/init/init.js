(function(_, window){
  let user = JSON.parse(sessionStorage.getItem('user'));
  if (user.username) {
    _.getID('user').text(user.username);
  } else {
    window.location.href = '/';
  }
})(_, window);
