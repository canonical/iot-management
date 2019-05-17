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

import React, { Component } from 'react'

class Pagination extends Component {

  constructor(props) {
    super(props)

    this.state = {
      page: 1,
      query: null,
      maxRecords: props.pageSize | 5,
    }
  }

  pageUp = () => {
    var pages = this.calculatePages();
    var page = this.state.page + 1;
    if (page > pages) {
        page = pages;
    }
    this.setState({page: page});
    this.signalPageChange(page);
  }

  pageDown = () => {
    var page = this.state.page - 1;
    if (page <= 0) {
        page = 1;
    }
    this.setState({page: page});
    this.signalPageChange(page);
  }

  signalPageChange(page) {
    // Signal the rows that the owner should display
    var startRow = ((page - 1) * this.state.maxRecords);

    this.props.pageChange(startRow, startRow + this.state.maxRecords);
  }

  calculatePages() {
    // Use the filtered row count when we a query has been entered
    var length = this.props.displayRows.length;

    var pages = parseInt(length / this.state.maxRecords, 10);
    if (length % this.state.maxRecords > 0) {
        pages += 1;
    }

    return pages;
  }

  renderPaging() {
    var pages = this.calculatePages();
    if (pages > 1) {
        return (
            <div className="u-float--right spacer">
                <button className="p-button--neutral small" href="" onClick={this.pageDown}>&laquo;</button>
                <span>&nbsp;{this.state.page} of {pages}&nbsp;</span>
                <button className="p-button--neutral small" href="" onClick={this.pageUp}>&raquo;</button>
            </div>
        );
    } else {
        return <div className="u-float--right" />;
    }
  }

  render() {
    return (
        <div className="col-12">
            {this.renderPaging()}
        </div>
    );
  }
}

export default Pagination
