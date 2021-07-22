/* eslint-disable import/no-duplicates */

declare namespace NodeJS {
  export interface ProcessEnv {
    IMPORTS_PATH?: string;
    NODE_ENV: 'test' | 'development' | 'production';
    POLL_PERIOD?: string;
    STORAGE_PATH: string;
    ROOT_DIR: string;
  }
}

declare module '*.modules.css' {
  const cssModules: { [key: string]: string };
  export default cssModules;
}

declare module '*.css' {}
