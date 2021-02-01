import axios from "axios";
import md5 from "js-md5";
import { AxiosResponse } from "../node_modules/_axios@0.21.1@axios/index";
import { timeAgo } from "./time-ago";
class KBComment {
  private container: HTMLElement | null;
  private theme?: string;
  private link?: string;
  private token?: string;
  constructor() {
    this.container = document.getElementById("kb-comment");
    if (this.container) {
      this.container.className = "kb-comment-container";
      this.theme = this.container.dataset.theme || "light";
      this.link = this.container.dataset.link || "";
      this.token = this.container.dataset.token || location.pathname;
      this.loadTheme();
      this.renderCommentForm();
      this.getList();
    } else {
      console.error("未设定渲染容器");
    }
  }

  private loadTheme() {
    return new Promise((resolve) => {
      const link = document.createElement("link");
      link.rel = "stylesheet";
      link.setAttribute("crossorigin", "anonymous");
      link.onload = resolve;
      link.href = `/themes/${this.theme}.css`;
      document.head.appendChild(link);
    });
  }

  private renderCommentForm() {
    var formContainer = document.createElement("form");
    formContainer.className = "kb-comment-form";
    formContainer.innerHTML = `
		<div class="user-info clearfix">
			<div class="input-col">
				<input type="text" maxlength="40" placeholder="昵称(必填)" />
			</div>
			<div class="input-col">
				<input type="email" placeholder="邮箱(必填)" />
			</div>
			<div class="input-col">
				<input type="url" maxlength="40" placeholder="网址" />
			</div>
		</div>
		<div class="message">
			<textarea row="6" placeholder="请输入你的留言"></textarea>
		</div>
		<div class="btn-group">
			<button type="submit">评论</button>
		</div>`;
    if (this.container) {
      this.container.appendChild(formContainer);
    }
  }

  private getList(page: number = 1) {
    var listContainer = document.createElement("div");
    listContainer.className = "kb-comment-list";
    axios.get(this.link + "/api-comment/page?token=" + this.token).then((res: AxiosResponse) => {
      if (res.data.code == 200) {
        if (this.container) {
          listContainer.innerHTML = this.renderCommentItem(res.data.data.records);
          this.container.appendChild(listContainer);
        }
      }
    });
  }

  private renderCommentItem(list: CommentItem[]): string {
    return list.reduce((html: string, item: CommentItem) => {
      html += `
				<div class="comment-item">
					<div class="comment-avatar">
						<img src="https://s.gravatar.com/avatar/${md5(item.mail)}?s=50&d=retro&r=g" />
					</div>
					<div class="comment-message clear-right">
						<div>
							<div class="comment-time">${timeAgo(item.createdAt)}</div>
							<div class="comment-nickname"><a href="${item.site}">${item.nickName}</a></div>
						</div>
						<div class="comment-content">${item.content}</div>
					</div>
					<div class="comment-replys">
						${Array.isArray(item.replys) ? this.renderCommentItem(item.replys) : ""}
					</div>
				</div>`;
      return html;
    }, "");
  }
}

new KBComment();

interface CommentItem {
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
