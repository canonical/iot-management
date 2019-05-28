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
import Header from './components/Header';
import Footer from './components/Footer';
import Navigation from './components/Navigation';
import AlertBox from './components/AlertBox';
import Accounts from './components/Accounts';
import AccountEdit from './components/AccountEdit';
import Users from './components/Users';
import UserEdit from './components/UserEdit';
import Device from './components/Device';
import Register from './components/Register';
import RegisterEdit from './components/RegisterEdit';
import DeviceSnaps from './components/DeviceSnaps';
import Groups from './components/Groups';
import If from './components/If';
import Index from './components/Index';
import api from './models/api';
import {T, isLoggedIn, parseRoute, sectionNavLinks, getAccount, saveAccount, formatError} from './components/Utils'
import './sass/App.css'

import createHistory from 'history/createBrowserHistory'
const history = createHistory()

class App extends Component {

  constructor(props) {
    super(props)
    this.state = {
      location: history.location,
      token: props.token || {},
      accounts: [],
      selectedAccount: getAccount() || {},
      devices: [],
      groups: [],
      clients: [],
      client: {},
      clientObject: {},
      snaps: [],
    }

    history.listen(this.handleNavigation.bind(this))
    this.getAccounts()
  }

  handleNavigation(location) {
    this.setState({ location: location })
    window.scrollTo(0, 0)
  }

  getAccounts() {
    if (isLoggedIn(this.props.token)) {
        api.accountsList().then(response => {
            var selectedAccount = this.state.selectedAccount;

            // Reset selected if we're not filtering accounts
            // if (!this.props.token.accountFilter) {
            //   selectedAccount = {code: 'all'}
            // }

            if ((!this.state.selectedAccount.code) && (!getAccount().code)) {
              // Set to the first in the account list
              if (response.data.organizations.length > 0) {
                selectedAccount = response.data.organizations[0]
                saveAccount(selectedAccount)
              }
            }

            this.setState({accounts: response.data.organizations, selectedAccount: selectedAccount})
            this.updateDataForRoute(selectedAccount, false)
        })
    }
  }

  getDevices(code) {
    api.devicesList(code).then(response => {
        this.setState({devices: response.data.devices})
    })
  }

  getClient (code, endpoint) {
    api.clientsDetail(code, endpoint).then(response => {
        this.setState({client: response.data})
    }).catch(e => {
        this.setState({message: formatError(e.response.data), client: {}});
    })
  }

  getGroups (code) {
    api.groupsList(code).then(response => {
        this.setState({groups: response.data.groups})
    }).catch(e => {
        this.setState({message: formatError(e.response.data), clients: []});
    })
  }

  getSnaps(code, endpoint) {
    api.snapsList(code, endpoint).then(response => {
        this.setState({snaps: response.data.snaps})
    })
    .catch(e => {
        this.setState({message: formatError(e.response), snaps: []});
    })
  }

  // Get the data that's conditional on the route
  updateDataForRoute(selectedAccount, accountChanged) {
      const r = parseRoute()

      // Devices and registrations are unique for an account. So changing account may need to trigger a page change
      if (accountChanged) {
        // If we are on a clients subsection and the account is changed, navigated to the clients list
        if ((r.section==='clients') && (r.sectionId)) {
          window.location = '/clients'
          return
        }
      }

      if(r.section==='devices') {this.getDevices(selectedAccount.code)}
      if((r.section==='devices') && (r.sectionId)) {this.getClient(selectedAccount.code, r.sectionId)}
      if((r.section==='devices') && (r.sectionId) && (r.subsection==='snaps')) {this.getSnaps(selectedAccount.code, r.sectionId)}
      if(r.section==='groups') {this.getGroups(selectedAccount.code)}
  }

  handleAccountChange = (account) => {
    saveAccount(account)
    this.setState({selectedAccount: account})

    this.updateDataForRoute(account, true)
  }

  renderAccounts(sectionId, subsection) {
    
    if (!sectionId) {
      return <Accounts token={this.props.token} />
    }

    switch(sectionId) {
      case 'new':
        return <AccountEdit token={this.props.token} />
      default:
        return <AccountEdit token={this.props.token} id={sectionId} />
    }
  }

