(function (window, JSON, _) {
    let costCenterCtrl = {
        apiUrl: '/api/cost-centers/',
        viewUrl: '#/centros-costo',
        page: 1,
        totalPages: 0,
        limit: 10,
        orderBy: 2,
        orderType: 'asc',
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
            let self = _.getCtrl(),
                data = JSON.parse(response.content).data,
                msgDiv = self.messageDiv.delClass('mo-mostrar');

            switch (response.status) {
            case _.STATUS_OK:
                _.getID('id').setValue(data.id);
                _.getID('code').setValue(data.code);
                _.getID('costCenter').setValue(data.costCenter);
                break;
            case _.STATUS_FORBIDDEN:
                msgDiv.text(data.message);
                break;
            default:
                msgDiv.text('Status no esperado: ' + response.status + ' ' + data.error + ' ' + data.message);
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
            _.execute({
                method: 'PUT',
                url: self.apiUrl + data.id,
                body: JSON.stringify(body),
                callback: self.updated
            });
        },
        updated: function (response) {
            let self = _.getCtrl(),
                msgDiv = self.messageDiv.delClass('no-mostrar'),
                data = JSON.parse(response.content).data;

            switch (response.status) {
            case _.STATUS_OK:
                msgDiv.text("Actualizado Correctamente");
                break;
            case _.STATUS_FORBIDDEN:
                msgDiv.text("No está autorizado para realizar esta acción");
                break;
            default:
                msgDiv.text('Status no esperado: ' + response.status + ' ' + data.error + ' ' + data.message);
            }

            self.form.reset();
            setTimeout(function () {
                window.location.hash = self.viewUrl;
            }, 3000);
        },
        create: function () {
            let self = this,
                data = _.serializeJSON(self.form),
                body = {data};
            _.execute({
              method: 'POST',
              url: self.apiUrl,
              body: JSON.stringify(body),
              callback: self.created
            });
        },
        created: function (response) {
            let self = _.getCtrl(),
                msgDiv = self.messageDiv.delClass('no-mostrar'),
                data = JSON.parse(response.content).data;

            switch (response.status) {
            case _.STATUS_CREATED:
                msgDiv.text('Centro de costo creado con el id: ' + data.id);
                self.form.reset();
                break;
            case _.STATUS_FORBIDDEN:
                msgDiv.text('No estás autorizado para esta acción');
                break;
            default:
                msgDiv.text('Status no esperado: ' + response.status + ' ' + data.error + ' ' + data.message);
            }
        },
        confirmDelete: function (id) {
            let self = this;
            if (confirm('Desea eliminar este centro?')) {
                self.delete(id);
            }
        },
        delete: function (id) {
            let self = this;
            _.execute({
                method: 'DELETE',
                url: self.apiUrl + id,
                callback: self.deleted
            });
        },
        deleted: function (response) {
            console.log(response.content);
            let self = _.getCtrl(),
                data = {},
                messageDiv = self.messageDiv.delClass('no-mostrar');

            switch (response.status) {
            case _.STATUS_NOCONTENT:
                messageDiv.text('Eliminado correctamente');
                self.list(self.loadTable);
                break;
            case _.STATUS_FORBIDDEN:
                messageDiv.text('No estás autorizado para realizar esta acción');
                break;
            default:
                data = JSON.parse(response.content).data;
                msgDiv.text('Status no esperado: ' + response.status + ' ' + data.error + ' ' + data.message);
            }
        },
        list: function (callback) {
            let self = this,
                body = _.pagination(this.page, this.limit, this.orderBy, this.orderType),
                request = {
                    method: 'GET',
                    url: self.apiUrl + body,
                    callback: callback
                };
            _.execute(request);
        },
        loadTable: function (response) {
            let self = _.getCtrl(),
                msgDiv = self.messageDiv,
                data = JSON.parse(response.content).data,
                campos = [],
                columns = [],
                actions = {};

            switch (response.status) {
            case _.STATUS_OK:
                fields = ['id', 'code', 'costCenter'];
                columns = ['id', 'codigo', 'nombre'];
                actions = {
                    delete: {
                        class: '.eliminar',
                        callback: function (e) {
                                    e.preventDefault();
                                    self.confirmDelete(e.target.dataset.id);
                        }
                    },
                    update: {
                        class: '.actualizar',
                        callback: function (e) {
                                    e.preventDefault();
                                    self.confirmUpdate(e.target.dataset.id);
                        }
                    }
                };

                self.totalPages = Math.ceil(data.length / self.limit);
                _.getID('page').get().setAttribute('max', self.totalPages);
                _.getID('totalPages').text('de ' + self.totalPages);
                _.fillRows('cuerpoTabla', 'plantilla', data, fields, columns, actions);
                break;
            case _.STATUS_FORBIDDEN:
                msgDiv.delClass('no-mostrar').text('No estás autorizado para listar este contenido');
                break;
            default:
                msgDiv.delClass('no-mostrar').text('Status no esperado: ' + response.status + ' ' + data.error + ' ' + data.message);
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
