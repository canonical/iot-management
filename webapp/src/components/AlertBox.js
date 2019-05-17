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


class AlertBox extends Component {
    render() {
        if (this.props.message) {
            var c = 'p-notification--';
            if (this.props.type) {
                c = c + this.props.type;
            } else {
                c = c + 'negative';
            }

            return (
                <div className={c}>
                    <p className="p-notification__response">{this.props.message}</p>
                </div>
            );
        } else {
            return <span />;
        }
    }
}

export default AlertBox;
