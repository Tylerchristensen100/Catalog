import { MantineProvider } from "@mantine/core";
import { AuthProvider } from "./global/auth/auth_provider";
import { DataProvider } from "./global/contexts/DataContext.jsx";
import { BrowserRouter } from "react-router-dom";
import { AppShell } from "@mantine/core";
import { Notifications } from "@mantine/notifications";

import "@mantine/core/styles.css";
import "@mantine/notifications/styles.css";
import "./App.css";

import theme from "./global/theme";
import Nav from "./global/components/nav/Nav";
import Router from "./global/Router.jsx";
import Footer from "./global/components/footer/Footer.jsx";

function App() {
  return (
    <>
      <MantineProvider theme={theme}>
        <AuthProvider>
          <DataProvider>
            <BrowserRouter basename="/admin/">
              <AppShell header={{ height: 120 }} withBorder={false}>
                <AppShell.Header>
                  <Nav />
                </AppShell.Header>
                <Notifications />
                <AppShell.Main>
                  <Router />
                </AppShell.Main>
              </AppShell>
            </BrowserRouter>
          </DataProvider>
        </AuthProvider>
      </MantineProvider>
      <Footer />
    </>
  );
}

export default App;
