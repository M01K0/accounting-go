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
            this.messageDiv = _.getID('message');
        },
        initCreate: function () {
            this.form = _.getID('frmCrearCentrosCosto').noSubmit().get();
            this.init();
        },
        initUpdate: function (id) {
            this.form = _.getID('frmActualizarCentrosCosto').noSubmit().get();
            this.init();
            this.getByID(id, this.fillData);
        },
        getByID: function(id, callback){
            let self = this,
                request = {
                method: 'GET',
                url: self.apiUrl + id,
                callback: callback
            };
            _.execute(request);
        },
        fillData: function (response) {
            switch (response.status) {
            case _.STATUS_OK:
                data = JSON.parse(response.content);
                _.getID('id').setValue(data.data.id);
                _.getID('code').setValue(data.data.code);
                _.getID('costCenter').setValue(data.data.costCenter);
                break;
            case _.STATUS_FORBIDDEN:
                self.messageDiv.text(data.data.message);
                break;
            default:
                self.messageDiv.text("Código de respuesta no esperado: "+data.status);
            }
        },
        confirmUpdate: function (id) {
            if (confirm('Desea actualizar este centro de costo?')) {
                window.location.hash = this.viewUrl + '/actualizar?id=' + id;
            }
        },
        update: function () {
            let self = this,
                data = _.serializeJSON(self.form),
                body = {};

            data.id = parseInt(data.id, 10);
            body = {data};
            _.ajax({
                method: 'PUT',
                url: self.apiUrl + data.id,
                body: JSON.stringify(body)
            }).then(function (response) {
                        self.updated(response);
                    },
                    function (error) {
                        console.log(error);
                    }
            );
        },
        updated: function (response) {
            let self = _.getCtrl(),
                data = JSON.parse(response.content);
            switch (response.status) {
            case _.STATUS_OK:
                self.messageDiv.delClass('no-mostrar').text("Actualizado Correctamente");
                break;
            case _.STATUS_FORBIDDEN:
                self.messageDiv.delClass('no-mostrar').text("No está autorizado para realizar esta acción");
                break;
            }
            self.form.reset();
            setTimeout(function () {
                window.location.hash = self.viewUrl;
            }, 3000);
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
                actions = {};

            if (response.status === 200) {
                data = JSON.parse(response.content);
                campos = ['id', 'code', 'costCenter'];
                columns = ['id', 'codigo', 'nombre'];
                actions = {
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
                _.fillRows('cuerpoTabla', 'plantilla', data.data, campos, columns, actions);
            } else if (data.tipo === _.MSG_ADVERTENCIA || data.tipo === _.MSG_ERROR) {
                self.messageDiv.delClass('no-mostrar').text(data.mensaje);
            } else if (data.tipo === _.MSG_NO_AUTENTICADO) {
                window.location.href = '/';
            }
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

    _.setCtrl('costCenter', costCenterCtrl);
})(window, JSON, _);
