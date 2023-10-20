/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */


export interface paths {
  "/app-configs": {
    get: operations["listAppConfig"];
  };
  "/app-configs/{id}": {
    /**
     * Find a AppConfig by ID
     * @description Finds the AppConfig with the requested ID and returns it.
     */
    get: operations["readAppConfig"];
  };
  "/health": {
    /** Simple endpoint to check if the server is up */
    get: operations["Health"];
  };
}

export type webhooks = Record<string, never>;

export interface components {
  schemas: {
    AppConfig: {
      /** Format: int64 */
      id: number;
      /** Format: date-time */
      created_at: string;
      /** Format: date-time */
      updated_at: string;
      /** Format: date-time */
      deleted_at?: string;
      app_name: string;
      environment: string;
      stack: string;
      key: string;
      value: string;
      /**
       * @default environment
       * @enum {string}
       */
      source: "stack" | "environment";
    };
    AppConfigList: {
      /** Format: int64 */
      id: number;
      /** Format: date-time */
      created_at: string;
      /** Format: date-time */
      updated_at: string;
      /** Format: date-time */
      deleted_at?: string;
      app_name: string;
      environment: string;
      stack: string;
      key: string;
      value: string;
      /**
       * @default environment
       * @enum {string}
       */
      source: "stack" | "environment";
    };
    AppConfigRead: {
      /** Format: int64 */
      id: number;
      /** Format: date-time */
      created_at: string;
      /** Format: date-time */
      updated_at: string;
      /** Format: date-time */
      deleted_at?: string;
      app_name: string;
      environment: string;
      stack: string;
      key: string;
      value: string;
      /**
       * @default environment
       * @enum {string}
       */
      source: "stack" | "environment";
    };
  };
  responses: {
    /** @description invalid input, data invalid */
    400: {
      content: {
        "application/json": {
          code: number;
          status: string;
          errors?: unknown;
        };
      };
    };
    /** @description insufficient permissions */
    403: {
      content: {
        "application/json": {
          code: number;
          status: string;
          errors?: unknown;
        };
      };
    };
    /** @description resource not found */
    404: {
      content: {
        "application/json": {
          code: number;
          status: string;
          errors?: unknown;
        };
      };
    };
    /** @description conflicting resources */
    409: {
      content: {
        "application/json": {
          code: number;
          status: string;
          errors?: unknown;
        };
      };
    };
    /** @description unexpected error */
    500: {
      content: {
        "application/json": {
          code: number;
          status: string;
          errors?: unknown;
        };
      };
    };
  };
  parameters: never;
  requestBodies: never;
  headers: never;
  pathItems: never;
}

export type $defs = Record<string, never>;

export type external = Record<string, never>;

export interface operations {

  listAppConfig: {
    parameters: {
      query: {
        /** @description what page to render */
        page?: number;
        /** @description item count to render per page */
        itemsPerPage?: number;
        app_name: string;
        environment: string;
        stack?: string;
      };
    };
    responses: {
      /** @description result AppConfig list */
      200: {
        content: {
          "application/json": components["schemas"]["AppConfigList"][];
        };
      };
      400: components["responses"]["400"];
      404: components["responses"]["404"];
      409: components["responses"]["409"];
      500: components["responses"]["500"];
    };
  };
  /**
   * Find a AppConfig by ID
   * @description Finds the AppConfig with the requested ID and returns it.
   */
  readAppConfig: {
    parameters: {
      path: {
        /** @description ID of the AppConfig */
        id: number;
      };
    };
    responses: {
      /** @description AppConfig with requested ID was found */
      200: {
        content: {
          "application/json": components["schemas"]["AppConfigRead"];
        };
      };
      400: components["responses"]["400"];
      404: components["responses"]["404"];
      409: components["responses"]["409"];
      500: components["responses"]["500"];
    };
  };
  /** Simple endpoint to check if the server is up */
  Health: {
    responses: {
      /** @description Server is reachable */
      200: {
        content: {
          "application/json": {
            status: string;
          };
        };
      };
      /** @description Server is not reachable */
      503: {
        content: never;
      };
    };
  };
}
