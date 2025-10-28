import { createContext } from "react";
import type { Client } from "@connectrpc/connect";
import type { TodoV1TodoService } from "@labset/template-api-web-sdk";

export interface ConnectApi {
  todoClient: Client<typeof TodoV1TodoService.TodoService>;
}

export const ConnectApiContext = createContext<ConnectApi | undefined>(
  undefined,
);
