import {
  Burger,
  Container,
  Drawer,
  Group,
  Breadcrumbs,
  Anchor,
  Flex,
} from "@mantine/core";
import { useAuth } from "../../auth/use_auth";
import classes from "./nav.module.css";
import { useDisclosure } from "@mantine/hooks";
import { useState } from "react";
import { Link, useLocation } from "react-router-dom";
import Logo from "../../../assets/logo.svg";

const Nav = () => {
  const { auth } = useAuth();
  const [opened, { open, close }] = useDisclosure(false);
  const location = useLocation();
  const crumbs = location.pathname.split("/").filter((x) => x);

  const links = [
    { link: "/", label: "Home" },
    { link: "/programs", label: "Programs" },
    { link: "/courses", label: "Courses" },
    { link: "/schools", label: "Schools" },
    auth
      ? { link: "/logout", label: "logout" }
      : { link: "/login", label: "login" },
  ];

  const [active, setActive] = useState(links[0].link);

  const items = links.map((link) => (
    <li key={link.label}>
      {" "}
      <Link
        to={link.link}
        className={classes.link}
        data-active={active === link.link || undefined}
        onClick={() => {
          setActive(link.link);
          close();
        }}
      >
        {link.label}
      </Link>
    </li>
  ));

  return (
    <>
      <header className={classes.header}>
        <Container size="xl" className={classes.inner}>
          <Link to="/">
            <Flex justify="space-between" align="center" gap="md">
              <img src={Logo} alt="Catalog Logo" className={classes.logo} />

              <h1>Catalog</h1>
            </Flex>
          </Link>

          <Group gap={5} visibleFrom="sm">
            <nav className={classes.nav}>
              <ul>{items}</ul>
            </nav>
          </Group>

          <Drawer
            opened={opened}
            onClose={close}
            title="Navigation"
            position="right"
            className={classes.drawer}
          >
            <ul>{items}</ul>
          </Drawer>

          <Burger
            opened={opened}
            onClick={open}
            hiddenFrom="sm"
            size="sm"
            className={classes.burger}
          />
        </Container>
      </header>
      <Breadcrumbs className={classes.breadcrumb} separator="/">
        {crumbs.map((item, i) => (
          <Anchor
            href={`/admin/${crumbs
              .slice(0, i + 1)
              .map((path) => `${path}/`)
              .join("")}`}
            key={i}
          >
            {decodeURIComponent(item)}
          </Anchor>
        ))}
      </Breadcrumbs>
    </>
  );
};

export default Nav;
