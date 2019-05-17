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
import moment from 'moment';
import api from '../models/api';
import SnapDialogBox from './SnapDialogBox';
import If from './If';
import AlertBox from './AlertBox';
import {T} from './Utils';

class DeviceSnaps extends Component {

    constructor(props) {
        super(props)
        this.state = {
            message: null,
            messageType: null,
            showInstall: false,
            snapName: '',
            snapSettings: null,
            settings: {},
            hideForm: false,
        };

    }

    findSnap(name) {
        return this.props.snaps.filter( s => s.name===name)[0]
    }

    handleRefreshList = (e) => {
        e.preventDefault()
        window.location.reload()
    }

    handleInstall= (e) => {
        e.preventDefault()
        this.setState({showInstall: true})
    }

    handleSnapInstall = (snap) => {
        api.snapsInstall(this.props.account.code, this.props.device, snap).then(response => {
            this.setState({
                message: 'Sent request to install snap: ' + snap,
                messageType: 'information',
            })
        })
    }

    handleSnapOnChange= (e) => {
        e.preventDefault();
        this.setState({snapName: e.target.value});
    }

    handleRefresh = (e) => {
        e.preventDefault()

        var snap = e.target.getAttribute('data-key')
        api.snapsUpdate(this.props.account.code, this.props.device, snap, 'refresh').then(response => {
            this.setState({message: 'Sent request to refresh snap: ' + snap, messageType: 'information'});
        })
    }

    handleToggle = (e) => {
        e.preventDefault()

        var snap = e.target.getAttribute('data-key')

        // Check if we need to activate or deactivate the snap
        var s = this.findSnap(snap)
        var action = 'enable';
        if (s.status === 'active') {
            action = 'disable';
        }

        api.snapsUpdate(this.props.account.code, this.props.device, snap, action).then(response => {
            this.setState({message: 'Sent request to enable/disable snap: ' + snap, messageType: 'information'});
        })
    }

    handleRemove = (e) => {
        e.preventDefault()
        var snap = e.target.getAttribute('data-key')
        api.snapsRemove(this.props.account.code, this.props.device, snap).then(response => {
            this.setState({message: 'Sent request to remove snap: ' + snap, messageType: 'information'});
        })
    }

    handleDialogCancel = (e) => {
        e.preventDefault();
        this.setState({showInstall: false});
    }

    handleShowSettings = (e) => {
        e.preventDefault();
        var snap = e.target.getAttribute('data-key');
        if (this.state.snapSettings === snap) {
            this.setState({snapSettings: null, settings: {}})
        } else {
            var s = this.findSnap(snap)
            var settings = JSON.stringify( JSON.parse(s.config), null, 2) // pretty print
            this.setState({snapSettings: snap, settings: settings})
        }
    }

    handleSettingsChange = (e) => {
        e.preventDefault();
        this.setState({settings: e.target.value})
    }

    handleSettingsUpdate = (e) => {
        e.preventDefault();
        var snap = e.target.getAttribute('data-key');

        api.snapsSettingsUpdate(this.props.account.code, this.props.device, snap, this.state.settings).then(response => {
            this.setState({snapSettings: null, message: 'Sent request to update snap: ' + snap,
                    messageType: 'information',
            })
        })
    }

    renderShowInstall() {
        if (this.state.showInstall) {
            return (
                <SnapDialogBox message={T('confirm-snap-install')} 
                    handleTextChange={this.handleSnapOnChange}
                    handleInstallClick={this.handleSnapInstall} handleCancelClick={this.handleDialogCancel} />
            );
        }
    }

    renderSettings(snap) {
        if (snap.name !== this.state.snapSettings) {
            return
        }

        return (
            <tr>
                <td colSpan="6">
                    <form>
                        <fieldset>
                            <label htmlFor={this.state.snapSettings}>
                                Settings for {this.state.snapSettings}:
                                <textarea className="col-12" rows="4" value={this.state.settings} data-key={this.state.snapSettings} onChange={this.handleSettingsChange} />
                            </label>
                        </fieldset>
                        <button className="p-button--brand" onClick={this.handleSettingsUpdate} data-key={snap.name}>{T('update')}</button>
                    </form>
                </td>
            </tr>
        )
    }

    render () {
        var tableWidth = '';
        var hide = '';
        if (this.state.showInstall) {
            tableWidth = ' col-8';
            hide = 'hide';
        }
        var d = this.props.device;
        if (!d.info) {return <div>Loading...</div>};

        return (
            <div className="row">
                <h1 className="tight">{d.device.accountCode} {d.device.name}</h1>
                <h4 className="subtitle">{d.device.serial}</h4>

                <section className="row spacer">
                    <div className="col-12">
                        <AlertBox message={this.state.message} type={this.state.messageType} />
                    </div>
                </section>

                <If cond={!this.state.hideForm}>
                    <section className="row spacer">
                        <div className={'p-card' + tableWidth}>
                            <h3 className="p-card__title">
                                {T('installed-snaps')}

                                <div className="u-float-right">
                                    <button onClick={this.handleRefreshList} className="p-button--brand" title={T('refresh-snap-list')}>
                                        <i className="fa fa-sync" aria-hidden="true" />
                                    </button>
                                    <button onClick={this.handleInstall} className="p-button--brand" title={T('install-snap')}>
                                        <i className="fa fa-plus" aria-hidden="true" />
                                    </button>
                                </div>
                            </h3>
                            <p title={T('last-update')}>{moment(d.device.lastRefresh).format('llll')}</p>

                            <table className="p-card__content">
                                <thead>
                                    <tr>
                                        <th className={hide} />
                                        <th className="small">{T('snap')}</th><th className="xsmall">{T('version')}</th><th>{T('summary')}</th><th className="xsmall">{T('status')}</th>
                                        <th className="xsmall">{T('settings')}</th>
                                    </tr>
                                </thead>
                                {this.props.snaps.map((s) => {
                                    var c = '';
                                    if (s.name === this.state.snapSettings) {
                                        c = 'borderless'
                                    }
                                    return (
                                        <tbody>
                                        <tr key={s.name} className={c}>
                                            <td className={hide}>
                                                <button onClick={this.handleRefresh}  className="small u-float" title={T("refresh")} data-key={s.name}>
                                                    <i className="fa fa-sync" aria-hidden="true" data-key={s.name} />
                                                </button>
                                                <button onClick={this.handleToggle}  className="small" title={T("toggle-status")} data-key={s.name}>
                                                    <i className="fa fa-plug" aria-hidden="true" data-key={s.name} />
                                                </button>
                                                <button onClick={this.handleRemove} className="small" title={T("remove")} data-key={s.name}>
                                                    <i className="fa fa-times" aria-hidden="true" data-key={s.name} />
                                                </button>
                                            </td>
                                            <td title={s.description}>{s.name}</td>
                                            <td>{s.version}</td>
                                            <td>{s.summary}</td>
                                            <td>{s.status}</td>
                                            <td>
                                                <button className="p-button--neutral small" title={T('view-settings')} data-key={s.name} onClick={this.handleShowSettings}>
                                                    <i className="fa fa-sliders-h" aria-hidden="true" data-key={s.name} />
                                                </button>
                                            </td>
                                        </tr>
                                        {this.renderSettings(s)}
                                        </tbody>
                                    )
                                })}
                            </table>

                        </div>
                        {this.renderShowInstall()}
                    </section>
                </If>
            </div>
        );
    }
}

export default DeviceSnaps