  renderUsers(sectionId, subsection) {

    if (!sectionId) {
      return <Users token={this.props.token} />
    }

    switch(sectionId) {
      case 'new':
        return <UserEdit token={this.props.token} />
      default:
        return <UserEdit token={this.props.token} id={sectionId} />
    }
  }

  renderDevices(sectionId, subsection) {

    if (!sectionId) {
      return <Register token={this.props.token} devices={this.state.devices} account={this.state.selectedAccount} />
    }

    switch(subsection) {
      case 'snaps':
        return <DeviceSnaps token={this.props.token} endpoint={sectionId} account={this.state.selectedAccount}
                  device={this.state.client} snaps={this.state.snaps} />
      default:
        return <Device token={this.props.token} endpoint={sectionId} message={this.state.message}
                  client={this.state.client} clientObject={this.state.clientObject} account={this.state.selectedAccount} />
    }
  }

  renderRegister(sectionId, subsection) {

    if (!sectionId) {
        return <Register token={this.props.token} devices={this.state.devices} account={this.state.selectedAccount} />
    }

    switch(sectionId) {
      case 'new':
        return <RegisterEdit token={this.props.token} account={this.state.selectedAccount} />
      default:
        return <RegisterEdit token={this.props.token} account={this.state.selectedAccount} id={sectionId} />
    }
  }

  renderGroups(sectionId, subsection) {

    if (!sectionId) {
        return <Groups token={this.props.token} groups={this.state.groups} account={this.state.selectedAccount} />
    }

    // switch(sectionId) {
    //   case 'new':
    //     return <RegisterEdit token={this.props.token} account={this.state.selectedAccount} />
    //   default:
    //     return <RegisterEdit token={this.props.token} account={this.state.selectedAccount} id={sectionId} />
    // }
  }

  renderSubnav(currentSection, sectionId, subsection) {
    var l = sectionNavLinks(currentSection, sectionId);
    if (l) {
      return (
        <div className="subnav">
          <nav className="p-navigation__nav p-navigation--light" role="menubar">
            <Navigation links={l} section={currentSection} sectionId={sectionId} subsection={subsection} token={this.props.token} />
          </nav>
        </div>
      );
    } else {
      return <span />
    }
  }

  render() {

    const r = parseRoute()
    console.log(r)
    console.log('---account', this.state.selectedAccount)

    return (
      <div className="App">
        <Header token={this.props.token} section={r.section}
          accounts={this.state.accounts} selectedAccount={this.state.selectedAccount}
          onAccountChange={this.handleAccountChange} />

        <section className="p-strip--image is-shallow snapcraft-banner-background">
          <div className="row">
            
          </div>
        </section>

        <If cond={isLoggedIn(this.props.token)}>
          {this.renderSubnav(r.section, r.sectionId, r.subsection)}
        </If>

        <If cond={isLoggedIn(this.props.token)}>
          <div className="page-content">
            {r.section===''? <Index token={this.props.token} account={this.state.selectedAccount} /> : ''}
            {r.section==='notfound'? <Index token={this.props.token} error /> : ''}
    
            {r.section==='devices'? this.renderDevices(r.sectionId, r.subsection) : ''}
            {r.section==='register'? this.renderRegister(r.sectionId, r.subsection) : ''}
            {r.section==='groups'? this.renderGroups(r.sectionId, r.subsection) : ''}
            {r.section==='accounts'? this.renderAccounts(r.sectionId, r.subsection) : ''}
            {r.section==='users'? this.renderUsers(r.sectionId, r.subsection) : ''}
          </div>
        </If>
        
        <If cond={!isLoggedIn(this.props.token)}>
          <div className="page-content">
              <div className="row">
                <AlertBox message={T('error-no-auth')} />
                <a href="/login" className="p-button--brand">{T('login')}</a>
              </div>
          </div>
        </If>

        <Footer />

      </div>
    );
  }
}


export default App;
