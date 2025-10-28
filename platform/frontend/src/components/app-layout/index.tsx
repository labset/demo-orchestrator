import { AppShell } from "@mantine/core";
import type { PropsWithChildren } from "react";
import { TopNavigation } from "./top-navigation.tsx";

const AppLayout = ({ children }: PropsWithChildren) => {
  return (
    <AppShell header={{ height: "3rem" }} padding="md">
      <AppShell.Header>
        <TopNavigation />
      </AppShell.Header>
      <AppShell.Main>{children}</AppShell.Main>
    </AppShell>
  );
};

export { AppLayout };
