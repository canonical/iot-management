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
import {Role} from './Constants'
import {T} from './Utils';

const linksStandard  = ['devices', 'groups'];
const linksAdmin     = ['devices', 'groups', 'accounts'];
const linksSuperuser = ['devices', 'groups', 'accounts', 'users'];

class Navigation extends Component {

    link(l) {
        if (this.props.sectionId) {
            // This is the secondary menu
            return '/' + this.props.section + '/' + this.props.sectionId + '/' + l;
        } else {
            return '/' + l;
        }
    }

    render() {

        var token = this.props.token
        var links;

        if (this.props.links) {
            links = this.props.links;
        } else {
            switch(token.role) {
                case Role.Admin:
                    links = linksAdmin;
                    break;
                case Role.Superuser:
                    links = linksSuperuser;
                    break;
                case Role.Standard:
                    links = linksStandard
                    break
                default:
                    links = []
            }
        }

        return (
          <ul className="p-navigation__links" role="menu">
              {links.map((l) => {
                  var active = '';
                  if ((this.props.section === l) || (this.props.subsection === l)) {
                      active = ' active'
                  }
                  return (
                    <li key={l} className={'p-navigation__link' + active} role="menuitem"><a href={this.link(l)}>{T(l)}</a></li>
                  )
              })}
          </ul>
        );
    }
}

export default Navigation;
