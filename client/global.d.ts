import { AxiosAdapter, AxiosInstance } from "../node_modules/_axios@0.21.1@axios/index";

declare global {
  export function md5(str: string): string;
  export const axios: any;
  export let returnCitySN: any;
  export interface KBCommentConfig {
    apiBase: string;
    token: string;
    theme: string;
  }

  export interface CommentItem {
    id: number;
    createdAt: string;
    articleToken: string;
    parentId: number;
    rId: number;
    nickName: string;
    mail: string;
    site: string;
    content: string;
    ip: string;
    replys: CommentItem[] | null;
  }
}
