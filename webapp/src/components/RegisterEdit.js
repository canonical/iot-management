/*
 * This file is part of the IoT Management Service
 * Copyright 2019 Canonical Ltd.
 *
 * This program is free software: you can redistribute it and/or modify it
 * under the terms of the GNU Affero General Public License version 3, as
 * published by the Free Software Foundation.
 *
 * This program is distributed in the hope that it will be useful, but WITHOUT
 * ANY WARRANTY; without even the implied warranties of MERCHANTABILITY,
 * SATISFACTORY QUALITY, or FITNESS FOR A PARTICULAR PURPOSE.
 * See the GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

import React, {Component} from 'react';
import api from '../models/api';
import AlertBox from './AlertBox';
import {T, isUserAdmin, formatError} from './Utils';


const statusOptions = [{id:1, name:'Waiting'}, {id:2, name:'Enrolled'}, {id:3, name:'Disabled'}]


class RegisterEdit extends Component {
    
    constructor(props) {
        super(props)
        this.state = {
            title: null,
            error: null,
            hideForm: false,
            deviceId: '',
            status: 0,
            statusTo: 0,
            device: {orgid: this.props.account.orgid},
        };
    }

    componentDidMount() {
        if (this.props.id) {
            this.setTitle('edit-device');
            this.getDevice(this.props.id);
        } else {
            this.setTitle('new-device');
        }
    }

    getDevice(id) {
        api.clientsGet(this.props.account.orgid, id).then(response => {
            this.setState({device: response.data.enrollment.device,
                status: response.data.enrollment.status, statusTo:response.data.enrollment.status,
                deviceId: response.data.enrollment.id});
        })
        .catch((e) => {
            this.setState({error: formatError(e.response.data), hideForm: true});
        })
    }

    setTitle(title) {
        this.setState({title: T(title)});
    }

    handleChangeBrand = (e) => {
        var device = this.state.device;
        device.brand = e.target.value;
        this.setState({device: device});
    }

    handleChangeModel = (e) => {
        var device = this.state.device;
        device.model = e.target.value;
        this.setState({device: device});
    }

    handleChangeSerial = (e) => {
        var device = this.state.device;
        device.serial = e.target.value;
        this.setState({device: device});
    }

    handleChangeStatus = (e) => {
        let status = parseInt(e.target.value, 10);
        this.setState({statusTo: status});
    }
    
    handleSaveClick = (e) => {
        e.preventDefault();

        if (this.props.id) {
            // Update the existing device
            api.clientsUpdate(this.props.account.orgid, this.state.deviceId, this.state.statusTo).then(response => {
                window.location = '/register';
            })
            .catch(e => {
                this.setState({error: formatError(e.response.data), hideForm: false});
            })
        } else {
            // Create a new device
            api.clientsNew(this.props.account.orgid, this.state.device).then(response => {
                window.location = '/register';
            })
            .catch(e => {
                this.setState({error: formatError(e.response.data), hideForm: false});
            })
        }
    }

    render () {
        if (!isUserAdmin(this.props.token)) {
            return (
                <div className="row">
                <AlertBox message={T('error-no-permissions')} />
                </div>
            )
        }

        if (this.state.hideForm) {
            return (
                <div className="row">
                <AlertBox message={this.state.error} />
                </div>
            )
        }

        let disabled = this.state.status ? true : false;

        return (
            <div className="row">
                <section className="row">
                    <h2>{this.state.title}</h2>
                    <div>{T('register-desc')}<br /></div>

                    <AlertBox message={this.state.error} />

                    <form>
                        <fieldset>
                            {
                                this.props.id ?
                                    <label htmlFor="id">{T('id')}:
                                        <input type="text" id="id"
                                               value={this.state.deviceId} disabled={true} />
                                    </label>
                                    :
                                    ''
                            }

                            <label htmlFor="brand">{T('brand')}:
                                <input type="text" id="brand" placeholder={T('brand-desc')}
                                    value={this.state.device.brand} onChange={this.handleChangeBrand} disabled={disabled} />
                            </label>
                            <label htmlFor="model">{T('model')}:
                                <input type="text" id="model" placeholder={T('model-desc')}
                                       value={this.state.device.model} onChange={this.handleChangeModel} disabled={disabled} />
                            </label>
                            <label htmlFor="serial">{T('serial')}:
                                <input type="text" id="serial" placeholder={T('serial-desc')}
                                       value={this.state.device.serial} onChange={this.handleChangeSerial} disabled={disabled} />
                            </label>

                            <label>
                                {this.state.status > 0 ?
                                    <select value={this.state.statusTo} onChange={this.handleChangeStatus}>
                                        {statusOptions.map(a => {
                                            return <option key={a.id} value={a.id}
                                                           selected={a.id === this.state.statusTo}>{a.name}</option>;
                                        })}
                                    </select>
                                    :
                                    ''
                                }
                            </label>
                        </fieldset>
                    </form>

                    <div>
                        <a href='/register' className="p-button--neutral">{T('cancel')}</a>
                        &nbsp;
                        <a href='/register' onClick={this.handleSaveClick} className="p-button--brand">{T('save')}</a>
                    </div>
                </section>
                <br />
            </div>
        )

    }
}

export default RegisterEdit;
