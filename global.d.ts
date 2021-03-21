declare function md5(str: string): string;
declare const axios: any;
declare interface YozCommentConfig {
  token: string;
  pageUrl: string;
  pageTitle: string;
}
declare interface Window {
  YozCommentConfig: YozCommentConfig;
}

declare interface YozCommentItem {
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
  replys: YozCommentItem[] | null;
}

declare interface YozCommentUserInfo {
  nickName: string;
  mail: string;
  site: string;
}
