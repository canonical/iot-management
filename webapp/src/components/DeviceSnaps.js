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

class DeviceRow extends Component {
    constructor(props){
        super(props)
        this.state = {
            snapSettings: null,
            settings: {},
            channel: this.props.snap.channel,
        };
    }

    handleToggle = (e) => {
        e.preventDefault()

        var snap = e.target.getAttribute('data-key')

        // Check if we need to activate or deactivate the snap
        var s = this.props.snap
        var action = 'enable';
        if (s.status === 'active') {
            action = 'disable';
        }

        api.snapsUpdate(this.props.account.orgid, this.props.device, snap, action).then(response => {
            this.props.handleMessage({message: 'Sent request to enable/disable snap: ' + snap, messageType: 'information'})
        })
    }

    handleRemove = (e) => {
        e.preventDefault()
        var snap = e.target.getAttribute('data-key')
        api.snapsRemove(this.props.account.orgid, this.props.device, snap).then(response => {
            this.props.handleMessage({message: 'Sent request to remove snap: ' + snap, messageType: 'information'})
        })
    }

    handleShowSettings = (e) => {
        e.preventDefault();
        var snap = e.target.getAttribute('data-key');
        if (this.state.snapSettings === snap) {
            this.setState({snapSettings: null, settings: {}})
        } else {
            var s = this.props.snap
            if (s.config.length===0) {
                s.config = '{}'
            }
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

        api.snapsSettingsUpdate(this.props.account.orgid, this.props.device, snap, this.state.settings).then(response => {
            this.setState({snapSettings: null})
            this.props.handleMessage({message: 'Sent request to update snap: ' + snap,
            messageType: 'information'})
        })
    }

    handleSnapRestart = (e) => {
        e.preventDefault();
        var snap = e.target.getAttribute('data-key');

        api.snapsRestart(this.props.account.orgid, this.props.device, snap).then(response => {
            this.setState({snapSettings: null})
            this.props.handleMessage({message: 'Sent restart to snap: ' + snap,
            messageType: 'information'})
        })
    }

    handleShowSnapshot = (e) => {
        e.preventDefault();
        var snap = e.target.getAttribute('data-key');
        this.setState({snapSnapshotDialog: snap})
    }

    handleSnapshotSend = (e) => {
        e.preventDefault();
        var snap = e.target.getAttribute('data-key');
        var url = this.state.snapshotUrl;
        var data = JSON.stringify({url: url});

        api.snapsSnapshot(this.props.account.orgid, this.props.device, snap, data).then(response => {
            this.setState({snapshotUrl: null})
            this.setState({snapSnapshotDialog: null})
            this.props.handleMessage({message: 'Requested shapshot of snap: ' + snap,
            messageType: 'information'})
        })
    }

    handleSnapshotCancel = (e) => {
        e.preventDefault();
        this.setState({snapSnapshotDialog: null})
    }

    handleSnapshotUpdateUrl = (e) => {
        e.preventDefault();
        this.setState({snapshotUrl: e.target.value})
    }

    renderSnapshotDialog(snap) {
        if (snap.name !== this.state.snapSnapshotDialog) {
            return
        }

        return (
            <tr>
                <td colSpan="6">
                    <form>
                        <fieldset>
                            <label htmlFor={this.state.snapSnapshotDialog}>
                                Upload url for {this.state.snapSnapshotDialog}:
                                <input type="text" rows="12" value={this.state.snapshotUrl} data-key={this.state.snapSnapshotDialog} onChange={this.handleSnapshotUpdateUrl}/>
                            </label>
                        </fieldset>
                        <button className="p-button--brand" onClick={this.handleSnapshotSend} data-key={snap.name}>{T('Send')}</button>
                        <button className="p-button--brand" onClick={this.handleSnapshotCancel} data-key={snap.name}>{T('cancel')}</button>
                    </form>
                </td>
            </tr>
        )
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
                        <button className="p-button--brand" onClick={this.handleSnapRestart} data-key={snap.name}>{T('snap-restart')}</button>
                    </form>
                </td>
            </tr>
        )
    }

    handleRefresh = (e) => {
        e.preventDefault()

        var snap = e.target.getAttribute('data-key')
        api.snapsUpdate(this.props.account.orgid, this.props.device, snap, 'refresh').then(response => {
            this.props.handleMessage({message: 'Sent request to refresh snap: ' + snap, messageType: 'information'})
        })
    }

