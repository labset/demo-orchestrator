import "@mantine/core/styles.css";
import { createTheme, MantineProvider } from "@mantine/core";
import { Pages } from "./pages";
import { ConnectApiProvider } from "./context-providers";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

const theme = createTheme({
  // You can customize your theme here
});

const queryClient = new QueryClient();

const App = () => {
  return (
    <MantineProvider theme={theme} defaultColorScheme="auto">
      <QueryClientProvider client={queryClient}>
        <ConnectApiProvider>
          <Pages />
        </ConnectApiProvider>
      </QueryClientProvider>
    </MantineProvider>
  );
};

export { App };
