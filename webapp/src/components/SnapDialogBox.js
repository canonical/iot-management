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
import Pagination from './Pagination';
import {LoadingImage} from './Constants';
import {T} from './Utils';
import filesize from 'filesize'

const PAGELENGTH = 5;

class SnapDialogBox extends Component {

    constructor(props) {
        super(props)
        this.state = {
            snapName: '',
            snaps: [],
            loadingSearch: false,

            page: 1,
            startRow: 0,
            endRow: PAGELENGTH,
        };
    }

    searchStore() {
        if (this.state.snapName.length===0) {
            return
        }
        this.setState({loadingSearch: true})
        api.storeSearch(this.state.snapName).then(response => {
            if ((response.data._embedded) && (response.data._embedded['clickindex:package'])) {
                this.setState({snaps: response.data._embedded['clickindex:package'], loadingSearch: false, message: null, messageType: null})
            }
        })
    }

    handleSearchChange = (e) => {
        e.preventDefault()
        this.setState({snapName: e.target.value})
    }

    handleKeyPress = (e) => {
        if (e.key === 'Enter') {
            this.handleSearchStore(e)
        }
    }

    handleSearchStore = (e) => {
        e.preventDefault()
        this.searchStore()
    }

    handleInstall = (e) => {
        e.preventDefault()
        var snap = e.target.getAttribute('data-key')

        this.props.handleInstallClick(snap);
    }

    handleRecordsForPage = (startRow, endRow) => {
        this.setState({startRow: startRow, endRow: endRow});
    }

    renderSnaps(snaps) {
        if (snaps.length > 0) {

            return (
                <div>
                    <p>{snaps.length} snaps found</p>
                    <table>
                        <tbody>
                        {snaps.slice(this.state.startRow, this.state.endRow).map(s => {
                            return (
                                <tr key={s.snap_id} title={s.description}>
                                    <td className="small">
                                        <button data-key={s.package_name} className="p-button--neutral small" title={T("install-on-device")} onClick={this.handleInstall}>
                                            <i data-key={s.package_name} className="fa fa-cloud-upload-alt" aria-hidden="true" />
                                        </button>
                                    </td>
                                    <td className="overflow">
                                        <b>{s.package_name}</b> {s.version}
                                    </td>
                                    <td className="overflow">
                                        {s.developer_name} ({filesize(s.binary_filesize)})
                                    </td>
                                </tr>
                            )
                        })}
                        </tbody>
                    </table>
                </div>
            )
        } else {
            return <div>No snaps found.</div>
        }
    }

    render() {
        var snaps = this.state.snaps

        if (this.props.message) {
            return (
                <div className="p-card col-4">
                    <h3 className="p-card__title">{T('snap-store')}</h3>
                    <p>
                        {this.state.loadingSearch ? <img src={LoadingImage} alt={T('loading')} /> : ''}
                        {this.props.message}
                    </p>
                    <p>
                        <form className="p-search-box">
                          <input className="p-search-box__input" type="search" name="snapname" onKeyPress={this.handleKeyPress} onChange={this.handleSearchChange} placeholder={T('search-store')} />
                          <button type="reset" className="p-search-box__reset" alt="reset"><i className="p-icon--close" /></button>
                          <button type="submit" onClick={this.handleSearchStore} className="p-search-box__button" alt="search"><i className="p-icon--search" /></button>
                        </form>
                    </p>

                    <Pagination displayRows={snaps}
                            pageSize={PAGELENGTH}
                            pageChange={this.handleRecordsForPage} />

                    {this.renderSnaps(snaps)}
                    <div>
                        <button onClick={this.props.handleCancelClick} className="p-button--brand">
                            {T('close')}
                        </button>
                    </div>
                </div>
            );
        } else {
            return <span />;
        }
    }
}

export default SnapDialogBox;
