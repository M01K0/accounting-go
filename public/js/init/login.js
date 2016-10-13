(function(window, _, CryptoJS){

  _.getID('frmLogin').noSubmit();

  let conexion = function(response){
    let div  = _.getID('mensaje');
    switch (response.status) {
    case 200:
      let data = JSON.parse(response.content).data;
      let user = JSON.stringify(data.user),
          tokenPost = data.tokenPost,
          tokenPut  = data.tokenPut,
          tokenDelete = data.tokenDelete,
          tokenGet = data.tokenGet;
      sessionStorage.setItem('user', user);
      sessionStorage.setItem('tokenPOST', tokenPost);
      sessionStorage.setItem('tokenPUT', tokenPut);
      sessionStorage.setItem('tokenDELETE', tokenDelete);
      sessionStorage.setItem('tokenGET', tokenGet);
      window.location.href = '/contabilidad';
      break;
    case 401:
      div.text(JSON.parse(response.content).data.message)
      div.delClass('no-mostrar');
    }
  };

  _.getID('btnLogin').click(function(e){
    e.preventDefault();
    let data = {
        data: {
          "email": _.getID('email').getValue(),
          "password": CryptoJS.SHA3(_.getID('password').getValue()).toString()
        }
    };

    _.ajax({
        method: 'POST',
        url: '/api/login',
        body: JSON.stringify(data)
    }).then(function (response) {
            conexion(response);
         },
         function(error){
            console.log("Error en la promesa:", error);
         }
    );
  });

})(window, _, CryptoJS);
