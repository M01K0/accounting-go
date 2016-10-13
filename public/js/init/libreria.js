'use strict';
(function (window, document) {
    function init() {
        let marco = null,
            vistaNoEncontrado = null,
            rutas = {},
            controladores = {},
            ctrl = null,
            singleton = {},
            libreria = {
                STATUS_OK: 200,
                STATUS_UNAUTHORIZED: 401,
                STATUS_FORBIDDEN: 403,
                STATUS_INTERNALSERVERERROR: 500,
                getID: function (id) {
                    var clone = {elemento: document.getElementById(id)};
                    clone = this.extender(clone, this);
                    return clone;
                },
                getElement: function (ele) {
                    var clone = {elemento: ele};
                    clone = this.extender(clone, this);
                    return clone;
                },
                get: function () {
                    return this.elemento;
                },
                click: function (funcion) {
                    this.elemento.addEventListener('click', funcion, false);
                    return this;
                },
                noSubmit: function () {
                    this.elemento.addEventListener('submit', function (e) {
                        e.preventDefault();
                    }, false);
                    return this;
                },
                addClass: function (clase) {
                    this.elemento.classList.add(clase);
                    return this;
                },
                delClass: function (clase) {
                    this.elemento.classList.remove(clase);
                    return this;
                },
                toggleClass: function (clase) {
                    this.elemento.classList.toggle(clase);
                    return this;
                },
                innerHTML: function (contenido) {
                    this.elemento.innerHTML = contenido;
                    return this;
                },
                text: function (contenido) {
                    this.elemento.textContent = contenido;
                    return this;
                },
                setValue: function (valor) {
                    this.elemento.value = valor;
                    return this;
                },
                getValue: function () {
                    return this.elemento.value;
                },
                getSingleton: function () {
                    return singleton;
                },
                setSingleton: function (objeto) {
                    singleton = objeto;
                },
                onEnterNext: function(){
                    this.elemento.addEventListener('keypress', function(e){
                        var indice = parseInt(e.target.getAttribute('tabindex'),10),
                            siguiente;
                        if (e.keyCode === 13){
                            e.preventDefault();
                            siguiente = e.target.parentNode.parentNode.querySelector('[tabindex="'+(indice+1)+'"]');
                            siguiente.focus();
                            if(siguiente.select){
                                siguiente.select();
                            }
                        }
                    }, false);
                    return this;
                },
                llenarFilas: function (cuerpoTabla, template, datos, campos, acciones) {
                    var cuerpo = document.getElementById(cuerpoTabla),
                            fila = document.getElementById(template),
                            frag = document.createDocumentFragment(),
                            i = 0, j = 0, maxDatos = datos.length, registro = {},
                            clon = null, maxCampos = campos.length, campo = null,
                            accion = null, btnAccion = null;

                    cuerpo.textContent = '';
                    for (; i < maxDatos; i = i + 1) {
                        registro = datos[i];
                        clon = fila.content.cloneNode(true);
                        for (; j < maxCampos; j = j + 1) {
                            campo = clon.querySelector('.' + campos[j]);
                            if (typeof registro[campos[j]] !== 'boolean') {
                                campo.textContent = registro[campos[j]];
                            } else {
                                campo.textContent = registro[campos[j]] ? 'Si' : 'No';
                            }
                        }
                        j = 0;

                        /**
                         * Accines a realizar
                         * El objeto debe tener la siguiente estructura
                         * {'nombre': {'clase', 'funcion'}}
                         */
                        for (accion in acciones) {
                            btnAccion = clon.querySelector(acciones[accion].clase);
                            btnAccion.dataset.idu = registro['id'];
                            btnAccion.addEventListener('click', acciones[accion].funcion, false);
                        }

                        frag.appendChild(clon);
                    }
                    cuerpo.appendChild(frag);
                },
                // Función que permite llenar una tabla
                // Se le debe enviar el id del tbody
                // el id del template
                // el objeto json con la data
                // el nombre de los campos del objeto json (en array)
                // el nombre de las clases de cada campo o columna (en array)
                // las acciones a realizar
                fillRows: function (table, template, data, fields, columns, actions) {
                    let body = document.getElementById(table),
                        row = document.getElementById(template),
                        fragment = document.createDocumentFragment(),
                        i = 0,
                        j = 0,
                        maxData = data.length,
                        registry = {},
                        clon = null,
                        maxFields = fields.length,
                        column = null,
                        action = null,
                        btnAction = null;

                    body.textContent = '';
                    for (; i < maxData; i = i + 1) {
                        registry = data[i];
                        clon = row.content.cloneNode(true);
                        for (; j < maxFields; j = j + 1) {
                            column = clon.querySelector('.' + columns[j]);
                            switch (typeof registry[fields[j]]) {
                              case 'boolean':
                                column.textContent = registry[fields[j]] ? 'Si' : 'No';
                                break;
                              case 'number':
                                column.textContent = registry[fields[j]].formatNumero();
                                break;
                              default:
                                column.textContent = registry[fields[j]];
                            }
                        }
                        j = 0;

                        /**
                         * Accines a realizar
                         * El objeto debe tener la siguiente estructura
                         * {'nombre': {'clase', 'funcion'}}
                         */
                        for (action in actions) {
                            btnAction = clon.querySelector(actions[action].clase);
                            btnAction.dataset.id = registry[fields[0]];
                            btnAction.addEventListener('click', actions[action].funcion, false);
                        }

                        fragment.appendChild(clon);
                    }
                    body.appendChild(fragment);
                },
                paginate: function (ctrl) {
                    let input = this.getID('page'),
                        page = parseInt(input.value(), 10),
                        limit = parseInt(this.getID('limit').value(), 10);
                    if (page <= ctrl.totalPages) {
                        if (page > 0) {
                            ctrl.page = page;
                        } else {
                            ctrl.page = 1;
                        }
                    } else {
                        ctrl.page = ctrl.totalPages;
                    }
                    input.setValue(ctrl.page);
                    if (limit > 0) {
                        ctrl.limit = limit;
                    } else {
                        ctrl.limit = 1;
                        this.getID('limit').setValue(1);
                    }
                    ctrl.orderBy = parseInt(this.getID('orderBy').value(), 10);
                    ctrl.orderType = this.getID('orderType').value();
                },
                movePaginate: function (ctrl, action) {
                    let page = this.getID('page');
                    switch (action) {
                        case 'next':
                            page.setValue(ctrl.page + 1);
                            break;
                        case 'previus':
                            page.setValue(ctrl.page - 1);
                            break;
                        case 'first':
                            page.setValue(1);
                            break;
                        case 'last':
                            page.setValue(ctrl.totalPages);
                    }
                },
                pagination: function (page, limit, orderBy, orderType) {
                    var data = new FormData();
                    data.append("page", page);
                    data.append("limit", limit);
                    data.append("orderBy", orderBy);
                    data.append("orderType", orderType);
                    return data;
                },
                /**
                 * poblarSelect permite poblar la información de un select
                 * @param {JSON} datos Objeto JSON que contiene los datos
                 * @param {string} tabla Nombre de la tabla que contiene los datos. Es el nombre que está en el select como AS ...
                 * @param {string} id Nombre de la columna que identifica el ID de la tabla y será el value del select
                 * @param {string} campo Nombre del campo que se desea mostrar en el select
                 * @param {string} select ID del select que se va a poblar
                 * @param {boolean} esInterno Identifica si la información que viene en el JSON tiene paginación o no.
                 * @returns {void} Pobla el select con la información obtenida.
                 */
                poblarSelect: function (datos, tabla, id, campo, select, esInterno) {
                    var data = JSON.parse(datos),
                            fragmento = document.createDocumentFragment(),
                            lista = null, i = 0, max = 0, opcion = null;
                    if (data.tipo === this.MSG_CORRECTO) {
                        if (esInterno) {
                            lista = data.objeto[tabla];
                        } else {
                            lista = data.objeto;
                        }
                        max = lista.length;
                        for (; i < max; i = i + 1) {
                            opcion = document.createElement('option');
                            opcion.setAttribute('value', lista[i][id]);
                            opcion.textContent = lista[i][campo];
                            fragmento.appendChild(opcion);
                        }
                        this.getID(select).get().appendChild(fragmento);
                    }
                },
                setCtrl: function (name, controller) {
                    controladores[name] = {'controlador': controller};
                },
                getCtrl: function () {
                    if (arguments.length === 0) {
                        return ctrl;
                    } else {
                        return controladores[arguments[0]].controlador;
                    }
                },
                enrutar: function (id) {
                    marco = document.getElementById(id);
                    return this;
                },
                ruta: function (url, plantilla, controller, carga) {
                    rutas[url] = {
                        'plantilla': plantilla,
                        'controlador': controller,
                        'carga': carga
                    };
                    return this;
                },
                manejadorRutas: function () {
                    let hash = window.location.hash.substring(1) || '/',
                        realHash = hash.split('?')[0],
                        id = _.getParameterByName('id'),
                        rutaDestino = rutas[realHash];

                    if (rutaDestino && rutaDestino.plantilla) {

                        if (rutaDestino.controlador) {
                            ctrl = controladores[rutaDestino.controlador].controlador;
                        }

                    } else {
                        rutaDestino = {};
                        rutaDestino.plantilla = vistaNoEncontrado;
                    }
                    rutaDestino.id = id;
                    _.cargaVista(rutaDestino);
                },
                cargaVista: function (rutaDestino) {
                    this.ajax({
                        method: 'GET',
                        url: rutaDestino.plantilla
                    }).then(function (data) {
                        marco.innerHTML = data.content;
                        if (typeof (rutaDestino.carga) === 'function') {
                            if (rutaDestino.id !== null && rutaDestino.id.length > 0) {
                                rutaDestino.carga(rutaDestino.id);
                            } else {
                                rutaDestino.carga();
                            }
                        }
                    }, function (error) {
                        console.log(error);
                    });
                },
                notFound: function (archivo) {
                    vistaNoEncontrado = archivo;
                    return this;
                },
                getParameterByName: function (name, url) {
                    // This line can be window.location.href
                    if (!url) url = window.location.hash;
                    name = name.replace(/[\[\]]/g, "\\$&");
                    var regex = new RegExp("[?&]" + name + "(=([^&#]*)|&|#|$)"),
                        results = regex.exec(url);
                    if (!results) return null;
                    if (!results[2]) return '';
                    return decodeURIComponent(results[2].replace(/\+/g, " "));
                },
                ajax: function (objeto) {
                    return new Promise(function (resolver, rechazar) {
                        let method = objeto.method,
                            url = objeto.url || '',
                            body = objeto.body || null,
                            token = '',
                            xhr = new XMLHttpRequest();

                        token = sessionStorage.getItem('token' + method);

                        xhr.open(method, url, true);
                        xhr.setRequestHeader('Content-Type', 'application/json');
                        xhr.setRequestHeader('Authorization', `Bearer ${token}`);
                        xhr.addEventListener('load', function () {
                            let response = {
                                status: this.status,
                                content: this.responseText
                            };
                            // Unauthorized status
                            if (this.status === _.STATUS_UNAUTHORIZED) {
                                alert('No estás autenticado. Por favor ingresa nuevamente.');
                                window.location.href = '/';
                                return;
                            }
                            resolver(response);
                        }, false);
                        xhr.addEventListener('error', function () {
                            rechazar(Error('Hubo un error en la red'));
                        }, false);
                        xhr.send(body);
                    });
                },
                /* Es un FACADE del ajax */
                execute: function (obj) {
                    this.ajax({
                        method: obj.method,
                        url: obj.url,
                        body: obj.body
                    }).then(function (response) {
                                obj.callback(response);
                            }, function (error) {
                                console.log(error);
                            });
                },
                extender: function (out) {
                    out = out || {};

                    for (var i = 1; i < arguments.length; i++) {
                        var obj = arguments[i];

                        if (!obj)
                            continue;

                        for (var key in obj) {
                            if (obj.hasOwnProperty(key)) {
                                if (typeof obj[key] === 'object')
                                    this.extender(out[key], obj[key]);
                                else
                                    out[key] = obj[key];
                            }
                        }
                    }
                    return out;
                },
                serializeJSON: function (form) {
                	if (!form || form.nodeName !== "FORM") {
                		return;
                	}

                	let i = 0, j = 0, obj = {}, name = '', value;

                	for (i = form.elements.length - 1; i >= 0; i = i - 1) {
                		if (form.elements[i].name === "") {
                			continue;
                		}
                		name = form.elements[i].name;
                		value = form.elements[i].value;

                		switch (form.elements[i].nodeName) {
                		case 'INPUT':
                			switch (form.elements[i].type) {
                			case 'checkbox':
                			    if (form.elements[i].checked) {
                			        obj[name] = true;
                			    } else {
                			        obj[name] = false;
                			    }
                			    break;
                			case 'radio':
                				if (form.elements[i].checked) {
                					obj[name] = value;
                				}
                				break;
                			case 'file':
                				break;
                			default:
                				obj[name] = value;
                				break;
                			}
                			break;
                		case 'TEXTAREA':
                			obj[name] = value;
                			break;
                		case 'SELECT':
                			switch (form.elements[i].type) {
                			case 'select-one':
                				obj[name] = value;
                				break;
                			case 'select-multiple':
                			    let myOptions = [];
                				for (j = form.elements[i].options.length - 1; j >= 0; j = j - 1) {
                					if (form.elements[i].options[j].selected) {
                						myOptions.push(encodeURIComponent(form.elements[i].options[j].value));
                					}
                				}
                				obj[name] = myOptions;
                				break;
                			}
                			break;
                		}
                	}
                	// return q.join("&");
                	return obj;
                }
            };

        /**
         * Da formato de moneda a los números
         * @param {int} decimales Cantidad de decimales que tiene el número
         * @param {char} digitos Separador de decimales
         * @param {char} temporal Separador de miles
         * @param {boolean} negativo Indica si el número negativo se envuelve (false) o se muestra el signo menos (true)
         * @returns {String} Numero en tipo Cadena de texto formateada
         */
        Number.prototype.formatNumero = function (decimales, negativo, digitos, temporal) {
            var n = this,
                c = isNaN(decimales = Math.abs(decimales)) ? 0 : decimales,
                d = typeof digitos === 'undefined' ? "." : digitos,
                t = typeof temporal === 'undefined' ? "," : temporal,
                si = n < 0 ? "-" : "",
                sf = "",
                i = parseInt(n = Math.abs(+n || 0).toFixed(c)) + "",
                j = (j = i.length) > 3 ? j % 3 : 0;
            if(!negativo){
                si = n < 0 ? "(" : "";
                sf = n < 0 ? ")" : "";
            }
            return si + (j ? i.substr(0, j) + t : "") + i.substr(j).replace(/(\d{3})(?=\d)/g, "$1" + t) + (c ? d + Math.abs(n - i).toFixed(c).slice(2) : "") + sf;
        };

        return libreria;
    };

    if (typeof window.libreria === "undefined") {
        window.libreria = window._ = init();
    } else {
        console.log("Ya está llamada");
    }

})(window, document);
