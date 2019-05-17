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


import React, { Component } from 'react';
import {T, isLoggedIn} from './Utils'
import AlertBox from './AlertBox'


class Index extends Component {

  constructor(props) {

    super(props)
    this.state = {
      token: props.token || {},
    }
  }

  renderUser() {
    if (isLoggedIn(this.props.token)) {
      return <div />
    } else {
      return (
        <div>
          <a href="/login" className="p-button--brand">{T('login')}</a>
        </div>
      )
    }
  }

  renderError() {
    if (this.props.error) {
      return (
        <AlertBox message={T('user-not-found')} />
      )
    }
  }

  render() {
    return (
        <div className="row">


          <section className="row">
            <h2>{T('title')}</h2>
            <h3>{this.props.account.name}</h3>
            {this.renderError()}
            <div>
              <div className="p-card">
                {T('site-description')}
              </div>
            </div>
          </section>

          <section className="row spacer">
            {this.renderUser()}
          </section>
        </div>
    );
  }
}

export default Index;