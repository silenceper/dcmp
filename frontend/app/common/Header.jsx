import React from 'react';
import { Navbar, Nav, NavItem, NavDropdown, MenuItem } from 'react-bootstrap';

class Header extends React.Component {
  render() {
    return (
      <Navbar inverse>
    <Navbar.Header>
      <Navbar.Brand>
        <a href="#">统一配置平台</a>
      </Navbar.Brand>
      <Navbar.Toggle />
    </Navbar.Header>
  </Navbar>
    )
  }
}

export default Header;
