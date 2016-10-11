(function (window, JSON, _) {
    var costCenterCtrl = {
        apiUrl: '/api/cost-centers/',
        viewUrl: '#/centros-costo',
        page: 1,
        totalPages: 0,
        limit: 10,
        orderBy: 2,
        orderType: 'acs',
        queryType: 1,
        form: null,
        messageDiv: null,
        init: function () {
            this.messageDiv = _.getID('mensaje');
        },
        initCreate: function () {
            this.form = _.getID('frmCrearCentrosCosto').noSubmit().get();
            this.init();
        },
        initUpdate: function (id) {
            this.form = _.getID('frmActualizarCentrosCosto').noSubmit().get();
            this.init();
            this.getByID(id);
        },
        getByID: function (id) {
            // TODO: Validar que tiene permisos para actualizar
            console.log("Este es el id:", id);
            // TODO: Traer la información del centro de costo y colocarlo
            /*
            _.getID('id').setValue(objeto.id);
            _.getID('codigo').setValue(objeto.codigo);
            _.getID('nombre').setValue(objeto.nombre);
            */
        },
        confirmUpdate: function (id) {
            if (confirm('Desea actualizar este centro de costo?')) {
                window.location.hash = this.viewUrl + '/actualizar?id=' + id;
            }
        },
        update: function () {
            let self = this,
                data = new FormData(this.form);
            _.ajax({
                method: 'PUT',
                url: self.apiUrl + id,
                body: data
            }).then(function (data) {
                        self.updated(data);
                    },
                    function (error) {
                        console.log(error);
                    }
            );
        },
        updated: function (response) {
            let self = this,
                data = JSON.parse(response);

            self.messageDiv.delClass('no-mostrar').text(data.mensaje);
            if (data.tipo === _.MSG_CORRECTO) {
                self.form.reset();
                setTimeout(function () {
                    window.location.hash = self.viewUrl;
                }, 3000);
            } else if (data.tipo === _.MSG_NO_AUTENTICADO) {
                window.location.href = '/';
            }
        },
        create: function () {
            let self = this,
                data = new FormData(this.form);
            _.ajax({
                method: 'POST',
                url: self.apiUrl,
                body: data
            }).then(function (data) {
                        self.created(data);
                    },
                    function (error) {
                        console.log(error);
                    }
            );
        },
        created: function (response) {
            let self = this,
                data = JSON.parse(response);

            self.messageDiv.delClass('no-mostrar').text(data.mensaje);
            if (data.tipo === _.MSG_CORRECTO) {
                self.form.reset();
            } else if (data.tipo === _.MSG_NO_AUTENTICADO) {
                window.location.href = '/';
            }
        },
        confirmDelete: function (id) {
            if (confirm('Desea eliminar este centro?')) {
                erase(id);
            }
        },
        erase: function (id) {
            let self = this;

            _.ajax({
                method: 'DELETE',
                url: self.apiUrl + id
            }).then(function (data) {
                        self.erased(data);
                    }, function (error) {
                        console.log(error);
                    }
            );
        },
        erased: function (response) {
            let self = this,
                data = JSON.parse(response);

            self.messageDiv.delClass('no-mostrar').text(data.mensaje);
            if (data.tipo === _.MSG_CORRECTO) {
                self.list(self.loadTable);
            } else if (data.tipo === _.MSG_NO_AUTENTICADO) {
                window.location.href = '/';
            }
        },
        list: function (callback) {
            let self = this,
                body = _.pagination(this.page, this.limit, this.orderBy, this.orderType);
            _.ajax({
                method: 'GET',
                url: self.apiUrl,
                body: body
            }).then(function (data) {
                        callback(data);
                    }, function (error) {
                        console.log(error);
                    });
        },
        loadTable: function (response) {
            let self = _.getCtrl(),
                data = {},
                campos = [],
                columns = [],
                acciones = {};

            if (response.status === 200) {
                data = JSON.parse(response.content);
                campos = ['id', 'code', 'costCenter'];
                columns = ['id', 'codigo', 'nombre'];
                acciones = {
                    eliminar: {
                        clase: '.eliminar',
                        funcion: function (e) {
                                    e.preventDefault();
                                    self.confirmDelete(e.target.dataset.id);
                        }
                    },
                    actualizar: {
                        clase: '.actualizar',
                        funcion: function (e) {
                                    e.preventDefault();
                                    self.confirmUpdate(e.target.dataset.id);
                        }
                    }
                };

                self.totalPages = Math.ceil(data.data.length / self.limit);
                _.getID('page').get().setAttribute('max', self.totalPages);
                _.getID('totalPages').text('de ' + self.totalPages);
                _.fillRows('cuerpoTabla', 'plantilla', data.data, campos, columns, acciones);
            } else if (data.tipo === _.MSG_ADVERTENCIA || data.tipo === _.MSG_ERROR) {
                self.messageDiv.delClass('no-mostrar').text(data.mensaje);
            } else if (data.tipo === _.MSG_NO_AUTENTICADO) {
                window.location.href = '/';
            }
        },
        buscarXCodigoOId: function(tipo, codigo, id, callback){
            var data = new FormData(),
                obj = {
                    url: 'SCentroCostoGetXCodigoXId',
                    datos: data,
                    callback: callback
                };
            data.append('codigo', codigo);
            data.append('id', id);
            _.ejecutar(obj);
        },
        paginate: function () {
            _.paginate(this);
            this.list(this.loadTable);
        },
        movePaginate: function (action) {
            _.movePaginate(this, action);
            this.paginate();
        }
    };

    _.controlador('costCenter', costCenterCtrl);
})(window, JSON, _);
