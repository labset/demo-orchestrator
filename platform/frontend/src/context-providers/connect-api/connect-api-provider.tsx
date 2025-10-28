import type { PropsWithChildren } from "react";

import { createConnectTransport } from "@connectrpc/connect-web";
import { ConnectApiContext } from "./connect-api-context";
import { createClient } from "@connectrpc/connect";
import { TodoV1TodoService } from "@labset/template-api-web-sdk";

export const ConnectApiProvider = ({ children }: PropsWithChildren) => {
  const transport = createConnectTransport({
    baseUrl: "http://localhost:8080/api",
    fetch: (input, init) => {
      return fetch(input, {
        ...init,
        credentials: "include",
        mode: "cors",
      });
    },
  });

  const todoClient = createClient(TodoV1TodoService.TodoService, transport);

  return (
    <ConnectApiContext.Provider value={{ todoClient }}>
      {children}
    </ConnectApiContext.Provider>
  );
};
