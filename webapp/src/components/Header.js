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
import Navigation from './Navigation';
import NavigationUser from './NavigationUser';

class Header extends Component {

    render() {
        return (
            <header className="p-navigation p-navigation--dark" role="banner">
                <div className="row">
                  <div className="p-navigation__banner">
                    <div className="p-navigation__logo">
                            <div className="nav_logo">
                                <a className="p-navigation__link" href="/">
                                    <img className="p-navigation__image" src="https://assets.ubuntu.com/v1/eb4e0ae3-iotdevice.svg" alt="IoT Management Service" />
                                    &nbsp;
                                    <h3>IoT Management</h3>
                                </a>
                            </div>

                    </div>
                    <a href="#navigation" className="p-navigation__toggle--open" title="menu" onClick={this.handleToggleMenu}>
                        <img src="/static/images/navigation-menu-plain.svg" width="30px" alt="menu" />
                    </a>
                    <nav className="p-navigation__nav">
                        <span className="u-off-screen"><a href="#navigation">Jump to site</a></span>
                        <Navigation section={this.props.section} token={this.props.token} />
                        <NavigationUser token={this.props.token} 
                            accounts={this.props.accounts} selectedAccount={this.props.selectedAccount}
                            onAccountChange={this.props.onAccountChange} />
                    </nav>
                  </div>
                </div>
            </header>
        )
    }
}

export default Header;
