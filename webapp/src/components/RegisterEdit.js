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


class RegisterEdit extends Component {
    
    constructor(props) {
        super(props)
        this.state = {
            title: null,
            error: null,
            hideForm: false,
            device: {active: true},
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
        api.devicesGet(this.props.account.code, id).then(response => {
            this.setState({device: response.data.device});
        })
        .catch((e) => {
            this.setState({error: formatError(e.response.data), hideForm: true});
        })
    }

    setTitle(title) {
        this.setState({title: T(title)});
    }

    handleChangeName = (e) => {
        var device = this.state.device;
        device.name = e.target.value;
        this.setState({device: device});
    }

    handleChangeActive = (e) => {
        var device = this.state.device;
        device.active = e.target.checked;
        this.setState({device: device});
    }

    handleSaveClick = (e) => {
        e.preventDefault();

        if (this.props.id) {
            // Update the existing device
            api.devicesUpdate(this.props.account.code, this.state.device).then(response => {
                window.location = '/devices';
            })
            .catch(e => {
                this.setState({error: formatError(e.response.data), hideForm: false});
            })
        } else {
            // Create a new device
            api.devicesNew(this.props.account.code, this.state.device).then(response => {
                window.location = '/devices';
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

        return (
            <div className="row">
                <section className="row">
                    <h2>{this.state.title}</h2>

                    <AlertBox message={this.state.error} />

                    <form>
                        <fieldset>
                            <label htmlFor="name">{T('name')}:
                                <input type="text" id="name" placeholder={T('device-name-desc')}
                                    value={this.state.device.name} onChange={this.handleChangeName}/>
                            </label>
                            <label htmlFor="active">{T('active')}
                                <input type="checkbox" id="active"
                                    checked={this.state.device.active} onChange={this.handleChangeActive}/>
                            </label>
                        </fieldset>
                    </form>

                    <div>
                        <a href='/devices' className="p-button--neutral">{T('cancel')}</a>
                        &nbsp;
                        <a href='/devices' onClick={this.handleSaveClick} className="p-button--brand">{T('save')}</a>
                    </div>
                </section>
                <br />
            </div>
        )

    }
}

export default RegisterEdit;
