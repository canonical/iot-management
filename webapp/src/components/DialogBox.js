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
import {T} from './Utils';

class DialogBox extends Component {

    render() {

        if (this.props.message) {
            return (
                <div className="p-card warning">
                    <p>{this.props.message}</p>
                    <div>
                        <button onClick={this.props.handleCancelClick} className="p-button--neutral small">
                            {T('cancel')}
                        </button>
                        <button onClick={this.props.handleYesClick} className="p-button--brand small">
                            {T('ok')}
                        </button>
                    </div>
                </div>
            );
        } else {
            return <span />;
        }
    }
}

export default DialogBox;