    handleSnapChannelSwitch = (e) => {
        e.preventDefault()

        var snap = e.target.getAttribute('data-key')
        var channel = this.state.channel ? this.state.channel : ""
        var data = JSON.stringify({data: channel})

        api.snapsUpdate(this.props.account.orgid, this.props.device, snap, 'switch', data).then(response => {
            this.props.handleMessage({message: 'Sent request to switch snap channel:' + snap + " to channel: " + channel, messageType: 'information'})
        })
    }

    handleChangeChannel = (e) => {
        e.preventDefault();
        this.setState({channel: e.target.value})
    }

    renderChannel(s) {
        
        return (
            <>
            <td>
                <form>
                <input data-key={s.name} type="text" list="channels" value={this.state.channel} onChange={this.handleChangeChannel}></input>
                <datalist id="channels">
                    <option key="latest/stable" value="latest/stable">latest/stable</option>
                    <option key="latest/candidate" value="latest/candidate">latest/candidate</option>
                    <option key="latest/edge" value="latest/edge">latest/edge</option>
                    <option key="latest/beta" value="latest/beta">latest/beta</option>
                </datalist>
                </form>
            </td>
            <td>
               <button onClick={this.handleSnapChannelSwitch}  className="small" title={T("switch")} data-key={s.name}>
                    <i className="fa fa-sync" aria-hidden="true" data-key={s.name} />
                </button>
            </td>
            </>
        )
    }

    render() {
        var s = this.props.snap;
        var c = this.props.class
        var hide = this.props.hide
        return (
            <>
            <tr key={s.name} className={c}>
                <td className={hide}>
                    <button onClick={this.handleRefresh}  className="small u-float" title={T("refresh")} data-key={s.name}>
                        <i className="fa fa-sync" aria-hidden="true" data-key={s.name} />
                    </button>
                    <button onClick={this.handleToggle}  className="small" title={T("toggle-status")} data-key={s.name}>
                        <i className="fa fa-plug" aria-hidden="true" data-key={s.name} />
                    </button>
                    <button onClick={this.handleShowSnapshot} className="small" title={T("snapshot-snap")} data-key={s.name}>
                        <i className="fa fa-camera" aria-hidden="true" data-key={s.name} />
                    </button>
                    <button onClick={this.handleRemove} className="small" title={T("remove")} data-key={s.name}>
                        <i className="fa fa-times" aria-hidden="true" data-key={s.name} />
                    </button>
                </td>
                <td title={s.description}>{s.name}</td>
                {this.renderChannel(s)}
                <td>{s.version}</td>
                <td>{s.status}</td>
                <td>
                    <button className="p-button--neutral small" title={T('view-settings')} data-key={s.name} onClick={this.handleShowSettings}>
                        <i className="fa fa-sliders-h" aria-hidden="true" data-key={s.name} />
                    </button>
                </td>
            </tr>
            {this.renderSnapshotDialog(s)}
            {this.renderSettings(s)}
            </>
        )
    }
}

class DeviceSnaps extends Component {

    constructor(props) {
        super(props)
        this.state = {
            snapName: '',
            message: null,
            messageType: null,
            showInstall: false,
            hideForm: false,
        };

    }

    handleMessage = (message) => {
        this.setState(message)
    }

    handleRefreshList = (e) => {
        e.preventDefault()
        api.snapsInstallRefresh(this.props.account.orgid, this.props.device).then(response => {
            this.handleMessage({
                message: 'Sent request to list snaps',
                messageType: 'information',
            })
        })
    }

    handleSnapOnChange= (e) => {
        e.preventDefault();
        this.setState({snapName: e.target.value});
    }

    handleDialogCancel = (e) => {
        e.preventDefault();
        this.setState({showInstall: false});
    }

    handleInstall= (e) => {
        e.preventDefault()
        this.setState({showInstall: true})
    }

    handleSnapInstall = (snap) => {
        api.snapsInstall(this.props.account.orgid, this.props.device, snap).then(response => {
            this.handleMessage({
                message: 'Sent request to install snap: ' + snap,
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

    render () {
        var tableWidth = '';
        var hide = '';
        if (this.state.showInstall) {
            tableWidth = ' col-8';
            hide = 'hide';
        }
        var d = this.props.device;
        if (!d.device) {return <div>Loading...</div>};

        return (
            <div className="row">
                <h1 className="tight">{d.device.orgid} {d.device.model}</h1>
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
                                        <th className="small">{T('snap')}</th><th className="small">{T('channel')}</th><th className="xsmall"></th><th className="xsmall">{T('version')}</th><th className="xsmall">{T('status')}</th>
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
                                        <DeviceRow 
                                            key={s.name} 
                                            snap={s} 
                                            class={c} 
                                            hide={hide} 
                                            device={this.props.device} 
                                            account={this.props.account}
                                            handleMessage={this.handleMessage}></DeviceRow>
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