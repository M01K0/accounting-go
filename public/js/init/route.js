/* global _ */
'use strict';
(function (window, _) {
    _.enrutar('vista')
            .notFound('/404.html')
            .ruta('/', 'views/index.html')
            .ruta('/perfiles',
                    'views/perfiles/listar.html',
                    'perfiles',
                    function () {
                        var ctrl = _.getCtrl();
                        ctrl.listar(ctrl.cargaLista);
                    })
            .ruta('/perfiles-crear',
                    'views/perfiles/crear.html',
                    'perfiles',
                    function () {
                        _.getID('frmCrearPerfil').noSubmit();
                    })
            .ruta('/perfiles-actualizar',
                    'views/perfiles/actualizar.html',
                    'perfiles',
                    function () {
                        _.getID('frmActualizarPerfil').noSubmit();
                    })
            .ruta('/perfiles-permisos',
                    'views/perfiles/permisos.html',
                    'objetosxperfil',
                    function () {
                        _.getCtrl().cargaPermisos();
                    })
            .ruta('/usuarios',
                    'views/usuarios/listar.html',
                    'usuarios',
                    function () {
                        var ctrl = _.getCtrl();
                        ctrl.listar(ctrl.cargaLista);
                    })
            .ruta('/usuarios-crear',
                    'views/usuarios/crear.html',
                    'usuarios',
                    function () {
                        _.getCtrl().inicio_crear();
                    })
            .ruta('/usuarios-actualizar',
                    'views/usuarios/actualizar.html',
                    'usuarios',
                    function () {
                        _.getCtrl().inicio_actualizar();
                    })
            .ruta('/usuarios-cambiar-clave',
                  'views/usuarios/cambiarclave.html',
                  'usuarios',
                  function(){
                      _.getCtrl().inicio_cambiar_clave();
                  }
            )
            .ruta('/tipo-identificacion',
                    'views/tipo_identificacion/listar.html',
                    'tipoIdentificacion',
                    function () {
                        _.getCtrl().inicio_listar();
                    })
            .ruta('/tipo-identificacion-crear',
                    'views/tipo_identificacion/crear.html',
                    'tipoIdentificacion',
                    function () {
                        _.getCtrl().inicio_crear();
                    })
            .ruta('/tipo-identificacion-actualizar',
                    'views/tipo_identificacion/actualizar.html',
                    'tipoIdentificacion',
                    function () {
                        _.getCtrl().inicio_actualizar();
                    })
            .ruta('/tipo-funcionario',
                    'views/tipo_funcionario/listar.html',
                    'tipoFuncionario',
                    function () {
                        _.getCtrl().inicio_listar();
                    })
            .ruta('/tipo-funcionario-crear',
                    'views/tipo_funcionario/crear.html',
                    'tipoFuncionario',
                    function () {
                        _.getCtrl().inicio_crear();
                    })
            .ruta('/tipo-funcionario-actualizar',
                    'views/tipo_funcionario/actualizar.html',
                    'tipoFuncionario',
                    function () {
                        _.getCtrl().inicio_actualizar();
                    })
            .ruta('/empresa',
                    'views/empresa/admin.html',
                    'empresa',
                    function () {
                        _.getCtrl().inicio();
                    })
            .ruta('/empresa-funcionarios',
                    'views/empresa_funcionario/listar.html',
                    'empresaFuncionario',
                    function () {
                        var ctrl = _.getCtrl();
                        ctrl.listar(ctrl.cargarTabla);
                    }
            )
            .ruta('/empresa-funcionarios-crear',
                    'views/empresa_funcionario/crear.html',
                    'empresaFuncionario',
                    function () {
                        _.getCtrl().inicio_crear();
                    }
            )
            .ruta('/empresa-funcionarios-actualizar',
                    'views/empresa_funcionario/actualizar.html',
                    'empresaFuncionario',
                    function () {
                        _.getCtrl().inicio_actualizar();
                    }
            )
            .ruta('/cuentas-puc',
                  'views/cuentaspuc/listar.html',
                  'cuentaspuc',
                  function(){
                      var ctrl = _.getCtrl();
                      ctrl.listar(ctrl.cargarTabla);
                  }
            )
            .ruta('/cuentas-puc-crear',
                  'views/cuentaspuc/crear.html',
                  'cuentaspuc',
                  function(){
                      _.getCtrl().inicio_crear();
                  }
            )
            .ruta('/cuentas-puc-actualizar',
                  'views/cuentaspuc/actualizar.html',
                  'cuentaspuc',
                  function(){
                      _.getCtrl().inicio_actualizar();
                  }
            )
            .ruta('/centros-costo',
                  'views/centroscosto/listar.html',
                  'costCenter',
                  function(){
                      let ctrl = _.getCtrl();
                      ctrl.init();
                      ctrl.list(ctrl.loadTable);
                  }
            )
            .ruta('/centros-costo/crear',
                  'views/centroscosto/crear.html',
                  'costCenter',
                  function(){
                      _.getCtrl().initCreate();
                  }
            )
            .ruta('/centros-costo/actualizar',
                  'views/centroscosto/actualizar.html',
                  'costCenter',
                  function(id){
                      _.getCtrl().initUpdate(id);
                  }
            )
            .ruta('/terceros',
                  'views/terceros/listar.html',
                  'terceros',
                  function(){
                      var ctrl = _.getCtrl();
                      ctrl.inicio();
                      ctrl.listar(ctrl.cargarTabla);
                  }
            )
            .ruta('/terceros/crear',
                  'views/terceros/crear.html',
                  'terceros',
                  function(){
                      _.getCtrl().inicio_crear();
                  }
            )
            .ruta('/terceros/actualizar',
                  'views/terceros/actualizar.html',
                  'terceros',
                  function(){
                      _.getCtrl().inicio_actualizar();
                  }
            )
            .ruta('/terceros/detalle',
                  'views/terceros/detalle.html',
                  'terceros',
                  function(){
                      _.getCtrl().inicio();
                  }
            )
            .ruta('/documento-contable',
                  'views/documento_contable/listar.html',
                  'documentoContable',
                  function(){
                      var ctrl = _.getCtrl();
                      ctrl.inicio();
                      ctrl.listar(ctrl.cargarTabla);
                  }
            )
            .ruta('/documento-contable/crear',
                  'views/documento_contable/crear.html',
                  'documentoContable',
                  function(){
                      _.getCtrl().inicio_crear();
                  }
            )
            .ruta('/documento-contable/actualizar',
                  'views/documento_contable/actualizar.html',
                  'documentoContable',
                  function(){
                      _.getCtrl().inicio_actualizar();
                  }
            )
            .ruta('/periodos/crear',
                  'views/periodo/crear.html',
                  'periodo',
                  function(){
                      _.getCtrl().inicio();
                  }
            )
            .ruta('/contabilidad/registro',
                  'views/contabilidad/movimientos/registrocontable.html',
                  'registroContable',
                  function(){
                      _.getCtrl().inicio_registro();
                  }
            )
            .ruta('/contabilidad/registro/detalle',
                  'views/contabilidad/movimientos/registrocontabledetalle.html',
                  'registroContable',
                  function(){
                      _.getCtrl().inicio_detalle();
                  }
            )
            .ruta('/contabilidad/consultar-documento',
                  'views/contabilidad/movimientos/consultadocumento.html',
                  'registroContable',
                  function(){
                      _.getCtrl().inicio_consultar();
                  }
            )
            .ruta('/contabilidad/editar-documento',
                  'views/contabilidad/movimientos/modificaregistrocontable.html',
                  'registroContable',
                  function(){
                      _.getCtrl().inicio_editar();
                  }
            )
            .ruta('/contabilidad/cierre',
                  'views/contabilidad/cierre.html',
                  'cierreContable',
                  function(){
                      _.getCtrl().inicio_cierre();
                  }
            )
            .ruta('/informes/detallados/movimientos/cuentas',
                  'views/informes/detallados/movimientos/cuentas.html',
                  'inf-mov-cuenta',
                  function () {
                    _.getCtrl().init();
                  }
            )
            .ruta('/informes/balance/general',
                  'views/informes/balances/general.html',
                  'inf-balance-general',
                  function () {
                    _.getCtrl().init();
                  }
            );

    window.addEventListener('load', _.manejadorRutas, false);
    window.addEventListener('hashchange', _.manejadorRutas, false);

})(window, _);
